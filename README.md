# gslf: Fiber Middleware for gsl
[![godoc](https://godoc.org/github.com/hashamali/gslf?status.svg)](http://godoc.org/github.com/hashamali/gslf)
[![sec](https://img.shields.io/github/workflow/status/hashamali/gslf/security?label=security&style=flat-square)](https://github.com/hashamali/gslf/actions?query=workflow%3Asecurity)
[![go-report](https://goreportcard.com/badge/github.com/hashamali/gslf)](https://goreportcard.com/report/github.com/hashamali/gslf)
[![license](https://badgen.net/github/license/hashamali/gslf)](https://opensource.org/licenses/MIT)

A [gsl](https://github.com/hashamali/gsl) [Fiber](https://github.com/gofiber/fiber) middlware.

## API

* `Middleware(logger gsl.Log)`: Creates a new Fiber middleware with the given logger.
* `Recover(c *fiber.Ctx, err error)`: Recovery handler on errors.