package swayipc

import (
	"encoding/json"
	"fmt"
)

// Event types.
const (
	EV_Workspace       int = 0x80000000
	EV_Output          int = 0x80000001
	EV_Mode            int = 0x80000002
	EV_Window          int = 0x80000003
	EV_BarconfigUpdate int = 0x80000004
	EV_Binding         int = 0x80000005
	EV_Shutdown        int = 0x80000006
	EV_Tick            int = 0x80000007
	EV_BarStateUpdate  int = 0x80000014
	EV_Input           int = 0x80000015
)

// Subscriber struct, use this to tell  swayipc which events you want to
// subscribe.
type Event struct {
	Workspace       bool
	Output          bool
	Mode            bool
	Window          bool
	BarconfigUpdate bool
	Binding         bool
	Shutdown        bool
	Tick            bool
	BarStateUpdate  bool
	Input           bool
}

// Workspace event response
type EventWorkspace struct {
	Change  string `json:"change"`
	Current *Node  `json:"workspace"`
	Old     *Node  `json:"old"`
}

// Output event response
type EventOutput struct {
	Change string `json:"change"`
}

// Mode event response
type EventMode struct {
	Change      string `json:"change"`
	PangoMarkup *Node  `json:"pango_markup"`
}

// Window event response
type EventWindow struct {
	Change    string `json:"change"`
	Container *Node  `json:"container"`
}

// BarConfig event response
type EventBarConfig struct {
	Change  string   `json:"change"`
	Binding *Binding `json:"binding"`
}

// Shutdown event response
type EventShutdown struct {
	Change string `json:"change"`
}

// Tick event response
type EventTick struct {
	First   bool   `json:"first"`
	Payload string `json:"payload"`
}

// BarState event response
type EventBarState struct {
	Id                string `json:"id"`
	VisibleByModifier bool   `json:"visible_by_modifier"`
}

// Input event response
type EventInput struct {
	Change string `json:"change"`
	Input  *Input `json:"input"`
}

// Subscribe  to  one or  more  events.  Fill the  swayipc.Event  object
// accordingly.
//
// Returns a response list containing  a response for every subscribed
// event.
func (ipc *SwayIPC) Subscribe(sub *Event) ([]*Response, error) {
	events := []string{}

	// looks ugly but makes it much more comfortable for the user
	if sub.Workspace {
		events = append(events, "workspace")
	}
	if sub.Output {
		events = append(events, "output")
	}
	if sub.Mode {
		events = append(events, "mode")
	}
	if sub.Window {
		events = append(events, "window")
	}
	if sub.BarconfigUpdate {
		events = append(events, "barconfig_update")
	}
	if sub.Binding {
		events = append(events, "binding")
	}
	if sub.Shutdown {
		events = append(events, "shutdown")
	}
	if sub.Tick {
		events = append(events, "tick")
	}
	if sub.BarStateUpdate {
		events = append(events, "bar_state_update")
	}
	if sub.Input {
		events = append(events, "input")
	}

	jsonpayload, err := json.Marshal(events)
	if err != nil {
		return nil, fmt.Errorf("failed to json marshal event list: %w", err)
	}

	err = ipc.sendHeader(SUBSCRIBE, uint32(len(jsonpayload)))
	if err != nil {
		return nil, err
	}

	err = ipc.sendPayload(jsonpayload)
	if err != nil {
		return nil, err
	}

	payload, err := ipc.readResponse()
	if err != nil {
		responses := []*Response{}
		if err := json.Unmarshal(payload.Payload, &responses); err != nil {
			return nil, fmt.Errorf("failed to unmarshal json response: %w", err)
		}

		return responses, err
	}

	// register the subscribed events
	ipc.Events = sub

	return nil, nil
}

// Event loop: Once you have subscribed to an event, you need to enter
// an event  loop over the  same running socket connection.  Sway will
// send event payloads for every subscribed event whenever it happens.
//
// You  supply the  loop a  generic callback  function, which  will be
// called every  time an event  occurs. The function will  receive the
// swayipc.RawResponse object for the event.  You need to Unmarshall the
// Payload field yourself.
//
// If your callback function returns  an error, the event loop returns
// with this error and finishes thus.
func (ipc *SwayIPC) EventLoop(callback func(event *RawResponse) error) error {
	for {
		payload, err := ipc.readResponse()
		if err != nil {
			return err
		}

		if err := callback(payload); err != nil {
			return err
		}
	}
}
