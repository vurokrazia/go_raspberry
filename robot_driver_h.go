package main
import (
	"fmt"
	"time"
	"github.com/stianeikeland/go-rpio"
)
func MotorStatus(motor_a, motor_b rpio.Pin){
	motor_a.High()
	motor_b.High()
	time.Sleep(1 * time.Second)
	motor_a.Low()
	motor_b.Low()
}
func main() {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}
	defer rpio.Close()

	// Motor A
	pin := rpio.Pin(26)
	pin_b := rpio.Pin(19)
	// Motor B
	pin_c := rpio.Pin(13)
	pin_d := rpio.Pin(6)

	pin.Output()
	pin_b.Output()
	pin_c.Output()
	pin_d.Output()

	for x := 0; x < 1; x++ {
		MotorStatus(pin,pin_d)
		MotorStatus(pin_b,pin_c)
	}
}

