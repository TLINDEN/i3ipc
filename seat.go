/*
Copyright © 2025 Thomas von Dein

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package i3ipc

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
func (ipc *I3ipc) GetSeats() ([]*Seat, error) {
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
