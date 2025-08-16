package swayipc

import (
	"net"
)

const (
	VERSION = "v0.2.0"

	IPC_HEADER_SIZE = 14
	IPC_MAGIC       = "i3-ipc"
	IPC_MAGIC_LEN   = 6
)

const (
	// message types
	RUN_COMMAND = iota
	GET_WORKSPACES
	SUBSCRIBE
	GET_OUTPUTS
	GET_TREE
	GET_MARKS
	GET_BAR_CONFIG
	GET_VERSION
	GET_BINDING_MODES
	GET_CONFIG
	SEND_TICK
	SYNC
	GET_BINDING_STATE

	GET_INPUTS = 100
	GET_SEATS  = 101
)

// This is the primary struct to work with the swayipc module.
type SwayIPC struct {
	socket     net.Conn
	SocketFile string // filename of the i3 IPC socket
	Events     *Event // store subscribed events, see swayipc.Subscribe()
}

// A rectangle struct, used at various places for geometry etc.
type Rect struct {
	X      int `json:"x"` // X coordinate
	Y      int `json:"y"` // Y coordinate
	Width  int `json:"width"`
	Height int `json:"height"`
}

// Stores responses retrieved via ipc
type Response struct {
	Success    bool   `json:"success"`
	ParseError bool   `json:"parse_error"`
	Error      string `json:"error"`
}

// Stores the user config for the WM
type Config struct {
	Config string `json:"config"`
}

// Stores the binding state
type State struct {
	Name string `json:"name"`
}

// Create a new swayipc.SwayIPC object.  Filename argument is optional and
// may denote  a filename or the  name of an environment  variable.
//
// By default and if nothing is  specified we look for the environment
// variable SWAYSOCK  and use  the file  it points  to as  unix domain
// socket to communicate with sway (and possible i3).
func NewSwayIPC(file ...string) *SwayIPC {
	ipc := &SwayIPC{}

	if len(file) == 0 {
		ipc.SocketFile = "SWAYSOCK"
	} else {
		ipc.SocketFile = file[0]
	}

	return ipc
}

// internal convenience wrapper
func (ipc *SwayIPC) get(command uint32) (*RawResponse, error) {
	err := ipc.sendHeader(command, 0)
	if err != nil {
		return nil, err
	}

	payload, err := ipc.readResponse()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
