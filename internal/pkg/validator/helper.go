package validator

func ValidateStruct(data interface{}) (validatorErr error) {
	validate := GetValidator()
	err := validate.Struct(data)

	return err
}
