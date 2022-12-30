package server

import "time"

/*
Most of this struct is redundant as the function AssignColor is used instead.
*/
type Colors struct {
	Reset      string // Resets terminal colour to default after 'text colouring'
	DarkGrey   string
	Red        string
	LightRed   string
	Green      string
	LightGreen string
	Orange     string
	Yellow     string
	Blue       string
	LightBlue  string
	Purple     string
	Black      string
	Cyan       string
	Magenta    string
}

/*
AssignColor takes an integer as input and assigns a colour based on integer values
0 - 9. For all other integer values, a default colour of 'black' is assigned.
The assigned color code is returned in string format.
*/
func AssignColor(n int) string {
	var color string
	switch n {
	case 0:
		color = "\033[1;30m" //Dark Gray
	case 1:
		color = "\033[0;31m" //Red
	case 2:
		color = "\033[1;31m" //Light Red
	case 3:
		color = "\033[0;32m" //Green
	case 4:
		color = "\033[1;32m" //Light Green
	case 5:
		color = "\033[0;33m" //Brown/orange
	case 6:
		color = "\033[1;33m" //Yellow
	case 7:
		color = "\033[0;34m" //Blue
	case 8:
		color = "\033[1;34m" //Light Blue
	case 9:
		color = "\033[0;35m" //Purple
	default:
		color = "\033[0;30m" //Black
	}
	return color
}

/*
FindIndex returns the index integer for the within the NumbTerminal slice for the element
matching the input integer.
*/
func FindIndex(element int) int {
	for i := 0; i < len(NumbTerminal); i++ {
		if element == NumbTerminal[i] {
			return i
		}
	}
	return -1
}

/*
RemoveElement takes an input element and removes its matching element in the global NumbTerminal slice.
*/
func RemoveElement(element int) {
	index := FindIndex(element)
	NumbTerminal = append(NumbTerminal[:index], NumbTerminal[index+1:]...)
}

/*
GetTime returns a formatted string of the current time.
*/
func GetTime() string {
	timeFormat := "2006-01-02 15:04:05"
	return time.Now().Format(timeFormat)
}

/*
PortAtoi takes a string representing a positive number and returns it
in integer form. No edge cases / errors are handled as it is nested within the
CheckPort function which contains prior checks.
*/
func PortAtoi(s string) int {
	nbr := 0
	for _, digit := range s {
		digit = digit - 48
		nbr = nbr*10 + int(digit)
	}
	return nbr
}
