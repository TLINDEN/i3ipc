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

package swayipc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Execute  the specified  global command[s]  (one or  more  commands can  be
// given) and returns a response list.
//
// Possible commands are all non-specific commands, see sway(5)
func (ipc *SwayIPC) RunGlobalCommand(command ...string) ([]Response, error) {
	return ipc.RunCommand(0, "", command...)
}

// Execute  the specified  container command[s]  (one or  more  commands can  be
// given) and returns a response list.
//
// Possible commands are all container-specific commands, see sway(5)
func (ipc *SwayIPC) RunContainerCommand(id int, command ...string) ([]Response, error) {
	return ipc.RunCommand(id, "con", command...)
}

// Execute  the specified  (target) command[s]  (one or  more  commands can  be
// given) and returns a response list.
//
// Possible commands are all container-specific commands, see sway(5).
//
// Target can be one of con, workspace, output, input, etc. see sway-ipc(7).
func (ipc *SwayIPC) RunCommand(id int, target string, command ...string) ([]Response, error) {
	if len(command) == 0 {
		return nil, errors.New("empty command arg")
	}

	commands := strings.Join(command, ",")

	if id > 0 {
		// a type specific command, otherwise global
		commands = fmt.Sprintf("[%s_id=%d] %s", target, id, commands)
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

	if err := json.Unmarshal(payload.Payload, &responses); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json response: %w", err)
	}

	if len(responses) == 0 {
		return nil, errors.New("got zero IPC response")
	}

	for _, response := range responses {
		if !response.Success {
			return responses, errors.New("one or more commands failed")
		}
	}

	return responses, nil
}
