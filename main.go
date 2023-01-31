package main

import (
	// "bufio"
	"fmt"
	// "strings"
	"unicode/utf8"
	// "log"
	// "os"
	// "reflect"
	// "strconv"
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

	// //Casting
	// cV1 := 1.5
	// cV2 := int(cV1)
	// pl(cV2)

	// //String to int
	// cV3 := "5000000"
	// cV4, err := strconv.Atoi(cV3)
	// pl(cV4, err, reflect.TypeOf(cV4))

	// //Int to string
	// cV5 := 5000000
	// cV6 := strconv.Itoa(cV5)
	// pl(cV6)

	// //String to float 
	// cV7 := "3.14"
	// if cV8, err := strconv.ParseFloat(cV7, 64); err == nil{
	// 	pl(cV8)
	// }

	//IF STATEMENT
	//Conditional Operators : < > >= <= == !=
	//Logical Operators : && || !

	// iAge := 8
	// if (iAge >= 1) && (iAge <= 18){
	// 	pl("Important Birthday")
	// }else if (iAge == 21) || (iAge == 50){
	// 	pl("Important Birthday")
	// }else if iAge >= 65 {
	// 	pl("Important Birthday")
	// }else{
	// 	pl("Not an Important Birthday")
	// }
	// pl("!true =", !true)

	//String
	// sV1 := "A word"
	// replacer := strings.NewReplacer("A", "Another")
	// sV2 := replacer.Replace((sV1))
	// pl(sV2)
	// pl("length :", len(sV2))
	// pl("Contains Another :", strings.Contains(sV2, "Another"))
	// pl("o index :", strings.Index(sV2, "o"))
	// pl("Replace :", strings.Replace(sV2, "o", "0", -1))
	// sV3 := "\nSome Words\n"
	// sV3 = strings.TrimSpace(sV3)
	// pl("split :", strings.Split("a-b-c-d", "-"))
	// pl("Lower :", strings.ToLower(sV2))
	// pl("Upper :", strings.ToUpper(sV2))
	// pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	// pl("suffix :", strings.HasSuffix("tacocat", "cat"))
	
	//Runes(Characters)
	rStr := "abcdefg"
	pl("Rune Count :", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr{
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}


