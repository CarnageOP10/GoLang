package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://ineuron.ai/one-neuron"

func main() {
	fmt.Println("urls in golang")

	fmt.Println(myUrl)

	result, err := url.Parse(myUrl)
	checkNilError(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Println(qparams)

	for index,value:=range qparams{
		fmt.Println(index,":", value)
	}

}

func checkNilError(err error){
	if err != nil{
		panic(err)
	}
}
