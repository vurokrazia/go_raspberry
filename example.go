package main
import (
	"fmt"
	"time"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin := rpio.Pin(18)
	pin_b := rpio.Pin(21)
	pin_c := rpio.Pin(14)
	pin.Output()
	pin_b.Output()
	pin_c.Output()

	for x := 0; x < 1000; x++ {
		pin.Toggle()
		pin_b.Toggle()
		pin_c.Toggle()
		time.Sleep(time.Second / 8)
	}
}

