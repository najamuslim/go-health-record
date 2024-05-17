package handler

import (
	"errors"
	"health-record/model/dto"
	"health-record/src/usecase"
	"log"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	iAuthUsecase usecase.AuthUsecaseInterface
}

func NewAuthHandler(iAuthUsecase usecase.AuthUsecaseInterface) AuthHandlerInterface {
	return &AuthHandler{iAuthUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request dto.RequestCreateUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Register bad request")
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	// Validate request payload
	err = ValidateRegisterRequest(request.Nip, request.Name, request.Password)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	// Check if email already exists
	exists, _ := h.iAuthUsecase.GetUserByNIP(request.Nip)
	if exists {
		log.Println("Register bad request ", err)
		c.JSON(409, gin.H{"status": "bad request", "message": "nip already exists"})
		return
	}

	token, userId, err := h.iAuthUsecase.Register(request)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	log.Println("Register successful")
	c.JSON(201, gin.H{
    "message": "User registered successfully",
    "data": gin.H{
			"userId": userId,
			"nip": request.Nip, 
			"name": request.Name, 
      "accessToken": token,
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.RequestAuth
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Login bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	err = ValidateLoginRequest(request.Nip, request.Password)
	if err != nil {
		log.Println("Login bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	token, userData, err := h.iAuthUsecase.Login(request)
	if err != nil {
		log.Println("Login bad request ", err)
		if err.Error() == "user not found" {
			c.JSON(404, gin.H{"status": "bad request", "message": "user not found"})
			return
		} 
		if err.Error() == "wrong password" {
			c.JSON(400, gin.H{"status": "bad request", "message": "wrong password"})
			return
		}
	}

	log.Println("Login successful")
	c.JSON(200, gin.H{
    "message": "User logged successfully",
    "data": gin.H{
			"nip": userData.Nip, 
			"name": userData.Name, 
      "accessToken": token,
		},
	})
}

// ValidateRegisterRequest validates the register user request payload
func ValidateRegisterRequest(nip int64, name, password string) error {
	// Validate email format
	if !isValidNip(nip) {
		return errors.New("nip must be in valid email format")
	}

	// Validate name length
	if len(name) < 5 || len(name) > 50 {
		return errors.New("name length must be between 5 and 50 characters")
	}

	// Validate password length
	if len(password) < 5 || len(password) > 15 {
		return errors.New("password length must be between 5 and 15 characters")
	}

	return nil
}

func ValidateLoginRequest(nip int64, password string) error {
	// Validate email format
	if !isValidNip(nip) {
		return errors.New("nip must be in valid email format")
	}

	// Validate password length
	if len(password) < 5 || len(password) > 15 {
		return errors.New("password length must be between 5 and 15 characters")
	}

	return nil
}

// Helper function to validate email format
// TODO fix to the correct nip validation
func isValidNip(nip int64) bool {
	// Regular expression pattern for email format
	nipRegex := `^615[12][2-9][0-9]{3}(0[1-9]|1[0-2])[0-9]{3}$`
	match, _ := regexp.MatchString(nipRegex, strconv.FormatInt(nip, 10))
	return match
}
