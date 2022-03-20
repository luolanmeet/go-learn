package hello

type HelloService struct {
}

func (p *HelloService) Hello(name string, reply *string) error {
	*reply = "hello " + name
	return nil
}
