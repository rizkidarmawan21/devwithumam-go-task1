package services

import (
	"codewithumam-go-task1/handlers/dto/response"
	"codewithumam-go-task1/repositories"
	"time"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReportByDateRange(start, end time.Time) (*response.ReportResponse, error) {
	return s.repo.GetReportByDateRange(start, end)
}

func (s *ReportService) GetReportHariIni() (*response.ReportResponse, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24*time.Hour - time.Nanosecond)
	return s.repo.GetReportByDateRange(startOfDay, endOfDay)
}
