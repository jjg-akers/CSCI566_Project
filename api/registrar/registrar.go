package registrar

// Registrar defines the methods necessary for a registrar
type Registrar interface{
	RegisterVoter([]byte) error
	CheckRegistration([]byte) (bool, error)
}