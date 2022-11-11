package validator

type Validator struct {
	Errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, messgae string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = messgae
	}
}

func (v *Validator) Check(ok bool, key, messgae string) {
	if !ok {
		v.AddError(key, messgae)
	}
}

// In returns true if a specific value in a list of strings.
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}
