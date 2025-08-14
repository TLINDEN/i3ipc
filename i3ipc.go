package i3ipc

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	VERSION = "v0.1.0"

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
	GET_INPUTS
	GET_SEATS
)

// module struct
type I3ipc struct {
	socket     net.Conn
	SocketFile string
}

// i3-ipc structs
type Rect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Response struct {
	Success    bool   `json:"success"`
	ParseError bool   `json:"parse_error"`
	Error      string `json:"error"`
}

func NewI3ipc(file string) *I3ipc {
	if file == "" {
		file = "SWAYSOCK"
	}
	return &I3ipc{SocketFile: file}
}

func (ipc *I3ipc) GetTree() (*Node, error) {
	err := ipc.sendHeader(GET_TREE, 0)
	if err != nil {
		return nil, err
	}

	payload, err := ipc.readResponse()
	if err != nil {
		return nil, err
	}

	node := &Node{}
	if err := json.Unmarshal(payload, &node); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return node, nil
}
