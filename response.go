package ammana

type SharedMeta struct {
	Hostname string `json:"hostname"`
	ClientIP string `json:"client_ip"`
}

type GenerateAccessTokenResponse struct {
	Data struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	} `json:"data"`
	Meta    SharedMeta `json:"meta"`
	Message string     `json:"message"`
	Type    string     `json:"type"`
}

type OrderResponse struct {
	Data struct {
		Url       string `json:"url"`
		ExpiresIn int `json:"expires_in"`
	} `json:"data"`
	Meta    SharedMeta `json:"meta"`
	Message string     `json:"message"`
	Type    string     `json:"type"`
}

type GetCardDetailResponse struct {
	Data struct {
		ID              string `json:"id"`
		Status          string `json:"status"`
		LimitAmount     int    `json:"limit_amount"`
		UsageAmount     int    `json:"usage_amount"`
		AvailableAmount int    `json:"available_amount"`
	} `json:"data"`
	Meta    SharedMeta `json:"meta"`
	Message string     `json:"message"`
	Type    string     `json:"type"`
}

type GetPaymentDetailResponse struct {
	Data struct {
		TransactionId string `json:"transaction_id"`
		Status        string `json:"status"`
		InvoiceCode   string `json:"invoice_code"`
		Total         int    `json:"total"`
		CreatedAt     string `json:"created_at"`
		PaidAt        string `json:"paid_at"`
	} `json:"data"`
	Meta    SharedMeta `json:"meta"`
	Message string     `json:"message"`
	Type    string     `json:"type"`
}

type SetVoidPaymentResponse struct {
	Data struct {
		TransactionId string `json:"transaction_id"`
		Status        string `json:"status"`
		InvoiceCode   string `json:"invoice_code"`
		Total         int    `json:"total"`
		CreatedAt     string `json:"created_at"`
		VoidAt        string `json:"void_at"`
	} `json:"data"`
	Meta    SharedMeta `json:"meta"`
	Message string     `json:"message"`
	Type    string     `json:"type"`
}
type SetSettlePaymentResponse struct {
	Data struct {
		TransactionId string `json:"transaction_id"`
		Status        string `json:"status"`
		InvoiceCode   string `json:"invoice_code"`
		Total         int    `json:"total"`
		CreatedAt     string `json:"created_at"`
		PaidAt        string `json:"paid_at"`
		SettledAt     string `json:"settled_at"`
	} `json:"data"`
	Meta    SharedMeta `json:"meta"`
	Message string     `json:"message"`
	Type    string     `json:"type"`
}
