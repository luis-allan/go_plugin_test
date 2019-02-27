package main

import (
   "fmt"
   "plugin-test/shared"
)

type _Greeting struct {
   name string
}

func (m *_Greeting) SetName(name string) {
   m.name = name
}

func (m *_Greeting) SayHello() {
   fmt.Printf("Hello %v...\n", m.name)
}

func NewGreeting(name string) shared.Greeting {
   g := &_Greeting{name:name}
   return g
}
