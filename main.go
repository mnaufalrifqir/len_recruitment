package main

import (
	"fmt"
	"len_recruitment/route"
	"time"
)

func main() {
	fmt.Println(time.Now())
	route := route.StartRoute()
	route.Start(":8000")
}
