package format

import (
	"fmt"
	"math"
	"strings"
	"time"
	"todo/todo"
	"github.com/fatih/color"
)

func Indent(msg string) string {
	var ret string
	lines := strings.Split(msg, "\n")
	for _, l := range lines{
		ret += formatIndent(l) + "\n"
	}
	return ret
}

func formatIndent(msg string) string{
	return fmt.Sprintf("%*s %s", 4, "", msg)
}

func RemovedMessage(msg string){
	color.Set(color.Bold)
	color.Red(Indent(msg))
}

func ShowWarningMessage(msg string){
	color.Set(color.Bold)
	color.Yellow(Indent(msg))
}

func ShowErrorMessage(msg string){
	color.Set(color.Bold)
	color.Red(Indent(msg))
}

func ShowDirectoryMessage(directory string){
	ShowInformationMessage((fmt.Sprintf("/%s", directory)))
}

func ShowInformationMessage(msg string){
	color.Set(color.Bold)
	color.RGB(255, 255, 255).Print(Indent(msg))
}

func ShowSuccessMessage(msg string){
	color.Set(color.Bold)
	color.RGB(0, 255, 20).Print(Indent(msg))
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


func ShowToDoListItems(tdl []todo.ToDoListItem){
	for _, td := range tdl{
		remainingTime := td.RemainingTime()
		// Want the colour to get progressively more red and less green until expiry
		if remainingTime <= time.Duration(0){
			showToDoListItemExpired(td)
		}else{
			showToDoListItemByRemainingTimeFraction(td, td.RemainingTimeFraction())
		}
	}
}

func showToDoListItemExpired(td todo.ToDoListItem){
	color.Red(Indent(fmt.Sprintf("[%d] %s EXPIRED", td.Id, td.Do)))
}


func showToDoListItemByRemainingTimeFraction(td todo.ToDoListItem, remainingTimeFraction float64 ){
	color.Set(color.Bold)
	remainingTime := td.RemainingTime()
	output := Indent(fmt.Sprintf("%s %s", td.String(), DurationHumanReadable(remainingTime)))
	switch{
		case remainingTimeFraction == 0:
			showToDoListItemExpired(td)
		case remainingTimeFraction < 0.33:
			color.RGB(252, 163, 8).Print(output)
		case remainingTimeFraction < 0.66:
			color.Yellow(output)
		default:
			color.RGB(0, 252, 8).Print(output)
	}
}

// Similar to function above but we normalise the remaining time fractions
func ShowToDoListItemsNormalised(tdl []todo.ToDoListItem){
	max := -math.MaxFloat64
	var remainingTimeFractions []float64

	for _, td := range tdl {
		tdRemainingTime := td.RemainingTime().Seconds()
		max = math.Max(tdRemainingTime, max)
		remainingTimeFractions = append(remainingTimeFractions, tdRemainingTime)
	}

	if max != 0 {
		for i, rtf := range remainingTimeFractions {
			// Normalise to 0-1 range
			remainingTimeFractions[i] = rtf / max
		}
	}

	// Define scaling range (avoid full red)
	const minScale, maxScale = 0.3, 1.0

	for i, rtf := range remainingTimeFractions {
		// Scale between 0.3 and 1.0
		if rtf != 0{
			rtf = minScale + (rtf * (maxScale - minScale))
		}

		showToDoListItemByRemainingTimeFraction(tdl[i], rtf)
	}
}
