package gslf

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/gofiber/fiber"
	"github.com/hashamali/gsl"
	"github.com/rs/zerolog"
)

// Middleware will return a new Fiber middleware for logging.
func Middleware(logger gsl.Log) func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		start := time.Now()
		l := new(c)
		defer l.send(c, logger, start)
		c.Next()
	}
}

// Recover will handle a recover error.
func Recover(c *fiber.Ctx, err error) {
	c.Locals(recoverErrKey, err)
}

type log struct {
	ID         string
	RemoteIP   string
	Host       string
	Method     string
	Path       string
	Protocol   string
	StatusCode int
	Latency    float64
	Error      error
}

func (l *log) MarshalZerologObject(zle *zerolog.Event) {
	zle.Str("id", l.ID)
	zle.Str("remote_ip", l.RemoteIP)
	zle.Str("host", l.Host)
	zle.Str("method", l.Method)
	zle.Str("path", l.Path)
	zle.Str("protocol", l.Protocol)
	zle.Int("status_code", l.StatusCode)
	zle.Float64("latency", l.Latency)

	if l.Error != nil {
		zle.Err(l.Error)
	}
}

func (l *log) send(c *fiber.Ctx, logger gsl.Log, start time.Time) {
	rErr := c.Locals(recoverErrKey)
	if rErr != nil {
		err, ok := rErr.(error)
		if !ok {
			err = fmt.Errorf("%v", rErr)
		}

		l.Error = err
	}

	l.StatusCode = c.Fasthttp.Response.StatusCode()
	l.Latency = float64(time.Since(start).Nanoseconds()) / 1000000.0

	switch {
	case rErr != nil || l.StatusCode >= 300:
		if logger != nil {
			logger.Errorw(l, "")
		}
	case l.StatusCode < 300:
		if logger != nil {
			logger.Infow(l, "")
		}
	}
}

func new(c *fiber.Ctx) *log {
	rid := getString(c.Fasthttp.Response.Header.Peek(fiber.HeaderXRequestID))
	return &log{
		ID:       rid,
		RemoteIP: c.IP(),
		Method:   c.Method(),
		Host:     c.Hostname(),
		Path:     c.Path(),
		Protocol: c.Protocol(),
	}
}

func getString(b []byte) string {
	/* #nosec */
	return *(*string)(unsafe.Pointer(&b))
}

const recoverErrKey = "rerr"
