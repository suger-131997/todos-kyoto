package main

import (
	"github.com/kyoto-framework/kyoto"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

type ComponentTodos struct {
	Todos      []Todo
	FilterType string

	NewTitle   string
	NextTodoId int
}

func (c *ComponentTodos) Init() {
	c.FilterType = "All"

	c.NewTitle = ""
	c.NextTodoId = 0
}

//<!-- {{ if or (or (eq $.FilterType "All") (and (eq $.FilterType "Active") $todo.Completed)) ((eq $.FilterType "Completed") $todo.Completed) }} -->
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
			c.FilterType = args[0].(string)
		},
	}
}
