package helpers

type HelperInterface interface {
	GenerateToken(catId int) (string, error)
	ValidateJWT(tokenString string) (*Claims, error)
}
