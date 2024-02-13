package report

import (
	"backend-comee/internal/models"
	"backend-comee/pkg/transactions/report"
	"github.com/jmoiron/sqlx"
)

type ServerReport struct {
	Event report.PortsServerReport
}

func NewServerReport(db *sqlx.DB, user *models.User, txID string) *ServerReport {
	repoReport := report.FactoryStorage(db, user, txID)
	return &ServerReport{
		Event: report.NewReportService(repoReport, user, txID),
	}
}
