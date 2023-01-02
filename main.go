package main

import (
	"fmt"
	"github.com/judy-c/code_helper/json_helper/go_helper"
)

func main(){
	res, err := go_helper.GetGoStructDefine("{\n\t\"Legal\": true,\n\t\"Students\": [{\n\t\t\t\"StudentsID\": \"7124324354\",\n\t\t\t\"Name\": \"小明\",\n\t\t\t\"Subjects\": [{\n\t\t\t\t\t\"Name\": \"语文\",\n\t\t\t\t\t\"Score\": 100,\n\t\t\t\t\t\"Percent\": 40.00\n\t\t\t\t},\n\t\t\t\t{\n\t\t\t\t\t\"Name\": \"数学\",\n\t\t\t\t\t\"Score\": 50,\n\t\t\t\t\t\"Percent\": 40.00\n\t\t\t\t}\n\t\t\t]\n\t\t},\n\t\t{\n\t\t\t\"StudentsID\": \"7124324354\",\n\t\t\t\"Name\": \"小红\",\n\t\t\t\"Subjects\": [{\n\t\t\t\t\"Name\": \"语文\",\n\t\t\t\t\"Score\": 100,\n\t\t\t\t\"Percent\": 40.00\n\t\t\t}]\n\t\t}\n\t]\n}")
	if err == nil {
		fmt.Println(res)
	}
}
