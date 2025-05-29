package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //remove the pwd
	Tags     []string `json:"tags,omitempty"` //if null dont show
}

func main() {
	fmt.Println("welcome to json")
	//encodeJson()
	decodeConsumeJson()
}

func encodeJson() {
	lcoCourses := []course{
		{"reactJS", 299, "learnCodeOnline", "abc123", []string{"web-dev", "js"}},
		{"NodeJS", 399, "learnCodeOnline", "abc143", []string{"web-dev", "js"}},
		{"fullStack", 499, "learnCodeOnline", "abc133", []string{}},
	}

	//package this as json data

	finalJson, err := json.Marshal(lcoCourses)
	CheckNilErr(err)

	finalJson2, err := json.MarshalIndent(lcoCourses, "", "\t")

	CheckNilErr(err)

	//fmt.Println(finalJson)
	fmt.Println(string(finalJson))
	fmt.Println(string(finalJson2)) //better o/p
}

func decodeConsumeJson() {
	jsonDataFromWeb := []byte(`
			{
                "coursename": "reactJS",
                "Price": 299,
                "website": "learnCodeOnline",
                "tags": ["web-dev","js"]
        	}
	`)

	var lcoCourse course

	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("josn was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
	} else {
		fmt.Println("not valid")
	}

	fmt.Printf("%#v\n", lcoCourse)

	//u just want to add data to key value and not struc

	var myonlinedata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myonlinedata)
	fmt.Println(myonlinedata)

	for key, val := range myonlinedata {
		fmt.Printf("key: %v and value: %v and type: %T\n", key, val, val)
	}
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
