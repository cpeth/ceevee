package main

import "time"

//Store handles and presists application state
type Store struct {
}

type profile struct {
	Name     string
	PassHash string
}

type target struct {
	Key    string
	Email  string
	Events []viewEvent
}

type viewEvent struct {
	EventDate time.Time
	IP        string
}

//NewStore creates a new Store with default parameters
func NewStore() Store {
	return Store{}
}
