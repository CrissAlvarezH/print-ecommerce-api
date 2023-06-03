package services

import (
	"github.com/CrissAlvarezH/print-ecommerce-api/internal/users/application/ports"
	users "github.com/CrissAlvarezH/print-ecommerce-api/internal/users/domain"
	"math/rand"
	"strconv"
)

type UserService struct {
	repo                    ports.UserRepository
	addressRepo             ports.AddressRepository
	verificationCodeManager ports.VerificationCodeManager
	passwordManager         ports.PasswordManager
}

func (s *UserService) List(
	filters map[string]string, scopes []users.ScopeName, limit int64, offset int64,
) ([]users.User, int64, error) {
	return s.repo.List(filters, scopes, limit, offset)
}

func (s *UserService) GetByID(ID users.UserID) (users.User, error) {
	return s.repo.GetByID(ID)
}

func (s *UserService) Add(
	name string, email string, password string, phone string, isActive bool, scopes []users.ScopeName,
) (users.User, error) {
	hashPassword, err := s.passwordManager.Encrypt(password)
	if err != nil {
		return users.User{}, err
	}
	return s.repo.Add(name, email, hashPassword, phone, isActive, scopes)
}

func (s *UserService) Update(
	ID users.UserID, name string, email string, phone string, isActive bool, scopes []users.ScopeName,
) (users.User, error) {
	return s.repo.Update(ID, name, email, phone, isActive, scopes)
}

func (s *UserService) Deactivate(ID users.UserID) error {
	return s.repo.Deactivate(ID)
}

func (s *UserService) AddAddress(
	ID users.UserID, department string, city string, address string, receiverPhone string,
	receiverName string,
) (users.Address, error) {
	createdAddress, err := s.addressRepo.Add(department, city, address, receiverPhone, receiverName)
	if err != nil {
		return users.Address{}, err
	}

	err = s.repo.AttachAddress(ID, createdAddress.ID)
	if err != nil {
		return users.Address{}, err
	}

	return createdAddress, nil
}

func (s *UserService) DeleteAddress(ID users.UserID, addressID users.AddressID) error {
	err := s.repo.DetachAddress(ID, addressID)
	if err != nil {
		return err
	}
	err = s.addressRepo.Delete(addressID)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateAddress(
	addressID users.AddressID, department string, city string, address string,
	receiverPhone string, receiverName string,
) (users.Address, error) {
	return s.addressRepo.Update(addressID, department, city, address, receiverPhone, receiverName)
}

func (s *UserService) SendVerificationCode(ID users.UserID) error {
	codeRangeMin := 1000
	codeRangeMax := 9999
	code := strconv.Itoa(rand.Intn(codeRangeMax-codeRangeMin) + codeRangeMin)
	err := s.verificationCodeManager.Send(code, ports.MessageProviderEmail)
	if err != nil {
		return err
	}

	err = s.repo.SaveVerificationCode(ID, code)
	return err
}

func (s *UserService) ValidateVerificationCode(ID users.UserID, code string) (bool, error) {
	return s.repo.ValidateVerificationCode(ID, code)
}
