package proxy

import (
	"crypto/tls"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c fiber.Ctx) bool

	// ModifyRequest allows you to alter the request
	//
	// Optional. Default: nil
	ModifyRequest fiber.Handler

	// ModifyResponse allows you to alter the response
	//
	// Optional. Default: nil
	ModifyResponse fiber.Handler

	// tls config for the http client.
	TLSConfig *tls.Config

	// Client is custom client when client config is complex.
	// Note that Servers, Timeout, WriteBufferSize, ReadBufferSize, TLSConfig
	// and DialDualStack will not be used if the client are set.
	Client *fasthttp.LBClient

	// Servers defines a list of <scheme>://<host> HTTP servers,
	//
	// which are used in a round-robin manner.
	// i.e.: "https://foobar.com, http://www.foobar.com"
	//
	// Required
	Servers []string

	// Timeout is the request timeout used when calling the proxy client
	//
	// Optional. Default: 1 second
	Timeout time.Duration

	// Per-connection buffer size for requests' reading.
	// This also limits the maximum header size.
	// Increase this buffer if your clients send multi-KB RequestURIs
	// and/or multi-KB headers (for example, BIG cookies).
	ReadBufferSize int

	// Per-connection buffer size for responses' writing.
	WriteBufferSize int

	// Attempt to connect to both ipv4 and ipv6 host addresses if set to true.
	//
	// By default client connects only to ipv4 addresses, since unfortunately ipv6
	// remains broken in many networks worldwide :)
	//
	// Optional. Default: false
	DialDualStack bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:           nil,
	ModifyRequest:  nil,
	ModifyResponse: nil,
	Timeout:        fasthttp.DefaultLBClientTimeout,
}

// configDefault function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Timeout <= 0 {
		cfg.Timeout = ConfigDefault.Timeout
	}

	// Set default values
	if len(cfg.Servers) == 0 && cfg.Client == nil {
		panic("Servers cannot be empty")
	}
	return cfg
}
