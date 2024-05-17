package handler

import (
	"errors"
	"health-record/model/dto"
	"health-record/src/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type NurseHandler struct {
	iNurseUsecase usecase.NurseUsecaseInterface
}

func NewNurseHandler(iNurseUsecase usecase.NurseUsecaseInterface) NurseHandlerInterface {
	return &NurseHandler{iNurseUsecase}
}

func (h *NurseHandler) RegisterNurse(c *gin.Context) {
	var request dto.RequestCreateNurse
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Register bad request")
		c.JSON(400, gin.H{"status": "bad request", "message": err})
		return
	}

	// Validate request payload
	err = ValidateRegisterNurseRequest(request.Nip, request.Name)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	// Check if email already exists
	exists, _ := h.iNurseUsecase.GetNurseByNIP(request.Nip)
	if exists {
		log.Println("Register bad request ", err)
		c.JSON(409, gin.H{"status": "bad request", "message": "nip already exists"})
		return
	}

	userId, err := h.iNurseUsecase.RegisterNurse(request)
	if err != nil {
		log.Println("Register bad request ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err})
		return
	}

	log.Println("Register successful")
	c.JSON(201, gin.H{
    "message": "Nurse registered successfully",
    "data": gin.H{
			"userId": userId,
			"nip": request.Nip, 
			"name": request.Name,
		},
	})
}

func ValidateRegisterNurseRequest(nip int64, name string) error {
	// Validate email format
	if !isValidNip(nip) {
		return errors.New("nip must be in valid email format")
	}

	// Validate name length
	if len(name) < 5 || len(name) > 50 {
		return errors.New("name length must be between 5 and 50 characters")
	}

	return nil
}

func ValidateLoginNurseRequest(nip int64, password string) error {
	// Validate email format
	if !isValidNip(nip) {
		return errors.New("email must be in valid email format")
	}

	// Validate password length
	if len(password) < 5 || len(password) > 15 {
		return errors.New("password length must be between 5 and 15 characters")
	}

	return nil
}
