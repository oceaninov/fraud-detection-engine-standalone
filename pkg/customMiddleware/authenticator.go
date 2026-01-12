package customMiddleware

import (
	"context"
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/environments"
	"gitlab.com/fds22/detection-sys/pkg/hashing"
	"gitlab.com/fds22/detection-sys/src/repositories/rprAuthentication"
	"net/http"
	"strconv"
	"strings"
)

func AuthBasic(envs *environments.Envs) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			// Get the Authorization header
			authHeader := ectx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return ectx.JSON(basicObject.NewResponseError("missing authorization header", http.StatusUnauthorized))
			}

			// Decode the Basic Auth credentials
			authHeaderParts := strings.SplitN(authHeader, " ", 2)
			if len(authHeaderParts) != 2 || authHeaderParts[0] != "Basic" {
				return ectx.JSON(basicObject.NewResponseError("invalid authorization header format", http.StatusUnauthorized))
			}

			decoded, err := base64.StdEncoding.DecodeString(authHeaderParts[1])
			if err != nil {
				return ectx.JSON(basicObject.NewResponseError("invalid authorization encoding", http.StatusUnauthorized))
			}

			// Extract username and password
			credentials := strings.SplitN(string(decoded), ":", 2)
			if len(credentials) != 2 {
				return ectx.JSON(basicObject.NewResponseError("invalid authorization format", http.StatusUnauthorized))
			}

			username, password := credentials[0], credentials[1]
			if !(username == envs.SettingsFDSBasicAuthUser &&
				password == envs.SettingsFDSBasicAuthPassword) {
				return ectx.JSON(basicObject.NewResponseError("unauthorized", http.StatusUnauthorized))
			}

			// Validate user credentials
			ctx := context.Background()
			requestId := ectx.Request().Header.Get(defaultHeaders.XRequestId)
			ctx = context.WithValue(ctx, defaultHeaders.XRequestId, requestId)

			// Replace the request context with the updated context
			ectx.SetRequest(ectx.Request().WithContext(ctx))

			// Call the next handler in the chain
			return next(ectx)
		}
	}
}

func AuthJWT(envs *environments.Envs, rprAuthentication rprAuthentication.Blueprint) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			// Get the Authorization header
			authHeader := ectx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return ectx.JSON(basicObject.NewResponseError("missing authorization header", http.StatusUnauthorized))
			}

			// Split the header to get the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return ectx.JSON(basicObject.NewResponseError("invalid authorization header format", http.StatusUnauthorized))
			}

			// Parse the token
			claims, err := hashing.VerifyJWT(tokenString, envs.JWTSecret)
			if err != nil {
				return ectx.JSON(basicObject.NewResponseError("invalid_token", http.StatusUnauthorized))
			}

			// User read existing
			ctx := context.Background()
			requestId := ectx.Request().Header.Get(defaultHeaders.XRequestId)
			ctx = context.WithValue(ctx, defaultHeaders.XRequestId, requestId)
			user, err := rprAuthentication.ReadRowUser(ctx, map[string]interface{}{
				"id":    claims["userId"].(string),
				"email": claims["email"].(string),
			})
			if err != nil {
				return ectx.JSON(basicObject.NewResponseError("unauthorized", http.StatusUnauthorized))
			}

			// Add user data to the context
			ctx = context.WithValue(ctx, defaultHeaders.Email, user.Email)
			ctx = context.WithValue(ctx, defaultHeaders.XUserId, user.Id)

			// Replace the request context with the updated context
			ectx.SetRequest(ectx.Request().WithContext(ctx))

			// Call the next handler in the chain
			return next(ectx)
		}
	}
}

func AuthJWTChecker(envs *environments.Envs, rprAuthentication rprAuthentication.Blueprint) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			// Get the Authorization header
			authHeader := ectx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return ectx.JSON(basicObject.NewResponseError("missing authorization header", http.StatusUnauthorized))
			}

			// Split the header to get the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return ectx.JSON(basicObject.NewResponseError("invalid authorization header format", http.StatusUnauthorized))
			}

			// Parse the token
			claims, err := hashing.VerifyJWT(tokenString, envs.JWTSecret)
			if err != nil {
				return ectx.JSON(basicObject.NewResponseError("invalid_token", http.StatusUnauthorized))
			}

			// User read existing
			ctx := context.Background()
			requestId := ectx.Request().Header.Get(defaultHeaders.XRequestId)
			ctx = context.WithValue(ctx, defaultHeaders.XRequestId, requestId)
			user, err := rprAuthentication.ReadRowUser(ctx, map[string]interface{}{
				"id":    claims["userId"].(string),
				"email": claims["email"].(string),
			})
			if err != nil {
				return ectx.JSON(basicObject.NewResponseError("unauthorized", http.StatusUnauthorized))
			}

			if strconv.Itoa(user.RoleId) != basicObject.RoleChecker {
				return ectx.JSON(basicObject.NewResponseError("not eligible", http.StatusUnauthorized))
			}

			// Add user data to the context
			ctx = context.WithValue(ctx, defaultHeaders.Email, user.Email)
			ctx = context.WithValue(ctx, defaultHeaders.XUserId, user.Id)

			// Replace the request context with the updated context
			ectx.SetRequest(ectx.Request().WithContext(ctx))

			// Call the next handler in the chain
			return next(ectx)
		}
	}
}

func AuthJWTMaker(envs *environments.Envs, rprAuthentication rprAuthentication.Blueprint) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			// Get the Authorization header
			authHeader := ectx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return ectx.JSON(basicObject.NewResponseError("missing authorization header", http.StatusUnauthorized))
			}

			// Split the header to get the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return ectx.JSON(basicObject.NewResponseError("invalid authorization header format", http.StatusUnauthorized))
			}

			// Parse the token
			claims, err := hashing.VerifyJWT(tokenString, envs.JWTSecret)
			if err != nil {
				return ectx.JSON(basicObject.NewResponseError("invalid_token", http.StatusUnauthorized))
			}

			// User read existing
			ctx := context.Background()
			requestId := ectx.Request().Header.Get(defaultHeaders.XRequestId)
			ctx = context.WithValue(ctx, defaultHeaders.XRequestId, requestId)
			user, err := rprAuthentication.ReadRowUser(ctx, map[string]interface{}{
				"id":    claims["userId"].(string),
				"email": claims["email"].(string),
			})
			if err != nil {
				return ectx.JSON(basicObject.NewResponseError("unauthorized", http.StatusUnauthorized))
			}

			if strconv.Itoa(user.RoleId) != basicObject.RoleMaker {
				return ectx.JSON(basicObject.NewResponseError("not eligible", http.StatusUnauthorized))
			}

			// Add user data to the context
			ctx = context.WithValue(ctx, defaultHeaders.Email, user.Email)
			ctx = context.WithValue(ctx, defaultHeaders.XUserId, user.Id)

			// Replace the request context with the updated context
			ectx.SetRequest(ectx.Request().WithContext(ctx))

			// Call the next handler in the chain
			return next(ectx)
		}
	}
}
