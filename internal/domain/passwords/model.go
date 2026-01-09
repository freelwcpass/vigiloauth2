package passwords

type PasswordPolicyResponse struct {
	RequireUpper  bool `json:"require_upper"`
	RequireNumber bool `json:"require_number"`
	RequireSymbol bool `json:"require_symbol"`
	MinLength     int  `json:"min_length"`
}