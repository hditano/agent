package main

import (
	"fmt"

	"github.com/hditano/agent/helper"
)

func main() {

	response, _ := helper.RequestData()
	fmt.Println(response)

}
