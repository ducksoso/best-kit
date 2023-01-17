package main

import "fmt"

func main() {
	anno := make(map[string]string)
	anno["1"] = "111"
	anno["2"] = "222"
	modifyMap(anno)
	fmt.Println(anno)

	num := []int{1, 2, 3, 4, 5, 5, 6, 7, 9, 10}
	fmt.Println(num[1:])
}

func modifyMap(anno map[string]string) {

	anno["1"] = "123"
	anno["2"] = "234"
}
