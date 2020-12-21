package api

// ConnManager is
type ConnManager interface {
	Connection() error
	GetToken() error
}
