package customer

import (
	"net/http"
	"strconv"
)

func GetAll(w http.ResponseWriter, r *http.Request, service interfaces.ICustomerService) {
	page, errPage := strconv.Atoi(r.URL.Query().Get("page"))
	if errPage != nil || page < 1 {
		page = 1
	}
	limit, errLimit := strconv.Atoi(r.URL.Query().Get("limit"))
	if errLimit != nil || limit < 1 {
		limit = 10
	}

	responseModel, err := service.GetAll(page, limit)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	JSON(w, http.StatusOK, true, responseModel, "")
}
