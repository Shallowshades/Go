package main

import "fmt"

var (
	coin  = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func value(v rune) int {
	switch {
	case v == 'e' || v == 'E':
		return 1
	case v == 'i' || v == 'I':
		return 2
	case v == 'o' || v == 'O':
		return 3
	case v == 'u' || v == 'U':
		return 4
	default:
		return 0
	}
}

func dispatchCoin() int {
	sum := 0
	for i, num := 0, 0; i < len(users); i++ {
		for _, v := range users[i] {
			num += value(v)
		}
		distribution[users[i]] = num
		sum += num
		num = 0
	}
	return coin - sum
}

func printDispatch() {
	for k, v := range distribution {
		fmt.Println(k, " : ", v)
	}
}

func main() {

	left := dispatchCoin()
	fmt.Println("rest = ", left)
	printDispatch()
}
