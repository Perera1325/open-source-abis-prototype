package storage

import (
	"sync"

	"github.com/Perera1325/open-source-abis-prototype/internal/models"
)

var (
	Users = make(map[string]models.User)
	Mu    sync.RWMutex
)
