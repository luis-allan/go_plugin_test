package shared

import (
   "fmt"
   "os"
   "plugin"
)

type Greeting interface {
   SetName (name string)
   SayHello ()
}

var (
   NewGreeting func(name string) Greeting
)

func LoadGreetingPlugin (path string) {
   plug, err := plugin.Open(path)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   _newGreeting, err := plug.Lookup("NewGreeting")
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   NewGreeting = _newGreeting.(func(name string) Greeting)
}