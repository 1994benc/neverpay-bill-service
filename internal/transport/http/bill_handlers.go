package http

import (
	"1994benc/neverpay-api/internal/bill"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetBill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := h.parseID(idStr)
	if err != nil {
		log.Printf("Error parsing UINT from ID: %s", err)
	}
	bill, err := h.BillService.GetBill(uint(id))
	if err != nil {
		log.Printf("Error retreiving bill by ID: %s", err)
	}

	fmt.Fprintf(w, "%+v", bill)
}

func (h *Handler) GetAllBills(w http.ResponseWriter, r *http.Request) {
	bills, err := h.BillService.GetAllBills()
	if err != nil {
		log.Printf("Error retrieving bills: %s", err)
	}
	fmt.Fprintf(w, "%+v", bills)
}

func (h *Handler) AddBill(w http.ResponseWriter, r *http.Request) {
	bill, err := h.BillService.AddBill(bill.Bill{
		Payer: "Ben",
	})
	if err != nil {
		fmt.Printf("Error adding bill: %s", err)
	}
	fmt.Fprintf(w, "%+v", bill)
}

func (h *Handler) UpdateBill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := h.parseID(idStr)
	if err != nil {
		log.Printf("Error parsing UINT from ID: %s", err)
	}
	bill, err := h.BillService.UpdateBill(id, bill.Bill{
		Payer: "Ben",
	})
	if err != nil {
		fmt.Printf("Error adding bill: %s", err)
	}
	fmt.Fprintf(w, "%+v", bill)
}

func (h *Handler) DeleteBill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := h.parseID(idStr)
	if err != nil {
		log.Printf("Error parsing UINT from ID: %s", err)
	}
	err = h.BillService.DeleteBill(id)
	if err != nil {
		fmt.Printf("Error adding bill: %s", err)
	}
	fmt.Fprintf(w, "%+v", err)
}

func (h *Handler) parseID(idStr string) (uint, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	return uint(id), err
}
