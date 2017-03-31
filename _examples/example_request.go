package main

import (
	"fmt"
	"github.com/robjporter/go-functions/request"
	//"os"
)

func main() {
	//gorequest.New().Get("http://msn.com/").End(printStatus)

	handler := request.New()
	handler.SetInsecureDefaults()
	handler.SetRecorder(true)
	//hander.OverrideDefaults()
	//path, _ := os.Getwd()
	//handler.SetRecorderPath(path + "/fixtures/test")
	//handler.SetRecorderRefresh(false)

	fmt.Println("FIRST            >", handler.GetRecorderMode())
	handler.Get("http://google.com/").End(printStatus)
	fmt.Println("Response Time:   >", handler.ResponseTime)
	fmt.Println("Counter:         >", handler.RequestCounter)
	fmt.Println("********************************************")

	fmt.Println("SECOND           >", handler.GetRecorderMode())
	handler.Get("https://10.52.208.160").End(printStatus)
	fmt.Println("Response Time:   >", handler.ResponseTime)
	fmt.Println("Counter:         >", handler.RequestCounter)
	fmt.Println("********************************************")

	fmt.Println("THIRD            >", handler.GetRecorderMode())
	handler.Get("http://msn.com").End(printStatus)
	fmt.Println("Response Time:   >", handler.ResponseTime)
	fmt.Println("Counter:         >", handler.RequestCounter)
	fmt.Println("********************************************")

	fmt.Println("FOURTH           >", handler.GetRecorderMode())
	handler.Post("https://10.52.208.71/nuova").Set("Content-Type", "application/xml").Send(`<aaaLogin inName="admin" inPassword="C1sco123"/>`).End(printStatus)
	fmt.Println("Response Time:   >", handler.ResponseTime)
	fmt.Println("Counter:         >", handler.RequestCounter)
	fmt.Println("********************************************")

	fmt.Println("FINSIHED")
	handler.Terminate()
}
func printStatus(resp request.Response, body string, errs []error) {
	fmt.Println("Response Status: >", resp.Status)
	fmt.Println("Response Code:   >", resp.StatusCode)
	if errs != nil {
		fmt.Println(errs)
	}
}
