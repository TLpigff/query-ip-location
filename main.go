package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ResponseBody struct {
	Resultcode string `json:"resultcode"`
	Reason     string `json:"reason"`
	Result     struct {
		Country  string `json:"Country"`
		Province string `json:"Province"`
		City     string `json:"City"`
		Isp      string `json:"Isp"`
	} `json:"result"`
	ErrorCode int `json:"error_code"`
}

type RequestBody struct {
	Ip string `json:"ip"`
	Key string `json:"key"`
}

const urlFmt = "http://apis.juhe.cn/ip/ipNew?ip=%s&key=efd5f8d1a4bbb72d960631a57ee8d5b8"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr,"must specified ip")
		os.Exit(1)
	}

	response, err := http.Get(fmt.Sprintf(urlFmt, os.Args[1]))
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var resp ResponseBody
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &resp)
	if err != nil {
		panic(err)
	}

	if resp.Resultcode != "200" {
		fmt.Fprintf(os.Stderr,resp.Reason)
		os.Exit(1)
	}

	fmt.Printf("IP所属国家:%s\n",resp.Result.Country)
	fmt.Printf("IP所属省份:%s\n",resp.Result.Province)
	fmt.Printf("IP所属城市:%s\n",resp.Result.City)
	fmt.Printf("ISP供应商:%s\n",resp.Result.Isp)
}
