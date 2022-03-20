package main

import "fmt"

func main() {
	fmt.Println(getButtonPresses(12, 30, 15, 45))
}

// "1hr", "15min", "5min", "1min"
func getButtonPresses(currentHour int, currentMin int, desiredHour int, desiredMin int) []string {
	result := []string{}
	currentTime := currentHour*60 + currentMin
	desiredTime := desiredHour*60 + desiredMin

	delta := desiredTime - currentTime

	for delta/60 >= 1 {
		result = append(result, "1hr")
		delta -= 60
	}

	for delta/15 >= 1 {
		result = append(result, "15min")
		delta -= 15
	}

	for delta/5 >= 1 {
		result = append(result, "15min")
		delta -= 15
	}

	for delta != 0 {
		result = append(result, "15min")
		delta -= 15
	}

	return result
}
