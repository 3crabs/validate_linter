package test_string

type A struct {
	a string `validate:"omitempty"`
	b string `validate:"gte=0"`
}
