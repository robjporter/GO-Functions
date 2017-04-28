package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/robjporter/go-functions/as"
	"github.com/robjporter/go-functions/yaml"
)

var cfg = New()

type Config struct {
	data        map[string]interface{}
	keyDelim    string
	configFiles []string
	mutex       sync.Mutex
}

func New() *Config {
	return &Config{
		data:     make(map[string]interface{}),
		keyDelim: ".",
	}
}

/////////////////////////////////////////////////////
// GET
/////////////////////////////////////////////////////
func (c *Config) GetString(key string) string {
	return as.ToString(c.Get(key))
}

func (c *Config) GetBool(key string) bool {
	return as.ToBool(c.Get(key))
}

func (c *Config) GetInt(key string) int {
	return as.ToInt(c.Get(key))
}

func (c *Config) GetFloat(key string) float64 {
	return as.ToFloat(c.Get(key))
}

func (c *Config) GetTime(key string) time.Time {
	return as.ToTime(c.Get(key))
}

func (c *Config) GetStringSlice(key string) []string {
	return as.ToStringSlice(c.Get(key))
}

func (c *Config) GetStringMap(key string) map[string]interface{} {
	return as.ToStringMap(c.Get(key))
}

func (c *Config) GetStringMapString(key string) map[string]string {
	return as.ToStringMapString(c.Get(key))
}

func (c *Config) Get(key string) interface{} {
	lkey := strings.ToLower(key)
	val := c.find(lkey)
	if val == nil {
		return nil
	}
	return val
}

/////////////////////////////////////////////////////
// SET
/////////////////////////////////////////////////////

/////////////////////////////////////////////////////
// RESET
/////////////////////////////////////////////////////
func (c *Config) Reset() {
	c.data = nil
	c.data = make(map[string]interface{})
	c.configFiles = []string{}
	c.keyDelim = "."
}

/////////////////////////////////////////////////////
// FILES
/////////////////////////////////////////////////////
func ReadFiles(files ...string) { cfg.ReadFiles(files...) }
func (c *Config) ReadFiles(files ...string) error {
	for _, file := range files {
		tmp := new(Config)
		tmp.keyDelim = c.keyDelim
		err := tmp.readFile(file)
		if err != nil {
			return err
		} else {
			c.configFiles = append(c.configFiles, file)
			c.merge(tmp.data)
		}
	}
	return nil
}

func (c *Config) GetConfigFiles() []string {
	return c.configFiles
}

/////////////////////////////////////////////////////
// STRING
/////////////////////////////////////////////////////
func ReadString(data string) { cfg.ReadString(data) }
func (c *Config) ReadString(data string) {
	tmp := new(Config)
	err := tmp.readBuffer(as.ToBytes(data))
	if err != nil {
		fmt.Printf("Cannot read config data [%s]: %s\n", data, err.Error())
	} else {
		c.merge(tmp.data)
	}
}

/////////////////////////////////////////////////////
// OUTPUT
/////////////////////////////////////////////////////
func AllSettings() map[string]interface{} { return cfg.AllSettings() }
func (c *Config) AllSettings() map[string]interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	allSettings := map[string]interface{}{}
	list(&allSettings, "", c.data)
	return allSettings
}

func (c *Config) SaveYaml(filename string) {
	out, err := yaml.Marshal(c.AllSettings())
	if err == nil {
		fp, err := os.Create(filename)
		if err == nil {
			defer fp.Close()
			_, err = fp.Write(out)
		}
	}
}

func (c *Config) SaveJson(filename string) {
	out, err := json.Marshal(c.AllSettings())
	if err == nil {
		fp, err := os.Create(filename)
		if err == nil {
			defer fp.Close()
			_, err = fp.Write(out)
		}
	}
}

/////////////////////////////////////////////////////
// INTERNAL
/////////////////////////////////////////////////////
func (c *Config) find(key string) interface{} {
	path := strings.Split(key, c.keyDelim)

	val := c.searchMap(c.data, path)
	if val != nil {
		return val
	}
	return nil
}

func (c *Config) searchMap(source map[string]interface{}, path []string) interface{} {
	var next interface{}
	var ok bool

	if len(path) == 0 {
		return source
	}

	next, ok = source[path[0]]
	if ok {
		if len(path) == 1 {
			return next
		}
		switch next.(type) {
		case []interface{}:
			if len(path[1:]) > 0 {
				if _, err := strconv.Atoi(path[1:][0]); err == nil {
					tmp := as.ToSlice(next)
					p := as.ToInt(path[1:])
					return c.searchMap(as.ToStringMap(tmp[p]), path[2:])
				}
			} else {
				return c.searchMap(as.ToStringMap(next), path[1:])
			}
		case map[interface{}]interface{}:
			return c.searchMap(as.ToStringMap(next), path[1:])
		case map[string]interface{}:
			return c.searchMap(next.(map[string]interface{}), path[1:])
		default:
			return nil
		}
	}
	return nil
}

