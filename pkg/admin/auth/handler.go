package auth

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taverok/lazyadmin/pkg/admin/page"
	"github.com/taverok/lazyadmin/pkg/middleware"
	"github.com/taverok/lazyadmin/pkg/utils"
)

type Handler struct {
	Service Service
	Prefix  string
}

func (it *Handler) Register(router *mux.Router) {
	router.HandleFunc(fmt.Sprintf("/%s/login", it.Prefix), it.form).Methods(http.MethodGet)
	router.HandleFunc(fmt.Sprintf("/%s/login", it.Prefix), it.authenticate).Methods(http.MethodPost)
}

func (it *Handler) form(w http.ResponseWriter, r *http.Request) {
	p := page.Page[page.Login]{}

	defer utils.HtmlTemplate("login").Execute(w, &p)
}

func (it *Handler) authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	request := NewLoginRequest(r.Form)
	jwt, err := it.Service.AuthJwt(request)
	if err != nil {
		// TODO: ошибку на фронт
		slog.Error(err.Error())
		return
	}

	cookie := http.Cookie{
		Name:  middleware.AuthKey,
		Value: jwt,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, fmt.Sprintf("/%s", it.Prefix), http.StatusTemporaryRedirect)
}
