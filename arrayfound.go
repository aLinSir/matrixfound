package main

import (
	"fmt"
	"strconv"
	"os"
)

var (
	UP     int = 1
	DOWN   int = 2
	LEFT   int = 3
	RIGHT  int = 4
	X, Y   int                           //记录当前空格位置
	RESULT string = "123456780"
	Status []string                      //记录移动的状态
	PATH   []int                         //记录行动路径
)

func getRowCol(ex [][]int) (row, col int) {//获取行列
	row = len(ex)
	col = len(ex[0])
	return
}

func getStatus(ex [][]int) (result string) {//获取当前状态
	row, col := getRowCol(ex)
	for i:=0; i<row; i++ {
		for j:=0 ;j<col; j++ {
			result = result + strconv.Itoa(ex[i][j])
		}
	}
	return
}

func findZero(ex [][]int) (int, int) {//寻找空格
	row, col := getRowCol(ex)
	for i:=0; i<row; i++ {
		for j:=0 ;j<col; j++ {
			if ex[i][j] == 0 {
				X = i
				Y = j
			}
		}
	}
	return X, Y
}

func canMove(ex [][]int, dir int) bool {//判断能否移动
	row, col := getRowCol(ex)
	switch dir {
	case UP:
		if X == 0 {
			return false
		}
	case DOWN:
		if X == row - 1 {
			return false
		}
	case LEFT:
		if Y == 0 {
			return false
		}
	case RIGHT:
		if Y == col - 1 {
			return false
		}
	}
	return true
}

func startMove(ex [][]int, dir int) {//前进
	switch dir {
	case UP:
		ex[X][Y], ex[X-1][Y] = ex[X-1][Y], ex[X][Y]
		X = X - 1
	case DOWN:
		ex[X][Y], ex[X+1][Y] = ex[X+1][Y], ex[X][Y]
		X = X + 1
	case LEFT:
		ex[X][Y], ex[X][Y-1] = ex[X][Y-1], ex[X][Y]
		Y = Y - 1
	case RIGHT:
		ex[X][Y], ex[X][Y+1] = ex[X][Y+1], ex[X][Y]
		Y = Y + 1
	}
	PATH = append(PATH, dir)
	fmt.Print("添加路径成功: ")
	printPath()
}

func backMove(ex [][]int, dir int) {//回退
	switch dir {
	case UP:
		ex[X][Y], ex[X+1][Y] = ex[X+1][Y], ex[X][Y]
		X = X + 1
	case DOWN:
		ex[X][Y], ex[X-1][Y] = ex[X-1][Y], ex[X][Y]
		X = X - 1
	case LEFT:
		ex[X][Y], ex[X][Y+1] = ex[X][Y+1], ex[X][Y]
		Y = Y + 1
	case RIGHT:
		ex[X][Y], ex[X][Y-1] = ex[X][Y-1], ex[X][Y]
		Y = Y - 1
	}
	PATH = PATH[:len(PATH) - 1]
	fmt.Print("回退路径成功: ")
	printPath()
}

func isContain(status string) bool {//查询是否已经出现过相同的状态
	for _,v := range Status {
		if status == v {
			return true
		}
	}
	return false
}

func search(ex [][]int, dir int) bool {//开始搜索
	if canMove(ex, dir) && len(PATH) < 20 {
		startMove(ex, dir)
		if getStatus(ex) == RESULT  {
			fmt.Println("得到答案: ")
			return true
		}

		status := getStatus(ex)
		if isContain(status) {
			backMove(ex, dir)
			return false
		} else {
			Status = append(Status, status)//记录此次移动后的状态
			fmt.Println("未包含:", Status, X, Y)
		}

		re := search(ex, UP) || search(ex, DOWN) || search(ex, LEFT) || search(ex, RIGHT)
		if re {
			return true
		} else {
			backMove(ex, dir)
			return false
		}
	}
	return false
}

func printPath() {
	var s []string
	for _, v := range PATH {
		switch v {
		case 1:
			s = append(s, "上")
		case 2:
			s = append(s, "下")
		case 3:
			s = append(s, "左")
		case 4:
			s = append(s, "右")
		}
	}
	fmt.Println(s)
}

func main() {
	ex1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{0, 7, 8},
	}
	ex2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{0, 7, 8},
	}
	ex3 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{0, 7, 8},
	}
	ex4 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{0, 7, 8},
	}
	X, Y = findZero(ex4)
	Status = append(Status, getStatus(ex4))
	if search(ex1, UP) {
		fmt.Println(ex1[0])
		fmt.Println(ex1[1])
		fmt.Println(ex1[2])
		fmt.Println(Status)
		printPath()
		os.Exit(0)
	}
	if search(ex2, DOWN) {
		fmt.Println(ex2[0])
		fmt.Println(ex2[1])
		fmt.Println(ex2[2])
		fmt.Println(Status)
		printPath()
		os.Exit(0)
	}
	if search(ex3, LEFT) {
		fmt.Println(ex3[0])
		fmt.Println(ex3[1])
		fmt.Println(ex3[2])
		fmt.Println(Status)
		printPath()
		os.Exit(0)
	}
	if search(ex4, RIGHT) {
		fmt.Println(ex4[0])
		fmt.Println(ex4[1])
		fmt.Println(ex4[2])
		fmt.Println(Status)
		printPath()
		os.Exit(0)
	}
	fmt.Println("No answer!")

}
















/*a := search(ex, UP)
		if a == 1 {
			return 1
		} else if a == 2 {
			backMove(ex, UP)
		}

		b := search(ex, DOWN)
		if b == 1 {
			return 1
		} else if b == 2 {
			backMove(ex, DOWN)
		}

		c := search(ex, LEFT)
		if c== 1 {
			return 1
		} else if c == 2 {
			backMove(ex, LEFT)
		}

		d := search(ex, RIGHT)
		if d == 1 {
			return 1
		} else if d == 2 {
			backMove(ex, RIGHT)
		}*/