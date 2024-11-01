package domain

type TokenClaim struct {
	Alg  string `json:"-"`
	User struct {
		ID string `json:"id"`
	}
}
