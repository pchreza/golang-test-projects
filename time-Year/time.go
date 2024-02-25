package main

import (
	"fmt"
	"time"
)

func main() {
	var untilStartDay = 0
	now := time.Now()
	yourYear := now.Year()
	yourMonth := int(now.Month())
	yourDay := now.Day()
	var yourMonth2 int = yourMonth - 1
	switch yourMonth2 {
	case 11:
		untilStartDay += 31
		fallthrough

	case 10:
		untilStartDay += 30
		fallthrough

	case 9:
		untilStartDay += 31
		fallthrough

	case 8:
		untilStartDay += 30
		fallthrough

	case 7:
		untilStartDay += 31
		fallthrough

	case 6:
		untilStartDay += 31
		fallthrough

	case 5:
		untilStartDay += 30
		fallthrough

	case 4:
		untilStartDay += 31
		fallthrough

	case 3:
		untilStartDay += 30
		fallthrough

	case 2:
		untilStartDay += 31
		fallthrough

	case 1:
		untilStartDay += 28
	default:
		untilStartDay += 0
	}
	fmt.Printf("%d days have passed since the beginning of year %d \n", untilStartDay+yourDay, yourYear)
	fmt.Printf("until end year %d and enter to %d : %d day Left ! \n", yourYear, yourYear+1, 365-(untilStartDay+yourDay))

}
