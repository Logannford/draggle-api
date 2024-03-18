package middleware

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
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

var identityKey = "id"

// User demo
type User struct {
  UserName  string
  FirstName string
  LastName  string
}

func helloHandler(c *gin.Context) {
  claims := jwt.ExtractClaims(c)
  user, _ := c.Get(identityKey)
  c.JSON(200, gin.H{
    "userID":   claims[identityKey],
    "userName": user.(*User).UserName,
    "text":     "Hello World.",
  })
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

		// add the token to the context
		c.Set("jwt", jwtToken)

		c.JSON(http.StatusOK, gin.H{
			"message": jwtToken,
		})

		// Continue down the chain to handler etc
		c.Next()
	}
}