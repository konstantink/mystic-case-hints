package main

type Level struct {
	Name  string     `json:"name"`
	Steps []Sublevel `json:"steps"`
}

type Sublevel struct {
	Name string `json:"name"`
	Tips []Hint `json:"tips"`
}

type Hint struct {
	Name string `json:"name"`
	Text string `json:"text"`
}
