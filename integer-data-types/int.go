package main

import "fmt"

// int8 -> the set of all signed 8-bit integers (-128 to 127)
// int16 -> the set of all signed 16-bit integers (-32768 to 32767)
// int32 -> the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64 -> the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

// uint8 -> the set of all unsigned 8-bit integers (0 to 255)
// uint16 -> the set of all unsigned 16-bit integers (0 to 65535)
// uint32 -> the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64 -> the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int -> either 32 or 64 bits
// uint -> either 32 or 64 bits
// rune -> alias for int32
// byte -> alias for uint8

func main() {
	var SmallPositiveNumber uint8
	SmallPositiveNumber = 42
	fmt.Println(SmallPositiveNumber) // 42

	var SmallNegativeNumber int8
	SmallNegativeNumber = -42
	fmt.Println(SmallNegativeNumber) // -42

	// type case from int8 to int16
	var SmallNegativeNumber16 int16
	SmallNegativeNumber16 = int16(SmallNegativeNumber)
	fmt.Println(SmallNegativeNumber16) // -42

	var myByte byte = 'B'
	fmt.Println(myByte) // 66
	//byte is mainly use to represent raw  data
	// for example, when you read a file, you read it as a byte stream
	// byte is also used to represent ASCII characters
	// since byte is an alias for uint8, it can only represent values from 0 to 255

	// rune is used to represent a Unicode code point int32
	var myRune rune = 'B'
	fmt.Println(myRune) // 66
	// rune is used to represent a Unicode code point
	myRune = '🐹'
	fmt.Println(myRune) // 128057
}
