// package main

// import (
// 	"draggle/api/controllers"
// 	"net/http"

// 	"github.com/gin-gonic/gin"

// 	"os"
// )

// var db = make(map[string]string)

// func setupRouter() *gin.Engine {
// 	// Disable Console Color
// 	// gin.DisableConsoleColor()

// 	// ':=' shorthand variable declaration
// 	r := gin.Default()

// 	// Ping test
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.String(http.StatusOK, "pong")

// 		jwt, _ := c.Get("jwt")

// 		c.JSON(http.StatusOK, gin.H{
// 			"jwt": jwt,
// 		})
// 	})

// 	// Get user value
// 	r.GET("/user/:name", func(c *gin.Context) {
// 		user := c.Params.ByName("name")
// 		c.String(http.StatusOK, "Hello %s", user)
// 		value, ok := db[user]
// 		if ok {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
// 		}
// 	})

// 	r.GET("/welcome/:id", func(c *gin.Context){
// 		id := c.Params.ByName("id")

// 		// use the id to auth the user from supabase
// 		c.String(http.StatusOK, "Welcome %s", id)
// 	})

// 	// Authorized group (uses gin.BasicAuth() middleware)
// 	// Same than:
// 	// authorized := r.Group("/")
// 	// authorized.Use(gin.BasicAuth(gin.Credentials{
// 	//	  "foo":  "bar",
// 	//	  "manu": "123",
// 	//}))
// 	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
// 		"foo":  "bar", // user:foo password:bar
// 		"manu": "123", // user:manu password:123
// 	}))

// 	/* example curl for /admin with basicauth header
// 	   Zm9vOmJhcg== is base64("foo:bar")

// 		curl -X POST \
// 	  	http://localhost:8080/admin \
// 	  	-H 'authorization: Basic Zm9vOmJhcg==' \
// 	  	-H 'content-type: application/json' \
// 	  	-d '{"value":"bar"}'
// 	*/
// 	authorized.POST("admin", func(c *gin.Context) {
// 		user := c.MustGet(gin.AuthUserKey).(string)

// 		// Parse JSON
// 		var json struct {
// 			Value string `json:"value" binding:"required"`
// 		}

// 		if c.Bind(&json) == nil {
// 			db[user] = json.Value
// 			c.JSON(http.StatusOK, gin.H{"status": "ok"})
// 		}
// 	})

// 	return r
// }

// // 'gin.HandlerFunc' is a type used to define middleware functions
// // it has access to incoming HTTP request and outgoing HTTP response
// // func Logger() gin.HandlerFunc {
// // 	return func(c *gin.Context){
// // 		start := time.Now()

// // 		// Process the request
// // 		c.Next()

// // 		end := time.Now()
// // 		latency := end.Sub(start)
// // 		fmt.Printf("[%s] %s - %v\n", c.Request.Method, c.Request.URL.Path, latency)
// // 	}
// // }

// func main() {
// 	// get the port from the env
// 	port := os.Getenv("PORT")
// 	if(port == ""){
// 		port = "8080"
// 	}

// 	r := setupRouter()

// 	// router.Use(Logger())

// 	public := r.Group("/api")

// 	// Apply the middleware to the router
// 	//r.Use(middleware.AuthMiddleware())

// 	public.POST("/register", controllers.Register)
// 	public.GET("/register", controllers.Register)

// 	// Listen and Server in 0.0.0.0:8080
// 	r.Run(port)
// }

package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

		// at the root path, render the index.html.tmpl template
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := map[string]string{
            "Region": os.Getenv("FLY_REGION"),
        }

				// print hello 
				fmt.Println("Hello, World! x10")

        t.ExecuteTemplate(w, "index.html.tmpl", data)
    })

		// test endpoint before starting stripe webhook integration 
		http.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})

    log.Println("listening on", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
