package controllers

import "net/http"

type BaseRailController interface {
	Handle(w http.ResponseWriter, r *http.Request, params map[string]string)
}
