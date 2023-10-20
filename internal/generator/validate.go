package generator

import "github.com/go-playground/validator/v10"

func (g *Generator) validate() error {
	v := validator.New(validator.WithRequiredStructEnabled())
	
	err := v.Struct(g)
	if err != nil {
		return err
	}

	return nil
}
