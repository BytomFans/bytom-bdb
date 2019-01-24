package bytom

// The Block Modle
type Bytom struct {
	provider ProviderInterface
}

// NewBytom
func NewBytom(provider ProviderInterface) *Bytom {
	bytom := new(Bytom)
	bytom.provider = provider
	return bytom
}

//create-key
func (btm *Bytom) CreateKey(alias, password string) (string, error) {

	params := make([]string, 2)
	params[0] = alias
	params[1] = password

	pointer := &RequestResult{}

	err := btm.provider.SendRequest(pointer, params)

	if err != nil {
		return "", err
	}

	return pointer.resultToString()

}

//create-account
func (btm *Bytom) CreateAccount() (string, error) {
	pointer := &RequestResult{}

	err := btm.provider.SendRequest(pointer, nil)

	if err != nil {
		return "", err
	}

	return pointer.resultToString()
}

//build-transaction
func (btm *Bytom) BuildTransaction() (string, error) {
	pointer := &RequestResult{}

	err := btm.provider.SendRequest(pointer, nil)

	if err != nil {
		return "", err
	}

	return pointer.resultToString()

}
