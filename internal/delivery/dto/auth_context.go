package dto
import("github.com/google/uuid")
type AuthContext struct {
	UserID uuid.UUID
	Role   string
}