package webhooks

type (
	Webhook struct {
		ID           string `json:"id"`
		Event        string `json:"event"`
		TargetURL    string `json:"targetUrl"`
		MajorVersion string `json:"majorVersion"`
		MinorVersion string `json:"minorVersion"`
		CreatedDate  string `json:"createdDate"`
		Status       string `json:"status"`
	}

	ListWebhooksResponse struct {
		Data []Webhook `json:"data"`
	}

	CreateWebhookSubscriptionResponse struct {
		ID string `json:"id"`
		// A unique string used to verify event signatures
		Secret string `json:"secret"`
	}
)
