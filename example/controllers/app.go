package controllers

import (
	"net/http"
	"time"
)

// App is a sample controller.
type App struct {
	*Controllers
}

// Before is a magic action that is started before every
// other action of the App controller to render current year.
func (c *App) Before() http.Handler {
	c.Context["year"] = time.Now().Year()
	return nil
}

// Index is an action that renders a home page.
func (c *App) Index() http.Handler {
	c.Context["adminlist"] = []struct {
		Karma        int
		RegisterDate time.Time
		Biography    string
	}{
		{Karma: 10, RegisterDate: time.Now().Local(), Biography: "Smth here..."},
		{Karma: 100, RegisterDate: time.Now().Local(), Biography: "Smth else here..."},
	}
	return c.Render()
}
