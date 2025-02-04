package main
import "fmt"
func main(){

	const format = "%c\n"
	var pieces = []rune{'r','k','b','q','k','b','k','r'}
	var board [8][8]rune 
	for i:=0;i<8;i++{
		board[0][i] = pieces[i]
	}
	for i:=0;i<8;i++{
		board[7][i] = pieces[i]-'a'+'A'
	}
	for i:=2; i<6;i++{
		for col:=range board[i] {
			board[i][col] = ' '
		}
	}
	for column := range board[1] {
		board[1][column] = 'p'
	}
	for column := range board[6] {
		board[6][column] = 'P'
	}


	for i:=0;i<8;i++{
		fmt.Printf(format,board[i])
	}
}