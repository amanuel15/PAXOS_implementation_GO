package main

import (
	Entity "./Entities"
	Paxos "./Runner"
)

func main() {
	n := Paxos.CreateNetwork(3, 3, 2, []Entity.Any{"first", "second", "go", "four", "me"})
	Paxos.RunInMain(n)
}
