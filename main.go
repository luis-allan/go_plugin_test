package main

import (
   "plugin-test/shared"
)

func main() {
   shared.LoadGreetingPlugin("plugin/hello.so")
   greeter1 := shared.NewGreeting("test1")
   greeter2 := shared.NewGreeting("test2")
   greeter2.SetName("test3")

   greeter1.SayHello()
   greeter2.SayHello()
}
