package swayipc

import (
	"encoding/json"
	"fmt"
)

// Information about a seat containing input devices
type Seat struct {
	Name         string   `json:"name"`
	Capabilities int      `json:"capabilities"`
	Focus        int      `json:"focus"`
	Devices      []*Input `json:"devices"`
}

// Get input seats
func (ipc *SwayIPC) GetSeats() ([]*Seat, error) {
	payload, err := ipc.get(GET_SEATS)
	if err != nil {
		return nil, err
	}

	seats := []*Seat{}
	if err := json.Unmarshal(payload.Payload, &seats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return seats, nil
}
