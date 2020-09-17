package main
import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	scanner.Scan()
	
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	pin_a := rpio.Pin(18)
	pin_b := rpio.Pin(14)
	pin_c := rpio.Pin(21)

	pin_a.Output()
	pin_b.Output()
	pin_c.Output()

	for {
		fmt.Println("Write a number instruction \t 1. On \t 2.Off")			
		option, _ := reader.ReadString('\n')
    // convert CRLF to LF
    option = strings.Replace(option, "\n", "", -1)

		switch option {
		case "0":
			pin_a.Low()
			pin_b.Low()
			pin_c.Low()
		case "1":
			pin_a.High()
			pin_b.High()
			pin_c.High()
		default:
			break
		}
	}	
}

