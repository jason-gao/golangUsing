package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main()  {
	resp, err1 := http.Get("http://homev5.yundun.cc/")
	//resp, err1 := http.Get("http://baidu.com/")
	fmt.Println(resp, err1)

	fmt.Println(os.Getenv("http_proxy"))

	params := map[string][]string{
		"message": {"test PostForm"},
	}
	resp1, code, err := PostForm("http://127.0.0.1:8888/v1/postForm", params)

	fmt.Println(resp1, code, err)

}


func PostForm(postUrl string, postData map[string][]string) (string, int, error) {
	resp, err := http.PostForm(postUrl, postData)
	if err != nil {
		fmt.Println(resp, err)
		return "", resp.StatusCode, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, err
	}

	return string(body), resp.StatusCode, err
}