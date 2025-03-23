package format

import (
	"fmt"
	"strings"
	"time"
	"github.com/fatih/color"
)

func Indent(msg string) string {
	var ret string
	lines := strings.Split(msg, "\n")
	for _, l := range lines{
		ret += fmt.Sprintf("%*s %s\n", 4, "", l)
	}
	return ret
}

func RemovedMessage(msg string){
	color.Set(color.Bold)
	color.Red(msg)
}

func ShowWarningMessage(msg string){
	color.Set(color.Bold)
	color.Yellow(Indent(msg))
}

func ShowErrorMessage(msg string){
	color.Set(color.Bold)
	color.Red(Indent(msg))
}

func ShowSuccessMessage(msg string){
	color.Set(color.Bold)
	color.Green(Indent(msg))
}

func DurationHumanReadable(d time.Duration) string{
	var parts []string
	
	day := time.Hour * 24
	days := int(d / day)
	afterDays := d - time.Duration(days)*day
	hours := int(afterDays / time.Hour)
	afterHours := afterDays - time.Duration(time.Hour)*time.Duration(hours)
	mins := int(afterHours / time.Minute)

	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if mins > 0 {
		parts = append(parts, fmt.Sprintf("%dm", mins))
	}

	return fmt.Sprintf("%s", parts)

}