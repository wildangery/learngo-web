package handlers

import (
	"net/http"

	"github.com/wildangery/learngo-web/packages/renders"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the Home page")
	renders.RenderHtml(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	/*res := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 2 is %d", res))*/
	renders.RenderHtml(w, "about.page.html")
}