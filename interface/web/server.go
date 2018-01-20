package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/Mushus/cms"
	"github.com/labstack/echo"
)

var htmlTemplates = map[string]string{
	"login": `<!doctype html>
<html lang="ja">
<head>
<title>Login</title>
</head>
<body>
<form action="{{.LoginPath}}" method="POST">
<input type="text" name="login">
<input type="password" name="password">
<button type="submit">login</button>
</form>
</body>
</html>`,
	"editor": `<!doctype html>
<html lang="ja">
<head>
<title></title>
</head>
<body>
<input type="text" name="path" disabled value="{{.ID}}">
<textarea></textarea>
</body>
</html>`,
}

type server struct {
	cms cms.CMS
}

// NewServer is create cms server
func NewServer(c cms.CMS) *server {
	return &server{
		cms: c,
	}
}

func (s server) Start() error {
	r := echo.New()
	r.Renderer = newWebRenderer()
	r.GET("/login", s.loginPageHandler)
	r.POST("/login", s.loginHandler)
	r.POST("/logout", s.logoutHandler)
	r.GET("/edit/:id", s.editHandler)
	r.PUT("/:id", s.saveHandler)
	r.POST("/:id", s.saveHandler)
	r.GET("/:id", s.showHandler)
	return r.Start(":8080")
}

func (s server) loginPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login", loginView{
		LoginPath: "/login",
	})
}

func (s server) loginHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login", loginView{
		LoginPath: "/login",
	})
}

func (s server) logoutHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login", loginView{
		LoginPath: "/login",
	})
}

func (s server) editHandler(c echo.Context) error {
	id := c.Param("id")
	// typ := c.FormValue("type")
	//cms.Edit(id, data)
	return c.Render(http.StatusOK, "editor", editView{
		ID: id,
	})
}

func (s server) saveHandler(c echo.Context) error {
	return nil
}
func (s server) showHandler(c echo.Context) error {
	return nil
}

// implements of echo.Renderer
type renderer struct {
	templates *template.Template
}

// create new renderer for echo
func newWebRenderer() *renderer {
	tmpls := make([]string, 0)
	for name, tmpl := range htmlTemplates {
		tmpls = append(tmpls, fmt.Sprintf(`{{define "%s"}}%s{{end}}`, name, tmpl))
	}
	tmplStr := strings.Join(tmpls, "")
	editorTemplate := template.Must(template.New("webpage").Parse(tmplStr))
	return &renderer{
		templates: editorTemplate,
	}
}

func (t *renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type (
	loginView struct {
		LoginPath string
	}
	editView struct {
		ID    string
		Input interface{}
	}
)
