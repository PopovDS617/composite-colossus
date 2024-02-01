package handlers

import (
	"context"
	"gateway/client"
	"net/http"
	"strconv"
)

type InvoiceHandler struct {
	client *client.HTTPClient
}

func NewInvoiceHandler(client *client.HTTPClient) *InvoiceHandler {
	return &InvoiceHandler{
		client: client,
	}
}

func (iv *InvoiceHandler) HandleGetInvoice(rw http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		writeJSON(rw, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
	values, ok := r.URL.Query()["obu"]

	if !ok || len(values) == 0 || len(values[0]) == 0 {
		return writeJSON(rw, http.StatusBadRequest, map[string]string{"error": "missing OBU ID"})

	}

	obuID, err := strconv.Atoi(values[0])

	if err != nil {
		return writeJSON(rw, http.StatusBadRequest, map[string]string{"error": "wrong OBU ID format"})

	}

	inv, err := iv.client.GetInvoice(context.Background(), obuID)

	if err != nil {
		return err
	}

	return writeJSON(rw, http.StatusOK, inv)
}
