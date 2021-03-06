package main

import "fmt"

type stack struct {
	i    int
	data [10]int
}


func (s *stack) push(k int) {
	if s.i + 1 > 10 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop(){
	s.i--
	fmt.Println(s.data[s.i])
}

func main() {
	var s stack
	s.push(25)
	fmt.Printf("stack %v\n", s)
	s.push(14)
	fmt.Printf("stack %v\n", s)
}

