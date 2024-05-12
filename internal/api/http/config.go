package http

import "github.com/Alp4ka/classifier-mango/internal/interactions/core"

type Config struct {
	// HTTP.
	Port      int
	RateLimit int

	// Security.
	APIKey string

	// Client.
	CoreClient core.Client
}
