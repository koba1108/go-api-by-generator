package openapi

import (
	"github.com/gin-gonic/gin"
	"github.com/go-api-by-generator/external"
	"log"
	"net/http"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func MessageMock(c *gin.Context) {
	// Push通知テスト用
	token := c.Param("targetToken")
	data := map[string]string{}
	_ = c.Bind(&data)
	log.Printf("token: %v", token)
	log.Printf("data: %v", data)
	result, err := external.SendFCM(token, data)
	if err != nil {
		// エラー制御は細かく刷る場合は、 IsInvalidArgument などで検査可能
		// https://godoc.org/firebase.google.com/go/messaging
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusOK, "OK", result)
	}
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},
	{
		"Index",
		http.MethodPost,
		"/message/:targetToken",
		MessageMock,
	},
	{
		"UsersGet",
		http.MethodGet,
		"/users",
		UsersGet,
	},
	{
		"UsersUserIdPost",
		http.MethodPost,
		"/users/:userId",
		UsersUserIdPost,
	},
}
