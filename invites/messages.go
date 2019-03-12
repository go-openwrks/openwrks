package invites

type (
	// CreateInviteRequestMetadata data used to provide further details about your customer to help personalise their journey through Flow.
	//m The following meta data fields are viewable on your customer's record in the Surface Dashboard when provided via the /invite endpoint
	CreateInviteRequestMetadata struct {
		Firstname string `json:"firstname,omitempty"`
		Surname   string `json:"surname,omitempty"`
		Email     string `json:"email,omitempty"`
		// Business name will also be displayed on the customer's e-statement - available to download through the Surface Dashboard.
		BusinessName string `json:"businessName,omitempty"`
	}

	// CreateInviteRequest information about the user and our reference to that user.
	CreateInviteRequest struct {
		// A unique string you use to identify your customer
		CustomerReference string `json:"customerReference"`
		// The URL that you want to redirect the customer back to. The redirectUrl must match the redirectUrl registered for your chosen Flow configuration
		RedirectURL string `json:"redirectUrl"`
		// Key/value data used to provide further details about your customer to help personalise their journey through Flow.
		// The following meta data fields are viewable on your customer's record in the Surface Dashboard when provided via the /invite endpoint
		Metadata *CreateInviteRequestMetadata `json:"metaData"`
	}

	CreateInviteResponse struct {
		// A unique identifier for your invitation
		InvitationID string `json:"invitationId"`
		// The customer reference provided as part of the request
		CustomerReference string `json:"customerReference"`
		// A unique identifier for your user
		UserID string `json:"userId"`
		// A unique invite URL - provide this to your user so that they can consent to share their account and transaction details
		FlowURL string `json:"flowUrl"`
		// The date and time that the flowUrl will no longer be valid
		ExpirationDate string `json:"expirationDate"`
	}
)
