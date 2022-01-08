package main

import (
	"github.com/kyoto-framework/kyoto"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

type FilterType int

const (
	All FilterType = iota
	Active
	Completed
)

func (f FilterType) String() string {
	switch f {
	case All:
		return "All"
	case Active:
		return "Active"
	case Completed:
		return "Completed"
	}

	return "Unknown"
}

type ComponentTodos struct {
	Todos         []Todo
	CurrentFilter FilterType

	NewTitle   string
	NextTodoId int

	All       FilterType
	Active    FilterType
	Completed FilterType
}

func (c *ComponentTodos) Init() {
	c.CurrentFilter = All

	c.NewTitle = ""
	c.NextTodoId = 0

	c.All = All
	c.Active = Active
	c.Completed = Completed
}

func (c *ComponentTodos) FilteredTodos() []Todo {
	filtered := make([]Todo, 0)

	if c.CurrentFilter == All {
		filtered = append(filtered, c.Todos...)
	} else if c.CurrentFilter == Active {
		for _, todo := range c.Todos {
			if !todo.Completed {
				filtered = append(filtered, todo)
			}
		}
	} else if c.CurrentFilter == Completed {
		for _, todo := range c.Todos {
			if todo.Completed {
				filtered = append(filtered, todo)
			}
		}
	}

	return filtered
}

func (c *ComponentTodos) Actions() kyoto.ActionMap {
	return kyoto.ActionMap{
		"Add": func(args ...interface{}) {
			c.Todos = append(c.Todos, Todo{c.NextTodoId, c.NewTitle, false})
			c.NewTitle = ""
			c.NextTodoId++
		},
		"Complete": func(args ...interface{}) {
			target := int(args[0].(float64))
			for i := 0; i < len(c.Todos); i++ {
				if c.Todos[i].Id == target {
					c.Todos[i].Completed = !c.Todos[i].Completed
				}
			}
		},
		"ChangeFilter": func(args ...interface{}) {
			c.CurrentFilter = FilterType(int(args[0].(float64)))
		},
	}
}
