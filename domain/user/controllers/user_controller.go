package controllers

import (
	"encoding/json"
	"net/http"
	"rest/app/helpers/responses"
	"rest/app/middleware"
	"rest/app/validators"
	"rest/domain/user/forms"
	"rest/domain/user/models"
	userResp "rest/domain/user/responses"
	"rest/domain/user/services"
	"rest/router/lang"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct {
	us services.IUserService
}

func NewController(uServ services.IUserService) *Controller {
	return &Controller{us: uServ}
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	params := responses.ParamsMapFromRequest(r)
	_, ok := params["all"]

	users := make([]models.User, 0)
	var err error

	if ok {
		users, err = c.us.GetUsers(params)
		if err != nil {
			responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
			return
		}
		responses.ResponseCollection(w, users)
		return
	}

	paginate, pgErr := c.us.Paginate(params)
	if pgErr != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}

	err = responses.CreateResponse(w, paginate, http.StatusOK)
	if err != nil {
		responses.GenerateErrorResponse(w, err.Error(), nil)
		return
	}
}

func (c *Controller) Profile(w http.ResponseWriter, r *http.Request) {
	profile := middleware.CtxValue(r.Context())
	if profile == nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.resp_error"), nil)
		return
	}
	responses.ResponseResource(w, userResp.BindingUserModelResponse(*profile))
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	userForm := forms.NewUser{}

	decodeErr := json.NewDecoder(r.Body).Decode(&userForm)

	if decodeErr != nil {
		responses.GenerateErrorResponse(w,
			"Ошибка десериализации тела запроса",
			map[string]interface{}{
				"message": decodeErr.Error(),
			})
		return
	}

	errsVal, errVal := validators.Validate(userForm)
	if errVal != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.validation_error"), errsVal)
		return
	}

	userModel, errRepo := c.us.AddUser(userForm)
	if errRepo != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}
	responses.ResponseResource(w, *userModel)
}

func (c *Controller) Show(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	idInt, errSyntax := strconv.Atoi(idString)
	if errSyntax != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.syntax_error"), nil)
		return
	}
	user, errRepo := c.us.ShowUser(uint(idInt))
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

	errRepo := c.us.DestroyUser(uint(idInt))
	if errRepo != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}
	responses.ResponseResource(w, "User was deleted successfully")
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	user := middleware.CtxValue(r.Context())
	userUpdateForm := forms.UpdateUser{Id: user.ID}

	decodeErr := json.NewDecoder(r.Body).Decode(&userUpdateForm)

	if decodeErr != nil {
		responses.GenerateErrorResponse(w,
			"Ошибка десериализации тела запроса",
			map[string]interface{}{
				"message": decodeErr.Error(),
			})
		return
	}

	errsVal, errVal := validators.Validate(userUpdateForm)
	if errVal != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.validation_error"), errsVal)
		return
	}

	userUpdate, errRepo := c.us.UpdateUser(user.ID, userUpdateForm)
	if errRepo != nil {
		responses.GenerateErrorResponse(w, lang.GetTranslator().Trans("messages.db_error"), nil)
		return
	}
	responses.ResponseResource(w, *userUpdate)
}
