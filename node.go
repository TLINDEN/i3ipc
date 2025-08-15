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

// A node  can be an output,  a workspace, a container  or a container
// containing a window.
type Node struct {
	Id                 int     `json:"id"`
	Type               string  `json:"type"` // output, workspace or container
	Name               string  `json:"name"` // workspace number or app name
	Nodes              []*Node `json:"nodes"`
	FloatingNodes      []*Node `json:"floating_nodes"`
	Focused            bool    `json:"focused"`
	Urgent             bool    `json:"urgent"`
	Sticky             bool    `json:"sticky"`
	Border             string  `json:"border"`
	Layout             string  `json:"layout"`
	Orientation        string  `json:"orientation"`
	CurrentBorderWidth int     `json:"current_border_width"`
	Percent            float32 `json:"percent"`
	Focus              []int   `json:"focus"`
	Window             int     `json:"window"` // wayland native
	X11Window          string  `json:"app_id"` // x11 compat
	Current_workspace  string  `json:"current_workspace"`
	Rect               Rect    `json:"rect"`
	WindowRect         Rect    `json:"window_rect"`
	DecoRect           Rect    `json:"deco_rect"`
	Geometry           Rect    `json:"geometry"`
}

var __focused *Node
var __currentworkspace string

// Get the whole information tree,  contains everything from output to
// containers as a tree of nodes.  Each node has a field 'Nodes' which
// points  to  a   list  subnodes.  Some  nodes  also   have  a  field
// 'FloatingNodes' which points to a list of floating containers.
//
// The top level node is the "root" node.
//
// Use the returned node oject to further investigate the wm setup.
func (ipc *I3ipc) GetTree() (*Node, error) {
	err := ipc.sendHeader(GET_TREE, 0)
	if err != nil {
		return nil, err
	}

	payload, err := ipc.readResponse()
	if err != nil {
		return nil, err
	}

	node := &Node{}
	if err := json.Unmarshal(payload.Payload, &node); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return node, nil
}

// Usually called  on the root  node, returns the container  which has
// currently the focus.
func (node *Node) FindFocused() *Node {
	searchFocused(node.Nodes)
	if __focused == nil {
		searchFocused(node.FloatingNodes)
	}

	return __focused
}

// internal recursive focus node searcher
func searchFocused(nodes []*Node) {
	for _, node := range nodes {
		if node.Focused {
			__focused = node
			return
		} else {
			searchFocused(node.Nodes)
			if __focused == nil {
				searchFocused(node.FloatingNodes)
			}
		}

	}
}

// Usually  called  on  the  root node,  returns  the  current  active
// workspace name.
func (node *Node) FindCurrentWorkspace() string {
	searchCurrentWorkspace(node.Nodes)
	return __currentworkspace
}

// internal recursive workspace node searcher
func searchCurrentWorkspace(nodes []*Node) {
	for _, node := range nodes {
		if node.Current_workspace != "" {
			__currentworkspace = node.Current_workspace
			return
		} else {
			searchCurrentWorkspace(node.Nodes)
		}
	}
}
