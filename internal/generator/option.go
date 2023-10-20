package generator

type GeneratorOption func(*Generator)

func WithProjectName(name string) GeneratorOption {
	return func(g *Generator) {
		g.projectName = name
	}
}

func WithProjectPath(path string) GeneratorOption {
	return func(g *Generator) {
		g.projectPath = path
	}
}

func WithBackendName(name string) GeneratorOption {
	return func(g *Generator) {
		g.backendName = name
	}
}

func WithBackendPath(path string) GeneratorOption {
	return func(g *Generator) {
		g.backendPath = path
	}
}

func WithFrontendName(name string) GeneratorOption {
	return func(g *Generator) {
		g.frontendName = name
	}
}

func WithFrontendPath(path string) GeneratorOption {
	return func(g *Generator) {
		g.frontendPath = path
	}
}

func WithGit() GeneratorOption {
	return func(g *Generator) {
		g.git = true
	}
}

func WithDagger() GeneratorOption {
	return func(g *Generator) {
		g.dagger = true
	}
}

func WithPulumi() GeneratorOption {
	return func(g *Generator) {
		g.pulumi = true
	}
}

func WithFly() GeneratorOption {
	return func(g *Generator) {
		g.fly = true
	}
}
