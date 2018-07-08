package main

import "net/http"

func adminIndexHandler(w http.ResponseWriter, r *http.Request) *appError {
	user := profileFromSession(r)
	if user == nil {
		http.Redirect(w, r, "/login?redirect=/admin/", http.StatusFound)
		return nil
	}
	return adminIndexTmpl.Execute(w, r, nil)
}
