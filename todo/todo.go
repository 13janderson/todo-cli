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

	Complete()
}

func (td *ToDoListItem) String() string{
	return fmt.Sprintf("[%d] %s by: %s", td.Id, td.Do, td.DoBy.Format(time.RFC822))
}

func (td ToDoListItem) RemainingTimeFraction() float64{
	// Fraction of the remaining time left on the task and the initial allowed time for the task
	// fmt.Println(td.String())
	allowedTime := (float64) (td.DoBy.Sub(td.CreatedAt).Seconds())
	// fmt.Println(allowedTime)
	remainingTime := (float64) (td.RemainingTime().Seconds())
	// fmt.Println(remainingTime)
	return remainingTime/allowedTime
}

func (td ToDoListItem) RemainingTime() time.Duration{
	return time.Until(td.DoBy)
}



