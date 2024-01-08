package g

import (
	"ats3fx/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/memory/v2"
	"go.uber.org/zap"
)

var (
	CONFIG     *config.Server
	Log        *zap.Logger
	APP        *fiber.App
	MemStorage = memory.New()
)
