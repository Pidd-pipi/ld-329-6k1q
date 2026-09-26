package errors

type BusinessError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e BusinessError) Error() string { return e.Message }

func New(code, message string) BusinessError {
	return BusinessError{Code: code, Message: message}
}

// 阶段陪练业务错误，错误码集中维护，禁止在业务代码中抛裸字符串。
var (
	ErrAppointmentNotFound  = New("APPOINTMENT_NOT_FOUND", "预约不存在")
	ErrAppointmentClosed    = New("APPOINTMENT_CLOSED", "预约已结项，历史记录不能改动")
	ErrNotMentor            = New("NOT_MENTOR", "只有导师可以执行该操作")
	ErrNotLearner           = New("NOT_LEARNER", "只有学员可以执行该操作")
	ErrTaskNotFound         = New("TASK_NOT_FOUND", "练习任务不存在")
	ErrTaskPendingReview    = New("TASK_PENDING_REVIEW", "任务已提交，等待导师处理，不能重复提交")
	ErrTaskNotAwaiting      = New("TASK_NOT_AWAITING_REVIEW", "任务当前不在待审核状态")
	ErrRejectReasonRequired = New("REJECT_REASON_REQUIRED", "驳回必须说明原因")
	ErrEmptyTaskTitle       = New("EMPTY_TASK_TITLE", "任务标题不能为空")
	ErrEmptySubmission      = New("EMPTY_SUBMISSION", "请填写练习说明或作品链接")
)
