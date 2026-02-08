package handlers

import (
	"codewithumam-go-task1/internal/client"
	"codewithumam-go-task1/services"
	"log"
	"net/http"
	"time"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

// GetReportHariIni handles GET /api/report/hari-ini
func (h *ReportHandler) GetReportHariIni(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		client.WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	report, err := h.service.GetReportHariIni()
	if err != nil {
		log.Printf("ReportHandler GetReportHariIni: %v", err)
		client.WriteJSON(w, http.StatusInternalServerError, "Failed to get report", nil)
		return
	}

	client.WriteJSON(w, http.StatusOK, "Report fetched successfully", report)
}

// GetReport handles GET /api/report?start_date=2026-01-01&end_date=2026-02-01
func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		client.WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	startStr := r.URL.Query().Get("start_date")
	endStr := r.URL.Query().Get("end_date")

	if startStr == "" || endStr == "" {
		client.WriteJSON(w, http.StatusBadRequest, "start_date and end_date are required (format: YYYY-MM-DD)", nil)
		return
	}

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid start_date format, use YYYY-MM-DD", nil)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		client.WriteJSON(w, http.StatusBadRequest, "Invalid end_date format, use YYYY-MM-DD", nil)
		return
	}

	if end.Before(start) {
		client.WriteJSON(w, http.StatusBadRequest, "end_date must be after or equal to start_date", nil)
		return
	}

	report, err := h.service.GetReportByDateRange(start, end)
	if err != nil {
		log.Printf("ReportHandler GetReport: %v", err)
		client.WriteJSON(w, http.StatusInternalServerError, "Failed to get report", nil)
		return
	}

	msg := "Report for date range " + startStr + " to " + endStr + " fetched successfully"
	client.WriteJSON(w, http.StatusOK, msg, report)
}
