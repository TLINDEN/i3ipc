package swayipc

import (
	"encoding/json"
	"fmt"
)

// An input (keyboard, mouse, whatever)
type Input struct {
	Identifier           string    `json:"identifier"`
	Name                 string    `json:"name"`
	Vendor               int       `json:"vendor"`
	Product              int       `json:"product"`
	Type                 string    `json:"type"`
	XkbActiveLayoutName  string    `json:"xkb_active_layout_name"`
	XkbLayoutNames       []string  `json:"xkb_layout_names"`
	XkbActiveLayoutIndex int       `json:"xkb_active_layout_index"`
	ScrollFactor         float32   `json:"scroll_factor"`
	Libinput             *LibInput `json:"libinput"`
}

// Holds the data associated with libinput
type LibInput struct {
	SendEvents        string    `json:"send_events"`
	Tap               string    `json:"tap"`
	TapButtonMap      string    `json:"tap_button_map"`
	TapDrag           string    `json:"tap_drag"`
	TapDragLock       string    `json:"tap_drag_lock"`
	AccelSpeed        float32   `json:"accel_speed"`
	AccelProfile      string    `json:"accel_profile"`
	NaturalScroll     string    `json:"natural_scroll"`
	LeftHanded        string    `json:"left_handed"`
	ClickMethod       string    `json:"click_method"`
	ClickButtonMap    string    `json:"click_button_map"`
	MiddleEmulation   string    `json:"middle_emulation"`
	ScrollMethod      string    `json:"scroll_method"`
	ScrollButton      int       `json:"scroll_button"`
	ScrollButtonLock  string    `json:"scroll_button_lock"`
	Dwt               string    `json:"dwt"`
	Dwtp              string    `json:"dwtp"`
	CalibrationMatrix []float32 `json:"calibration_matrix"`
}

// A key binding
type Binding struct {
	Command        string   `json:"command"`
	EventStateMask []string `json:"event_state_mask"`
	InputCode      int      `json:"input_code"`
	Symbol         string   `json:"symbol"`
	InputType      string   `json:"input_type"`
}

// Get a list of all currently supported inputs
func (ipc *SwayIPC) GetInputs() ([]*Input, error) {
	payload, err := ipc.get(GET_INPUTS)
	if err != nil {
		return nil, err
	}

	inputs := []*Input{}
	if err := json.Unmarshal(payload.Payload, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return inputs, nil
}
