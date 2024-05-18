package handler

import (
	"errors"
	"health-record/model/dto"
	"health-record/src/usecase"
	"log"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/relvacode/iso8601"
)

type PatientHandler struct {
	iPatientUsecase usecase.PatientUsecaseInterface
}

func NewPatientHandler(iPatientUsecase usecase.PatientUsecaseInterface) PatientHandlerInterface {
	return &PatientHandler{iPatientUsecase}
}

func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var request dto.RequestCreatePatient
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("Register patient bad request >> ShouldBindJSON")
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	err = ValidateRegisterPatientRequest(request)

	if err != nil {
		log.Println("Register patient bad request >> ValidateRegisterPatientRequest")
		c.JSON(400, gin.H{"status": "bad request", "message": err.Error()})
		return
	}

	//error 409
	exist, err := h.iPatientUsecase.GetPatientByIdentityNumber(request.IdentityNumber)
	if(exist) {
		log.Println("Register patient bad request >> GetPatientByIdentityNumber")
		c.JSON(409, gin.H{"status": "bad request", "message": "IdentityNumber already registered"})
		return
	}

	//create
	err = h.iPatientUsecase.RegisterPatient(request)
	
	//error 500
	if err != nil {
		log.Println("Register patient error ", err)
		c.JSON(500, gin.H{"status": "internal server error", "message": err.Error()})
		return
	}

	//success 201
	c.JSON(201, gin.H{
		"message": "Medical patient successfully added",
	})
}

func ValidateRegisterPatientRequest(request dto.RequestCreatePatient) error {
	if !is16DigitInteger(request.IdentityNumber) {
		return errors.New("invalid IdentityNumber")
	}

	if !validatePhoneNumber(request.PhoneNumber) {
		return errors.New("invalid PhoneNumber")
	}

	if !validateName(request.Name) {
		return errors.New("invalid Name")
	}

	if !validateDateFormat(request.BirthDate) {
		return errors.New("invalid BirthDate")
	}

	if !validateGender(request.Gender) {
		return errors.New("invalid Gender")
	}

	if !isValidURL(request.IdentityCardScanImg) {
		return errors.New("invalid IdentityCardScanImg")
	}

	return nil
}

func is16DigitInteger(num int) bool {
    if num >= 1000000000000000 && num <= 9999999999999999 {
        return true
    }
    return false
}

func validatePhoneNumber(phone string) bool {
    if !strings.HasPrefix(phone, "+62") {
        return false
    }

    length := len(phone)
    if length < 10 || length > 15 {
        return false
    }

    return true
}

func validateName(name string) bool {
    length := len(name)
    if length < 3 || length > 30 {
        return false
    }

    return true
}

func validateDateFormat(date string) bool {
	_, err := iso8601.Parse([]byte(date))

	return err == nil
}

func validateGender(gender string) bool {
	return gender == "male" || gender == "female"
}

func isValidURL(str string) bool {
	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || (u.Scheme != "http" && u.Scheme != "https") {
		return false
	}
	return true
}