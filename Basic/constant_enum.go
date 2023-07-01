package main

import "fmt"

const (
	red    = 1
	yellow = 2
	blue   = 3
	black  = 4
)

const (
	red1 = iota
	yellow1
	blue1
	black1
)

func main() {
	const i int16 = 42
	fmt.Printf("%v, %T\n", i, i)

	fmt.Printf("%v, %T\n", red, red)
	fmt.Printf("%v, %T\n", yellow, yellow)
	fmt.Printf("%v, %T\n", blue, blue)
	fmt.Printf("%v, %T\n", black, black)

	fmt.Printf("%v, %T\n", red1, red1)
	fmt.Printf("%v, %T\n", yellow1, yellow1)
	fmt.Printf("%v, %T\n", blue1, blue1)
	fmt.Printf("%v, %T\n", black1, black1)

}

/**
1. constant - phải cấp phát một giá trị ngay tại thời điểm biên dịch chương trình
2. enum - tap hop nhieu du lieu constant
3. tự đông khởi tạo giá trị hằng số Enum bằng iota. lưu ý: các giá trị khởi tạo trong iota có tác dụng trong 1 block, và giá trị khởi tạo mặc định ban đầu là 0
*/
