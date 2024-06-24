package chubbyFiber

import (
	"github.com/gofiber/fiber/v3"
	chubby "github.com/whiteboxsolutions/Chubby"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip middleware.
	// Optional. Default: nil
	Next func(fiber.Ctx) bool

	// SuccessHandler defines a function which is executed for a valid key.
	// Optional. Default: nil
	SuccessHandler fiber.Handler

	// ErrorHandler defines a function which is executed for an invalid key.
	// It may be used to define a custom error.
	// Optional. Default: 401 Invalid or expired key
	ErrorHandler fiber.ErrorHandler

	Roll        chubby.Roll
	Requirement chubby.Roll
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	SuccessHandler: func(c fiber.Ctx) error {
		return c.Next()
	},
	ErrorHandler: func(c fiber.Ctx, err error) error {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	},
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.SuccessHandler == nil {
		cfg.SuccessHandler = ConfigDefault.SuccessHandler
	}
	if cfg.ErrorHandler == nil {
		cfg.ErrorHandler = ConfigDefault.ErrorHandler
	}
	if cfg.Requirement == (chubby.Roll{}) {
		panic("fiber: chubby requires a roll to compare against")
	}

	return cfg
}
