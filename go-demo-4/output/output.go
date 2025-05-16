package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {

	switch t := value.(type) {
	case string:
		color.Red(t)
	case error:
		color.Red(t.Error())
	case int:
		color.Red("Error code: %d", t)
	default:
		color.Red("Something went wrong...")
	}
}

func Success() {
	color.Green("Success!")
}
