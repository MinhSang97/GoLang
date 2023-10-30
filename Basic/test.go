package main

import (
	"fmt"
)

func main() {
	nums := map[string]string{
		"key1": "ádj",
		"key2": "jhabsd",
		"key3": "absdjasbd",
	}

	for key, value := range nums {
		fmt.Println(key, value)
	}
}
