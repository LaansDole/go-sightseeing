package main

import (
	"fmt"
)

func main() {
	human(2)
}

//        *
//        *
//        *
//        * (2*n+1)
//* * * * + * * * * 4*n+1
//      *   *   loop 0 left 6 = 2n - 2 right 10 = 2n + 2
//    *       *  loop 1 left 4 = 2n - 2 - i * 2 right 12 = 2n + 2 + i * 2
//  *           loop 2 left 2 = 2n - 2 - i * 2

func human(n int) {
	head(n)
	body(n)
	leg(n)
}

func head(n int) {
	for i := 0; i < n; i++ {
		s := ""
		for j := 0; j < 2*n+1; j++ {
			if j != 2*n {
				s = s + " "
			} else {
				s = s + "*"
			}
		}
		fmt.Println(s)
	}
}
func body(n int) {
	s := ""
	for i := 0; i < 4*n+1; i++ {
		if i == 2*n {
			s = s + "+"
		} else {
			if i%2 == 0 {
				s = s + "*"
			} else {
				if i%2 == 1 {
					s = s + " "
				}
			}
		}
	}
	fmt.Println(s)
}
func leg(n int) {
	s := ""
	counter := 2*n - 1
	for i := 0; i < n; i++ {
		for j := 0; j < 2*n; j++ {
			if j == counter-1 {
				s = s + "*"
			} else {
				s = s + " "
			}
		}
		for j := 2*n + 1; j > 0; j-- {
			if j == counter {
				s = s + "*"
			} else {
				s = s + " "
			}
		}
		fmt.Println(s)
		counter = counter - 2
		s = ""
	}
}
