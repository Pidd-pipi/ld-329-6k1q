package repository

import (
	"sync"

	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

var (
	coachingMu       sync.RWMutex
	coachingSessions []model.CoachingSession
	skillWallEntries []model.SkillWallEntry
	nextTaskID       int
	nextWallEntryID  int
)

func init() {
	seedCoaching()
}

// ListCoachingSessions 返回全部阶段陪练预约的深拷贝
func ListCoachingSessions() []model.CoachingSession {
	coachingMu.RLock()
	defer coachingMu.RUnlock()
	out := make([]model.CoachingSession, len(coachingSessions))
	for i, s := range coachingSessions {
		out[i] = cloneCoachingSession(s)
	}
	return out
}

// FindCoachingSession 按 ID 查找预约，返回深拷贝
func FindCoachingSession(id int) (model.CoachingSession, bool) {
	coachingMu.RLock()
	defer coachingMu.RUnlock()
	for _, s := range coachingSessions {
		if s.ID == id {
			return cloneCoachingSession(s), true
		}
	}
	return model.CoachingSession{}, false
}

// NextCoachingTaskID 生成全局唯一的任务 ID
func NextCoachingTaskID() int {
	coachingMu.Lock()
	defer coachingMu.Unlock()
	id := nextTaskID
	nextTaskID++
	return id
}

// AddCoachingTask 向进行中的预约追加任务
func AddCoachingTask(sessionID int, task model.PracticeTask) bool {
	coachingMu.Lock()
	defer coachingMu.Unlock()
	for i, s := range coachingSessions {
		if s.ID == sessionID && s.Status == constants.CoachingStatusOpen {
			coachingSessions[i].Tasks = append(coachingSessions[i].Tasks, task)
			return true
		}
	}
	return false
}

// UpdateCoachingTask 覆盖更新预约中的某个任务
func UpdateCoachingTask(sessionID int, task model.PracticeTask) bool {
	coachingMu.Lock()
	defer coachingMu.Unlock()
	for i, s := range coachingSessions {
		if s.ID != sessionID {
			continue
		}
		for j, t := range s.Tasks {
			if t.ID == task.ID {
				coachingSessions[i].Tasks[j] = task
				return true
			}
		}
	}
	return false
}

// CloseCoachingSession 将预约标记为已结项
func CloseCoachingSession(sessionID int, closedAt string) bool {
	coachingMu.Lock()
	defer coachingMu.Unlock()
	for i, s := range coachingSessions {
		if s.ID == sessionID && s.Status == constants.CoachingStatusOpen {
			coachingSessions[i].Status = constants.CoachingStatusClosed
			coachingSessions[i].ClosedAt = closedAt
			return true
		}
	}
	return false
}

// AddSkillWallEntries 写入技能墙条目并分配 ID
func AddSkillWallEntries(entries []model.SkillWallEntry) {
	coachingMu.Lock()
	defer coachingMu.Unlock()
	for _, e := range entries {
		e.ID = nextWallEntryID
		nextWallEntryID++
		skillWallEntries = append(skillWallEntries, e)
	}
}

// ListSkillWall 返回指定用户的技能墙条目
func ListSkillWall(user string) []model.SkillWallEntry {
	coachingMu.RLock()
	defer coachingMu.RUnlock()
	out := make([]model.SkillWallEntry, 0)
	for _, e := range skillWallEntries {
		if user == "" || e.User == user {
			out = append(out, e)
		}
	}
	return out
}

func cloneCoachingSession(s model.CoachingSession) model.CoachingSession {
	tasks := make([]model.PracticeTask, len(s.Tasks))
	for i, t := range s.Tasks {
		tasks[i] = clonePracticeTask(t)
	}
	s.Tasks = tasks
	return s
}

func clonePracticeTask(t model.PracticeTask) model.PracticeTask {
	history := make([]model.TaskEvent, len(t.History))
	copy(history, t.History)
	t.History = history
	if t.Submission != nil {
		submission := *t.Submission
		t.Submission = &submission
	}
	return t
}
