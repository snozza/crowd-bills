package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
)

type BillHandler struct {
	*httprouter.Router
	Logger *log.Logger
}

// NewBillHandler returnsa new instance of BillHandler
func NewBillHandler() *BillHandler {
	h := &BillHandler{
		Router: httprouter.New(),
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
	h.GET("/api/bills", h.handleGetBills)
	return h
}

func (h *BillHandler) handleGetBills(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
		Date string
	}{
		"andrew",
		"somedate",
	}
	encodeJSON(w, data, h.Logger)
}
