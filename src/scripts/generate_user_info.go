package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const letterBytes = "0123456789"
const (
	letterIdxBits = 4
	letterIdxMask = 1<<letterIdxBits - 1
)

var headerNums = [...]string{"139", "138", "137", "136", "135", "134", "159", "158", "157", "150", "151", "152", "188", "187", "182", "183", "184", "178", "130", "131", "132", "156", "155", "186", "185", "176", "133", "153", "189", "180", "181", "177"}
var headerNumsLen = len(headerNums)

const (
	headerIdxBits = 6
	headerIdxMask = 1<<headerIdxBits - 1
)

func getHeaderIdx(cache int64) int {
	for cache > 0 {
		idx := int(cache & headerIdxMask)
		if idx < headerNumsLen {
			return idx
		}
		cache >>= headerIdxBits
	}
	return rand.Intn(headerNumsLen)
}

func randomPhone(generator *rand.Rand) string {
	b := make([]byte, 12)
	cache := generator.Int63()
	headerIdx := getHeaderIdx(cache)
	for i := 0; i < 3; i++ {
		b[i] = headerNums[headerIdx][i]
	}
	for i := 3; i < 12; {
		if cache == 0 {
			cache = generator.Int63()
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
		cache >>= letterIdxBits
	}
	return string(b)
}

func main() {
	before := time.Now()
	fmt.Println(before.Format("2006-01-02 15:04:05"))

	var wg sync.WaitGroup
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	loop := 1000000000 / nCPU

	phoneList := list.New()
	for i := 0; i < nCPU; i++ {
		//避免种子一样，生成的随机数也一样
		generator := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*1000)))
		wg.Add(1)
		go func(generator *rand.Rand, i int) {
			goPhoneList := list.New()
			for j := 0; j < loop; j++ {
				goPhoneList.PushBack(randomPhone(generator))
			}
			defer wg.Done()
			defer phoneList.PushBack(goPhoneList)
		}(generator, i)
	}
	wg.Wait()

	after := time.Now()
	fmt.Println(after.Format("2006-01-02 15:04:05"))
	fmt.Println(after.Sub(before))
	fmt.Println(phoneList.Len())
}
