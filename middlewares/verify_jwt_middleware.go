package middlewares

import (
	"github.com/SawitProRecruitment/UserService/consts"
	"github.com/SawitProRecruitment/UserService/responses"
	"github.com/SawitProRecruitment/UserService/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type VerifyJwtMiddleware struct {
	authenticationService services.AuthenticationServiceInterface
}

func (v *VerifyJwtMiddleware) getWhiteListRoute() map[string]string {

	return map[string]string{
		"/users/register": "POST",
		"/users/login":    "POST",
		"/":               "GET",
	}
}

func (v *VerifyJwtMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		request := c.Request()

		isUrlAllowed := v.isUrlAllowed(request.Method, request.URL.Path)

		if isUrlAllowed {
			return next(c)
		}

		isTokenAllowed, userId := v.isTokenAllowed(request.Header.Get("Authorization"))

		if isTokenAllowed == false {
			return c.JSON(http.StatusBadRequest, responses.BadRequestResponse{
				ErrorMessage: "Your request is made with invalid credential",
			})
		}

		c.Set(consts.ContextAuthorizedUsedId, *userId)

		return next(c)
	}
}

func (v *VerifyJwtMiddleware) isUrlAllowed(requestMethod string, url string) bool {

	whiteListRoute := v.getWhiteListRoute()

	if val, ok := whiteListRoute[url]; ok && requestMethod == val {
		return true
	}

	return false
}

func (v *VerifyJwtMiddleware) isTokenAllowed(jwtToken string) (bool, *int64) {

	tokenString := strings.Replace(jwtToken, "Bearer ", "", -1)

	authorizeResult, err := v.authenticationService.Authorize(tokenString)

	if err != nil || authorizeResult.IsAuthorized == false {
		return false, nil
	}

	return true, &authorizeResult.UserId
}

func NewVerifyJwtMiddleware(svc services.Services) VerifyJwtMiddleware {

	return VerifyJwtMiddleware{
		authenticationService: svc.Authentication,
	}
}
