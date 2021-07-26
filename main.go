package main

import (
	"fmt"
	"strconv"
	"strings"
	// "string"
)

type Dimension struct {
	width, height int64
}

func main() {
	str := "250x300"
	splitted := strings.Split(str, "x")
	fmt.Println(splitted)

	split(splitted)
}

func dimension(v1, v2 int64) (dimension Dimension, err bool) {
	dimension = Dimension{width:v1, height: v2}
	fmt.Printf("value is %v, type is %T", dimension,dimension)
	return dimension, err
}

func split(splitted []string){
	var value1, value2 string
	for i := 0; i <= 1; i++ {
		value1 = splitted[0]
		value2 = splitted[1]
	}
 
	v1, err := strconv.ParseInt(value1, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	v2, err := strconv.ParseInt(value2, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n%v\n", v1, v2)

	dimension(v1, v2)
}
