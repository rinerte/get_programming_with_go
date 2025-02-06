package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
    "os/exec"
)
const (
	width = 120
	height = 30
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
			b = dead
			if u[y][x] {
				b = alive
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
	
	x:= (j + width) % width
	y:= (i + height) % height
	return u[y][x]
}

func (u Universe) Neighbours(i,j int) int{
	
	count := 0
	for row := i - 1; row < i+2; row++ {
		for col := j - 1; col < j+2; col++ {
			if (col != j || row != i) && u.Alive(row, col) {
				count++
			}
		}
	}
	return count
}

func (u Universe) Next (j,i int) bool {
	n:=u.Neighbours(i,j)
	return n == 3 || n == 2 && u.Alive(i, j)
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
		time.Sleep(time.Second/24)
		a, b = b, a 
	}	
}