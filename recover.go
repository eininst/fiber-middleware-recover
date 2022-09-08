package recovers

import (
	"github.com/eininst/flog"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

type ErrorHandler func(r interface{}) *fiber.Error

type Config struct {
	StackTraceBufLen int
	Handler          ErrorHandler
}

var DefaultConfig = Config{
	StackTraceBufLen: 1280,
	Handler: func(r interface{}) *fiber.Error {
		return fiber.NewError(fiber.StatusInternalServerError)
	},
}

func stackTraceHandler(e interface{}, bufLen int) {
	buf := make([]byte, bufLen)
	buf = buf[:runtime.Stack(buf, false)]
	flog.Errorf("panic: %v\n%s\n", e, buf)
}

func New(config ...Config) fiber.Handler {
	var cfg = DefaultConfig
	var errhandler ErrorHandler
	if len(config) > 0 {
		cfg = config[0]
		if cfg.Handler == nil {
			errhandler = DefaultConfig.Handler
		}
		if cfg.StackTraceBufLen == 0 {
			cfg.StackTraceBufLen = DefaultConfig.StackTraceBufLen
		}
	}

	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				stackTraceHandler(r, cfg.StackTraceBufLen)

				if errhandler != nil {
					err = errhandler(r)
				} else {
					err = fiber.NewError(fiber.StatusInternalServerError)
				}
			}
		}()
		return c.Next()
	}
}
