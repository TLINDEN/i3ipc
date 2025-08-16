package swayipc

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
func (ipc *SwayIPC) Connect() error {
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
func (ipc *SwayIPC) Close() {
	ipc.socket.Close()
}

func (ipc *SwayIPC) sendHeader(messageType uint32, len uint32) error {
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

func (ipc *SwayIPC) sendPayload(payload []byte) error {
	_, err := ipc.socket.Write(payload)

	if err != nil {
		return fmt.Errorf("failed to send payload to IPC socket %w", err)
	}

	return nil
}

func (ipc *SwayIPC) readResponse() (*RawResponse, error) {
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
