package main

import (
	"os"
)

// HandleUser takes the supplied username and applies it to the connection state
// The password and all account data is overwritten when this is done
func HandleUser(username string, state *ConnState) error {
	state.Username = username
	state.Password = ""
	return nil
}

// HandlePassword sets the provided password for the connection.
// Returns errNoUsername if no username has been set.
func HandlePassword(password string, state *ConnState) error {
	if state.Username == "" {
		return errNoUsername
	}
	state.Password = password
	if Authenticate(state) {
		return nil
	}
	return errAuthentication
}

// HandleCWD changes the CWD for state to the provided directory
func HandleCWD(directory string, state *ConnState) error {
	err := os.Chdir(directory)
	if err == nil {
		state.Directory = directory
	}
	return err
}

//
