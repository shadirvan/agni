package common

const (
	RouteBeacon = "/api/v1/beacon"
	RouteTasks  = "/api/v1/tasks"
	RouteResult = "/api/v1/result"
)

type BeaconRequest struct {
	ID         string `json:"id"`
	Hostname   string `json:"hostname"`
	Username   string `json:"username"`
	OS         string `json:"os"`
	Arch       string `json:"arch"`
	PID        int    `json:"pid"`
	InternalIP string `json:"internal_ip"`
}

type Task struct {
	ID      string `json:"id"`
	Command string `json:"command"`
	Args    string `json:"args"`
	Timeout int    `json:"timeout"`
}

type TaskResponse struct {
	TaskID  string `json:"task_id"`
	Success bool   `json:"success"`
	Output  string `json:"output"`
	Error   string `json:"error"`
}
