package main
import(
	"fmt"
	"encoding/json"
)

type location struct {
	Name string `json:"NAMEE"`
	Lat float64 `json:"LATITUTDDE"`
	Long float64 `json: "LUNGITUDUI"`
	}
	

func main(){
	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}
	var bytes,_ = json.MarshalIndent(locations, "", " ")
	fmt.Println(string(bytes))
}