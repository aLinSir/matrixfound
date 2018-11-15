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

func initArray() (ex [][]int) {
	ex = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{0, 7, 8},
	}
	X, Y = findZero(ex)
	Status = append(Status, getStatus(ex))
	return
}

func main() {
	fmt.Println("No answer!")
	for i:=1; i<5; i++ {
		ex := initArray()
		if search(ex, i) {
			fmt.Println(ex[0])
			fmt.Println(ex[1])
			fmt.Println(ex[2])
			fmt.Println(Status)
			printPath()
			os.Exit(0)
		}
	}
	fmt.Println("No answer!")
}