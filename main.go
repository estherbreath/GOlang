package main

import (
	// "bufio"
	// "crypto/rand"
	// "bufio"
	"fmt"
	// "log"
	// "os"
	// "strconv"
	// "strings"
	// "time"
	// "math/rand"
	// "math"
	// "time"
	// "time"
	// "strings"
	// "unicode/utf8"
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
	// rStr := "abcdefg"
	// pl("Rune Count :", utf8.RuneCountInString(rStr))
	// for i, runeVal := range rStr{
	// 	fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	// }

	//Working with Time
	// now := time.Now()
	// pl(now.Year(), now.Month(), now.Day())
	// pl(now.Hour(), now.Minute(), now.Second())

	//Basic Operators
	// pl("5 + 4", 5+4)
	// pl("5 - 4", 5-4)
	// pl("5 * 4", 5*4)
	// pl("5 / 4", 5/4)
	// pl("5 % 4", 5%4)
 
	//Precision with Float values
	// pl("Float Precision =", 0.1111111111111 + 0.1111111111111)

	// //Math Functions
	// pl("Abs(-10) =", math.Abs(-10))
	// pl("Pow(4, 2) =", math.Pow(4, 2))
	// pl("Sqrt(16) =", math.Sqrt(16))
	// pl("cbrt(8) =", math.Cbrt(8))
	// pl("ceil(4.4) =", math.Ceil(4.4))
	// pl("Floor(4.4) =", math.Floor(4.4))
	// pl("Round(4.4) =", math.Round(4.4))
	// pl("Log2(8) =", math.Log2(8))
	// pl("Log10(100) =", math.Log10(100))
	// //The Log of e to the power of 2
	// pl("Log(7.389) =", math.Log(math.Exp(2)))
	// pl("Max(5,4) =", math.Max(5, 4))
	// pl("Min(5,4) =", math.Min(5, 4))
	// // Convert 90 degrees to radius
	// r90 := 90 * math.Pi / 180
	// d90 := r90 * (180 / math.Pi)
	// fmt.Printf("%f radius = %f degrees\n", r90, d90)
	// pl("Sin(90) =", math.Sin(r90))
	// //There are also functions for Cos, Tan, Acos, Asin
	// // Atan, Asinh, Acosh, Atanh, Cosh, Sinh, Sincos

	//Formated Print

	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	// fmt.Printf("%s %d %c %f %t %o %x\n",
	// 	"Stuff", 1, 'A', 3.14, true, 1, 1) 
	// fmt.Printf("%9f\n", 3.14)	
	// fmt.Printf("%.2f\n", 3.141592)	
	// fmt.Printf("%9.f\n", 3.141592)
	
	// sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	// pl(sp1)

	//FOR LOOPS
	//for initialization; conditions; postStatement {BODY}
	// for x := 1; x <= 5; x++{
	// 	pl(x)
	// }

	// for x := 5; x >= 1; x--{
	// 	pl(x)
	// }

	// fx := 0
	// for fx < 5 {
	// 	pl(fx)
	// 	fx++
	// }

	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randNum := rand.Intn(50) + 1
	// for true {
	// 	fmt.Print("Guess a number between 0 and 50 :")
	// 	pl("Random Number is :", randNum)
	// 	reader := bufio.NewReader(os.Stdin)
	// 	guess, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	guess = strings.TrimSpace(guess)
	// 	iGuess, err := strconv.Atoi(guess)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if iGuess > randNum {
	// 		pl("Pick a lower value")
	// 	}else if iGuess < randNum {
	// 		pl("Pick a higher value")
	// 	}else{
	// 		pl("You Guess it")
	// 		break
	// 	}
	// }

	//for Arrays
	// aNums := []int{1, 2, 3}
	// for _, num := range aNums {
	// 	pl(num)
	// }
	
	// var arr1 [5]int
	// arr1[0] = 1
	// arr2 := [5]int{1,2,3,4,5}
	// pl("index 0 :", arr2[0])
	// pl("Arr length :", len(arr2))
	// for i := 0; i < len(arr2); i++{
	// 	pl(arr2[i])
	// }
	// for i, v := range arr2 {
	// 	fmt.Printf("%d : %d\n", i, v)
	// }
	// arr3 := [2][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// }
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		pl(arr3[i][j])
	// 	} 
	// }

	//Slices
	aStr1 := "adcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune array : %d\n", v)
	}
	//Byte array to string
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("i am a string :", bStr)
	
	//var name []datatype
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "simulated"
	sl1[4] = "universe"
	pl("Slice size :", len(sl1))
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}
	for _, x := range sl1 {
		pl(x)
	}
	sArr := [5]int{1,2,3,4,5}
	sl3 := sArr[0:2]
	pl("1st 3 :", sArr[:3])
	pl("last 3 :", sArr[2:])
	sArr[0] = 10
	pl("sl3 :", sl3)
	sl3[0] = 1
	pl("sArr :", sArr)

	//append value to array
	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)
	//Empty array
	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])


	//FUNCTIONS
	// func funcName(parameters) returnType {BODY}	


	
}


