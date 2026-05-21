package ports

import identity "backgammon-teacher/internal/modules/identity/domain"

type UserRepository interface {
	Save(user identity.User) error
	Get(id string) (identity.User, error)
}
