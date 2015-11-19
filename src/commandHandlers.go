package main

// HandleUser takes the supplied username and applies it to the connection state
func HandleUser(username string, state *ConnState) error {
	state.Username = username
	state.Password = ""
	return nil
}

// HandlePassword sets the provided password for the connection.
// Throws errNoUsername if no username has been set.
func HandlePassword(password string, state *ConnState) error {
	if state.Username == "" {
		return errNoUsername
	}
	return nil
}
