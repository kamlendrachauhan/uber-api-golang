package uber

import (
	"encoding/json"
	"strconv"
)

// List of time estimates
type TimeEstimates struct {
	StartLatitude  float64
	StartLongitude float64
	Times          []TimeEstimate `json:"times"`
}

// Uber time estimate
type TimeEstimate struct {
	ProductId   string `json:"product_id"`
	DisplayName string `json:"display_name"`
	Estimate    int    `json:"estimate"`
}

func convertToMins(estimate int) int {
	return estimate / 60
}

// Internal method that implements the Getter interface
func (te *TimeEstimates) get(c *Client) error {
	timeEstimateParams := map[string]string{
		"start_latitude":  strconv.FormatFloat(te.StartLatitude, 'f', 2, 32),
		"start_longitude": strconv.FormatFloat(te.StartLongitude, 'f', 2, 32),
	}

	data := c.getRequest("estimates/time", timeEstimateParams)
	if e := json.Unmarshal(data, &te); e != nil {
		return e
	}

	return nil
}
