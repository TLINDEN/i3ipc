package i3ipc

import (
	"encoding/json"
	"fmt"
)

func (ipc *I3ipc) GetWorkspaces() ([]*Node, error) {
	payload, err := ipc.get(GET_WORKSPACES)
	if err != nil {
		return nil, err
	}

	nodes := []*Node{}
	if err := json.Unmarshal(payload.Payload, &nodes); err != nil {
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
	if err := json.Unmarshal(payload.Payload, &marks); err != nil {
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
	if err := json.Unmarshal(payload.Payload, &modes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return modes, nil
}

func (ipc *I3ipc) GetBindingState() (*State, error) {
	payload, err := ipc.get(GET_BINDING_STATE)
	if err != nil {
		return nil, err
	}

	state := &State{}
	if err := json.Unmarshal(payload.Payload, &state); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return state, nil
}

func (ipc *I3ipc) GetConfig() (string, error) {
	payload, err := ipc.get(GET_CONFIG)
	if err != nil {
		return "", err
	}

	config := &Config{}
	if err := json.Unmarshal(payload.Payload, &config); err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return config.Config, nil
}

func (ipc *I3ipc) SendTick(payload string) error {
	err := ipc.sendHeader(SEND_TICK, uint32(len(payload)))
	if err != nil {
		return err
	}

	err = ipc.sendPayload([]byte(payload))
	if err != nil {
		return err
	}

	return nil
}
