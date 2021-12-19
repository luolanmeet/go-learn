package mock

type RetrieverB struct {
	Contents string
}

func (r RetrieverB) Get(url string) string {
	return r.Contents
}

func (r RetrieverB) Post(url string) string {
	return r.Contents
}
