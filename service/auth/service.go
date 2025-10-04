package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

var Store *sessions.CookieStore

func InitGothic() {
	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	gothic.Store = Store

	goth.UseProviders(
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			os.Getenv("GITHUB_CALLBACK_URL"),
			"user:email", "read:user",
		),
	)
}
