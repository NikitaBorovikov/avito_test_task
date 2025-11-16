package apperrors

import (
	"errors"
	"net/http"
)

// Errors for API responses
const (
	ErrorCodeTeamExists  = "TEAM_EXISTS"
	ErrorCodePRExists    = "PR_EXISTS"
	ErrorCodePRMerged    = "PR_MERGED"
	ErrorCodeNotAssigned = "NOT_ASSIGNED"
	ErrorCodeNoCandidate = "NO_CANDIDATE"
	ErrorCodeNotFound    = "NOT_FOUND"

	ErrorCodeInternalError = "INTERNAL_ERROR"
)

// Error messages
const (
	ErrTeamExistsMsg = "team_name already exists"
	ErrPRIDExistsMsg = "PR id already exists"

	ErrUserNotFoundMsg = "user not found"
	ErrTeamNotFoundMsg = "team not found"
	ErrPRNotFoundMsg   = "PR not found"

	ErrPRMergedMsg    = "cannot reassign on merged PR"
	ErrNotAssignedMsg = "reviewer is not assigned to this PR"
	ErrNoCandidateMsg = "no active replacement candidate in team"

	ErrInternalMsg = "internal server error"
)

// App errors
var (
	ErrDuplicatePRID     = errors.New("pull request with this ID already exists")
	ErrDuplicateTeamName = errors.New("team with that name already exists")

	ErrPRNotFound   = errors.New("pull request with this ID is not found")
	ErrTeamNotFound = errors.New("team with this ID is not found")
	ErrUserNotFound = errors.New("user with this ID is not found")

	ErrAlreadyMerged = errors.New("already merged")
	ErrNoCandidate   = errors.New("there aren't available candidates")
	ErrNotAssigned   = errors.New("user is not assigned to this pull request")
)

type ErrorInfo struct {
	Code     string
	Msg      string
	HttdCode int
}

func HandleError(err error) ErrorInfo {
	switch err {
	case ErrDuplicateTeamName:
		return ErrorInfo{ErrorCodeTeamExists, ErrTeamExistsMsg, http.StatusBadRequest}
	case ErrDuplicatePRID:
		return ErrorInfo{ErrorCodePRExists, ErrPRIDExistsMsg, http.StatusConflict}
	case ErrPRNotFound:
		return ErrorInfo{ErrorCodeNotFound, ErrPRNotFoundMsg, http.StatusNotFound}
	case ErrTeamNotFound:
		return ErrorInfo{ErrorCodeNotFound, ErrTeamNotFoundMsg, http.StatusNotFound}
	case ErrUserNotFound:
		return ErrorInfo{ErrorCodeNotFound, ErrUserNotFoundMsg, http.StatusNotFound}
	case ErrAlreadyMerged:
		return ErrorInfo{ErrorCodePRMerged, ErrPRMergedMsg, http.StatusConflict}
	case ErrNoCandidate:
		return ErrorInfo{ErrorCodeNoCandidate, ErrNoCandidateMsg, http.StatusConflict}
	case ErrNotAssigned:
		return ErrorInfo{ErrorCodeNotAssigned, ErrNotAssignedMsg, http.StatusConflict}
	default:
		return ErrorInfo{ErrorCodeInternalError, ErrInternalMsg, http.StatusInternalServerError}
	}
}
