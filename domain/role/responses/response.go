package responses

import "rest/domain/role/models"

type SingleResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func BindSingleResponse(model models.Role) *SingleResponse {
	return &SingleResponse{
		ID:    model.ID,
		Title: model.Name,
	}
}
