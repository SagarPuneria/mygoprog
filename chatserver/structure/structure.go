package structure

import (
	"net"
)

// ID struct stores client ID and group ID
type ID struct {
	ClientID int
	GroupID  int
}

// DataID is embeded of struct ID and Data
type DataID struct {
	ID
	Data string
}

// ServerStruct is embeded of struct DataID and ClientAddr
type ServerStruct struct {
	DataID
	ClientAddr net.Addr
}
