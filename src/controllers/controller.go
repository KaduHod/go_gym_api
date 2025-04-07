package controllers

import (
	"encoding/json"
	"net/http"
)
type Controller struct {
}

type MetaData struct {
	TotalItens int `json:"total_itens"`
}

type Response[T any] struct {
	Status   string  `json:"status"`
	Data     T       `json:"data"`
	MetaData MetaData `json:"metadata"`
}
type ResponseSwegger struct {
	Status   string  `json:"status"`
    Data     interface {}       `json:"data"`
	MetaData MetaData `json:"metadata"`
}

func SuccessResponse[T any](w http.ResponseWriter, data T, totalItems int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response[T]{
		Status: "success",
		Data:   data,
		MetaData: MetaData{
			TotalItens: totalItems,
		},
	})
}

func InternalServerErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Response[string]{
		Status: "fail",
		Data:   err.Error(),
	})
}

func EmptyResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(Response[interface{}]{
		Status: "success",
		Data:   nil,
	})
}
