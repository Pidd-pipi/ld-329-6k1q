package constants

// 阶段陪练预约状态
const (
	CoachingStatusOpen   = "进行中"
	CoachingStatusClosed = "已结项"
)

// 练习任务状态
const (
	TaskStatusPending   = "待提交"
	TaskStatusSubmitted = "待审核"
	TaskStatusApproved  = "已通过"
	TaskStatusRejected  = "已驳回"
)

// 任务历史事件动作
const (
	TaskEventCreated   = "创建任务"
	TaskEventSubmitted = "提交作业"
	TaskEventApproved  = "审核通过"
	TaskEventRejected  = "驳回作业"
	TaskEventClosed    = "预约结项"
)

// 导师审核动作
const (
	ReviewActionApprove = "approve"
	ReviewActionReject  = "reject"
)

// 技能墙条目类型
const (
	SkillWallMentor  = "指导成果"
	SkillWallLearner = "学习成果"
)

// 时间展示格式
const CoachingTimeLayout = "2006-01-02 15:04"
