package main

import "fmt"

func main() {
	fmt.Println("maps in go")
	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["ruby"] = "RubyLang"
	languages["JAVA"] = "java"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	//deleting
	delete(languages, "ruby")
	fmt.Println(languages)

	//loops 
	for key, value := range languages{
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

}
