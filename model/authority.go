package model

type Authority struct {
	ID         int
	MasterID   string
	OverseerID string
	WorkerID   string
	ProjectID  int
}
