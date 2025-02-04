package main
import (
	"fmt"
	//"sort"
)
func main(){
	//sort.StringSlice(planets).Sort()
	//sort.Strings(planets)

	// CAPACITY = 5
	 	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"} 
	// CAPACITY = 10
	 	dwarfs2 := append(dwarfs1, "Orcus")
		dwarfs3 := append(dwarfs2, "Third")
		fmt.Println(dwarfs1)
		fmt.Println(dwarfs2)
		dwarfs2[0] = "CHANGED"
		
		fmt.Println(dwarfs1)
		fmt.Println(dwarfs2)
		fmt.Println(dwarfs3)
	fmt.Println();
		dwarfs3[0] ="Third CHANGES?"

		fmt.Println(dwarfs1)
		fmt.Println(dwarfs2)
		fmt.Println(dwarfs3)
}