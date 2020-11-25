package ammana

type AmmanaService interface {
	// Get Authentication Method
	GetAuth() (DataAuth GenerateAccessTokenResponse, err error)
	// Get Card Detail Method
	GetCardDetail(request GetCardDetailRequest) (resp GetCardDetailResponse, err error)
	// Get Order URL Redirect to Ammana Method
	GetOrderURL(request OrderRequest) (resp OrderResponse, err error)
	// Get Detail Payment Method
	GetPaymentDetail(invoiceCode string) (resp GetPaymentDetailResponse, err error)
	// Set Order to Settlement Method
	SetSettleOrder(invoiceCode string) (resp SetSettlePaymentResponse, err error)
	// Set Order to Void Method
	SetVoidOrder(invoiceCode string) (resp SetVoidPaymentResponse, err error)
}
