package controllers

import (
	"net/http"

	"github.com/jgsheppa/fletters/views"
)

func NewStatic() *Static {
	return &Static{
		Home:     views.NewView("bootstrap", http.StatusFound, "static/home"),
		Contact:  views.NewView("bootstrap", http.StatusFound, "static/contact"),
		About:    views.NewView("bootstrap", http.StatusFound, "static/about"),
		NotFound: views.NewView("bootstrap", http.StatusNotFound, "static/404"),
	}
}

type Static struct {
	Home     *views.View
	NotFound *views.View
	Contact  *views.View
	About    *views.View
}
