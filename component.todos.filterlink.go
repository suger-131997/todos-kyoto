package main

type ComponentTodosFilterLink struct {
	Type          FilterType
	CurrentFilter *FilterType
}

func (c ComponentTodosFilterLink) IsActive() bool {
	return c.Type == *c.CurrentFilter
}
