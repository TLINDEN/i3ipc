package i3ipc

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

func (node *Node) FindFocused() *Node {
	searchFocused(node.Nodes)
	if __focused == nil {
		searchFocused(node.FloatingNodes)
	}

	return __focused
}

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

func (node *Node) FindCurrentWorkspace() string {
	searchCurrentWorkspace(node.Nodes)
	return __currentworkspace
}

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
