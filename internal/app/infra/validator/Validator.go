package validator

type Validator interface {
	ValidateStruct(interface{}) error
}

func NewValidator() Validator {
	return &validator{}
}

type validator struct{}

func (*validator) ValidateStruct(interface{}) error {
	return nil
}
