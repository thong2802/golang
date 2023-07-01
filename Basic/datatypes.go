package main

import "fmt"

func main() {
	var a bool = true
	var b int8 = -1
	var c int8 = 9

	fmt.Printf("%v, %T\n", a, a) // -> v ~ value | T ~ Type
	fmt.Printf("%v, %T\n", b, b) // -> v ~ value | T ~ Type

	fmt.Printf("%v, %T\n", b+c, b+c)
	fmt.Printf("%v, %T\n", b-c, b-c)
	fmt.Printf("%v, %T\n", b*c, b*c)
	fmt.Printf("%v, %T\n", b/c, b/c)
	fmt.Printf("%v, %T\n", b%c, b%c)

	// BitWise
	/* AND
	0 1 0 1 AND -> Bitwise AND cho kết quả là đúng (1) nếu cả hai bit trong cùng một cột là 1
	0 1 1 0
	--------
	0 1 0 0
	*/

	/* OR -> Trong 1 cột có ít nhất 1 số 1 thì cho kết quả là 1, còn lại là 0:
	0 1 0 1 OR
	0 1 1 0
	-----------
	0 1 1 1
	*/

	/* XOR -> XOR cho kết quả là đúng (1) nếu một và chỉ một trong các toán hạng của nó là đúng (1). Nếu cả 2 đều đúng hoặc không đúng thì nó cho kết quả là 0
	0 1 1 0 XOR
	0 0 1 1
	-------
	0 1 0 1
	*/

	/* NOT ->  lật từng bit từ 0 thành bit 1 hoặc ngược lại.
	0 1 0 0 NOT
	-------
	1 0 1 1
	*/

	/* >> ->  signed shift left operator, dịch bit qua bên phải : * 2^n
	a = 8 = 0000 0100 = 2 ^ 3 => a >> 3 = 8 * 2^3 = 64 [ or 0000 0100 = 0010 0000] = 64
	*/

	var d int8 = 10                    // 1010
	var f int8 = 3                     // 0011
	fmt.Printf("%v, %T\n", d&f, d&f)   // Bitwise AND -> 1010 & 0011 = 0010 = 2
	fmt.Printf("%v, %T\n", d|f, d|f)   // Bitwise OR  -> 1010 | 0011 = 1011 = 11
	fmt.Printf("%v, %T\n", d&^f, d&^f) // AND NOT or bit clear -> 1010 & ^0011 = 1010 & 1100 = 1000 = 8
	fmt.Printf("%v, %T\n", d^f, d^f)   // Bitwise XOR -> 1010 ^ 0011 = 1001 = 9
	fmt.Printf("%v, %T\n", ^d, ^d)     // Bitwise NOT -> 1010 = 0101 = 5
	fmt.Printf("%v, %T\n", d>>3, d>>3) // Signed shift left   = 0000 1010 >> 3 = 0101 0000 or 10 * 2 ^ 3 = 80
	fmt.Printf("%v, %T\n", d<<3, d<<3) // Signed shift right  = 0000 1010 << 3 = 0000 0001 or 10 / 2 ^ 3 = 1

	var m string = "hello world"
	var n string = "Dev"
	fmt.Printf("%v, %T\n", m+" "+n, m+n)

	var j string = "Hello World"
	var h []byte = []byte(j)
	fmt.Printf("%v, %T\n", h, h)

	var l rune = 'á'
	fmt.Printf("%v, %T\n", l, l)
}

/**
Primitive data types (kieu du lieu nguyen thuy)
	1. Boolean - luan li
	2. Numerics
		a. Integer
			- signed integer (so nguyen co dau) - int8  (8bit)  int16  (16bit)  int32  (32bit)  int64  (64bit)
			- unsigned integer (so nguyen khong dau)  - uint8 (8bit)  uint16 (16bit)  uint32 (32bit)  uint64 (64bit)
		b.Float : float32, float64
		c.Complex :complex64, complex128
	3.Text
		a. String
		b. Byte - alias unicode (UTF-8) - alias uint8
		c. Rune - UTF-32 - alias int32

*/
