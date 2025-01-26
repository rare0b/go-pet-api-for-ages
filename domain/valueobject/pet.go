package valueobject

import (
	"errors"
)

type PetStatus string

const (
	PetStatusAvailable   PetStatus = "available"
	PetStatusUnavailable PetStatus = "unavailable"
)

func NewPetStatus(status string) (PetStatus, error) {
	switch PetStatus(status) {
	case PetStatusAvailable, PetStatusUnavailable:
		return PetStatus(status), nil
	default:
		return "", errors.New("invalid pet status")
	}
}
