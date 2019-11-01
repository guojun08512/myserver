package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/gddo/httputil"
	"github.com/labstack/echo/v4"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

const (
	defaultContentTypeOffer = "application/json"
	acceptContentTypeKey    = "offer-content-type"
)

var errNotAuthorized = errors.New("no or invalid jwt token provided. You are not authorized")

type AcceptOptions struct {
	DefaultContentTypeOffer string
	Offers                  []string
}

var contentTypeOffers = []string{
	echo.MIMEApplicationJSON,
	echo.MIMETextHTML,
	echo.MIMETextPlain,
}

// AuthMiddleware is middleware used for each request. Includes functionality that validates the JWT tokens and user
// permissions.
func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := getToken(c)
			if err != nil {
				return c.String(http.StatusUnauthorized, err.Error())
			}

			// Validate token
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// All ok, continue
				username, okUsername := claims["username"]
				fmt.Println(username, okUsername)
				//roles, okPerms := claims["roles"]
				//if okUsername && okPerms && roles != nil {
				//	// Look through the perms until we find that the user has this permission
				//	err := rl.CheckRole(roles, c.Request().Method, c.Path())
				//	if err != nil {
				//		return c.String(http.StatusForbidden, fmt.Sprintf("Permission denied for user %s. %v", username, err))
				//	}
				//}
				return next(c)
			}
			return c.String(http.StatusUnauthorized, errNotAuthorized.Error())
		}
	}
}

func getToken(c echo.Context) (*jwt.Token, error) {
	// Get the token
	jwtRaw := c.Request().Header.Get("Authorization")
	split := strings.Split(jwtRaw, " ")
	if len(split) != 2 {
		return nil, errNotAuthorized
	}
	jwtString := split[1]

	// Parse token
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Accept ...
func Accept(args ...AcceptOptions) echo.MiddlewareFunc {
	var opts AcceptOptions
	if len(args) > 0 {
		opts = args[0]
	}
	if opts.DefaultContentTypeOffer == "" {
		opts.DefaultContentTypeOffer = defaultContentTypeOffer
	}
	if opts.Offers == nil {
		opts.Offers = contentTypeOffers
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			contentTypeOffer := httputil.NegotiateContentType(c.Request(), opts.Offers, opts.DefaultContentTypeOffer)
			c.Set(acceptContentTypeKey, contentTypeOffer)
			return next(c)
		}
	}
}
