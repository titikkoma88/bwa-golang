package main

import "fmt"

func main()  {
	// hitung rata-rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var total int

	for _, score := range scores {
		total = total + score
	}

	lenght := len(scores)
	average := float64(total) / float64(lenght)

	fmt.Println(average)

	// masukan nilai yang lebih besar atau sama dengan 90 kedalam goodScore
	scores2 := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScore []int

	for _, score := range scores2 {
		if score >= 90 {
			goodScore = append(goodScore, score)
		}
	}

	fmt.Println(goodScore)

}