package main

import (
	// "bufio"
	"fmt"
	// "log"
	// "os"
	// "reflect"
)

var pl = fmt.Println

func main() {
	// // pl("Hello Go")
	// pl("what is your name?")
	// reader := bufio.NewReader(os.Stdin)
	// name, err := reader.ReadString('\n')
	// if err == nil {
	// 	pl("Hello", name)
	// }else{
	// 	log.Fatal(err)
	// }

	// var vName string =  "Bosslady"
	// var v1, v2 = 1.2, 3.4
	// var v3 = "hello"
	// v4 := 2.4
	// v4 := 5.4

	// int, float64, bool, string, rune
	//default type 0, 0.0, true, ""

	// pl(reflect.TypeOf(25))
	// pl(reflect.TypeOf(3.14))
	// pl(reflect.TypeOf(true))
	// pl(reflect.TypeOf("Hello"))
	// pl(reflect.TypeOf('🚓'))

	//Casting
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)
}

