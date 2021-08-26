package api

import "exercises/models/database"

type CreateUserRequest struct {
	Username string
	Password string
	Email    string `json:"email" validate:"required,email"`
	Names    string
}

type UpdateUserRequest struct {
	Email    *string
	Names    *string
	Password *string
}

type UserResponse struct {
	ID    int
	Email string
	Names string
}

func UserResponseFromDBModel(u database.UserDBModel) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Names: u.Names,
	}
}

func UserDBModelFromCreateRequest(r CreateUserRequest) (database.UserDBModel, error) {
	hashedPassword, err := HashPassword(r.Password)
	if err != nil {
		return database.UserDBModel{}, err
	}
	return database.UserDBModel{
		Username:     r.Username,
		Email:        r.Email,
		PasswordHash: hashedPassword,
		Names:        r.Names,
	}, nil
}
