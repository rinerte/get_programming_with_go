package main

import (
	"fmt"
	"math"
)

type gps struct{
	current location
	destination location
	world
}

func (g gps) distance() float64{
	return g.world.distance(g.destination, g.current)
}

func (g gps) message() string{
	return fmt.Sprintf("To reach the distination remains %2.6v kilometers", g.distance())
}
type world struct {
	radius float64
}
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}
type location struct{
	name string
	long float64
	lat float64
}

type rover struct {
	gps
}
func (l location) description(){
	fmt.Printf("Name: %s, Long: %d , Lat: %d", l.name, l.long, l.lat)
}
func main(){
	mars:=gps{current:location{name:"Bradbury Landing",lat:-4.5895, long: 137.4417},destination:location{name:"Elysium Planitia",lat:4.5,long:135.9}, world:world{radius:3389.5}}

	curiosity:=rover{mars}
	fmt.Println(curiosity.message())
}