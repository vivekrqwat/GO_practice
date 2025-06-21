package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (t *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*t = append(*t, todo)

}

func (t *Todos) ValidateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		er := errors.New("invalid index")
		fmt.Println("invalid index")
		return er
	}
	return nil

}

func (t *Todos) del(index int) {
	err := t.ValidateIndex(index)
	if err != nil {
		fmt.Println("Enter valid index")
		return
	}
	*t = append((*t)[:index], (*t)[index+1:]...)
}

func (t *Todos) toggle(index int) error {
	ti := *t
	err := ti.ValidateIndex(index)
	if err != nil {
		fmt.Println("Enter valid index")
		return err
	}
	isCompleted := ti[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		ti[index].CompletedAt = &completionTime

	}
	ti[index].Completed = !isCompleted
	return nil

}

func (t *Todos) edit(index int, title string) error {
	ti := *t
	err := ti.ValidateIndex(index)
	if err != nil {
		fmt.Println("Enter valid index")
		return err
	}

	ti[index].Title = title
	return nil

}

func (t *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for i, td := range *t {
		completedAt := ""
		completed := "no"
		if td.Completed {
			completed = "yes"
			if td.CompletedAt != nil {
				completedAt = td.CompletedAt.Format(time.RFC3339)
			}
		}

		table.AddRow(strconv.Itoa(i), td.Title, completed, td.CreatedAt.Format(time.RFC3339), completedAt)

	}
	table.Render()
}
