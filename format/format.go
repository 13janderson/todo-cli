package format

import (
	"fmt"
	"strings"
	"todo/todo"
	"math"
	"github.com/fatih/color"
	"time"
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
	c := color.RGB(int ((1 - remainingTimeFraction)*255), int((remainingTimeFraction)*255), 0)
	c.Printf(Indent(fmt.Sprintf("%s %s", td.String(), DurationHumanReadable(remainingTime))))
}

// Similar to function above but we normalise the remaining time fractions
func ShowToDoListItemsNormalised(tdl []todo.ToDoListItem){
	max := -math.MaxFloat64
	var remainingTimeFractions []float64
	for _, td := range tdl{
		tdRemainingTime := td.RemainingTimeFraction()
		max = math.Max(tdRemainingTime, max)
		remainingTimeFractions = append(remainingTimeFractions, tdRemainingTime)
	}

	if max != 0{
		// Normalise between 0 and 1 to use entire colour spectrum properly
		for i, rt := range remainingTimeFractions{
			remainingTimeFractions[i] = rt / max
		}
	}

	// Use remaining time fractions to display like normal
	for i, rtf := range remainingTimeFractions{
		showToDoListItemByRemainingTimeFraction(tdl[i], rtf)
	}
}
