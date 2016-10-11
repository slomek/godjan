package main

import "github.com/slomek/godjan"

func main() {
	godjan.New()
	godjan.AddPlugin(MyPlugin{"admin"})
	godjan.Start()
}

type MyPlugin struct {
	ID string
}

func (p MyPlugin) Name() string {
	return p.ID
}

func (MyPlugin) Models() []godjan.Model {
	return []godjan.Model{
		MyModel{"user"},
	}
}

type MyModel struct {
	Name string
}

func (m MyModel) DbID() string {
	return m.Name
}
