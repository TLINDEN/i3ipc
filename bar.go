package i3ipc

import (
	"encoding/json"
	"fmt"
)

type Gaps struct {
	Top    int `json:"top"`
	Right  int `json:"right"`
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
}

type Colors struct {
	Background              string `json:"background"`
	Statusline              string `json:"statusline"`
	Separator               string `json:"separator"`
	FocusedBackground       string `json:"focused_background"`
	FocusedStatusline       string `json:"focused_statusline"`
	FocusedSeparator        string `json:"focused_separator"`
	FocusedWorkspaceBorder  string `json:"focused_workspace_border"`
	FocusedWorkspaceBg      string `json:"focused_workspace_bg"`
	FocusedWorkspaceText    string `json:"focused_workspace_text"`
	InactiveWorkspaceBorder string `json:"inactive_workspace_border"`
	InactiveWorkspaceBg     string `json:"inactive_workspace_bg"`
	InactiveWorkspaceText   string `json:"inactive_workspace_text"`
	Active_workspaceBorder  string `json:"active_workspace_border"`
	Active_workspaceBg      string `json:"active_workspace_bg"`
	Active_workspaceText    string `json:"active_workspace_text"`
	Urgent_workspaceBorder  string `json:"urgent_workspace_border"`
	Urgent_workspaceBg      string `json:"urgent_workspace_bg"`
	Urgent_workspaceText    string `json:"urgent_workspace_text"`
	BindingModeBorder       string `json:"binding_mode_border"`
	BindingModeBg           string `json:"binding_mode_bg"`
	BindingModeText         string `json:"binding_mode_text"`
}

type Bar struct {
	Id                   string  `json:"id"`
	Mode                 string  `json:"mode"`
	Position             string  `json:"position"`
	Status_command       string  `json:"status_command"`
	Font                 string  `json:"font"`
	Gaps                 *Gaps   `json:"gaps"`
	Height               int     `json:"bar_height"`
	StatusPadding        int     `json:"status_padding"`
	StatusEdgePadding    int     `json:"status_edge_padding"`
	WorkspaceButtons     bool    `json:"workspace_buttons"`
	WorkspaceMinWidth    int     `json:"workspace_min_width"`
	BindingModeIndicator bool    `json:"binding_mode_indicator"`
	Verbose              bool    `json:"verbose"`
	PangoMarkup          bool    `json:"pango_markup"`
	Colors               *Colors `json:"colors"`
}

func (ipc *I3ipc) GetBars() ([]string, error) {
	payload, err := ipc.get(GET_BAR_CONFIG)
	if err != nil {
		return nil, err
	}

	bars := []string{}
	if err := json.Unmarshal(payload, &bars); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return bars, nil
}

func (ipc *I3ipc) GetBar(id string) (*Bar, error) {
	err := ipc.sendHeader(GET_BAR_CONFIG, uint32(len(id)))
	if err != nil {
		return nil, err
	}

	err = ipc.sendPayload([]byte(id))
	if err != nil {
		return nil, fmt.Errorf("failed to send get_bar_config payload: %w", err)
	}

	payload, err := ipc.readResponse()
	if err != nil {
		return nil, err
	}

	bar := &Bar{}
	if err := json.Unmarshal(payload, &bar); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return bar, nil
}
