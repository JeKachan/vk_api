package vkapi

type Params struct {
	methodName  string
	parameters  string
	accessToken string
	version     string
}

func NewParams() *Params {
	model := new(Params)
	model.version = "5.52"
	return model
}
