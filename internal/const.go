package internal

import (
	"errors"
)

// Errors
var (
	ErrWrongOTP                    = errors.New("wrong one time password")
	ErrRoleMismatch                = errors.New("mismatch role")
	ErrInvalidCredentials          = errors.New("invalid email or password")
	ErrChangePassword              = errors.New("password is identical")
	ErrIDNumberAlreadyExists       = errors.New("patient with this id number already exists")
	ErrIdentifierAlreadyExists     = errors.New("the same identifier already exists")
	ErrSpecializationAlreadyExists = errors.New("this specialization already exists")
	ErrNotFound                    = errors.New("not found")
	ErrNoInput                     = errors.New("no input")
	ErrBlockedAccount              = errors.New("account is blocked")
	ErrPassword                    = errors.New("password not strong")
	ErrFailedChangePassword        = errors.New("failed to change password")
	ErrEmailVerified               = errors.New("email is verified already")
	ErrNotEmailVerified            = errors.New("email is not verified")
	ErrSpecializationNotFound      = errors.New("specialization does not exist")
	ErrSpecializationIsUsed        = errors.New("specialization is used by doctor")
	ErrAccountDeleted              = errors.New("account has been deleted")
	ErrTokenExpired                = errors.New("token is expired")
	ErrInvalidDayCode              = errors.New("day code is invalid")
	ErrDayCodeAlreadyExists        = errors.New("day code with this language is already exist")
	ErrLanguageCodeAlreadyExists   = errors.New("parameter with this code is already exist")
	ErrNotAuthorized               = errors.New("not authorized")
	ErrAddressOneIsRequired        = errors.New("failed to delete address. one address is required")
	ErrInvalidResponse             = errors.New("invalid response")
	ErrInvalidRequest              = errors.New("invalid_request")
)

type ParameterType string

const (
	AllParameterType       ParameterType = "all"
	DoctorParameterType    ParameterType = "doctor"
	MerchantParameterType  ParameterType = "merchant"
	MedicplusParameterType ParameterType = "medicplus"
)

type ReceiverType string

const (
	ToDoctor    ReceiverType = "doctor"
	ToMerchant  ReceiverType = "merchant"
	ToMedicplus ReceiverType = "medicplus"
)

type ReportContext string

const (
	Consultation ReportContext = "consultation"
	Purchase     ReportContext = "purchase"
)

// session types
const (
	LoginSessionType         = "login"
	OTPSessionType           = "otp"
	ResetPasswordSessionType = "reset_password"
)

// role types
type RoleTypes string

const (
	Patient RoleTypes = "patient"
	Doctor  RoleTypes = "doctor"
	Admin   RoleTypes = "admin"
)

// gender types
type GenderTypes string

const (
	Male   GenderTypes = "male"
	Female GenderTypes = "female"
)

// Language Code
type LanguageCode string

const (
	BahasaIndonesia LanguageCode = "id"
	English         LanguageCode = "en"
)
