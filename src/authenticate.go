package main

// Authenticate determines whether user has valid credentials to log in. True if
// they do, false if they do not.
// TODO
func Authenticate(state *ConnState) bool {
	state.Authenticated = true
	return true
}
