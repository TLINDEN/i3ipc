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
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
)

// Contains a raw json response, not marshalled yet.
type RawResponse struct {
	PayloadType int
	Payload     []byte
}

// Connect to unix domain ipc socket.
func (ipc *I3ipc) Connect() error {
	if !fileExists(ipc.SocketFile) {
		ipc.SocketFile = os.Getenv(ipc.SocketFile)
		if ipc.SocketFile == "" {
			return fmt.Errorf("socket file %s doesn't exist", ipc.SocketFile)
		}
	}

	conn, err := net.Dial("unix", ipc.SocketFile)
	if err != nil {
		return err
	}

	ipc.socket = conn

	return nil
}

// Close the socket.
func (ipc *I3ipc) Close() {
	ipc.socket.Close()
}

func (ipc *I3ipc) sendHeader(messageType uint32, len uint32) error {
	sendPayload := make([]byte, IPC_HEADER_SIZE)

	copy(sendPayload, []byte(IPC_MAGIC))
	binary.LittleEndian.PutUint32(sendPayload[IPC_MAGIC_LEN:], len)
	binary.LittleEndian.PutUint32(sendPayload[IPC_MAGIC_LEN+4:], messageType)

	_, err := ipc.socket.Write(sendPayload)

	if err != nil {
		return fmt.Errorf("failed to send header to IPC socket %w", err)
	}

	return nil
}

func (ipc *I3ipc) sendPayload(payload []byte) error {
	_, err := ipc.socket.Write(payload)

	if err != nil {
		return fmt.Errorf("failed to send payload to IPC socket %w", err)
	}

	return nil
}

func (ipc *I3ipc) readResponse() (*RawResponse, error) {
	// read header
	buf := make([]byte, IPC_HEADER_SIZE)

	_, err := ipc.socket.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read header from ipc socket: %s", err)
	}

	if string(buf[:6]) != IPC_MAGIC {
		return nil, errors.New("got invalid response from IPC socket")
	}

	payloadLen := binary.LittleEndian.Uint32(buf[6:10])
	if payloadLen == 0 {
		return nil, errors.New("got empty payload response from IPC socket")
	}

	payloadType := binary.LittleEndian.Uint32(buf[10:])

	// read payload
	payload := make([]byte, payloadLen)

	_, err = ipc.socket.Read(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to read payload from IPC socket: %s", err)
	}

	return &RawResponse{PayloadType: int(payloadType), Payload: payload}, nil
}
