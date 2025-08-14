package i3ipc

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

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

func (ipc *I3ipc) readResponse() ([]byte, error) {
	// read header
	buf := make([]byte, IPC_HEADER_SIZE)

	_, err := ipc.socket.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read header from ipc socket: %s", err)
	}

	if string(buf[:6]) != IPC_MAGIC {
		return nil, fmt.Errorf("got invalid response from IPC socket")
	}

	payloadLen := binary.LittleEndian.Uint32(buf[6:10])

	if payloadLen == 0 {
		return nil, fmt.Errorf("got empty payload response from IPC socket")
	}

	// read payload
	payload := make([]byte, payloadLen)

	_, err = ipc.socket.Read(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to read payload from IPC socket: %s", err)
	}

	return payload, nil
}
