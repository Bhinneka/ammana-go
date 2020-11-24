package ammana

//ServiceResult
type ServiceResult struct {
	Result interface{}
	Error  error
}

type AmmanaService interface {
	// Get Authentication Method
	GetAuth() (DataAuth GenerateAccessTokenResponse, err error)
	// Get Card Detail Method
	GetCardDetail(request GetCardDetailRequest) ServiceResult
	// Get Order URL Redirect to Ammana Method
	GetOrderURL(request OrderRequest) ServiceResult
	// Get Detail Payment Method
	GetPaymentDetail(invoiceCode string) ServiceResult
	// Set Order to Settlement Method
	SetSettleOrder(invoiceCode string) ServiceResult
	// Set Order to Void Method
	SetVoidOrder(invoiceCode string) ServiceResult
}
