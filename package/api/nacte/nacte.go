package nacte

func GetPaymentAPIKey(paymentRef string) string {
	if apiKey, exists := paymentAPIKeys[paymentRef]; exists {
		return apiKey
	}
	return ""
}
