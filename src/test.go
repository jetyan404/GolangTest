// package main

// //导入包
// import (
// 	"fmt"
// 	"testing"
// )

// //主函数，入口函数

// func square(num int, out chan<- int) {
// 	out <- num * num
// }

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5}
// 	results := make(chan int)

// 	for _, num := range numbers {
// 		go square(num, results)
// 	}

// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println(<-results)
// 	}
// }

// func TestSelect(t *testing.T) {
// 	ch1 := make(chan int, 1)
// 	ch2 := make(chan int, 1)

//		select {
//		case <-ch1:m
//			fmt.Println("ch1")
//		case <-ch2:
//			fmt.Println("ch2")
//		default:
//		}
//	}
//
// test2.go
package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {

	s := bytesToHex([]byte{0x3f, 0x44, 0x05})
	fmt.Println(s)

}

func bytesToHex(bytes []byte) string {
	var sb []string
	for _, aByte := range bytes {
		hex := hex.EncodeToString([]byte{aByte & 255})
		sb = append(sb, hex)
	}
	return strings.Join(sb, "")
}
