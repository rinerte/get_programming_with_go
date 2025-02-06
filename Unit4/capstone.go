package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
    "os/exec"
)
const (
	width = 80
	height = 15
	alive = '*'
	dead = ' '
)
type Universe [][]bool

func NewUniverse() Universe {
	u:= make(Universe, height)
	for row:= range u {
		u[row] = make([]bool, width)
	}
	return u;
}

// func (u Universe) Show(){
// 	for i:=0;i<height; i++ {
// 		for j:=0;j<width;j++{
// 			if (u[i][j]){
// 				fmt.Printf("%c",alive)
// 			} else {
// 				fmt.Printf("%c",dead)
// 			}
			
// 		}
// 		fmt.Print("|\n")
// 	}
// }
func (u Universe) Show() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print(u.String())
}
func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width+1)*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}
func (u Universe) Seed(){

	rand.Seed(time.Now().UnixNano())

	for i:=0;i<height; i++ {
		for j:=0;j<width;j++{
			switch n:=rand.Intn(4);n{
				case 0: u[i][j] = true
				default: u[i][j] = false
			}
		}
	}
}
func (u Universe) Alive(i,j int) bool{
	if(i==-1){
		i = height - 1
	}
	if(j==-1){
		j = width - 1
	}
	if(i==height){
		i = 0
	}
	if(j==width){
		j = 0
	}
	return u[i][j]
}

func (u Universe) Neighbours(i,j int) int{
	count:=0
	for row:=i-1; row<i+2; row++{
		for col:=j-1;col<j+2;col++{
			if((col!=j || row!=i)&&u.Alive(i,j)){
				count++
			}
		}
	}
	return count
}

func (u Universe) Next (j,i int) bool {
	count:=u.Neighbours(i,j)
	alive:=u.Alive(i,j)

	if(alive){
		if(count!=2 && count!=3) {
			return false
		}
	} else{
		if (count==3){
			return true;
		}
	}
	return alive
}

func Step(a,b Universe){
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}
func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}
func main(){
	a:=NewUniverse()	
	b:=NewUniverse()

	a.Seed();
	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second)
		a, b = b, a 
	}	
}