package router

type Router interface{}

type nspRouter struct {
}

func NewRouter() Router {
	return &nspRouter{}
}
