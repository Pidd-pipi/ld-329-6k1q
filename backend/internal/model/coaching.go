package model

// TaskSubmission 学员的一次作业提交（说明或作品链接）
type TaskSubmission struct {
	Note        string `json:"note"`
	Link        string `json:"link"`
	SubmittedAt string `json:"submittedAt"`
}

// TaskEvent 任务历史事件，结项后作为不可改动的历史保留
type TaskEvent struct {
	Action string `json:"action"`
	Actor  string `json:"actor"`
	Detail string `json:"detail"`
	At     string `json:"at"`
}

// PracticeTask 导师布置的一项练习任务
type PracticeTask struct {
	ID           int             `json:"id"`
	Title        string          `json:"title"`
	Requirement  string          `json:"requirement"`
	Status       string          `json:"status"`
	Submission   *TaskSubmission `json:"submission,omitempty"`
	RejectReason string          `json:"rejectReason,omitempty"`
	History      []TaskEvent     `json:"history"`
}

// CoachingSession 由已确认预约扩展而来的阶段陪练
type CoachingSession struct {
	ID       int            `json:"id"`
	Mentor   string         `json:"mentor"`
	Learner  string         `json:"learner"`
	Skill    string         `json:"skill"`
	Time     string         `json:"time"`
	Place    string         `json:"place"`
	Status   string         `json:"status"`
	Tasks    []PracticeTask `json:"tasks"`
	ClosedAt string         `json:"closedAt,omitempty"`
}

// SkillWallEntry 结项后写入双方技能墙的记录
type SkillWallEntry struct {
	ID        int    `json:"id"`
	User      string `json:"user"`
	Kind      string `json:"kind"`
	Skill     string `json:"skill"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	SessionID int    `json:"sessionId"`
	CreatedAt string `json:"createdAt"`
}
