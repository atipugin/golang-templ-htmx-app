package handler

import (
	"net/http"

	"github.com/atipugin/golang-templ-htmx-app/internal/domain"
	"github.com/atipugin/golang-templ-htmx-app/internal/view/users"
)

type usersHandler struct {
}

func (h usersHandler) newUser(w http.ResponseWriter, r *http.Request) error {
	return users.NewUser(users.NewUserProps{}).Render(r.Context(), w)
}

type createUserRequest struct {
	Email string `schema:"email"`
}

func (r createUserRequest) Validate() domain.ValidationErrors {
	errs := domain.ValidationErrors{}

	if r.Email == "" {
		errs.Add("email", "is required")
	}

	if errs.Any() {
		return errs
	}

	return nil
}

func (h usersHandler) createUser(w http.ResponseWriter, r *http.Request) error {
	params, err := decodeValid[createUserRequest](r)
	if err != nil {
		if valErrs, ok := err.(domain.ValidationErrors); ok {
			return users.NewUser(users.NewUserProps{Email: params.Email, Errors: valErrs}).Render(r.Context(), w)
		}

		return err
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

	return nil
}
