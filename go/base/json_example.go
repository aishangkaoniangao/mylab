package main

import (
	"encoding/json"
	"fmt"
)

//Name必须是大写，如果是小写则json编码的时候不会被编码
type Person struct {
	Name       string `json:"jname"`
	Age        int32  `json:"jage"`
	Height     int32  `json:"height,omitempty"`
	HeightUnit string `json:",omitempty"`
	Weight     int32  `json:"weight,omitempty"`
	WeightUnit string `json:",omitempty"`
	Nick       string `json:"nick"`
	Hide       bool   `json:"-"`
}

//json-example: http://golang.org/pkg/encoding/json/
func main() {
	//Person be converted to json
	var p Person
	p.Name = "test-name"
	p.Age = 100
	p.Height = 160
	p.HeightUnit = "meter"
	p.Weight = 100
	p.Hide = true

	var p_str []byte
	p_str, _ = json.Marshal(p)
	fmt.Println(p, string(p_str))

	//json be converted to Person
	var p1 Person
	json.Unmarshal(p_str, &p1)
	fmt.Println(p1, p1.Nick, p1.Name)
}
