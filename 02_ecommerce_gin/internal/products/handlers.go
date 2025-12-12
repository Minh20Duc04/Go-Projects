package products

import (
	"net/http"
	"github.com/Minh20Duc04/Go-Projects/internal/json"
)

type handler struct { //nhan request xong chạy -> service để xử lý
	service Service
}

func NewHandler(service Service) *handler { //giong constructor, khai bao de chay
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := struct{
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)
}
