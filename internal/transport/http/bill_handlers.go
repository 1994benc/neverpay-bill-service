package http

import (
	"1994benc/neverpay-api/internal/bill"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func (h *Handler) GetBill(w http.ResponseWriter, r *http.Request) {
	h.commonHeadersSetUp(w)
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := h.parseID(idStr)
	if err != nil {
		log.Error("Error parsing UINT from ID: %s", err)
		http.Error(w, "Error parsing UINT from ID", http.StatusInternalServerError)
		return
	}
	bill, err := h.BillService.GetBill(uint(id))
	if err != nil {
		log.Error("Error retreiving bill by ID: %s", err)
		http.Error(w, "Error retreiving bill by ID", http.StatusInternalServerError)
		return
	}

	err = bill.ToJSON(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllBills(w http.ResponseWriter, r *http.Request) {
	h.commonHeadersSetUp(w)
	bills, err := h.BillService.GetAllBills()
	if err != nil {
		log.Error("Error retrieving bills: %s", err)
		http.Error(w, "Error retrieving bills", http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(bills)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AddBill(w http.ResponseWriter, r *http.Request) {
	h.commonHeadersSetUp(w)
	var bill bill.Bill
	err := bill.FromJSON(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}
	bill, err = h.BillService.AddBill(bill)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding bill: %s", err), http.StatusInternalServerError)
		return
	}
	err = bill.ToJSON(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateBill(w http.ResponseWriter, r *http.Request) {
	h.commonHeadersSetUp(w)
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := h.parseID(idStr)
	if err != nil {
		log.Error("Error parsing UINT from ID: %s", err)
		http.Error(w, "Error parsing ID", http.StatusBadRequest)
		return
	}
	var bill bill.Bill
	err = bill.FromJSON(r.Body)
	if err != nil {
		log.Error("Error decoding JSON: %s", err)
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}
	bill, err = h.BillService.UpdateBill(id, bill)
	if err != nil {
		log.Error("Error adding bill: %s", err)
		http.Error(w, "Error adding bill", http.StatusInternalServerError)
		return
	}
	err = bill.ToJSON(w)
	if err != nil {
		log.Error("Error encoding JSON")
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteBill(w http.ResponseWriter, r *http.Request) {
	h.commonHeadersSetUp(w)
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := h.parseID(idStr)
	if err != nil {
		log.Error("Error parsing UINT from ID: %s", err)
		http.Error(w, "Error parsing ID", http.StatusBadRequest)
		return
	}
	err = h.BillService.DeleteBill(id)
	if err != nil {
		log.Error("Error deleting bill: %s", err)
		http.Error(w, "Error deleting bill", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
