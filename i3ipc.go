package i3ipc

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
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

type Config struct {
	Config string `json:"config"`
}

func NewI3ipc(file ...string) *I3ipc {
	ipc := &I3ipc{}

	if len(file) == 0 {
		ipc.SocketFile = "SWAYSOCK"
	} else {
		ipc.SocketFile = file[0]
	}

	return ipc
}

func (ipc *I3ipc) get(command uint32) ([]byte, error) {
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

func (ipc *I3ipc) RunGlobalCommand(command ...string) ([]Response, error) {
	return ipc.RunCommand(0, command...)
}

func (ipc *I3ipc) RunCommand(id int, command ...string) ([]Response, error) {
	if len(command) == 0 {
		return nil, fmt.Errorf("empty command arg")
	}

	commands := strings.Join(command, ",")

	if id > 0 {
		commands = fmt.Sprintf("[con_id=%d] %s", id, commands)
	}

	err := ipc.sendHeader(RUN_COMMAND, uint32(len(commands)))
	if err != nil {
		return nil, fmt.Errorf("failed to send run_command to IPC %w", err)
	}

	err = ipc.sendPayload([]byte(commands))
	if err != nil {
		return nil, fmt.Errorf("failed to send switch focus command: %w", err)
	}

	payload, err := ipc.readResponse()
	if err != nil {
		return nil, err
	}

	responses := []Response{}

	if err := json.Unmarshal(payload, &responses); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json response: %w", err)
	}

	if len(responses) == 0 {
		return nil, fmt.Errorf("got zero IPC response")
	}

	for _, response := range responses {
		if !response.Success {
			return responses, fmt.Errorf("one or more commands failed")
		}
	}

	return responses, nil
}

func (ipc *I3ipc) GetWorkspaces() ([]*Node, error) {
	payload, err := ipc.get(GET_WORKSPACES)
	if err != nil {
		return nil, err
	}

	nodes := []*Node{}
	if err := json.Unmarshal(payload, &nodes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return nodes, nil
}

func (ipc *I3ipc) GetMarks() ([]string, error) {
	payload, err := ipc.get(GET_MARKS)
	if err != nil {
		return nil, err
	}

	marks := []string{}
	if err := json.Unmarshal(payload, &marks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return marks, nil
}

func (ipc *I3ipc) GetBindingModes() ([]string, error) {
	payload, err := ipc.get(GET_BINDING_MODES)
	if err != nil {
		return nil, err
	}

	modes := []string{}
	if err := json.Unmarshal(payload, &modes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return modes, nil
}

func (ipc *I3ipc) GetConfig() (string, error) {
	payload, err := ipc.get(GET_CONFIG)
	if err != nil {
		return "", err
	}

	config := &Config{}
	if err := json.Unmarshal(payload, &config); err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return config.Config, nil
}
