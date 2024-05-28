package middleware

import "github.com/labstack/echo"

type GoMiddleware struct {
}

func (g *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(ctx)
	}
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
