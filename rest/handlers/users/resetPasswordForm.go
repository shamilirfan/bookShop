package users

import (
	"html/template"
	"net/http"
)

func (h *Handler) Reset_Password_Form(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token := r.URL.Query().Get("token")

	tmpl, err := template.ParseFiles("./templates/reset_form.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	// Set Content-Type to HTML
	w.Header().Set("Content-Type", "text/html")

	err = tmpl.Execute(w, token)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
