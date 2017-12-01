package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)

	fmt.Println(convertTime(input))
}

func convertTime(time string) string {
	trimmedTime := time[:8] // HH:MM:SS

	splittedTime := strings.Split(trimmedTime, ":")
	standardHour, _ := strconv.Atoi(splittedTime[0])

	meridiem := string(time[8:]) //AM:PM
	var militaryTime bytes.Buffer
	switch meridiem {
	case "AM":
		if standardHour == 12 {
			militaryTime.WriteString("00" + ":" + splittedTime[1] + ":" + splittedTime[2])
			return militaryTime.String()
		} else {
			// Nothing changed
			return trimmedTime
		}
	case "PM":
		if standardHour == 12 {
			return trimmedTime
		} else {
			militaryHour := standardHour + 12
			militaryTime.WriteString(strconv.Itoa(militaryHour) + ":" + splittedTime[1] + ":" + splittedTime[2])
			return militaryTime.String()
		}
	default:
		return ""
	}
}
