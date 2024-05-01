package main

import (
	"math/rand"
	"sort"
	"sync"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GenRandomSortedSlice(size int) []int {
	a := make([]int, 0, size)
	mi := -100_000
	mx := 100_000
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(mx-mi)+mi)
	}
	sort.Ints(a)
	return a
}

func BatchSearch(arr []int, pill int) int {
	var (
		wg  sync.WaitGroup
		ans int
	)
	batchSize := len(arr) / 10
	out := make(chan int, 10)
	done := make(chan struct{})
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, arr []int, num int) {
			defer wg.Done()
			// early exit
			if arr[0] > pill {
				return
			}
			if arr[len(arr)-1] < pill {
				return
			}
			// normal search
			left, right := 0, len(arr)
			for left < right-1 {
				mid := (right + left) / 2
				if arr[mid] == pill {
					out <- num*batchSize + mid
					return
				}
				if arr[mid] > pill {
					right = mid
				} else {
					left = mid
				}
			}
			if arr[left] == pill {
				out <- num*batchSize + left
			}
			return
		}(&wg, arr[i*batchSize:(i+1)*batchSize], i)
	}

	go func() {
		for res := range out {
			if ans == 0 {
				ans = res
			} else {
				ans = minInt(ans, res)
			}
		}
		close(done)
	}()
	wg.Wait()
	close(out)
	<-done
	//fmt.Printf("Found answer: %d\n", ans)
	return ans
}

func NormalBinSearch(arr []int, pill int) int {
	left, right := 0, len(arr)-1
	for left < right-1 {
		mid := (right + left) / 2
		if arr[mid] == pill {
			left = mid
			break
		}
		if arr[mid] > pill {
			right = mid
		} else {
			left = mid
		}
	}
	if arr[left] == pill {
		//fmt.Printf("Found in array: %d!\n", left)
		return left
	}
	//fmt.Printf("not in array!\n")
	return -1
}

func main() {
	a := GenRandomSortedSlice(1_000_000)
	//BatchSearch(a, 1)
	NormalBinSearch(a, 1)
}
