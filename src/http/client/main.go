package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func testGet() {
	url := "http://apis.juhe.cn/simpleWeather/query?city=广州&key=f9a2b2df932ffa624665533c769aea7d"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp : %v\n", string(b))
}

func testGet2() {
	// 将参数变为变量，而不是在url中
	params := url.Values{}
	Url, err := url.Parse("http://apis.juhe.cn/simpleWeather/query")
	if err != nil {
		return
	}
	params.Set("city", "广州")
	params.Set("key", "f9a2b2df932ffa624665533c769aea7d")
	// 如果参数中有中文参数，此方法会进行URLEncode
	Url.RawQuery = params.Encode()

	urlPath := Url.String()
	fmt.Println(urlPath)

	resp, err := http.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp : %v\n", string(b))
}

func testParseJson() {
	type result struct {
		Args    string            `json:"args"`
		Headers map[string]string `json:"headers"`
		Origin  string            `json:"origin"`
		Url     string            `json:"url"`
	}

	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
	// 解析响应到结构体中
	var res result
	_ = json.Unmarshal(body, &res)
	fmt.Printf("%#v", res)
}

func testAddHeader() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://httpbin.org/get", nil)
	req.Header.Add("name", "cck")
	req.Header.Add("email", "cck@163.com")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func testPost() {
	urlstr := "http://apis.juhe.cn/simpleWeather/query"
	values := url.Values{}
	values.Set("key", "f9a2b2df932ffa624665533c769aea7d")
	values.Set("city", "广州")

	resp, err := http.PostForm(urlstr, values)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp : %v\n", string(body))
}

func testPost2() {
	// 不推荐，不太直观
	values := url.Values{
		"name": {"cck"},
		"age":  {"50"},
	}
	reqBody := values.Encode()
	// strings.NewReader 把字符串转成 io 流
	resp, _ := http.Post("http://httpbin.org/post", "text/html", strings.NewReader(reqBody))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp : %v\n", string(body))
}

func testPostJson() {
	// 创建一个 map
	data := make(map[string]interface{})
	data["name"] = "cck"
	data["age"] = 50
	bytesData, _ := json.Marshal(data)
	resp, _ := http.Post("http://httpbin.org/post", "text/html", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp : %v\n", string(body))
}

func testClient() {
	// 用 client 主要是可以主动设置一些属性
	client := http.Client{
		Timeout: time.Second * 5,
	}
	url := "http://apis.juhe.cn/simpleWeather/query?city=广州&key=f9a2b2df932ffa624665533c769aea7d"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Add("referer", "http://apis.juhe.cn/")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp : %v\n", string(body))
}

func main() {
	//testGet()
	//testGet2()
	//testParseJson()
	//testAddHeader()

	//testPost()
	//testPost2()
	//testPostJson()

	testClient()
}
