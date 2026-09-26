package repository

import (
	"sort"
	"strings"
	"sync"
	"time"

	"cyskillswap/internal/errors"
	"cyskillswap/internal/model"
)

// coachingStore 保存阶段陪练的可变状态，所有读写都经过互斥锁。
type coachingStore struct {
	mu           sync.Mutex
	appointments []model.CoachingAppointment
	skillWall    []model.SkillWallEntry
	nextTaskID   int
}

var coaching = newCoachingStore()

func newCoachingStore() *coachingStore {
	s := &coachingStore{nextTaskID: 1}
	s.appointments = []model.CoachingAppointment{
		{
			ID: 1, Mentor: "孟野", Learner: "林澈", Skill: "民谣吉他陪练",
			Time: "周六 10:00", Place: "东校区湖边",
		},
		{
			ID: 2, Mentor: "林澈", Learner: "孟野", Skill: "毕业照人像摄影",
			Time: "周六 14:00", Place: "东校区湖边",
		},
		{
			ID: 3, Mentor: "周芮", Learner: "林澈", Skill: "Python 数据分析",
			Time: "周二 19:30", Place: "线上会议室",
			Closed: true, ClosedAt: "2026-09-20 21:00",
		},
	}
	// 预约 1：进行中的吉他陪练，覆盖四种任务状态。
	s.appointments[0].Tasks = []model.CoachingTask{
		s.seedTask(1, "爬格子与右手分解基本功", "每天 20 分钟，录制 1 分钟连贯练习视频。",
			model.TaskStatusApproved, &model.TaskSubmission{Note: "已能稳定 80bpm，视频见链接。", Link: "https://example.com/video/basic", SubmittedAt: "2026-09-21 20:12"}, ""),
		s.seedTask(1, "C 大调开放和弦转换练习", "C-G-Am-F 循环 4 遍不中断，提交音频。",
			model.TaskStatusSubmitted, &model.TaskSubmission{Note: "录了三遍，这版最稳。", Link: "https://example.com/audio/chords", SubmittedAt: "2026-09-25 22:40"}, ""),
		s.seedTask(1, "弹唱《童年》完整录制", "完整弹唱一遍，节奏型自定，提交视频链接。",
			model.TaskStatusRejected, &model.TaskSubmission{Note: "第一次完整录下来了。", Link: "https://example.com/video/tongnian", SubmittedAt: "2026-09-24 21:05"},
			"副歌节奏不稳，第 2 段抢拍，建议跟节拍器 60bpm 重录。"),
		s.seedTask(1, "扫弦节奏型练习", "下 下上 上下上，配合和弦转换，提交练习说明。", model.TaskStatusTodo, nil, ""),
	}
	// 预约 2：林澈作为导师的摄影陪练。
	s.appointments[1].Tasks = []model.CoachingTask{
		s.seedTask(2, "三分法构图练习", "拍摄 9 张不同主体的三分法构图样片。",
			model.TaskStatusApproved, &model.TaskSubmission{Note: "样片已整理成相册。", Link: "https://example.com/album/rule-of-thirds", SubmittedAt: "2026-09-22 18:30"}, ""),
		s.seedTask(2, "逆光人像外拍实践", "傍晚逆光人像 3 张，附曝光参数说明。", model.TaskStatusTodo, nil, ""),
	}
	// 预约 3：已结项，任务全部通过，历史不可改动。
	s.appointments[2].Tasks = []model.CoachingTask{
		s.seedTask(3, "pandas 数据清洗练习", "清洗给定 CSV，输出去重和缺失值处理说明。",
			model.TaskStatusApproved, &model.TaskSubmission{Note: "处理报告在 notebook 里。", Link: "https://example.com/nb/clean", SubmittedAt: "2026-09-18 20:00"}, ""),
		s.seedTask(3, "问卷数据可视化报告", "用 matplotlib 出 3 张图并写结论。",
			model.TaskStatusApproved, &model.TaskSubmission{Note: "图表和结论已提交。", Link: "https://example.com/nb/viz", SubmittedAt: "2026-09-20 19:30"}, ""),
	}
	s.skillWall = []model.SkillWallEntry{
		{User: "林澈", Skill: "Python 数据分析", Partner: "周芮", Note: "完成 2 项阶段任务：数据清洗、可视化报告", ClosedAt: "2026-09-20 21:00"},
		{User: "周芮", Skill: "Python 数据分析", Partner: "林澈", Note: "陪练 2 项阶段任务并全部确认通过", ClosedAt: "2026-09-20 21:00"},
	}
	return s
}

func (s *coachingStore) seedTask(appointmentID int, title, requirement, status string, sub *model.TaskSubmission, rejectReason string) model.CoachingTask {
	t := model.CoachingTask{
		ID: s.nextTaskID, AppointmentID: appointmentID, Title: title,
		Requirement: requirement, Status: status, Submission: sub, RejectReason: rejectReason,
	}
	s.nextTaskID++
	return t
}

// ListCoachingAppointments 按视角（导师/学员）返回用户相关的预约及任务。
func ListCoachingAppointments(user, role string) []model.CoachingAppointment {
	coaching.mu.Lock()
	defer coaching.mu.Unlock()
	result := make([]model.CoachingAppointment, 0)
	for _, a := range coaching.appointments {
		if role == "mentor" && a.Mentor != user {
			continue
		}
		if role == "learner" && a.Learner != user {
			continue
		}
		result = append(result, a)
	}
	return result
}

