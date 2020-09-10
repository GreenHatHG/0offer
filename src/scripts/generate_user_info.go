package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func getPhone() string {
	headerNums := [...]string{"139", "138", "137", "136", "135", "134", "159", "158", "157", "150", "151", "152", "188", "187", "182", "183", "184", "178", "130", "131", "132", "156", "155", "186", "185", "176", "133", "153", "189", "180", "181", "177"}
	headerNumsLen := len(headerNums)
	header := headerNums[rand.Intn(headerNumsLen)]
	body := fmt.Sprintf("%08d", rand.Intn(99999999))
	phone := header + body
	return phone
}

func main() {
	newSet := mapset.NewSet()
	var wg sync.WaitGroup
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	for i := 1; i < 100; i++ {
		go func(i int) {
			wg.Add(1)
			for j := 1; j <= 100; j++ {
				newSet.Add(getPhone())
			}
			defer wg.Done()
			defer fmt.Println(strconv.Itoa(i), "Done")
		}(i)
	}
	wg.Wait()
	fmt.Println(len(newSet.ToSlice()))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
