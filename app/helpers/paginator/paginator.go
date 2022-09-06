package paginator

import (
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
)

type Paginator struct {
	Meta MetaData    `json:"meta"`
	Data interface{} `json:"data"`
}

type MetaData struct {
	PerPage  int   `json:"perPage"`
	LastPage int   `json:"lastPage"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
}

func (m *MetaData) GetOffset() int {
	return (m.GetPage() - 1) * m.GetPerPage()
}

func (m *MetaData) GetPerPage() int {
	if m.PerPage == 0 {
		m.PerPage = 10
	}
	return m.PerPage
}

func (m *MetaData) GetPage() int {
	if m.Page == 0 {
		m.Page = 1
	}
	return m.Page
}

// Paginate
// call this method when you're filtering and sorting your model
func Paginate(r *http.Request, paginator *Paginator, dbQuery *gorm.DB) func(db *gorm.DB) *gorm.DB {
	pageReq := r.URL.Query().Get("page")
	if pageReq != "" {
		pageInt, err := strconv.Atoi(pageReq)
		if err == nil {
			paginator.Meta.Page = pageInt
		}
	} else {
		paginator.Meta.Page = paginator.Meta.GetPage()
	}
	perPageReq := r.URL.Query().Get("perPage")
	if perPageReq != "" {
		perPageInt, err := strconv.Atoi(perPageReq)
		if err == nil {
			paginator.Meta.PerPage = perPageInt
		}
	} else {
		paginator.Meta.PerPage = paginator.Meta.GetPerPage()
	}
	var totalRows int64
	dbQuery.Count(&totalRows)
	paginator.Meta.Total = totalRows
	if totalRows == 0 {
		paginator.Data = []string{}
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(paginator.Meta.GetPerPage())))

	paginator.Meta.LastPage = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return dbQuery.Offset(paginator.Meta.GetOffset()).Limit(paginator.Meta.GetPerPage())
	}
}

func PaginateFromMap(params map[string]string, paginator *Paginator, dbQuery *gorm.DB) func(db *gorm.DB) *gorm.DB {
	_, wpOk := params["withPaginate"]

	pageReq, ok := params["page"]
	if ok {
		pageInt, err := strconv.Atoi(pageReq)
		if err == nil {
			paginator.Meta.Page = pageInt
		}
	} else {
		paginator.Meta.Page = paginator.Meta.GetPage()
	}
	perPageReq, ok := params["perPage"]
	if ok {
		perPageInt, err := strconv.Atoi(perPageReq)
		if err == nil {
			paginator.Meta.PerPage = perPageInt
		}
	} else {
		paginator.Meta.PerPage = paginator.Meta.GetPerPage()
	}
	var totalRows int64
	dbQuery.Count(&totalRows)

	paginator.Meta.Total = totalRows
	if totalRows == 0 {
		paginator.Data = []string{}
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(paginator.Meta.GetPerPage())))
	paginator.Meta.LastPage = totalPages

	return func(db *gorm.DB) *gorm.DB {
		if !wpOk {
			return dbQuery
		}
		return dbQuery.Offset(paginator.Meta.GetOffset()).Limit(paginator.Meta.GetPerPage())
	}
}
