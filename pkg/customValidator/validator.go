package customValidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewCustomValidator() (*CustomValidator, error) {
	customValidator := validator.New()
	_ = customValidator.RegisterValidation("Urls", cValidatorUrls)
	_ = customValidator.RegisterValidation("Date", cValidatorDate)
	_ = customValidator.RegisterValidation("DateTime", cValidatorDateTime)
	_ = customValidator.RegisterValidation("AlphaNumerics", cValidatorAlphaNumerics)
	_ = customValidator.RegisterValidation("DecimalNumbers", cValidatorDecimalNumbers)
	_ = customValidator.RegisterValidation("EmailAddress", cValidatorEmailAddress)
	_ = customValidator.RegisterValidation("PhoneNumber", cValidatorPhoneNumber)
	_ = customValidator.RegisterValidation("PersonalAddress", cValidatorPersonalAddress)
	_ = customValidator.RegisterValidation("AlphabetOnly", cValidatorAlphabetOnly)
	return &CustomValidator{validator: customValidator}, nil
}

// cValidatorUrls validates if a string is a valid URL
func cValidatorUrls(fl validator.FieldLevel) bool {
	urlRegex := `^(http|https):\/\/[^\s$.?#].[^\s]*$`
	regex := regexp.MustCompile(urlRegex)
	return regex.MatchString(fl.Field().String())
}

// cValidatorDate validates if a string is a valid date (format: YYYY-MM-DD)
func cValidatorDate(fl validator.FieldLevel) bool {
	dateRegex := `^\d{4}-\d{2}-\d{2}$`
	regex := regexp.MustCompile(dateRegex)
	return regex.MatchString(fl.Field().String())
}

// cValidatorDateTime validates if a string is a valid date-time (format: YYYY-MM-DDTHH:MM:SS)
func cValidatorDateTime(fl validator.FieldLevel) bool {
	dateTimeRegex := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`
	regex := regexp.MustCompile(dateTimeRegex)
	return regex.MatchString(fl.Field().String())
}

// cValidatorAlphaNumerics validates if a string contains only alphanumeric characters
func cValidatorAlphaNumerics(fl validator.FieldLevel) bool {
	alphaNumericRegex := `^[a-zA-Z0-9]+$`
	regex := regexp.MustCompile(alphaNumericRegex)
	return regex.MatchString(fl.Field().String())
}

// cValidatorDecimalNumbers validates if a string is a valid decimal number
func cValidatorDecimalNumbers(fl validator.FieldLevel) bool {
	decimalNumberRegex := `^\d+(\.\d+)?$`
	regex := regexp.MustCompile(decimalNumberRegex)
	return regex.MatchString(fl.Field().String())
}

// cValidatorEmailAddress validates if a string is a valid email address
func cValidatorEmailAddress(fl validator.FieldLevel) bool {
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	regex := regexp.MustCompile(emailRegex)
	return regex.MatchString(fl.Field().String())
}

// cValidatorPhoneNumber validates if a string is a valid phone number
func cValidatorPhoneNumber(fl validator.FieldLevel) bool {
	phoneNumberRegex := `^\+?[1-9]\d{1,14}$`
	regex := regexp.MustCompile(phoneNumberRegex)
	return regex.MatchString(fl.Field().String())
}

// cValidatorPersonalAddress validates if a string is a valid personal address (example regex, adjust as needed)
func cValidatorPersonalAddress(fl validator.FieldLevel) bool {
	addressRegex := `^[a-zA-Z0-9\s,.'-]{3,}$`
	regex := regexp.MustCompile(addressRegex)
	return regex.MatchString(fl.Field().String())
}

// cValidatorAlphabetOnly validates if a string contains only alphabetic characters
func cValidatorAlphabetOnly(fl validator.FieldLevel) bool {
	alphabetOnlyRegex := `^[a-zA-Z]+$`
	regex := regexp.MustCompile(alphabetOnlyRegex)
	return regex.MatchString(fl.Field().String())
}
