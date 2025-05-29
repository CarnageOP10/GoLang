package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web verb")

	//PerformGetReq()
	//PerformPostJsonReq()
	PerformPostFormreq()

}

func PerformGetReq() {
	const myUrl = "http://localhost:3635/get"

	resp, err := http.Get(myUrl)
	CheckNilErr(err)

	defer resp.Body.Close()

	fmt.Println("Status Code: ", resp.StatusCode)
	fmt.Println("content length: ", resp.ContentLength)

	content, err := ioutil.ReadAll(resp.Body)

	CheckNilErr(err)

	// fmt.Println(string(content))
	var responseString strings.Builder
	bytecount, _ := responseString.Write(content)

	fmt.Println(bytecount)
	fmt.Println(responseString.String())

}

func PerformPostJsonReq() {
	const myUrl = "http://localhost:3635/post"

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"lets go with golang",
			"price" : 0,
			"name" : "parth"
		}
	`)

	resp, err := http.Post(myUrl, "applpicationi/json", requestBody)

	CheckNilErr(err)

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	CheckNilErr(err)

	fmt.Println(string(content))
}

func PerformPostFormreq() {
	const myUrl = "http://localhost:3635/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "Singh")
	data.Add("email", "gg@gg.ai")

	resp, _ := http.PostForm(myUrl, data)

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(content))
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
