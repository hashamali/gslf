# slfmw: Fiber Middleware for sl

An [sl](https://github.com/hashamali/sl) [Fiber](https://github.com/gofiber/fiber) middlware.

## API

* `Middleware(logger sl.Log)`: Creates a new Fiber middleware with the given logger.
* `Recover(c *fiber.Ctx, err error)`: Recovery handler on errors.