package main

import (
	"github.com/kyoto-framework/kyoto"
)

type ComponentCounter struct {
	Count int
}

func (c *ComponentCounter) Actions() kyoto.ActionMap {
	return kyoto.ActionMap{
		"Increment": func(args ...interface{}) {
			c.Count++
		},
		"Decrement": func(args ...interface{}) {
			c.Count--
		},
	}
}
