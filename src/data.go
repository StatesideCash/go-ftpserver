package main

import (
	"errors"
)

// ConnState tracks the state of the connection
type ConnState struct {
	Username      string
	Password      string
	Authenticated bool
}

var (
	errNoUsername = errors.New("Password given but no username supplied")
)
