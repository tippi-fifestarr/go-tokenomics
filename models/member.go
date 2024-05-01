package models

// Member represents a user with tokens and an immutable wallet address.
type Member struct {
	Username      string
	Password      string       // In the real application, this would be a hash. (or better yet authentication from a wallet provider)
	WalletAddress string       // In the real application, this would be a public key.
	CeptorWallet  CeptorWallet // Tokens owned by the member.
}

// CeptorWallet represents the tokens owned by a Member.
type CeptorWallet struct {
	GamesXP     int
	ArtXP       int
	TechXP      int
	ArtTokens   int
	GamesTokens int
	TechTokens  int
}
