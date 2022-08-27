package actuator

import (
	"fmt"
	"testing"
)

func TestSliceExtend(t *testing.T) {
	capacity := 0
	list := make([]int, 0)
	for i := 0; i < 4096; i++ {
		list = append(list, i)
		if capacity != cap(list) {
			times := float64(cap(list)) / float64(capacity)
			differ := cap(list) - capacity
			capacity = cap(list)
			if times == 2.0 {
				fmt.Printf("capacity is: %d \t times: %.2f \n", capacity, times)
			} else {
				fmt.Printf("capacity is: %d \t times: %.2f \t differ: %d \n", capacity, times, differ)
			}
		}
	}
}
