package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

// A.51 Web Server
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}

func webServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "john wick",
			"Message": "have a nice day",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil) // listen on port 8080 and serve requests
}

// A.52 URL Parsing
func URLParsing() {
	var urlString = "http://kalipare.com:80/hello?name=john wick&age=27"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme) // http
	fmt.Printf("host: %s\n", u.Host)       // kalipare.com:80
	fmt.Printf("path: %s\n", u.Path)       // /hello

	var name = u.Query()["name"][0] // john wick
	var age = u.Query()["age"][0]   // 27
	fmt.Printf("name: %s, age: %s\n", name, age)
}

type User struct {
	FullName string `json:"Name"`
	Age      int
}

// A.53 JSON Data
func mainJSONData() {
	// A.53.1 decode to struct
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)

	fmt.Println("================================================================")

	// A.53.2 decode to map[string]interface{}

	// map[string]interface{}
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)
	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	fmt.Println("================================================================")

	// interface{}
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

	fmt.Println("================================================================")

	// A.53.3 decode to array object
	var jsonArrayString = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`
	var dataArray []User
	var errArray = json.Unmarshal([]byte(jsonArrayString), &dataArray)
	if errArray != nil {
		fmt.Println(errArray.Error())
		return
	}

	fmt.Println("user 1:", dataArray[0].FullName)
	fmt.Println("user 2:", dataArray[1].FullName)

	fmt.Println("================================================================")

	// A.53.4 decode to json string
	var objectEncode = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonDataEncode, errEncode = json.Marshal(objectEncode)
	if errEncode != nil {
		fmt.Println(errEncode.Error())
		return
	}

	var jsonStringEncode = string(jsonDataEncode)
	fmt.Println(jsonStringEncode)
}
