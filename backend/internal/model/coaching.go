package model

// 阶段陪练任务状态：待提交 -> 待审核 -> 已完成，导师驳回后回到已驳回可重新提交。
const (
	TaskStatusTodo      = "todo"
	TaskStatusSubmitted = "submitted"
	TaskStatusApproved  = "approved"
	TaskStatusRejected  = "rejected"
)

type TaskSubmission struct {
	Note        string `json:"note"`
	Link        string `json:"link"`
	SubmittedAt string `json:"submittedAt"`
}

type CoachingTask struct {
	ID            int             `json:"id"`
	AppointmentID int             `json:"appointmentId"`
	Title         string          `json:"title"`
	Requirement   string          `json:"requirement"`
	Status        string          `json:"status"`
	Submission    *TaskSubmission `json:"submission,omitempty"`
	RejectReason  string          `json:"rejectReason,omitempty"`
}

type CoachingAppointment struct {
	ID       int            `json:"id"`
	Mentor   string         `json:"mentor"`
	Learner  string         `json:"learner"`
	Skill    string         `json:"skill"`
	Time     string         `json:"time"`
	Place    string         `json:"place"`
	Closed   bool           `json:"closed"`
	ClosedAt string         `json:"closedAt,omitempty"`
	Tasks    []CoachingTask `json:"tasks"`
}

// Progress 返回任务完成进度，用于页面展示待办和结项状态。
func (a CoachingAppointment) Progress() (done int, total int) {
	total = len(a.Tasks)
	for _, t := range a.Tasks {
		if t.Status == TaskStatusApproved {
			done++
		}
	}
	return done, total
}

type SkillWallEntry struct {
	User     string `json:"user"`
	Skill    string `json:"skill"`
	Partner  string `json:"partner"`
	Note     string `json:"note"`
	ClosedAt string `json:"closedAt"`
}

type AddTaskRequest struct {
	User        string `json:"user"`
	Title       string `json:"title"`
	Requirement string `json:"requirement"`
}

type SubmitTaskRequest struct {
	User string `json:"user"`
	Note string `json:"note"`
	Link string `json:"link"`
}

type ReviewTaskRequest struct {
	User    string `json:"user"`
	Approve bool   `json:"approve"`
	Reason  string `json:"reason"`
}
