package controllers

import (
	"encoding/json"
	"net/http"
	"rest/app/helpers/responses"
	"rest/app/middleware"
	"rest/app/validators"
	"rest/domain/user/forms"
	"rest/domain/user/repositories"
	"rest/router/lang"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct {
	repo repositories.UserRepository
}

func NewController(repo repositories.UserRepository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) Profile(w http.ResponseWriter, r *http.Request) {
	profile := middleware.CtxValue(r.Context())
	if profile != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.resp_error"), nil)
		return
	}
	responses.ResponseResource(w, *profile)
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	userForm := forms.NewUser{}

	_ = json.NewDecoder(r.Body).Decode(&userForm)

	errsVal, errVal := validators.Validate(userForm)
	if errVal != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.validation_error"), errsVal)
		return
	}

	userModel, errRepo := c.repo.Create(userForm)
	if errRepo != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}
	responses.ResponseResource(w, *userModel)
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	params := responses.ParamsMapFromRequest(r)
	_, ok := params["all"]
	if ok {
		users, errRepo := c.repo.List(params)
		if errRepo != nil {
			responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
			return
		}

		responses.ResponseCollection(w, users)
	} else {
		users, errRepo := c.repo.Paginate(params)
		if errRepo != nil {
			responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
			return
		}
		errorRs := responses.CreateResponse(w, users, http.StatusOK)
		if errorRs != nil {
			responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.resp_error"), nil)
		}
	}
}

func (c *Controller) Show(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	idInt, errSyntax := strconv.Atoi(idString)
	if errSyntax != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.syntax_error"), nil)
		return
	}
	user, errRepo := c.repo.ShowModel(uint(idInt))
	if errRepo != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}
	responses.ResponseResource(w, *user)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	idInt, errSyntax := strconv.Atoi(idString)
	if errSyntax != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.syntax_error"), nil)
		return
	}

	errRepo := c.repo.Delete(uint(idInt))
	if errRepo != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}
	responses.ResponseResource(w, "User was deleted successfully")
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	user := middleware.CtxValue(r.Context())
	userUpdateForm := forms.UpdateUser{Id: user.ID}

	_ = json.NewDecoder(r.Body).Decode(&userUpdateForm)

	errsVal, errVal := validators.Validate(userUpdateForm)
	if errVal != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.validation_error"), errsVal)
		return
	}

	userUpdate, errRepo := c.repo.UpdateModel(user.ID, userUpdateForm)
	if errRepo != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}
	responses.ResponseResource(w, *userUpdate)
}
