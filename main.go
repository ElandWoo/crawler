package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://thepaper.cn/")
	if err != nil {
		fmt.Println("Fetch url error:%v.", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error status code:#{resp.StatusCode}")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read content failed:%v.", err)
		return
	}

	fmt.Println("Body:", string(body))
}
