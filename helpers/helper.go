package helpers

type HelperInterface interface {
	GenerateToken(userId string) (string, error)
	ValidateJWT(tokenString string) (*Claims, error)
}
