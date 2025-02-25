package main 


import "fmt"

type turtle struct{
	location location
}

type location struct{
	x,y int
}

func (t *turtle) left(){
	t.location.x--
}
func (t *turtle) right(){
	t.location.x++
}
func (t *turtle) up(){
	t.location.y--
}
func (t *turtle) down(){
	t.location.y++
}

func main(){
	t:=turtle{}

	fmt.Printf("%v %v \n",t.location.x,t.location.y)

	t.right()
	t.up()

	fmt.Printf("%v %v \n",t.location.x,t.location.y)
}