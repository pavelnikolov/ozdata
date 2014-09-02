package main

import (
	"flag"
	"fmt"
	"github.com/tomjowitt/ozdata/lib"
)

func main() {

	var postcode = flag.Int64("p", 0, "A valid Australian postcode")
	flag.Parse()

	fmt.Println("")
	fmt.Println("Welcome to Ozdata")
	fmt.Println("")

	suburbData, err := ozdata.NewSuburbData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Getting suburb By postcode:", *postcode)
	subByPostcode, err := suburbData.GetSuburbByPostcode(*postcode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(subByPostcode)
	fmt.Println("")

	return

}
