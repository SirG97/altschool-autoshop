package structs
import "fmt"

type Car struct {
	// id string
	model          string
	color          string
	speed          uint
	fourWheelDrive bool
}

// // Get details of the car
func (car *Car) GetDetails() string {
	return fmt.Sprintf("Model: %v\nColor: %v\nSpeed: %v\nFourWheelDrive: %v\n", car.model, car.color, car.speed, car.fourWheelDrive)
}

func (car *Car) SetDetails(model string, color string, speed uint, fourWheelDrive bool) {
	car.model = model
	car.color = color
	car.speed = speed
	car.fourWheelDrive = fourWheelDrive
}

func (car Car) GetName() string {
	return "The name of the car is " + car.model
}