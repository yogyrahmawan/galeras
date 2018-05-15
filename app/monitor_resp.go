package app

var (
	NumberOfStartedNode = 0
)

// MonitorResp represent monitor resp
type MonitorResp struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// String stringify struct
func (m *MonitorResp) String() string {
	return m.Name + " " + m.Value
}
