package main

import (
	"fmt"
	"os"
	"strings"
)

// HandleUser takes the supplied username and applies it to the connection state
// The password and all account data is overwritten when this is done
func HandleUser(username string, state *ConnState) error {
	fmt.Println("Called HandleUser")
	state.Username = username
	state.Password = ""
	return nil
}

// HandlePassword sets the provided password for the connection.
// Returns errNoUsername if no username has been set.
func HandlePassword(password string, state *ConnState) error {
	fmt.Println("Called HandlePassword")
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
	fmt.Println("Called HandleCWD")
	if err := os.Chdir(directory); err != nil {
		return err
	}
	state.Directory = directory
	return nil
}

// HandleCDUP changes the user directory to one level higher
func HandleCDUP(state *ConnState) error {
	fmt.Println("Called HandleCDUP")
	cwd := state.Directory
	dirlist := strings.Split(cwd, string(os.PathSeparator))
	newdir := ""

	// Take the last directory segment off
	for i := 0; i < len(dirlist)-1; i++ {
		newdir += string(os.PathSeparator)
		newdir += dirlist[i]
	}

	if err := os.Chdir(newdir); err != nil {
		return err
	}
	state.Directory = newdir
	return nil
}

//