func (s *coachingStore) findAppointment(id int) (*model.CoachingAppointment, error) {
	for i := range s.appointments {
		if s.appointments[i].ID == id {
			return &s.appointments[i], nil
		}
	}
	return nil, errors.ErrAppointmentNotFound
}

// AddCoachingTask 导师为未结项预约追加练习任务。
func AddCoachingTask(appointmentID int, req model.AddTaskRequest) (model.CoachingAppointment, error) {
	coaching.mu.Lock()
	defer coaching.mu.Unlock()
	a, err := coaching.findAppointment(appointmentID)
	if err != nil {
		return model.CoachingAppointment{}, err
	}
	if a.Mentor != req.User {
		return model.CoachingAppointment{}, errors.ErrNotMentor
	}
	if a.Closed {
		return model.CoachingAppointment{}, errors.ErrAppointmentClosed
	}
	if strings.TrimSpace(req.Title) == "" {
		return model.CoachingAppointment{}, errors.ErrEmptyTaskTitle
	}
	a.Tasks = append(a.Tasks, coaching.seedTask(a.ID, strings.TrimSpace(req.Title), strings.TrimSpace(req.Requirement), model.TaskStatusTodo, nil, ""))
	return *a, nil
}

// SubmitCoachingTask 学员提交练习说明或作品链接；等待导师处理时禁止重复提交。
func SubmitCoachingTask(taskID int, req model.SubmitTaskRequest) (model.CoachingAppointment, error) {
	coaching.mu.Lock()
	defer coaching.mu.Unlock()
	a, t, err := locateTask(taskID)
	if err != nil {
		return model.CoachingAppointment{}, err
	}
	if a.Learner != req.User {
		return model.CoachingAppointment{}, errors.ErrNotLearner
	}
	if a.Closed {
		return model.CoachingAppointment{}, errors.ErrAppointmentClosed
	}
	if t.Status == model.TaskStatusSubmitted {
		return model.CoachingAppointment{}, errors.ErrTaskPendingReview
	}
	if t.Status == model.TaskStatusApproved {
		return model.CoachingAppointment{}, errors.ErrAppointmentClosed
	}
	if strings.TrimSpace(req.Note) == "" && strings.TrimSpace(req.Link) == "" {
		return model.CoachingAppointment{}, errors.ErrEmptySubmission
	}
	t.Submission = &model.TaskSubmission{
		Note:        strings.TrimSpace(req.Note),
		Link:        strings.TrimSpace(req.Link),
		SubmittedAt: time.Now().Format("2006-01-02 15:04"),
	}
	t.RejectReason = ""
	t.Status = model.TaskStatusSubmitted
	return *a, nil
}

// ReviewCoachingTask 导师确认通过或驳回；全部通过后预约结项并写入双方技能墙。
func ReviewCoachingTask(taskID int, req model.ReviewTaskRequest) (model.CoachingAppointment, error) {
	coaching.mu.Lock()
	defer coaching.mu.Unlock()
	a, t, err := locateTask(taskID)
	if err != nil {
		return model.CoachingAppointment{}, err
	}
	if a.Mentor != req.User {
		return model.CoachingAppointment{}, errors.ErrNotMentor
	}
	if a.Closed {
		return model.CoachingAppointment{}, errors.ErrAppointmentClosed
	}
	if t.Status != model.TaskStatusSubmitted {
		return model.CoachingAppointment{}, errors.ErrTaskNotAwaiting
	}
	if req.Approve {
		t.Status = model.TaskStatusApproved
		t.RejectReason = ""
		closeIfFinished(a)
	} else {
		if strings.TrimSpace(req.Reason) == "" {
			return model.CoachingAppointment{}, errors.ErrRejectReasonRequired
		}
		t.Status = model.TaskStatusRejected
		t.RejectReason = strings.TrimSpace(req.Reason)
	}
	return *a, nil
}

func locateTask(taskID int) (*model.CoachingAppointment, *model.CoachingTask, error) {
	for i := range coaching.appointments {
		a := &coaching.appointments[i]
		for j := range a.Tasks {
			if a.Tasks[j].ID == taskID {
				return a, &a.Tasks[j], nil
			}
		}
	}
	return nil, nil, errors.ErrTaskNotFound
}

// closeIfFinished 任务全部通过时结项，并把成果写入双方技能墙。
func closeIfFinished(a *model.CoachingAppointment) {
	done, total := a.Progress()
	if total == 0 || done < total {
		return
	}
	a.Closed = true
	a.ClosedAt = time.Now().Format("2006-01-02 15:04")
	note := "完成全部阶段任务：" + taskTitles(a)
	coaching.skillWall = append(coaching.skillWall,
		model.SkillWallEntry{User: a.Learner, Skill: a.Skill, Partner: a.Mentor, Note: note, ClosedAt: a.ClosedAt},
		model.SkillWallEntry{User: a.Mentor, Skill: a.Skill, Partner: a.Learner, Note: "陪练全部阶段任务并确认通过", ClosedAt: a.ClosedAt},
	)
}

func taskTitles(a *model.CoachingAppointment) string {
	titles := make([]string, 0, len(a.Tasks))
	for _, t := range a.Tasks {
		titles = append(titles, t.Title)
	}
	return strings.Join(titles, "、")
}

// ListSkillWall 返回用户的技能墙记录，按结项时间倒序。
func ListSkillWall(user string) []model.SkillWallEntry {
	coaching.mu.Lock()
	defer coaching.mu.Unlock()
	entries := make([]model.SkillWallEntry, 0)
	for _, e := range coaching.skillWall {
		if e.User == user {
			entries = append(entries, e)
		}
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].ClosedAt > entries[j].ClosedAt })
	return entries
}
