package format

import(
	"fmt"
	"time"
	"todo/todo"
	"github.com/fatih/color"
)

func indent() string {
	return fmt.Sprintf("%*s", 4, "")
}

func RemovedMessage(msg string){
	color.Set(color.Bold)
	color.Red(msg)
}


func ShowToDoListItem(td todo.ToDoListItem){
	remainingTime := td.RemainingTime()
	color.Set(color.Bold)
	if remainingTime <= time.Duration(0){
		color.Red("%s [%d] %s EXPIRED", indent(), td.Id, td.Do)
	}else{
		color.Green("%s [%d] %s %s", indent(), td.Id, td.Do, DurationHumanReadable(remainingTime))
	}
}

func ShowWarningMessage(msg string){
	color.Set(color.Bold)
	color.Yellow(msg)
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