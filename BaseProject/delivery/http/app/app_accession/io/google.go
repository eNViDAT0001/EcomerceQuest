package io

type GoogleProfile struct {
	Email      string `json:"email"`
	Verified   bool   `json:"email_verified"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Locale     string `json:"locale"`
	Name       string `json:"name"`
	Picture    string `json:"picture"`
	Sub        string `json:"sub"`
}