func (c *Config) searchMapWithPath(config map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return config
	}
	for i := len(path); i > 0; i-- {
		prefixKey := strings.ToLower(strings.Join(path[0:i], c.keyDelim))
		prefixKey = strings.TrimSpace(prefixKey)
		next, ok := c.data[prefixKey]
		if ok {
			if i == len(path) {
				return next
			}
			var val interface{}
			switch next.(type) {
			case map[interface{}]interface{}:
				val = c.searchMapWithPath(as.ToStringMap(next), path[i:])
			case map[string]interface{}:
				val = c.searchMapWithPath(next.(map[string]interface{}), path[i:])
			default:
			}
			if val != nil {
				return val
			}
		}
	}
	return nil
}

func (c *Config) isPathShadowedInDeepMap(config map[string]interface{}, path []string) interface{} {
	return nil
}

func (c *Config) readBuffer(buff []byte) error {
	return yaml.Unmarshal(buff, &c.data)
}

func (c *Config) readFile(filename string) error {
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return c.readBuffer(buff)
}

func castToMapStringInterface(src map[interface{}]interface{}) map[string]interface{} {
	tgt := map[string]interface{}{}
	for k, v := range src {
		tgt[fmt.Sprintf("%v", k)] = v
	}
	return tgt
}

func (c *Config) merge(src map[string]interface{}) {
	for sourceKey, sourceValue := range src {
		targetKey := keyExists(sourceKey, c.data)
		if targetKey == "" {
			c.data[sourceKey] = sourceValue
		}
		targetValue, ok := c.data[targetKey]
		if !ok {
			c.data[sourceKey] = sourceValue
		}

		sourceValueType := as.OfType(sourceValue)
		targetValueType := as.OfType(targetValue)
		if sourceValueType != targetValueType {
			continue
		}

		switch ttv := targetValue.(type) {
		case map[interface{}]interface{}:
			stv := castToMapStringInterface(ttv)
			c.merge(stv)
		case map[string]interface{}:
			c.merge(sourceValue.(map[string]interface{}))
		default:
			c.data[targetKey] = sourceValue
		}
	}
}

func keyExists(k string, m map[string]interface{}) string {
	lk := strings.ToLower(k)
	for mk := range m {
		lmk := strings.ToLower(mk)
		if lmk == lk {
			return mk
		}
	}
	return ""
}

func (c *Config) IsSet(key string) bool {
	lkey := strings.ToLower(key)
	val := c.find(lkey)
	return val != nil
}

func list(result *map[string]interface{}, prefix string, data interface{}) {
	m, ok := data.(map[string]interface{})
	if ok {
		for key, value := range m {
			nprefix := ""
			if prefix == "" {
				nprefix = fmt.Sprintf("%s", key)
			} else {
				nprefix = fmt.Sprintf("%s.%s", prefix, key)
			}
			if value != "#EMPTYVALUE#" {
				list(result, nprefix, value)
			}
		}
	} else {
		(*result)[prefix] = data
	}
}

func (c *Config) GetKeys() []string {
	m := map[string]bool{}
	m = c.flattenAndMergeMap(m, c.data, "")

	// convert set of paths to list
	a := []string{}
	for x := range m {
		a = append(a, x)
	}
	return a
}

func (c *Config) flattenAndMergeMap(shadow map[string]bool, m map[string]interface{}, prefix string) map[string]bool {
	if shadow != nil && prefix != "" && shadow[prefix] {
		// prefix is shadowed => nothing more to flatten
		return shadow
	}
	if shadow == nil {
		shadow = make(map[string]bool)
	}

	var m2 map[string]interface{}
	if prefix != "" {
		prefix += c.keyDelim
	}
	for k, val := range m {
		fullKey := prefix + k
		switch val.(type) {
		case map[string]interface{}:
			m2 = val.(map[string]interface{})
		case map[interface{}]interface{}:
			m2 = as.ToStringMap(val)
		default:
			// immediate value
			shadow[strings.ToLower(fullKey)] = true
			continue
		}
		// recursively merge to shadow map
		shadow = c.flattenAndMergeMap(shadow, m2, fullKey)
	}
	return shadow
}
