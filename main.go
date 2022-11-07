package main
import (
		"myautoshop/structs"
		"fmt"
		
)

func main() {
	fmt.Println("Welcome to my Autoshop")
	p := structs.Car{}
	// p.id = "1"
	// p.model = "Audi"
	// p.color = "Black"
	// p.fourWheelDrive = true
	// p.speed = 13
	p.SetDetails("Range Rover", "black", 400, true)
	
	// fmt.Println(p.GetName())
	// fmt.Println(p.GetDetails())
	fmt.Println(p)
}