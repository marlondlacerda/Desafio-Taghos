package util

import "github.com/gofrs/uuid"

// IsValidUUID verifica se o valor fornecido é um UUID válido.
func IsValidUUID(id string) bool {
	_, err := uuid.FromString(id)
	return err == nil
}
