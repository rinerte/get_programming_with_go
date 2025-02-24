package main

import(
	"fmt"
	"math/rand"
	"time"
)
type animal struct{
	name string
	food []string
	movings []string
}

type movingEater interface{
	eat()
	move()
}

func (a animal) String() string{
	return fmt.Sprintf("%v",a.name)
}
func (a animal) eat(){
	index:= rand.Intn(len(a.food))
	fmt.Printf("%v ест %v \n",a.name,a.food[index])
}
func (a animal) move(){
	index:= rand.Intn(len(a.movings))
	fmt.Printf("%v %v\n",a.name,a.movings[index])
}

func do(a movingEater){
	index:= rand.Intn(10)
	if(index>7){
		a.move()
	} else{
		a.eat()
	}
}
func main(){
	rand.Seed(time.Now().UnixNano())

	hedgehog:= animal{"Ёж",[]string{"гриб","червяка","соседа"},[]string{"бегает по кругу","сидит на месте","чешет за ухом"}}
	lion:= animal{"Лев",[]string{"конфеты","мороженое","чипсы"},[]string{"ползает","ходит на двух лапах","кувыркается"}}
	bunny:= animal{"Заяц",[]string{"кровавую плоть","живых детей"},[]string{"прыгает","телепортируется в клетки к соседям"}}

	animals:=[]animal{hedgehog,lion,bunny}
	
	for time:=0;time<24;time++{
		if(time>8 && time<19){
			index:= rand.Intn(3)
			fmt.Printf("Время:%v:00, ",time)
			do(animals[index])	
		} else if(time>=19){
			fmt.Printf("Время:%v:00, - все оставшиеся в живых спят\n",time)
		} else {
			fmt.Printf("Время:%v:00, - все спят\n",time)
		}
	}
}