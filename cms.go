package cms

import (
	"net/http"
	"strings"
	"time"

	"github.com/Mushus/cms/handler"
	"github.com/gin-gonic/gin"
)

// New is create new CMS instance
func New() CMS {
	return &cms{}
}

// CMS is simple content management system.
type CMS interface {
	Start()
}

type cms struct {
}

func (c cms) Start() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      cmsRouter{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}

type cmsRouter struct{}

func (h cmsRouter) ServeHTTP(http.ResponseWriter, *http.Request) {

}

func (c cms) createRouter() http.Handler {
	AppHandler := handler.App{}
	AuthHandler := handler.Auth{App: AppHandler}
	BackendHandler := handler.Backend{App: AppHandler}
	EditorHandler := handler.Editor{App: AppHandler}
	FrontendHandler := handler.Frontend{App: AppHandler}

	router := gin.New()
	router.GET("/admin/login", AuthHandler.LoginPage)
	router.POST("/admin/login", AuthHandler.Login)

	adminRouter := router.Group("/admin/")
	adminRouter.GET("/logout", AuthHandler.Logout)
	adminRouter.GET("/edit", EditorHandler.Edit)
	adminRouter.POST("/edit", EditorHandler.Save)

	router.NoRoute(func(c *gin.Context) {
		url := c.Request.URL.String()
		if strings.HasPrefix(url, "/admin/") {
			// admin
			BackendHandler.List(c)
		} else {
			FrontendHandler.Show(c)
		}
	})

	return router
}
