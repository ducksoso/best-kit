package main

import (
	"fmt"
	"time"
)

func main() {
	anno := make(map[string]string)
	anno["1"] = "111"
	anno["2"] = "222"
	modifyMap(anno)
	fmt.Println(anno)

	num := []int{1, 2, 3, 4, 5, 5, 6, 7, 9, 10}
	fmt.Println(num[1:]) // 左开右闭

	nn := float64(8) / 1024 / 1024
	fmt.Println(nn)

	fmt.Printf("%f\n", 1.25621772313)
	fmt.Printf("%q", "123")

	fmt.Println()
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC().Format("2006-01-02"))

}

func modifyMap(anno map[string]string) {
	anno["1"] = "123"
	anno["2"] = "234"
}
