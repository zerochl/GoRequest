package GoRequest

type InitRequestEntity struct {
	BaseUrl          string `json:"base_url"`
	HeaderJson       string `json:"header_json"`
	PoolMaxIdel      int    `json:"pool_max_idel"`
	PoolCore         int    `json:"pool_core"`
	RequestTimeOut   int    `json:"request_time_out"`
	RequestKeepAlive int    `json:"request_keep_alive"`
}
