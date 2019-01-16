package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 每个数量都在指定区间范围内
// 总数量固定, 个数不固定
func genRandQuantityList1() {
	minQ := float64(100)
	maxQ := float64(1000)
	totalQ := float64(1000000)
	restQ := totalQ
	for restQ > 0 {
		q := rand.Float64()*(maxQ-minQ) + minQ
		if q > restQ {
			q = restQ
		}
		restQ -= q
		fmt.Println(q)
	}
}

// 每个数量都在指定区间范围内
// 总数量固定, 个数固定
func genRandQuantityList2() {
	minQ := float64(100)
	maxQ := float64(1000)
	totalQ := float64(1000000)
	count := 20

	rdLis := make([]float64, count)
	sumRd := float64(0)
	for i := 0; i < count; i++ {
		rdLis[i] = rand.Float64()
		sumRd += rdLis[i]
	}

	rate := (totalQ - minQ*float64(count)) / (maxQ - minQ) / sumRd
	for _, rd := range rdLis {
		q := rd*rate*(maxQ-minQ) + minQ
		fmt.Println(q)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	genRandQuantityList1()
	fmt.Println("=================")
	genRandQuantityList2()
}
