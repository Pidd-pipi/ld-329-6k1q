package errors

import "net/http"

// 业务错误码，前后端共用同一套语义
const (
	CodeInvalidRequest       = "INVALID_REQUEST"
	CodeSessionNotFound      = "SESSION_NOT_FOUND"
	CodeTaskNotFound         = "TASK_NOT_FOUND"
	CodeSessionClosed        = "SESSION_CLOSED"
	CodeNotMentor            = "NOT_MENTOR"
	CodeNotLearner           = "NOT_LEARNER"
	CodeTaskPendingReview    = "TASK_PENDING_REVIEW"
	CodeTaskNotSubmittable   = "TASK_NOT_SUBMITTABLE"
	CodeTaskNotReviewable    = "TASK_NOT_REVIEWABLE"
	CodeRejectReasonRequired = "REJECT_REASON_REQUIRED"
	CodeInternal             = "INTERNAL_ERROR"
)

func New(code, message string) BusinessError {
	return BusinessError{Code: code, Message: message}
}

// HTTPStatus 把业务错误码映射为 HTTP 状态码
func HTTPStatus(err BusinessError) int {
	switch err.Code {
	case CodeSessionNotFound, CodeTaskNotFound:
		return http.StatusNotFound
	case CodeNotMentor, CodeNotLearner:
		return http.StatusForbidden
	case CodeSessionClosed, CodeTaskPendingReview, CodeTaskNotSubmittable, CodeTaskNotReviewable:
		return http.StatusConflict
	case CodeInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
