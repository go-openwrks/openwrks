package accounts

type (
	// ConsentStatus ...
	ConsentStatus string
)

const (
	// ConsentStatusAuthorised means the account is authorised to be queried.
	ConsentStatusAuthorised ConsentStatus = "Authorised"
	// ConsentStatusRevoked means the account is revoked by the bank.
	ConsentStatusRevoked ConsentStatus = "Revoked"
	// ConsentStatusExpired means the account has expired and needs reconnecting.
	ConsentStatusExpired ConsentStatus = "Expired"
)
