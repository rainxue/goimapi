package main

import (
	"fmt"
	api "github.com/rainxue/goimapi/api"
)

func main() {
	fmt.Println("hello goimapi.")
	api.InitRouters()
}
