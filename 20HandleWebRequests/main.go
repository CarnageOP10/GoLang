package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://ineuron.ai/"

func main() {
	fmt.Println("welcome to web request")

	resp, err := http.Get(url)
	checkNilError(err)
	fmt.Println(resp)

	defer resp.Body.Close()

	dataBytes, err := ioutil.ReadAll(resp.Body)

	checkNilError(err)

	fmt.Println(string(dataBytes))
}

func checkNilError(err error){
	if err != nil{
		panic(err)
	}
}
