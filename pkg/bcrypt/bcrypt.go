package bcrypt

import "golang.org/x/crypto/bcrypt"

type Interface interface {
	GenerateFromPassword(password string)(string,error)
	CompareAndHashPasswrord(hashedPassword, password string) error
}

type bcryptImp struct{
	cost int
}

func Init() Interface {
	return &bcryptImp{cost: bcrypt.DefaultCost}
}

func (b *bcryptImp) GenerateFromPassword(password string) (string, error){
	passwordHash, err :=  bcrypt.GenerateFromPassword([]byte(password),b.cost)

	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func (b *bcryptImp) CompareAndHashPasswrord(hashedPassword, password string) error {
	err:= bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}
	return nil
}