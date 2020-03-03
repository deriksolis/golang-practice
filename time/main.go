package main

import (
	"fmt"
	"time"
)

func main() {

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)

	fmt.Println(now)
}
