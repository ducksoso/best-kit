package main

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"strings"
)

func main() {

	token, err := url.QueryUnescape("os%2FgfQUlHJzXZTR2%2F20metSEqx6Np8O5CW3RT80lVFFDSuQ4Xd01qwzEaE5waEHS")
	if err != nil {
		panic(err)
	}

	token = strings.Replace(token, " ", "+", -1)

	ss := fmt.Sprintf("%x", sha256.Sum256([]byte(token)))
	fmt.Println(ss)

}

// 9465875e969db6cdd51b641d26e048c668482633704245fa6bbb0f30362b124f
