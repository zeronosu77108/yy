/*
1000円を両替
硬貨は最大で15枚
*/
package main
import (
	"fmt"
	"reflect"
)

var div []int

func main() {
	div = []int{5, 2, 5}
	coins := []int{0, 0, 0, 2}
	var comb [][]int
	comb = change(coins, comb )
	fmt.Println("count = ", len(comb))
	fmt.Println("comb = ", (comb))
}

func change(coins []int, comb [][]int) [][]int {
	// 15 枚のチェック
	var s int 
	s = sum(coins)
	if ( s > 15 ) {
		return comb
	}

	comb = append( comb, coins)
	if ( s == 15 ) {
		return comb
	}
	// fmt.Println("s = ", s)

	for i:=1; i<=3; i++ {

		if coins[i] <= 0 {
			continue
		}
		
		coins2 := make([]int, len(coins))
		copy(coins2, coins)

		coins2[i]--;
		coins2[i-1] += div[i-1]
		
		if include(comb, coins2) {
			continue
		}

		comb = change(coins2, comb)

	}
	return comb
}

func sum (x []int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func include (x [][]int, y []int) bool {
	for _, v := range x {
		if ( reflect.DeepEqual(v, y) ) {
			return true
		}
	}
	return false
}