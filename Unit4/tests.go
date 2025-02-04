package main
import "fmt"
func main(){
	var board [4][4]string
	board[0][0] = "r"
	board[0][3] = "r"
	for column := range board[1] {
		board[1][column] = "p"
	}
	for i:=0;i<4;i++{
		fmt.Println(board[i])
	}
}