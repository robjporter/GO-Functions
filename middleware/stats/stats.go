package stats

import (
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/mssola/user_agent"
	"gopkg.in/kataras/iris.v6"
)

var lastSampleTime time.Time
var mem runtime.MemStats
var lastPauseNs uint64 = 0
var lastNumGc uint32 = 0

type (
	Stats struct {
		Uptime           time.Time      `json:"uptime"`
		RequestCount     uint64         `json:"requestCount"`
		Statuses         map[string]int `json:"statuses"`
		Methods          map[string]int `json:"methods"`
		Paths            map[string]int `json:"paths"`
		Platform         map[string]int `json:"platforms"`
		OS               map[string]int `json:"os"`
		BrowserName      map[string]int `json:"browsername"`
		BrowserVersion   map[string]int `json:"browserversion"`
		Mobiles          map[string]int `json:"mobiles"`
		Bots             map[string]int `json:"bots"`
		mutex            sync.RWMutex
		GoVersion        string    `json:"goversion"`
		GoOs             string    `json:"goos"`
		GoArch           string    `json:"goarch"`
		CpuNum           int       `json:"cpunum"`
		GoroutineNum     int       `json:"goroutinenum"`
		Gomaxprocs       int       `json:"gomaxprocs"`
		CgoCallNum       int64     `json:"cgocallnum"`
		MemoryAlloc      uint64    `json:"memoryalloc"`
		MemoryTotalAlloc uint64    `json:"memorytotalalloc"`
		MemorySys        uint64    `json:"memorysys"`
		MemoryLookups    uint64    `json:"memorylookups"`
		MemoryMallocs    uint64    `json:"memorymallocs"`
		MemoryFrees      uint64    `json:"memoryfrees"`
		StackInUse       uint64    `json:"memorystack"`
		HeapAlloc        uint64    `json:"heapalloc"`
		HeapSys          uint64    `json:"heapsys"`
		HeapIdle         uint64    `json:"heapidle"`
		HeapInuse        uint64    `json:"heapinuse"`
		HeapReleased     uint64    `json:"heapreleased"`
		HeapObjects      uint64    `json:"heapobjects"`
		GcNext           uint64    `json:"gcnext"`
		GcLast           uint64    `json:"gclast"`
		GcNum            uint32    `json:"gcnum"`
		GcPerSecond      float64   `json:"gcpersecond"`
		GcPausePerSecond float64   `json:"gcpausepersecond"`
		GcPause          []float64 `json:"gcpause"`
	}
)

func (s *Stats) Serve(ctx *iris.Context) {
	defer func() {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		ua := user_agent.New(string(ctx.Request.UserAgent()))
		status := strconv.Itoa(ctx.ResponseWriter.StatusCode())
		browserName, browserVersion := ua.Browser()
		plat := "API"
		if ua.Platform() != "" {
			plat = ua.Platform()
		}
		os := "NA"
		if ua.OS() != "" {
			os = ua.OS()
		}
		s.RequestCount++
		s.Statuses[status]++
		s.Methods[ctx.Method()]++
		s.Platform[plat]++
		s.OS[os]++
		s.BrowserName[browserName]++
		s.BrowserVersion[browserName+" "+browserVersion]++
		s.Mobiles[strconv.FormatBool(ua.Mobile())]++
		s.Bots[strconv.FormatBool(ua.Bot())]++
		s.GoroutineNum = runtime.NumGoroutine()
		s.CgoCallNum = runtime.NumCgoCall()
		runtime.ReadMemStats(&mem)
		s.MemoryAlloc = mem.Alloc
		s.MemoryTotalAlloc = mem.TotalAlloc
		s.MemorySys = mem.Sys
		s.MemoryLookups = mem.Lookups
		s.MemoryMallocs = mem.Mallocs
		s.MemoryFrees = mem.Frees
		s.StackInUse = mem.StackInuse
		s.HeapAlloc = mem.HeapAlloc
		s.HeapSys = mem.HeapSys
		s.HeapIdle = mem.HeapIdle
		s.HeapInuse = mem.HeapInuse
		s.HeapReleased = mem.HeapReleased
		s.HeapObjects = mem.HeapObjects
		s.GcNext = mem.NextGC
		s.GcLast = mem.LastGC
		s.GcNum = mem.NumGC
		var gcPerSecond float64
		var gcPausePerSecond float64

		countGc := int(mem.NumGC - lastNumGc)

		if lastPauseNs > 0 {
			pauseSinceLastSample := mem.PauseTotalNs - lastPauseNs
			gcPausePerSecond = float64(pauseSinceLastSample) / float64(time.Millisecond)
		}
		if lastNumGc > 0 {
			diff := float64(countGc)
			diffTime := time.Now().Sub(lastSampleTime).Seconds()
			gcPerSecond = diff / diffTime
		}

		gcPause := make([]float64, countGc)

		if countGc > 256 {
			// lagging GC pause times
			countGc = 256
		}

		for i := 0; i < countGc; i++ {
			idx := int((mem.NumGC-uint32(i))+255) % 256
			pause := float64(mem.PauseNs[idx])
			gcPause[i] = pause / float64(time.Millisecond)
		}

		s.GcPerSecond = gcPerSecond
		s.GcPausePerSecond = gcPausePerSecond
		s.GcPause = gcPause

		lastNumGc = mem.NumGC
		lastSampleTime = time.Now()
	}()
	ctx.Next()
}

func New() *Stats {
	return &Stats{
		Uptime:         time.Now(),
		Statuses:       make(map[string]int),
		Methods:        make(map[string]int),
		Paths:          make(map[string]int),
		Platform:       make(map[string]int),
		OS:             make(map[string]int),
		BrowserName:    make(map[string]int),
		BrowserVersion: make(map[string]int),
		Mobiles:        make(map[string]int),
		Bots:           make(map[string]int),
		GoVersion:      runtime.Version(),
		GoOs:           runtime.GOOS,
		GoArch:         runtime.GOARCH,
		CpuNum:         runtime.NumCPU(),
		GoroutineNum:   0,
		Gomaxprocs:     runtime.GOMAXPROCS(0),
		CgoCallNum:     0,
	}
}

func (s *Stats) Handle(ctx *iris.Context) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	ctx.Render("application/json", s, iris.RenderOptions{"charset": "UTF-8"})
}
