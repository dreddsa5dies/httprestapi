package apiserver

// APIserver ...
type APIserver struct{}

// New ...
func New() *APIserver {
	return &APIserver{}
}

// Start ...
func (s *APIserver) Start() error {
	return nil
}
