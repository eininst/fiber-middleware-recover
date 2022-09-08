package recovers

import (
	"fmt"
	"github.com/eininst/flog"
	"github.com/gofiber/fiber/v2"
	"runtime"
)

var rlog flog.Interface

func init() {
	logf := "${level} %s[${pid}]%s ${time} ${msg}"
	rlog = flog.New(flog.Config{
		Format: fmt.Sprintf(logf, flog.Red, flog.Reset),
	})
}

type ErrorHandler func(c *fiber.Ctx, r interface{}) *fiber.Error

type Config struct {
	StackTraceBufLen int
	Handler          ErrorHandler
}

var DefaultConfig = Config{
	StackTraceBufLen: 1280,
	Handler: func(c *fiber.Ctx, r interface{}) *fiber.Error {
		return fiber.NewError(fiber.StatusInternalServerError)
	},
}

func stackTraceHandler(e interface{}, bufLen int) {
	buf := make([]byte, bufLen)
	buf = buf[:runtime.Stack(buf, false)]
	rlog.Errorf("panic: %v\n%s\n", e, buf)
}

func New(config ...Config) fiber.Handler {
	var cfg = DefaultConfig
	if len(config) > 0 {
		cfg = config[0]
		if cfg.Handler == nil {
			cfg.Handler = DefaultConfig.Handler
		}
		if cfg.StackTraceBufLen == 0 {
			cfg.StackTraceBufLen = DefaultConfig.StackTraceBufLen
		}
	}

	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				stackTraceHandler(r, cfg.StackTraceBufLen)

				err = cfg.Handler(c, r)
			}
		}()
		return c.Next()
	}
}
