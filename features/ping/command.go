package ping

// Command asks the server to mark a client as alive.
type Command struct {
	ClientID string `json:"client_id"`
}
