package tushare

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFutBasic(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	params = map[string]string{
		"exchange": "CFFEX",
	}
	// Check params
	resp, err := client.FutBasic(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestFutDaily(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	resp, err := client.FutDaily(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestFutHolding(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	params = map[string]string{
		"trade_date": "20250325",
		"symbol":     "EG2509",
	}
	fields = []string{
		"trade_date",
		"symbol",
		"broker",
		"vol",
		"vol_chg",
		"long_hld",
		"long_chg",
		"short_hld",
		"short_chg",
	}

	// Check params
	resp, err := client.FutHolding(params, fields)

	if err != nil {
		if resp.Code == -2002 {
			ast.Equal(err.Error(), "Your point is not enough to use this api")
		}
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}
