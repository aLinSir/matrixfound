package main

import "fmt"

var (
	DOWN  int = 1
	RIGHT int = 2
	Row   int
	Col   int
	X, Y  int =0, 0
	Numbers []*Number
	EX [][]int
)

type Number struct {
	Px, Py     int
	LastNumber int
	Dir        string
	NowNumber  int
	Sum        int
}

func isExist(dir int) bool {
	switch dir {
	case 1:
		if X == Row - 1 {
			return false
		}
	case 2:
		if Y == Col - 1 {
			return false
		}
	}
	return true
}

func startAdd(dir, sum int) {
	switch dir {
	case 1:
		number := Number{X+1, Y, EX[X][Y], "下", EX[X+1][Y], sum + EX[X+1][Y]}
		Numbers = append(Numbers, &number)
	case 2:
		number := Number{X, Y+1, EX[X][Y], "右", EX[X][Y+1], sum + EX[X][Y+1]}
		Numbers = append(Numbers, &number)
	}
}

func add() {
	for l:=0; l<len(Numbers); l++ {
		number := Numbers[l]
		X, Y= number.Px, number.Py
		sum := number.Sum
		for j:=1; j<3; j++ {
			if isExist(j) {
				startAdd(j, sum)
			}
		}
	}
}

func searchMax() (max, n int) {
	for i:=1; i<len(Numbers); i++ {
		number := Numbers[i]
		if number.Sum > max {
			max = number.Sum
			n = i
		}
	}
	return max, n
}

func printPath(n int) {
	var s []string
	lastNumber := Numbers[n]
	s = append(s, lastNumber.Dir)
	for i:=n-1; i>=0; i-- {
		numberMax := Numbers[n]
		number := Numbers[i]
		if number.Sum + numberMax.NowNumber == numberMax.Sum && number.NowNumber == numberMax.LastNumber && number.Dir != "" {
			s = append(s, number.Dir)
			n = i
		}
	}
	var news []string
	for i:=len(s)-1; i>=0; i-- {
		news = append(news, s[i])
	}
	fmt.Println(news)
}

func main() {
	EX = [][]int{
		{300, 500, 560, 400, 160},
		{1000, 100, 200, 340, 690},
		{600, 500, 500, 460, 320},
		{300, 400, 250, 210, 760},
	}
	/*EX = [][]int{
		{300, 500, 2560, 400},
		{1000, 100, 200, 340},
		{600, 500, 500, 460},
		{300, 400, 250, 210},
		{860, 690, 320, 760},
	}*/
	Row, Col = len(EX), len(EX[0])
	number := Number{Px:0, Py:0, NowNumber:EX[0][0], Sum:EX[0][0]}
	Numbers = append(Numbers, &number)
	add()
	max, n := searchMax()
	fmt.Println(max)
	printPath(n)
}