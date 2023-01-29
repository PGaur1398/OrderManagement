package routes

import (
	"OrderManagement/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type routesHandler struct {
	routesService RoutesService
}

// Adding Order with status and Items
func (rh routesHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest models.OrderRequest
	err := json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil {
		WriteResponse(w, models.BAD_REQUEST_RESPONSE_CODE, models.ErrorResponse{Message: models.BAD_REQUEST, Error: err.Error(), Success: false})
		return
	}
	err = validator.New().Struct(orderRequest)
	if err != nil {
		errMessage := StructValidatorErrorHandling(err)
		WriteResponse(w, models.BAD_REQUEST_RESPONSE_CODE, models.ErrorResponse{Message: models.BAD_REQUEST, Error: errMessage, Success: false})
		return
	}
	response, err := rh.routesService.AddOrder(r.Context(), orderRequest)
	if err != nil {
		WriteResponse(w, models.FAILED_RESPONSE_CODE, models.ErrorResponse{Message: models.FAILED, Error: err.Error(), Success: false})
		return
	}
	WriteResponse(w, http.StatusOK, models.SuccessResponse{Message: models.SUCCESS, Response: response, Success: true})
}

// // Adding item to particular orderId
func (rh routesHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var itemRequest models.ItemRequest
	orderId := r.URL.Query().Get("orderId")
	err := json.NewDecoder(r.Body).Decode(&itemRequest)
	if err != nil {
		WriteResponse(w, models.BAD_REQUEST_RESPONSE_CODE, models.ErrorResponse{Message: models.BAD_REQUEST, Error: err.Error(), Success: false})
		return
	}
	err = validator.New().Struct(itemRequest)
	if err != nil {
		errMessage := StructValidatorErrorHandling(err)
		WriteResponse(w, models.BAD_REQUEST_RESPONSE_CODE, models.ErrorResponse{Message: models.BAD_REQUEST, Error: errMessage, Success: false})
		return
	}
	orderDetail, err := rh.routesService.ValidateOrderId(r.Context(), orderId)
	if err != nil {
		WriteResponse(w, models.FAILED_RESPONSE_CODE, models.ErrorResponse{Message: models.FAILED, Error: err.Error(), Success: false})
		return
	}
	err = rh.routesService.AddItem(r.Context(), orderDetail, itemRequest)
	if err != nil {
		WriteResponse(w, models.FAILED_RESPONSE_CODE, models.ErrorResponse{Message: models.FAILED, Error: err.Error(), Success: false})
		return
	}
	WriteResponse(w, http.StatusOK, models.SuccessResponse{Message: models.SUCCESS, Response: "Item Added Successfully", Success: true})
}

// Change the status of orders to payment_success
func (rh routesHandler) GenerateInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	orderDetail, err := rh.routesService.ValidateOrderId(r.Context(), orderId)
	if err != nil {
		WriteResponse(w, models.FAILED_RESPONSE_CODE, models.ErrorResponse{Message: models.FAILED, Error: err.Error(), Success: false})
		return
	}
	err = rh.routesService.GenerateInvoice(r.Context(), orderDetail)
	if err != nil {
		WriteResponse(w, models.FAILED_RESPONSE_CODE, models.ErrorResponse{Message: models.FAILED, Error: err.Error(), Success: false})
		return
	}
	WriteResponse(w, http.StatusOK, models.SuccessResponse{Message: models.SUCCESS, Response: "Invoice Generated", Success: true})

}

// Get all items in sorted order
func (rh routesHandler) GetAllOrderItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	orderDetail, err := rh.routesService.ValidateOrderId(r.Context(), orderId)
	if err != nil {
		WriteResponse(w, models.FAILED_RESPONSE_CODE, models.ErrorResponse{Message: models.FAILED, Error: err.Error(), Success: false})
		return
	}
	response, err := rh.routesService.GetAllOrderItems(r.Context(), orderDetail)
	if err != nil {
		WriteResponse(w, models.FAILED_RESPONSE_CODE, models.ErrorResponse{Message: models.FAILED, Error: err.Error(), Success: false})
		return
	}
	WriteResponse(w, http.StatusOK, models.SuccessResponse{Message: models.SUCCESS, Response: response, Success: true})
}

// Get Orders based on status
func (rh routesHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orderStatus := r.URL.Query().Get("status")
	responses, err := rh.routesService.GetOrders(r.Context(), orderStatus)
	fmt.Println(responses)
	if err != nil {
		WriteResponse(w, models.FAILED_RESPONSE_CODE, models.ErrorResponse{Message: models.FAILED, Error: err.Error(), Success: false})
		return
	}
	WriteResponse(w, http.StatusOK, models.SuccessResponse{Message: models.SUCCESS, Response: responses, Success: true})

}

func (rh routesHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]
	_, err := rh.routesService.ValidateOrderId(r.Context(), orderId)
	if err != nil {
		WriteResponse(w, models.FAILED_RESPONSE_CODE, models.ErrorResponse{Message: models.FAILED, Error: err.Error(), Success: false})
		return
	}
	err = rh.routesService.DeleteOrder(r.Context(), orderId)
	if err != nil {
		WriteResponse(w, models.FAILED_RESPONSE_CODE, models.ErrorResponse{Message: models.FAILED, Error: err.Error(), Success: false})
		return
	}
	WriteResponse(w, http.StatusOK, models.SuccessResponse{Message: models.SUCCESS, Response: "Order Deleted.", Success: true})
}

// Writes the response as a Standard API JSON response with a response code
func WriteResponse(w http.ResponseWriter, responseCode int, response interface{}) {
	w.Header().Set(models.CONTENT_TYPE, models.APPLICATION_JSON)
	w.WriteHeader(responseCode)
	json.NewEncoder(w).Encode(response)
}

func NewRoutesHandler(router *mux.Router, routesService RoutesService) RoutesHandler {
	handler := routesHandler{routesService: routesService}
	// All routes
	router.HandleFunc("/add/order", handler.AddOrder).Methods("POST")
	router.HandleFunc("/add/item", handler.AddItem).Methods("POST")
	router.HandleFunc("/generate/invoice/{orderId}", handler.GenerateInvoice).Methods("GET")
	router.HandleFunc("/get/order/{orderId}", handler.GetAllOrderItems).Methods("GET")
	router.HandleFunc("/get/orders", handler.GetOrders).Methods("GET")
	router.HandleFunc("/order/{orderId}/remove", handler.DeleteOrder).Methods("DELETE")

	return routesHandler{routesService: routesService}
}
