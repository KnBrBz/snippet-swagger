package setup

type S struct {
}

func New() *S {
	return &S{}
}

func (*S) RestAllowedOrigins() string {
	return "*"
}

func (*S) Host() string {
	return "localhost:8300"
}
