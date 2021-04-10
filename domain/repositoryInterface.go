package domain

import "github.com/google/uuid"

type Repository interface {
	GetById(id uuid.UUID) (Entity, error)
	Create(entity Entity) error
	Update(entity Entity) error
	Delete(entity Entity) error
}
