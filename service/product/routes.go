package product

import (
	"fmt"
	"net/http"

	"github.com/georgs1xth/APIBACKEND/types"
	"github.com/georgs1xth/APIBACKEND/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload for %v", errors))
		return
	}
	// add product
	err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Quantity:    payload.Quantity,
		Price:       payload.Price,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error at creating product %v", err))
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
