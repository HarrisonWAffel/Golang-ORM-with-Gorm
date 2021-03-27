package shared

import "github.com/google/uuid"

type Repository interface {
	GetById(id uuid.UUID) (Entity, error)
	Create(entity BaseEntity) (Entity, error)
	Update(entity BaseEntity) (Entity, error)
	Delete(entity BaseEntity) error
}
