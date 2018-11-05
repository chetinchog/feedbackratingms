package rules

type serviceImpl struct {
	dao        Dao
	secService security.Service
}

// Service es la interfaz ue define el servicio
type Service interface {
	SignUp(rule *SignUpRequest) (string, error)
	SignIn(login string, password string) (string, error)
	Get(ruleID string) (*Rule, error)
	ChangePassword(ruleID string, current string, newPassword string) error
	Grant(ruleID string, permissions []string) error
	Revoke(ruleID string, permissions []string) error
	Granted(ruleID string, permission string) bool
	Disable(ruleID string) error
	Enable(ruleID string) error
	Rules() ([]*Rule, error)
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

// SignUp is the controller to signup new rules
func (s serviceImpl) SignUp(rule *SignUpRequest) (string, error) {
	newRule := NewRule()
	newRule.Login = rule.Login
	newRule.Name = rule.Name
	newRule.SetPasswordText(rule.Password)

	newRule, err := s.dao.Insert(newRule)
	if err != nil {
		if db.IsUniqueKeyError(err) {
			return "", ErrLoginExist
		}
		return "", err
	}

	token, err := s.secService.Create(newRule.ID)
	if err != nil {
		return "", nil
	}

	return token.Encode()
}

// SignIn is the controller to sign in rules
func (s serviceImpl) SignIn(login string, password string) (string, error) {
	rule, err := s.dao.FindByLogin(login)
	if err != nil {
		return "", err
	}

	if !rule.Enabled {
		return "", errors.Unauthorized
	}

	if err = rule.ValidatePassword(password); err != nil {
		return "", err
	}

	token, err := s.secService.Create(rule.ID)
	if err != nil {
		return "", nil
	}

	return token.Encode()
}

// Get wrapper para obtener un usuario
func (s serviceImpl) Get(ruleID string) (*Rule, error) {
	return s.dao.FindByID(ruleID)
}

// ChangePassword cambiar la contraseña del usuario indicado
func (s serviceImpl) ChangePassword(ruleID string, current string, newPassword string) error {
	rule, err := s.dao.FindByID(ruleID)
	if err != nil {
		return err
	}

	if err = rule.ValidatePassword(current); err != nil {
		return err
	}

	if err = rule.SetPasswordText(newPassword); err != nil {
		return err
	}

	_, err = s.dao.Update(rule)

	return err
}

// Grant Le habilita los permisos enviados por parametros
func (s serviceImpl) Grant(ruleID string, permissions []string) error {
	rule, err := s.dao.FindByID(ruleID)
	if err != nil {
		return err
	}

	for _, value := range permissions {
		rule.Grant(value)
	}
	_, err = s.dao.Update(rule)

	return err
}

// Revoke Le revoca los permisos enviados por parametros
func (s serviceImpl) Revoke(ruleID string, permissions []string) error {
	rule, err := s.dao.FindByID(ruleID)
	if err != nil {
		return err
	}

	for _, value := range permissions {
		rule.Revoke(value)
	}
	_, err = s.dao.Update(rule)

	return err
}

//Granted verifica si el usuario tiene el permiso
func (s serviceImpl) Granted(ruleID string, permission string) bool {
	usr, err := s.dao.FindByID(ruleID)
	if err != nil {
		return false
	}

	return usr.Granted(permission)
}

//Disable deshabilita un usuario
func (s serviceImpl) Disable(ruleID string) error {
	usr, err := s.dao.FindByID(ruleID)
	if err != nil {
		return err
	}

	usr.Enabled = false

	_, err = s.dao.Update(usr)

	return err
}

//Enable habilita un usuario
func (s serviceImpl) Enable(ruleID string) error {
	usr, err := s.dao.FindByID(ruleID)
	if err != nil {
		return err
	}

	usr.Enabled = true
	_, err = s.dao.Update(usr)

	return err
}

// Rules wrapper para obtener todos los usuarios
func (s serviceImpl) Rules() ([]*Rule, error) {
	return s.dao.FindAll()
}
