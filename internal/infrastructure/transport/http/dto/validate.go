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
	if r.TeamName == "" || len(r.TeamName) > maxTeamNameLen {
		return ErrInvalidTeamName
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

func (m *Member) Validate() error {
	if m.UserID == "" || len(m.UserID) > maxUserIDLen {
		return ErrInvalidUserID
	}

	if m.Username == "" || len(m.Username) > maxUsernameLen {
		return ErrInvalidUsername
	}
	return nil
}
