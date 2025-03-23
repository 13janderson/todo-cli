package todo

import (
	"time"
	"fmt"
	"todo/format"
	"github.com/fatih/color"
)

type ToDoListItem struct {
	Id        int       `json:"id" db:"id"`
	Do        string    `json:"do" db:"do"`
	DoBy time.Time `json:"doBy" db:"doBy"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
}

type ToDoList interface {
	// Init new to do list
	Init()
	// Writes all list items in readable format to stdout
	List() error
	Add(item ToDoListItem)
	// Remove specific item
	Remove(item ToDoListItem)
	// Removes most recent to do added
	Pop()

	Complete()
}

func (td *ToDoListItem) String() string{
	return fmt.Sprintf("[%d] %s by: %s", td.Id, td.Do, td.DoBy.Format(time.RFC822))
}

func (td ToDoListItem) RemainingTimeFraction() float64{
	// Fraction of the remaining time left on the task and the initial allowed time for the task
	allowedTime := (float64) (td.DoBy.Sub(td.CreatedAt))
	remainingTime := (float64) (td.RemainingTime())
	return remainingTime/allowedTime
}

func (td ToDoListItem) RemainingTime() time.Duration{
	return time.Until(td.DoBy)
}

func (td ToDoListItem) ShowToDoListItem(){
	remainingTime := td.RemainingTime()
	remainingTimeFraction := td.RemainingTimeFraction()
	color.Set(color.Bold)
	// Want the colour to get progressively more red and less green until expiry
	if remainingTime <= time.Duration(0){
		color.Red(format.Indent(fmt.Sprintf("[%d] %s EXPIRED", td.Id, td.Do)))
	}else{
		c := color.RGB(int ((1 - remainingTimeFraction)*255), int((remainingTimeFraction)*255), 0)
		c.Printf(format.Indent(fmt.Sprintf("%s %s", td.String(), format.DurationHumanReadable(remainingTime))))
	}
}




