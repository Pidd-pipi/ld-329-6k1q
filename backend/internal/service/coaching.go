package service

import (
	"fmt"
	"strings"
	"time"

	"cyskillswap/internal/constants"
	berrors "cyskillswap/internal/errors"
	"cyskillswap/internal/logger"
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

// CoachingSessions 返回全部阶段陪练预约
func CoachingSessions() []model.CoachingSession {
	return repository.ListCoachingSessions()
}

// SkillWall 返回指定用户的技能墙
func SkillWall(user string) []model.SkillWallEntry {
	return repository.ListSkillWall(user)
}

// AddCoachingTask 导师为进行中的预约追加练习任务
func AddCoachingTask(sessionID int, mentor, title, requirement string) (model.CoachingSession, error) {
	session, err := mustOpenSession(sessionID)
	if err != nil {
		return model.CoachingSession{}, err
	}
	if mentor != session.Mentor {
		return model.CoachingSession{}, berrors.New(berrors.CodeNotMentor, "只有导师可以为该预约添加练习任务")
	}
	task := model.PracticeTask{
		ID:          repository.NextCoachingTaskID(),
		Title:       strings.TrimSpace(title),
		Requirement: strings.TrimSpace(requirement),
		Status:      constants.TaskStatusPending,
		History: []model.TaskEvent{
			{Action: constants.TaskEventCreated, Actor: mentor, Detail: strings.TrimSpace(title), At: nowText()},
		},
	}
	if !repository.AddCoachingTask(sessionID, task) {
		return model.CoachingSession{}, berrors.New(berrors.CodeSessionClosed, "预约已结项，不能追加任务")
	}
	logger.Info("coaching task added", "session", sessionID, "task", task.ID)
	return currentSession(sessionID)
}

// SubmitCoachingTask 学员提交作业说明或作品链接；等待导师处理时禁止重复提交
func SubmitCoachingTask(sessionID, taskID int, learner, note, link string) (model.CoachingSession, error) {
	session, err := mustOpenSession(sessionID)
	if err != nil {
		return model.CoachingSession{}, err
	}
	if learner != session.Learner {
		return model.CoachingSession{}, berrors.New(berrors.CodeNotLearner, "只有学员可以提交该预约的作业")
	}
	task, found := findTask(session, taskID)
	if !found {
		return model.CoachingSession{}, berrors.New(berrors.CodeTaskNotFound, "练习任务不存在")
	}
	switch task.Status {
	case constants.TaskStatusSubmitted:
		return model.CoachingSession{}, berrors.New(berrors.CodeTaskPendingReview, "任务正在等待导师处理，不能重复提交")
	case constants.TaskStatusApproved:
		return model.CoachingSession{}, berrors.New(berrors.CodeTaskNotSubmittable, "任务已通过，无需再次提交")
	}
	now := nowText()
	task.Status = constants.TaskStatusSubmitted
	task.RejectReason = ""
	task.Submission = &model.TaskSubmission{
		Note:        strings.TrimSpace(note),
		Link:        strings.TrimSpace(link),
		SubmittedAt: now,
	}
	task.History = append(task.History, model.TaskEvent{
		Action: constants.TaskEventSubmitted, Actor: learner, Detail: firstNonEmpty(note, link), At: now,
	})
	repository.UpdateCoachingTask(sessionID, task)
	logger.Info("coaching task submitted", "session", sessionID, "task", taskID)
	return currentSession(sessionID)
}

// ReviewCoachingTask 导师审核作业；全部任务通过后预约自动结项并写入双方技能墙
func ReviewCoachingTask(sessionID, taskID int, mentor, action, reason string) (model.CoachingSession, error) {
	session, err := mustOpenSession(sessionID)
	if err != nil {
		return model.CoachingSession{}, err
	}
	if mentor != session.Mentor {
		return model.CoachingSession{}, berrors.New(berrors.CodeNotMentor, "只有导师可以审核该预约的作业")
	}
	task, found := findTask(session, taskID)
	if !found {
		return model.CoachingSession{}, berrors.New(berrors.CodeTaskNotFound, "练习任务不存在")
	}
	if task.Status != constants.TaskStatusSubmitted {
		return model.CoachingSession{}, berrors.New(berrors.CodeTaskNotReviewable, "任务不在待审核状态")
	}
	now := nowText()
	switch action {
	case constants.ReviewActionApprove:
		task.Status = constants.TaskStatusApproved
		task.History = append(task.History, model.TaskEvent{
			Action: constants.TaskEventApproved, Actor: mentor, Detail: "审核通过", At: now,
		})
	case constants.ReviewActionReject:
		if strings.TrimSpace(reason) == "" {
			return model.CoachingSession{}, berrors.New(berrors.CodeRejectReasonRequired, "驳回时必须填写原因，方便学员修改后重新提交")
		}
		task.Status = constants.TaskStatusRejected
		task.RejectReason = strings.TrimSpace(reason)
		task.History = append(task.History, model.TaskEvent{
			Action: constants.TaskEventRejected, Actor: mentor, Detail: task.RejectReason, At: now,
		})
	default:
		return model.CoachingSession{}, berrors.New(berrors.CodeInvalidRequest, "不支持的审核动作")
	}
	repository.UpdateCoachingTask(sessionID, task)
	if action == constants.ReviewActionApprove {
		closeSessionIfAllApproved(sessionID, now)
	}
	return currentSession(sessionID)
}

// mustOpenSession 校验预约存在且未结项；结项后的历史不可改动
func mustOpenSession(sessionID int) (model.CoachingSession, error) {
	session, ok := repository.FindCoachingSession(sessionID)
	if !ok {
		return model.CoachingSession{}, berrors.New(berrors.CodeSessionNotFound, "陪练预约不存在")
	}
	if session.Status == constants.CoachingStatusClosed {
		return model.CoachingSession{}, berrors.New(berrors.CodeSessionClosed, "预约已结项，历史记录不可改动")
	}
	return session, nil
}

// closeSessionIfAllApproved 全部任务通过时结项，并向导师和学员的技能墙各写入一条记录
func closeSessionIfAllApproved(sessionID int, now string) {
	session, ok := repository.FindCoachingSession(sessionID)
	if !ok || len(session.Tasks) == 0 {
		return
	}
	for _, t := range session.Tasks {
		if t.Status != constants.TaskStatusApproved {
			return
		}
	}
	if !repository.CloseCoachingSession(sessionID, now) {
		return
	}
	repository.AddSkillWallEntries([]model.SkillWallEntry{
		{
			User: session.Mentor, Kind: constants.SkillWallMentor, Skill: session.Skill,
			Title:     fmt.Sprintf("指导完成「%s」阶段陪练", session.Skill),
			Detail:    fmt.Sprintf("%d 项练习任务全部通过，学员%s完成本阶段学习。", len(session.Tasks), session.Learner),
			SessionID: sessionID,
			CreatedAt: now,
		},
		{
			User: session.Learner, Kind: constants.SkillWallLearner, Skill: session.Skill,
			Title:     fmt.Sprintf("完成「%s」阶段学习", session.Skill),
			Detail:    fmt.Sprintf("通过 %d 项练习任务：%s。", len(session.Tasks), taskTitles(session.Tasks)),
			SessionID: sessionID,
			CreatedAt: now,
		},
	})
	logger.Info("coaching session closed", "session", sessionID)
}

func findTask(session model.CoachingSession, taskID int) (model.PracticeTask, bool) {
	for _, t := range session.Tasks {
		if t.ID == taskID {
			return t, true
		}
	}
	return model.PracticeTask{}, false
}

func currentSession(sessionID int) (model.CoachingSession, error) {
	session, ok := repository.FindCoachingSession(sessionID)
	if !ok {
		return model.CoachingSession{}, berrors.New(berrors.CodeSessionNotFound, "陪练预约不存在")
	}
	return session, nil
}

func taskTitles(tasks []model.PracticeTask) string {
	titles := make([]string, 0, len(tasks))
	for _, t := range tasks {
		titles = append(titles, t.Title)
	}
	return strings.Join(titles, "、")
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func nowText() string {
	return time.Now().Format(constants.CoachingTimeLayout)
}
