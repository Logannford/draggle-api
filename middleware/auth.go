package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/appleboy/gin-jwt/v2"
)

type ResponseOptions struct {
	Code    int
	Message interface{}
}

func responseWithError(c *gin.Context, opts ResponseOptions) {
	// default to 401 if no code is passed in
	if(opts.Code == 0){
		opts.Code = http.StatusUnauthorized
	}
	c.AbortWithStatusJSON(opts.Code, gin.H{"error": opts.Message})
}

// Middleware to check if the user is authorized
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		// grab the jwt token from the header
		jwtToken := c.GetHeader("Authorization")

		// if the token is empty, return an error
		if jwtToken == "" {
			// code 401 is the default
			responseWithError(c, ResponseOptions{
				Message: "No token provided",
			})		

			return
		}

		// Continue down the chain to handler etc
		c.Next()
	}
}