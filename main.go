package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("Homepage has %d links!\n", numLinks)

	numLinks = bytes.Count(body, []byte("<a"))
	fmt.Printf("Homepage has %d links!\n", numLinks)

	exist := strings.Contains(string(body), "疫情")
	fmt.Printf("是否存在疫情,:%v\n", exist)

	exist = bytes.Contains(body, []byte("疫情"))
	fmt.Printf("是否存在疫情:%v\n", exist)
}
