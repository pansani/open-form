package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/occult/pagode/pkg/context"
	"github.com/occult/pagode/pkg/msg"
	"github.com/romsar/gonertia/v2"
)

func InertiaProps() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Get authenticated user
			user := ctx.Get(context.AuthenticatedUserKey)

			// Collect errors by type
			flash := make(map[string][]string)
			for _, typ := range []msg.Type{
				msg.TypeSuccess,
				msg.TypeInfo,
				msg.TypeWarning,
				msg.TypeDanger,
			} {
				messages := msg.Get(ctx, typ)
				if len(messages) > 0 {
					flash[string(typ)] = messages
				}
			}

			// Set Inertia props
			newCtx := gonertia.SetProps(ctx.Request().Context(), map[string]any{
				"flash": flash,
				"auth": map[string]any{
					"user": user,
				},
				"sidebarOpen": false,
			})

			// Replace request context
			ctx.SetRequest(ctx.Request().WithContext(newCtx))

			return next(ctx)
		}
	}
}
