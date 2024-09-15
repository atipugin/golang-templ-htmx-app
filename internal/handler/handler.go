package handler

import (
	"log/slog"
	"net/http"

	"github.com/atipugin/golang-templ-htmx-app/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

type Dependencies struct {
	AssetsFS http.FileSystem
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

type validator interface {
	Validate() domain.ValidationErrors
}

func RegisterRoutes(r *chi.Mux, deps Dependencies) {
	home := homeHandler{}

	r.Get("/", handler(home.handleIndex))
	r.Get("/about", handler(home.handleAbout))

	users := usersHandler{}

	r.Route("/users", func(s chi.Router) {
		s.Get("/new", handler(users.newUser))
		s.Post("/", handler(users.createUser))
	})

	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(deps.AssetsFS)))
}

func handler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handleError(w, r, err)
		}
	}
}

func handleError(w http.ResponseWriter, _ *http.Request, err error) {
	slog.Error("error during request", slog.String("err", err.Error()))
	http.Error(w, "Something went wrong", http.StatusInternalServerError)
}

func decodeValid[T validator](r *http.Request) (T, error) {
	var val T
	if err := r.ParseForm(); err != nil {
		return val, err
	}

	if err := schema.NewDecoder().Decode(&val, r.PostForm); err != nil {
		return val, err
	}

	if err := val.Validate(); err != nil {
		return val, err
	}

	return val, nil
}
