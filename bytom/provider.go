package bytom

type ProviderInterface interface {
	SendRequest(v interface{}, params interface{}) error
	Close() error
}
