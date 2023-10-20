package generator

// paths must end with a slash
type Generator struct {
	projectName string `validate:"required,notblank"`
	projectPath string `validate:"required,notblank,endswith=/"`

	backendName string `validate:"required,notblank"`
	backendPath string `validate:"required,notblank,endswith=/"`

	frontendName string `validate:"required,notblank"`
	frontendPath string `validate:"required,notblank,endswith=/"`

	git    bool
	dagger bool
	pulumi bool
	fly    bool
}

func New(options ...GeneratorOption) (*Generator, error) {
	g := &Generator{}

	for _, option := range options {
		option(g)
	}

	if err := g.validate(); err != nil {
		return nil, err
	}

	return g, nil
}
