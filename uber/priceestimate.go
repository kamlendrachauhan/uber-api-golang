package uber

// List of price estimates
type PriceEstimates struct {
	Prices []PriceEstimate `json:"prices"`
}

// Uber price estimate
type PriceEstimate struct {
	ProductId       string  `json:"product_id"`
	CurrencyCode    string  `json:"currency_code"`
	DisplayName     string  `json:"display_name"`
	Estimate        string  `json:"estimate"`
	LowEstimate     int     `json:"low_estimate"`
	HighEstimate    int     `json:"high_estimate"`
	SurgeMultiplier float32 `json:"surge_multiplier"`
	Duration        int     `json:"duration"`
	Distance        float32 `json:"distance"`
}
