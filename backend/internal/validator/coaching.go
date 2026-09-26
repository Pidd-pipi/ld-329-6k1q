package validator

import (
	"strings"

	berrors "cyskillswap/internal/errors"
)

// AddTaskRequest 导师添加练习任务
type AddTaskRequest struct {
	Mentor      string `json:"mentor"`
	Title       string `json:"title"`
	Requirement string `json:"requirement"`
}

// SubmitTaskRequest 学员提交作业说明或作品链接
type SubmitTaskRequest struct {
	Learner string `json:"learner"`
	Note    string `json:"note"`
	Link    string `json:"link"`
}

// ReviewTaskRequest 导师审核作业（approve / reject）
type ReviewTaskRequest struct {
	Mentor string `json:"mentor"`
	Action string `json:"action"`
	Reason string `json:"reason"`
}

func CheckAddTask(req AddTaskRequest) error {
	if strings.TrimSpace(req.Mentor) == "" {
		return berrors.New(berrors.CodeInvalidRequest, "缺少导师身份")
	}
	if strings.TrimSpace(req.Title) == "" {
		return berrors.New(berrors.CodeInvalidRequest, "任务标题不能为空")
	}
	return nil
}

func CheckSubmitTask(req SubmitTaskRequest) error {
	if strings.TrimSpace(req.Learner) == "" {
		return berrors.New(berrors.CodeInvalidRequest, "缺少学员身份")
	}
	if strings.TrimSpace(req.Note) == "" && strings.TrimSpace(req.Link) == "" {
		return berrors.New(berrors.CodeInvalidRequest, "请填写提交说明或作品链接")
	}
	return nil
}

func CheckReviewTask(req ReviewTaskRequest) error {
	if strings.TrimSpace(req.Mentor) == "" {
		return berrors.New(berrors.CodeInvalidRequest, "缺少导师身份")
	}
	if strings.TrimSpace(req.Action) == "" {
		return berrors.New(berrors.CodeInvalidRequest, "缺少审核动作")
	}
	return nil
}
