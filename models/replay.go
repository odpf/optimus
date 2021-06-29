package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	// ReplayStatusAccepted worker picked up the request
	ReplayStatusAccepted   = "accepted"
	ReplayStatusInProgress = "inprogress"
	// ReplayStatusFailed worker fail while processing the replay request
	ReplayStatusFailed  = "failed"  // end state
	ReplayStatusSuccess = "success" // end state
)

type ReplayMessage struct {
	Type    string
	Message string
}

type ReplayWorkerRequest struct {
	ID         uuid.UUID
	Job        JobSpec
	Start      time.Time
	End        time.Time
	Project    ProjectSpec
	DagSpecMap map[string]JobSpec
}

type ReplaySpec struct {
	ID        uuid.UUID
	Job       JobSpec
	StartDate time.Time
	EndDate   time.Time
	Status    string
	Message   ReplayMessage
}
