package handlers

import (
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/xin2025/stock-screener/models"
)

type Condition struct {
	Field    string  `json:"field"`
	Operator string  `json:"operator"`
	Value    float64 `json:"value"`
}

type ScreenRequest struct {
	Conditions []Condition `json:"conditions"`
	SortBy     string      `json:"sortBy"`
	SortOrder  string      `json:"sortOrder"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
}

func ScreenHandler(c *fiber.Ctx, stocks []models.Stock) error {
	var req ScreenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON payload",
		})
	}

	filtered := filterStocks(stocks, req.Conditions)

	sortStocks(&filtered, req.SortBy, req.SortOrder)

	paginated, total := paginate(filtered, req.Page, req.Limit)

	return c.JSON(fiber.Map{
		"data":  paginated,
		"total": total,
	})
}
func filterStocks(allStocks []models.Stock, conditions []Condition) []models.Stock {
	var results []models.Stock

OuterLoop:
	for _, st := range allStocks {

		for _, cond := range conditions {
			if !matchCondition(st, cond) {

				continue OuterLoop
			}
		}
		results = append(results, st)
	}
	return results
}

func matchCondition(st models.Stock, cond Condition) bool {
	var fieldValue float64

	switch strings.ToLower(cond.Field) {
	case "marketcap":
		fieldValue = st.MarketCap
	case "peratio", "pe":
		fieldValue = st.PE
	case "roe":
		fieldValue = st.ROE
	case "debttoequity":
		fieldValue = st.DebtToEquity
	case "dividendyield":
		fieldValue = st.DividendYield
	case "revenuegrowth":
		fieldValue = st.RevenueGrowth
	case "epsgrowth":
		fieldValue = st.EPSGrowth
	case "currentratio":
		fieldValue = st.CurrentRatio
	case "grossmargin":
		fieldValue = st.GrossMargin
	default:

		return true
	}

	switch cond.Operator {
	case ">":
		return fieldValue > cond.Value
	case "<":
		return fieldValue < cond.Value
	case "=":
		return fieldValue == cond.Value
	default:
		return false
	}
}

func sortStocks(stocks *[]models.Stock, sortBy, sortOrder string) {
	sort.Slice(*stocks, func(i, j int) bool {
		s1 := (*stocks)[i]
		s2 := (*stocks)[j]

		var v1, v2 float64

		switch strings.ToLower(sortBy) {
		case "marketcap":
			v1, v2 = s1.MarketCap, s2.MarketCap
		case "pe", "peratio":
			v1, v2 = s1.PE, s2.PE
		case "roe":
			v1, v2 = s1.ROE, s2.ROE
		case "debttoequity", "debtequity", "debtequityratio":
			v1, v2 = s1.DebtToEquity, s2.DebtToEquity
		case "dividendyield":
			v1, v2 = s1.DividendYield, s2.DividendYield
		case "revenuegrowth":
			v1, v2 = s1.RevenueGrowth, s2.RevenueGrowth
		case "epsgrowth":
			v1, v2 = s1.EPSGrowth, s2.EPSGrowth
		case "currentratio":
			v1, v2 = s1.CurrentRatio, s2.CurrentRatio
		case "grossmargin":
			v1, v2 = s1.GrossMargin, s2.GrossMargin
		default:

			v1, v2 = s1.MarketCap, s2.MarketCap
		}

		if strings.ToLower(sortOrder) == "desc" {
			return v1 > v2
		}
		return v1 < v2
	})
}

func paginate(stocks []models.Stock, page, limit int) ([]models.Stock, int) {
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	start := (page - 1) * limit
	if start >= len(stocks) {
		return []models.Stock{}, len(stocks)
	}
	end := start + limit
	if end > len(stocks) {
		end = len(stocks)
	}
	return stocks[start:end], len(stocks)
}
