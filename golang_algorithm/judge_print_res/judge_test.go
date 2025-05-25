package judge_print_res

import (
	"fmt"
	"testing"
)

func TestFunc1(t *testing.T) {
	s := []int{1, 2, 3}

	for i := range s {
		go fmt.Println(i)
		defer fmt.Println(i)
	}
}

// 打印的序列中一定会出现2 1 0，还可能会有0 1 2随机的不按顺序的插在前面说的2 1 0序列中
