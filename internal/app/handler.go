package app

import (
	"log"

	"github.com/jobquestvault/platform-go-challenge/internal/sys/config"
)

type (
	Handler struct {
		log log.Logger
		cfg config.Config
	}
)
