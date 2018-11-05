package rates

import (
	"github.com/nmarsollier/authgo/security"
	"github.com/nmarsollier/authgo/tools/db"
	"github.com/nmarsollier/authgo/tools/errors"
)

type serviceImpl struct {
	dao        Dao
	secService security.Service
}

// Service es la interfaz ue define el servicio
type Service interface {
	SignUp(rate *SignUpRequest) (string, error)
	SignIn(login string, password string) (string, error)
	Get(rateID string) (*Rate, error)
	ChangePassword(rateID string, current string, newPassword string) error
	Grant(rateID string, permissions []string) error
	Revoke(rateID string, permissions []string) error
	Granted(rateID string, permission string) bool
	Disable(rateID string) error
	Enable(rateID string) error
	Rates() ([]*Rate, error)
}

// NewService retorna una nueva instancia del servicio
func NewService() (Service, error) {
	secService, err := security.NewService()
	if err != nil {
		return nil, err
	}

	dao, err := newDao()
	if err != nil {
		return nil, err
	}

	return serviceImpl{
		dao:        dao,
		secService: secService,
	}, nil
}

// MockedService permite mockear el servicio
func MockedService(fakeDao Dao, fakeTRepo security.Service) Service {
	return serviceImpl{
		dao:        fakeDao,
		secService: fakeTRepo,
	}
}

// SignUpRequest es un nuevo usuario
type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Login    string `json:"login" binding:"required"`
}

// SignUp is the controller to signup new rates
func (s serviceImpl) SignUp(rate *SignUpRequest) (string, error) {
	newRate := NewRate()
	newRate.Login = rate.Login
	newRate.Name = rate.Name
	newRate.SetPasswordText(rate.Password)

	newRate, err := s.dao.Insert(newRate)
	if err != nil {
		if db.IsUniqueKeyError(err) {
			return "", ErrLoginExist
		}
		return "", err
	}

	token, err := s.secService.Create(newRate.ID)
	if err != nil {
		return "", nil
	}

	return token.Encode()
}

// SignIn is the controller to sign in rates
func (s serviceImpl) SignIn(login string, password string) (string, error) {
	rate, err := s.dao.FindByLogin(login)
	if err != nil {
		return "", err
	}

	if !rate.Enabled {
		return "", errors.Unauthorized
	}

	if err = rate.ValidatePassword(password); err != nil {
		return "", err
	}

	token, err := s.secService.Create(rate.ID)
	if err != nil {
		return "", nil
	}

	return token.Encode()
}

// Get wrapper para obtener un usuario
func (s serviceImpl) Get(rateID string) (*Rate, error) {
	return s.dao.FindByID(rateID)
}

// ChangePassword cambiar la contraseña del usuario indicado
func (s serviceImpl) ChangePassword(rateID string, current string, newPassword string) error {
	rate, err := s.dao.FindByID(rateID)
	if err != nil {
		return err
	}

	if err = rate.ValidatePassword(current); err != nil {
		return err
	}

	if err = rate.SetPasswordText(newPassword); err != nil {
		return err
	}

	_, err = s.dao.Update(rate)

	return err
}

// Grant Le habilita los permisos enviados por parametros
func (s serviceImpl) Grant(rateID string, permissions []string) error {
	rate, err := s.dao.FindByID(rateID)
	if err != nil {
		return err
	}

	for _, value := range permissions {
		rate.Grant(value)
	}
	_, err = s.dao.Update(rate)

	return err
}

// Revoke Le revoca los permisos enviados por parametros
func (s serviceImpl) Revoke(rateID string, permissions []string) error {
	rate, err := s.dao.FindByID(rateID)
	if err != nil {
		return err
	}

	for _, value := range permissions {
		rate.Revoke(value)
	}
	_, err = s.dao.Update(rate)

	return err
}

//Granted verifica si el usuario tiene el permiso
func (s serviceImpl) Granted(rateID string, permission string) bool {
	usr, err := s.dao.FindByID(rateID)
	if err != nil {
		return false
	}

	return usr.Granted(permission)
}

//Disable deshabilita un usuario
func (s serviceImpl) Disable(rateID string) error {
	usr, err := s.dao.FindByID(rateID)
	if err != nil {
		return err
	}

	usr.Enabled = false

	_, err = s.dao.Update(usr)

	return err
}

//Enable habilita un usuario
func (s serviceImpl) Enable(rateID string) error {
	usr, err := s.dao.FindByID(rateID)
	if err != nil {
		return err
	}

	usr.Enabled = true
	_, err = s.dao.Update(usr)

	return err
}

// Rates wrapper para obtener todos los usuarios
func (s serviceImpl) Rates() ([]*Rate, error) {
	return s.dao.FindAll()
}
