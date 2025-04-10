package tushare

import "fmt"

// 期货相关接口

// FutBasic 期货日线数据
/**
exchange	str	Y	交易所代码 CFFEX-中金所 DCE-大商所 CZCE-郑商所 SHFE-上期所 INE-上海国际能源交易中心 GFEX-广州期货交易所
*/
func (api *TuShare) FutBasic(params map[string]string, fields []string) (*APIResponse, error) {
	//Check params
	_, hasTsCode := params["exchange"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}

	body := map[string]interface{}{
		"api_name": "fut_basic",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)

}

// FutDaily 期货日线数据
func (api *TuShare) FutDaily(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["trade_date"]
	if !hasTsCode {
		return nil, fmt.Errorf("trade_date is a required argument")
	}
	_, hasTsCode = params["ts_code"]
	if !hasTsCode {
		return nil, fmt.Errorf("ts_code is a required argument")
	}
	//params = map[string]string{
	//	"trade_date": "20250325",
	//	"symbol":     "EG2509",
	//}

	body := map[string]interface{}{
		"api_name": "fut_daily",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)

}

// fut_holding 每日持仓排名
func (api *TuShare) FutHolding(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "fut_holding",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)

}
