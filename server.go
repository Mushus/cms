package cms

import (
	"github.com/labstack/echo"
)

type server struct {
	cms *cms
}

func NewServer(c *cms) *server {
	return &server{
		cms: c,
	}
}

func (s server) Start() error {
	r := echo.New()
	r.GET("/editor/:id", editorHandler)
	r.PUT("/:id", editHandler)
	r.POST("/:id", editHandler)
	r.GET("/:id", showHandler)
	return r.Start(":80")
}

func editHandler(c echo.Context) error {
	// id := c.Param("id")
	// typ := c.FormValue("type")
	// cms.Edit(id, data)
	return nil
}

func editorHandler(c echo.Context) error {
	return nil
}
func showHandler(c echo.Context) error {
	return nil
}
