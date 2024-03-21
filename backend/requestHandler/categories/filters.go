package categories

import "net/http"

func (c *CategoryClient) GetFilterCategories(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
}
