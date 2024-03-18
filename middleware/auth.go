package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseOptions struct {
	Code    int
	Message interface{}
}

func ResponseWithError(c *gin.Context, opts ResponseOptions) {
	// default to 401 if no code is passed in
	if(opts.Code == 0){
		opts.Code = http.StatusUnauthorized
	}
	c.AbortWithStatusJSON(opts.Code, gin.H{"error": opts.Message})
}

// Middleware to check if the user is authorized
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		// grab the jwt token from the header
		jwtToken := c.GetHeader("Authorization")

		// if the token is empty, return an error
		if jwtToken == "" {
			// code 401 is the default
			ResponseWithError(c, ResponseOptions{
				Message: "No token provided",
			})		

			return
		}

		//jwt.ErrEmptyAuthHeader = "No token"

		// Continue down the chain to handler etc
		c.Next()
	}
}