package ceevee

import (
	"net/http"
	"time"
)

type Profile struct {
	Name string
}

type Target struct {
	Key    string
	Email  string
	Events []ViewEvent
}

type ViewEvent struct {
	EventDate time.Time
	IP        string
}

//Serve starts the web server taking the port to run on as a parameter
func Serve(port string) {
	http.ListenAndServe(port, nil)
}
