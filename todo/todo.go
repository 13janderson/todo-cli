package todo

import (
	"time"
	"fmt"
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

	SelectWithId(int) ([]ToDoListItem, error)

	Complete()
}

func (td *ToDoListItem) String() string{
	return fmt.Sprintf("[%d] %s by: %s", td.Id, td.Do, td.DoBy.Format(time.RFC822))
}

func (td ToDoListItem) RemainingTimeFraction() float64{
	// Fraction of the remaining time left on the task and the initial allowed time for the task
	allowedTime := (float64) (td.DoBy.Sub(td.CreatedAt).Seconds())
	remainingTime := (float64) (td.RemainingTime().Seconds())
	return remainingTime/allowedTime
}

func (td ToDoListItem) RemainingTime() time.Duration{
	until := time.Until(td.DoBy)
	if until < 0{
		return time.Duration(0)
	}else{
		return until
	}
}



