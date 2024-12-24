
package main

import (
	"apiTest/verification"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

// تعریف ساختار برای ذخیره داده‌ها
type Reservation struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	NationalID  string `json:"national_id"`
	TicketCount int    `json:"ticket_count"`
}

var reservations = make([]Reservation, 0)

func saveReservation(res Reservation) {
	reservations = append(reservations, res)
}

func deleteReservation(nationalID string) {
	for index, res := range reservations {
		if res.NationalID == nationalID {
			// حذف رزرو با استفاده از slicing
			reservations = append(reservations[:index], reservations[index+1:]...)
			fmt.Println("Reservation deleted successfully.")
			return
		}
	}
	fmt.Println("Reservation not found.")
}

func updateReservation(nationalID string, updatedRes Reservation) bool {
	for index, res := range reservations {
		if res.NationalID == nationalID {
			// بروزرسانی رزرو
			reservations[index] = updatedRes
			fmt.Println("Reservation updated successfully.")
			return true
		}
	}
	return false
}

func reservationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var res Reservation

		err := json.NewDecoder(r.Body).Decode(&res)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		fmt.Println(res, reflect.TypeOf(res), reflect.TypeOf(res.FirstName), reflect.TypeOf(res.LastName), reflect.TypeOf(res.Email), reflect.TypeOf(res.NationalID), reflect.TypeOf(res.TicketCount))

		errors := verification.VerificationData(res.FirstName, res.LastName, res.Email, res.NationalID, res.TicketCount)
		if len(errors) > 0 {
			fmt.Println(errors)
			http.Error(w, fmt.Sprintf("Validation errors: %v", errors), http.StatusBadRequest)
			return
		}

		saveReservation(res)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Reservation successful"))

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reservations)

	case http.MethodDelete:

		nationalID := r.URL.Query().Get("national_id")
		fmt.Println(nationalID, r.URL)
		if nationalID == "" {
			http.Error(w, "Missing National ID", http.StatusBadRequest)
			return
		}
		deleteReservation(nationalID)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Reservation deleted"))

	case http.MethodPut:

		nationalID := r.URL.Query().Get("national_id")
		if nationalID == "" {
			http.Error(w, "Missing National ID", http.StatusBadRequest)
			return
		}

		var updatedRes Reservation
		err := json.NewDecoder(r.Body).Decode(&updatedRes)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if !updateReservation(nationalID, updatedRes) {
			http.Error(w, "Reservation not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Reservation updated"))
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please send data to '/reserve' for reservation.")
}

func main() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/reserve", reservationHandler)

	// راه‌اندازی سرور HTTP
	fmt.Println("Server is running on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
