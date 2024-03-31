package main

type Level struct {
	Name  string
	Steps []Sublevel
}

type Sublevel struct {
	Name string
	Tips []Hint
}

type Hint struct {
	Name string
	Text string
}
