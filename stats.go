package nicehash

type GlobalStats struct {
	Algo                  uint8 `json:"algo"`
	ProfitabilityAboveBtc float32 `json:"profitability_above_btc,string"`
	ProfitabilityAboveLtc float32 `json:"profitability_above_ltc,string"`
	Price                 float64 `json:"price,string"`
	ProfitabilityBtc      float32 `json:"profitability_btc,string"`
	ProfitabilityLtc      float32 `json:"profitability_ltc,string"`
	Speed                 float64 `json:"speed,string"`
}

func (client *NicehashClient) GetStatsGlobalCurrent() ([]GlobalStats, error) {
	stats := &struct {
		Result struct {
			       Stats []GlobalStats `json:"stats"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"stats.global.current", Algo:AlgoTypeMAX, Location:LocationMAX}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Stats, err
	}

	return stats.Result.Stats, nil
}

func (client *NicehashClient) GetStatsGlobalDay() ([]GlobalStats, error) {
	stats := &struct {
		Result struct {
			       Stats []GlobalStats `json:"stats"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"stats.global.24h", Algo:AlgoTypeMAX, Location:LocationMAX}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Stats, err
	}

	return stats.Result.Stats, nil
}

type ProviderStats struct {
	Algo          uint8 `json:"algo"`
	Balance       float64 `json:"balance,string"`
	AcceptedSpeed float64 `json:"accepted_speed,string"`
	RejectedSpeed float64 `json:"rejected_speed,string"`
}

func (client *NicehashClient) GetStatsProvider(addr string) ([]ProviderStats, error) {
	stats := &struct {
		Result struct {
			       Stats []ProviderStats `json:"stats"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"stats.provider", Algo:AlgoTypeMAX, Location:LocationMAX, Addr:addr}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Stats, err
	}

	return stats.Result.Stats, nil
}

