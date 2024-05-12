package app

import "github.com/Alp4ka/classifier-mango/internal/interactions/core"

type Config struct {
	// HTTP.
	HTTPPort      int
	HTTPRateLimit int

	// Security.
	APIKey string

	// Core client.
	CoreClient core.Client
}
