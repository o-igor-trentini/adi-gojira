package gjmodels

type User struct {
	AccountID    string  `json:"accountId"`
	Self         string  `json:"self"`
	EmailAddress string  `json:"emailAddress"`
	AvatarURLs   Avatars `json:"avatarUrls"`
	DisplayName  string  `json:"displayName"`
	Active       bool    `json:"active"`
	TimeZone     string  `json:"timeZone"`
	AccountType  string  `json:"accountType"`
}
