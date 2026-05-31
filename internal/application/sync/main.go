package sync

import (
	"time"
)

type SyncService struct{}

type SaveFile struct {
	Path      string
	BirthDate time.Time
	Ayanansa  string
	Latitude  float64
	Longitude float64
	Timezone  float64
}

func NewSyncService() *SyncService {
	return &SyncService{}
}
