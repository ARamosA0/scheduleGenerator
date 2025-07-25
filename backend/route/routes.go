package route

import (
	"backend.com/backend/internal/request"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")
	request.RegisterTeacherRoutes(api)
	request.RegisterRoomRoutes(api)
	request.RegisterAssigmentsRoutes(api)
	request.RegisterGroupRoutes(api)
	request.RegisterRuleRoutes(api)
	request.RegisterSubjectRoutes(api)
	request.RegisterTemplatesRoutes(api)
	request.RegisterAssigmentsRoutes(api)
}
