package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	// SessionKey is key of session with logged in admin user information
	SessionKey = "github.com/mushus/cms/server/middreware/auth"

	// ContextKey is key of context with logged in admin user information
	ContextKey = "github.com/mushus/cms/server/middreware/auth"
)

// Admin is struct having loggedin admin user information
type Auth interface {
	Login(user interface{})
	User() (user interface{})
	Logout()
}

type auth struct {
	user   interface{}
	isSave bool
}

func (a *auth) Login(user interface{}) {
	a.user = user
	a.isSave = true
}
func (a *auth) User() interface{} {
	return a.user
}
func (a *auth) Logout(user interface{}) {
	a.user = nil
	a.isSave = true
}

// Option is this Middreware Option.
// It use for "Required" function.
type Option struct {
	SessionKey string
	ContextKey string
	ErrorFunc  gin.HandlerFunc
}

// AdminPage is Middreware for checking whether it is a request sent by the logged in user or not.
// It required the middreware "github.com/gin-contrib/sessions".
func AdminPage(opt Option) (getAuth func(ctx *gin.Context) Auth, adminPageMW gin.HandlerFunc, requiredMW gin.HandlerFunc) {
	if opt.SessionKey == "" {
		opt.SessionKey = SessionKey
	}
	if opt.ContextKey == "" {
		opt.ContextKey = ContextKey
	}

	getAuth = func(ctx *gin.Context) Auth {
		return ctx.MustGet(opt.ContextKey).(Auth)
	}

	adminPageMW = func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		admin, ok := sess.Get(opt.SessionKey).(auth)
		if !ok {
			admin = auth{
				isSave: true,
			}
		}
		ctx.Set(opt.ContextKey, &admin)
		ctx.Next()
		if admin.isSave {
			sess.Set(opt.SessionKey, admin)
		}
	}

	requiredMW = func(ctx *gin.Context) {
		user := getAuth(ctx).User()
		if user == nil {
			opt.ErrorFunc(ctx)
		}
		ctx.Next()
	}
	return
}
