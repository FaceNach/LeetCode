package main

import "fmt"

func largestGoodInteger(num string) string {
	if num == "" {
		return "Cannot pass empty string"

	}

	out := ""

	for i := 0; i+2 < len(num); i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			partial := num[i : i+3]
			if partial > out {
				out = partial
			}
		}
	}

	return out
}

func main() {

	text := "556667789999"

	result := largestGoodInteger(text)

	fmt.Println(result)
}
