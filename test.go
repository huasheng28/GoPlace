package main

import "fmt"

func zip(a, b []string) ([][]string, error) {

	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}

	r := make([][]string, len(a), len(a))

	for i, e := range a {
		r[i] = []string{e, b[i]}
	}

	return r, nil
}

func main() {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	b := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println(zip(a, b))
}