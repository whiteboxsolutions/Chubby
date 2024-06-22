package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	chubby "github.com/whiteboxsolutions/Chubby"
)

type RollChecker func(fiber.Ctx) bool

func New(config ...Config) fiber.Handler {
	// Init config
	cfg := configDefault(config...)

	// Return middleware handler
	return func(c fiber.Ctx) error {
		// Filter request to skip middleware
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Extract and verify key
		allowed := chubby.HasRoll(cfg.Roll.Value, cfg.Requirement)
		if !allowed {
			return cfg.ErrorHandler(c, fmt.Errorf("unauthorized"))
		} else {
			return cfg.SuccessHandler(c)
		}
	}
}
