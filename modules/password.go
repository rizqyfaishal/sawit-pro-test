package modules

type PasswordAuthInterface interface {
	GenerateHashedPassword(password string) (string, error)
	CompareHashedPassword(hashedPassword, password string) (bool, error)
}
