package format

import (
	"fmt"
	"github.com/fatih/color"
	"math"
	"os"
	"path"
	"strings"
	"time"
	"todo/todo"
)

func Indent(msg string) string {
	var ret string
	lines := strings.Split(msg, "\n")
	for _, l := range lines {
		ret += formatIndent(l) + "\n"
	}
	return ret
}

func IndentN(msg string, d int) string {
	var ret string
	lines := strings.Split(msg, "\n")
	for _, l := range lines {
		ret += formatIndentN(l, d) + "\n"
		fmt.Printf("IndentNret: \n%s", ret)
	}
	return ret
}

func formatIndent(msg string) string {
	return fmt.Sprintf("%*s %s", 4, "", msg)
}

func formatIndentN(msg string, d int) string {
	// fmt.Printf("formatIndentN depth: %d", d)
	indent := int(math.Min(float64(d), float64(1)))
	ret := fmt.Sprintf("%*s %s", 4*indent, "", msg)
	fmt.Printf("formatindentNret: \n%s", ret)
	return ret
}

func RemovedMessage(msg string) {
	color.Set(color.Bold)
	color.Red(Indent(msg))
}

func ShowWarningMessage(msg string) {
	color.Set(color.Bold)
	color.Yellow(Indent(msg))
}

func ShowErrorMessage(msg string) {
	color.Set(color.Bold)
	color.Red(Indent(msg))
}

func ShowCwdMessage(depth int) {
	fmt.Printf("ShowCwdMessage: %d\n", depth)
	cwd, _ := os.Getwd()
	ShowDirectoryMessage(formatIndentN(path.Base(cwd), depth))
}

func ShowDirectoryMessage(directory string) {
	ShowInformationMessage((fmt.Sprintf("%s", directory)))
}

func ShowInformationMessage(msg string) {
	color.Set(color.Bold)
	color.RGB(255, 255, 255).Print(Indent(msg))
}

func ShowSuccessMessage(msg string) {
	color.Set(color.Bold)
	color.RGB(0, 255, 20).Print(Indent(msg))
}

func DurationHumanReadable(d time.Duration) string {
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

func ShowToDoListItems(tdl []todo.ToDoListItem, depth int) {
	for _, td := range tdl {
		remainingTime := td.RemainingTime()
		// Want the colour to get progressively more red and less green until expiry
		if remainingTime <= time.Duration(0) {
			showToDoListItemExpired(td, depth)
		} else {
			showToDoListItemByRemainingTimeFraction(td, td.RemainingTimeFraction(), depth)
		}
	}
}

func showToDoListItemExpired(td todo.ToDoListItem, depth int) {
	color.Red(IndentN(fmt.Sprintf("[%d] %s EXPIRED", td.Id, td.Do), depth))
}

// hate having to pass the depth around everywher... this feels fucking terrible
// what if we just made the format stateful and aware of how the cwd is changing over time within a
// single run of the program

func showToDoListItemByRemainingTimeFraction(td todo.ToDoListItem, remainingTimeFraction float64, depth int) {
	fmt.Printf("showToDoListItemByRemainingTimeFraction depth: %d\n", depth)
	color.Set(color.Bold)
	remainingTime := td.RemainingTime()
	output := IndentN(fmt.Sprintf("%s %s", td.String(), DurationHumanReadable(remainingTime)), depth)
	switch {
	case remainingTimeFraction == 0:
		showToDoListItemExpired(td, depth)
	case remainingTimeFraction < 0.33:
		color.RGB(255, 165, 0).Print(output)
	case remainingTimeFraction < 0.66:
		color.RGB(255, 255, 0).Print(output)
	default:
		color.RGB(0, 255, 0).Print(output)
	}
}

// Similar to function above but we normalise the remaining time fractions
func ShowToDoListItemsNormalised(tdl []todo.ToDoListItem, depth int) {
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
		if rtf != 0 {
			rtf = minScale + (rtf * (maxScale - minScale))
		}

		showToDoListItemByRemainingTimeFraction(tdl[i], rtf, depth)
	}
}
