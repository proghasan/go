package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"coursename"`
	Author   string   `json:"author"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []Course{
		{"Web development using PHP", "Author 1", "123", []string{"PHP", "Laravel"}},
		{"Web development using Python", "Author 2", "22", []string{"Python", "Django"}},
		{"Web development using Java", "Author 2", "333", []string{"Java", "Spring"}},
		{"Web development using Go", "Author 1", "44", []string{"GO", "Gin"}},
		{"Other course", "Mr. X", "445", nil},
	}

	jsonData, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}

func DecodeJson() {
	jsonFromWeb := []byte(`
		 {                                                    
			"coursename": "Web development using PHP",   
			"author": "Author 1",                        
			"tags": [                                    
					"PHP",                               
					"Laravel"                            
			]                                            
        }
	`)

	var course Course

	checkValid := json.Valid(jsonFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonFromWeb, &course)
		fmt.Printf("%#v\n", course)

	} else {
		fmt.Println("JSON is not valid")
	}

	// some time we need to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	// loop
	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v also data Type %T\n", key, value, value)
	}
}
