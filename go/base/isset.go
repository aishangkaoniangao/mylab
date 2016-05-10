package main

import (
	"fmt"
)

func main() {
	//test array exists
	int_arr := [3]int{1, 2, 3}
	if len(int_arr) <= 1 {
		fmt.Println("int_arr[1] error!")
	}
	if len(int_arr) <= 3 {
		fmt.Println("int_arr[3] error!")
	}

	//test map exists
	vmap := make(map[string]string, 3)
	vmap["k1"] = "v1"
	vmap["k2"] = "v2"
	vmap["k3"] = "v3"
	//vmap["k4"] = "v4"

	if _, ok := vmap["k1"]; !ok {
		fmt.Println("vmap[k1] error")
	}
	if _, ok := vmap["k4"]; !ok {
		fmt.Println("vmap[k4] not exist")
	}

	//convert
	map1 := make(map[string]interface{}, 3)
	map1["abc"] = "abc"
	//exist or not
	if _, ok := map1["abc"]; !ok {
		fmt.Println("map1[abc] not exist")
	} else {
		fmt.Println("map1[abc] exist")
	}

	//test convert ok
	if _, ok := map1["abc"].(string); !ok {
		fmt.Println("map1[abc] convert string fail")
	} else {
		fmt.Println("map1[abc] convert string ok")
	}

	//test convert fail
	if _, ok := map1["abc"].(int); !ok {
		fmt.Println("map1[abc] convert int fail")
	} else {
		fmt.Println("map1[abc] convert int ok")
	}
}
