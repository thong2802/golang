package main

import "fmt"

// I. Define variable : variable là một vùng nhớ được cáp phát by programming, using để save một value of a variable
//nào đó, tại một thời điểm nhất định, và nó được access qua một cái name đã được declare cho variable đó.

// II. Global scope variable && block scope(local variable)
// note: declare global scope chỉ có thể khai báo đầy đủ : var name_variable datatype

// III. Shadow: khi declare mot biến variable global scope, truy cập biến này thông qua một block scope và khai báo overload biến này
// trong block scope thì value của biến này sẽ theo value của biến đã được ghi đè trong block scope
// Tuy nhiên khi thoát khỏi block scope thì value của global scope vẫn quay lại value cũ.

// IV: Declare and not use -> will delete

// V: Naming conversion ->camel (con lac da) -> example : var numberOfDayPerYear = 365

/*
global scope
*/
var test int = 10
var (
	n int    = 10
	m int    = 20
	h string = "abc"
)

/*
	Export global scope -> declaration in capital letters -> package outside can access
*/

var N int = 10

func main() {
	// Access global variable from block scope
	fmt.Println(test)
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(h)

	//1. Single variable
	var number int
	number = 10
	fmt.Println(number)
	fmt.Printf("%v, %T ", number, number)

	var number1 int = 11
	fmt.Println(number1)

	// 1.1 Type Inference : kieu du lieu tu suy luan
	var email = "thongnd@bap.jp"
	fmt.Println(email)

	//2. multiple variables
	/*
		2.1. Declare multiple variables the same data type
	*/
	var a, b int
	a = 1
	b = 2
	fmt.Println(a, b)

	var a1, b1 int = 3, 4
	fmt.Println(a1, b1)

	// 2.1.1 Type Inference
	var a2, b2 = 5, 6
	fmt.Println(a2, b2)

	/*
		2.2 Declare multiple variables different data type
	*/
	var (
		name    string
		address string
		age     int
	)
	name = "thong"
	address = "HCM"
	age = 10

	fmt.Println(name, address, age)

	var (
		name1    string = "thongn1"
		address1 string = "HCM"
		age1     int    = 11
	)

	fmt.Println(name1, address1, age1)

	// 2.2.1 Type Inference
	var name2, address2, age2 = "thong2", "HCM2", 12

	fmt.Println(name2, address2, age2)

	// 3. Declare shorthand
	language := "golang"
	fmt.Println(language)
}
