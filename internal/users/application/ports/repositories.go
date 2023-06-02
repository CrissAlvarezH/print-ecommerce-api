package ports

import users "github.com/CrissAlvarezH/print-ecommerce-api/internal/users/domain"

type UserRepository interface {
	List(filters map[string]string, scopes []users.ScopeName, limit int64, offset int64) ([]users.User, int64, error)
	GetByID(ID users.UserID) (users.User, error)
	Add(name string, email string, phone string, isActive bool, scopes []users.ScopeName) (users.User, error)
	Update(
		ID users.UserID, name string, email string, phone string, isActive bool, scopes []users.ScopeName,
	) (users.User, error)
	Deactivate(ID users.UserID) error

	ListAddress(ID users.UserID) (users.Address, error)
	AttachAddress(ID users.UserID, addressID users.AddressID) error
	DetachAddress(ID users.UserID, addressID users.AddressID) error

	SaveVerificationCode(ID users.UserID, code string) error
	ValidateVerificationCode(ID users.UserID, code string) (bool, error)
}

type AddressRepository interface {
	Add(
		department string, city string, address string, receiverPhone string, receiverName string,
	) (users.Address, error)
	Update(
		ID users.AddressID, department string, city string, address string,
		receiverPhone string, receiverName string,
	) (users.Address, error)
	Delete(ID users.AddressID) error
}
