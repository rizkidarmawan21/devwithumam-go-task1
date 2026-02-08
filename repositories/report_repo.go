package repositories

import (
	"codewithumam-go-task1/handlers/dto/response"
	"database/sql"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

// GetReportByDateRange returns total revenue, total transactions count,
// and best-selling product (by quantity) within the given date range.
// Both start and end are inclusive; end is treated as end of day (23:59:59).
func (r *ReportRepository) GetReportByDateRange(start, end time.Time) (*response.ReportResponse, error) {
	endOfDay := end.Truncate(24*time.Hour).Add(24*time.Hour - time.Nanosecond)

	var totalRevenue, totalTransaksi sql.NullInt64
	err := r.db.QueryRow(`
		SELECT
			COALESCE(SUM(total_amount), 0),
			COUNT(*)
		FROM transactions
		WHERE created_at >= $1 AND created_at <= $2
	`, start, endOfDay).Scan(&totalRevenue, &totalTransaksi)
	if err != nil {
		return nil, err
	}

	resp := &response.ReportResponse{
		TotalRevenue:   int(totalRevenue.Int64),
		TotalTransaksi: int(totalTransaksi.Int64),
		ProdukTerlaris: response.ProdukTerlaris{Nama: "", QtyTerjual: 0},
	}

	var productName sql.NullString
	var qtyTerjual sql.NullInt64
	err = r.db.QueryRow(`
		SELECT td.product_name, SUM(td.quantity) AS qty_terjual
		FROM transaction_details td
		INNER JOIN transactions t ON t.id = td.transaction_id
		WHERE t.created_at >= $1 AND t.created_at <= $2
		GROUP BY td.product_id, td.product_name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`, start, endOfDay).Scan(&productName, &qtyTerjual)
	if err == nil && productName.Valid {
		resp.ProdukTerlaris = response.ProdukTerlaris{
			Nama:       productName.String,
			QtyTerjual: int(qtyTerjual.Int64),
		}
	}
	// If no rows (err == sql.ErrNoRows) or no details, resp.ProdukTerlaris stays default

	return resp, nil
}
