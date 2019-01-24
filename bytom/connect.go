package bytom

//web struct
type Web struct {
	Provider ProviderInterface
	Bytom    *Bytom
}

func NewWeb(provider ProviderInterface) *Web {
	web := new(Web)
	web.Provider = provider
	web.Bytom = NewBytom(provider)
	return web
}

var connection = NewWeb(NewHTTPProvider(10))

//connection
func ReturnConnection() *Web {
	return connection
}
