package apperrors

import "avitoTestTask/internal/infrastructure/repository/postgres"

// Errors for API responses
const (
	ErrorCodeTeamExists  = "TEAM_EXISTS"
	ErrorCodePRExists    = "PR_EXISTS"
	ErrorCodePRMerged    = "PR_MERGED"
	ErrorCodeNotAssigned = "NOT_ASSIGNED"
	ErrorCodeNoCandidate = "NO_CANDIDATE"
	ErrorCodeNotFound    = "NOT_FOUND"

	ErrorTeamExistsMsg = "team with that name already exists"
)

func HandleError(err error) (string, string) {
	switch err {
	case postgres.ErrDublicateTeamName:
		return ErrorCodeTeamExists, ErrorTeamExistsMsg
	default:
		return ErrorCodeNotFound, err.Error()
	}
}
