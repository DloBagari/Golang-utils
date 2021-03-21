package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := `Administrators, Domain Admins (AAA, AAA A), ddfffffd(BB), ddddddddddd  (EE,EEE, EEE)`
	re := regexp.MustCompile(`(\s*\([\d\w,.\s]+\))`)

	m := re.FindAllStringSubmatch(s, -1)
	fmt.Println(m)
	for _, k := range m {
		for _, j := range k {
			fmt.Println(j)
		}
	}
	fmt.Printf("Capture value: %s", m[0][1])
}
