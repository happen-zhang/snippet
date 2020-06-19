package kit

type Service interface {
	Say() string
}

type service struct{}

func (svc *service) Say() string {
	return "helloworld!"
}

func NewService() Service {
	return &service{}
}
