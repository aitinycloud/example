package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

type Person struct {
	Name string `json:"name,omitempty"`
	Age  uint8  `json:"age,omitempty"`
	//Address *Address `json:"address,omitempty"`
}

type Address struct {
	Street   string `json:"street,omitempty"`
	Postcode string `json:"postcode,omitempty"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(p Person) string {
	//return fmt.Sprintf("Hello %s", p)
	return fmt.Sprintf("Hello name:%s (Age: %d)!", p.Name, p.Age)
}
