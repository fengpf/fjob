package main

import (
	_ "fjob/work/routers"
	"fmt"
	"time"
)

func main() {
	// beego.Run()

	a := time.Now()
	location, _ := time.LoadLocation("Local")
	fmt.Println(a.In(location).Unix())
}
