package dto

import "errors"

const (
	maxTeamNameLen         = 255
	maxAmountOfTeamMembers = 100
	maxUsernameLen         = 255
	maxUserIDLen           = 64
)

var (
	ErrInvalidTeamName        = errors.New("invalid team name")
	ErrInvalidAmountOfMembers = errors.New("invalid amount of members")
	ErrInvalidUsername        = errors.New("invalid username")
	ErrInvalidUserID          = errors.New("invalid user id")
)

func (r *CreateTeamRequest) Validate() error {
	if err := ValidateTeamName(r.TeamName); err != nil {
		return err
	}

	if len(r.Members) == 0 || len(r.Members) > maxAmountOfTeamMembers {
		return ErrInvalidAmountOfMembers
	}

	for _, member := range r.Members {
		if err := member.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (r *SetUserActiveRequest) Validate() error {
	if err := ValidateUserID(r.UserID); err != nil {
		return err
	}
	return nil
}

func (m *Member) Validate() error {
	if err := ValidateUserID(m.UserID); err != nil {
		return err
	}
	if err := ValidateUsername(m.Username); err != nil {
		return err
	}
	return nil
}

func ValidateTeamName(name string) error {
	if name == "" || len(name) > maxTeamNameLen {
		return ErrInvalidTeamName
	}
	return nil
}

func ValidateUserID(userID string) error {
	if userID == "" || len(userID) > maxUserIDLen {
		return ErrInvalidUserID
	}
	return nil
}

func ValidateUsername(username string) error {
	if username == "" || len(username) > maxUsernameLen {
		return ErrInvalidUsername
	}
	return nil
}
