package main

import (
	"strconv"
	"fmt"
	"os"
)

var (
	UP     int = 1
	DOWN   int = 2
	LEFT   int = 3
	RIGHT  int = 4
	X, Y   int                           //记录当前空格位置
	RESULT string = "123456780"
	PATH   []string                      //记录状态
	AllStatus []*Status                  //记录移动前后状态和方向信息
)

type Status struct {
	NowStatus   string                    //移动后状态
	Dir         int
	LastStatus  string                    //移动前状态
}

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
}

func isContain(status string) bool {//查询是否已经出现过相同的状态
	for _,v := range PATH {
		if status == v {
			return true
		}
	}
	return false
}

func printPath(nowstatus string, l int, laststatus string) {
	//确定最后的位置
	var k int
	for i:=l; l<=len(AllStatus); i++ {
		status := AllStatus[i]
		if status.NowStatus == nowstatus && status.LastStatus == laststatus {
			k = i
			break
		}
	}

	//开始往前寻找路径
	var path []int
	path = append(path, AllStatus[k].Dir)
	for i:=k-1; i>=0; i-- {
		if i == 0 {
			break
		}
		statusNow := AllStatus[k]
		statusOld := AllStatus[i]
		if statusNow.LastStatus == statusOld.NowStatus {
			k = i
			path = append(path, AllStatus[k].Dir)
		}
	}

	//将路径转换成文字字符串打印
	var s []string
	for i:=len(path)-1; i>=0; i-- {
		switch path[i] {
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

func reset(status string) [][]int {
	var ex [][]int
	var e []int
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			v, _ := strconv.Atoi(string(status[i*3+j]))
			e = append(e, v)
		}
		ex = append(ex, e)
		e = make([]int,0)
	}
	X, Y = findZero(ex)
	return ex
}

func hirosearch() bool {//开始搜索
	for l:=0; l<len(AllStatus); l++ {
		status := AllStatus[l]
		if status.NowStatus == RESULT {
			fmt.Println("得到答案: ")
			return true
		}
		ex := reset(status.NowStatus)

		for i:=1; i<5; i++ {//朝四个方向移动
			if canMove(ex, i) {
				laststatus := status.NowStatus
				startMove(ex, i)
				nowstatus := getStatus(ex)
				if nowstatus == RESULT {
					fmt.Println("得到答案: ")
					AllStatus = append(AllStatus, &Status{nowstatus, i, laststatus})
					printPath(nowstatus, l, laststatus)
					return true
				}

				if isContain(nowstatus) {
					backMove(ex, i)
					continue
				}

				PATH = append(PATH, nowstatus)
				AllStatus = append(AllStatus, &Status{nowstatus, i, laststatus})
				//context, _ := json.Marshal(&AllStatus)
				//fmt.Println(string(context))
				backMove(ex, i)
			}
		}
	}
	return false
}



func main() {
	ex := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{8, 7, 0},
	}
	/*ex := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 5, 8},
	}*/
	status := getStatus(ex)
	PATH = append(PATH, status)
	AllStatus = append(AllStatus, &Status{NowStatus:status})
	X, Y = findZero(ex)
	if hirosearch() {
		os.Exit(0)
	}
	fmt.Println("No answer!")

}


