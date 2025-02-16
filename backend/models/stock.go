package models

type Stock struct {
	Ticker        string  `json:"ticker"`
	MarketCap     float64 `json:"marketCap"`
	PE            float64 `json:"peRatio"`
	ROE           float64 `json:"roe"`
	DebtToEquity  float64 `json:"debtToEquity"`
	DividendYield float64 `json:"dividendYield"`
	RevenueGrowth float64 `json:"revenueGrowth"`
	EPSGrowth     float64 `json:"epsGrowth"`
	CurrentRatio  float64 `json:"currentRatio"`
	GrossMargin   float64 `json:"grossMargin"`
}
