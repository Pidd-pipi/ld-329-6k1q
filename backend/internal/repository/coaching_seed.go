package repository

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/model"
)

// seedCoaching 初始化阶段陪练示例数据：
// 进行中的摄影陪练（覆盖待提交/待审核/已驳回/已通过四种状态）、
// 已结项的吉他陪练（含双方技能墙记录）、未布置任务的 Python 陪练。
func seedCoaching() {
	nextTaskID = 7
	nextWallEntryID = 3

	coachingSessions = []model.CoachingSession{
		{
			ID: 1, Mentor: "林澈", Learner: "孟野", Skill: "毕业照人像摄影",
			Time: "周六 10:00", Place: "东校区湖边", Status: constants.CoachingStatusOpen,
			Tasks: []model.PracticeTask{
				{
					ID: 1, Title: "光线与构图练习", Requirement: "在户外完成 3 组不同光线条件下的人像构图练习",
					Status: constants.TaskStatusApproved,
					Submission: &model.TaskSubmission{
						Note: "提交了清晨、正午、逆光三组练习，逆光组用了反光板补光。",
						Link: "https://works.campus.example/mengye/light-composition", SubmittedAt: "2026-09-12 21:08",
					},
					History: []model.TaskEvent{
						{Action: constants.TaskEventCreated, Actor: "林澈", Detail: "光线与构图练习", At: "2026-09-10 10:00"},
						{Action: constants.TaskEventSubmitted, Actor: "孟野", Detail: "提交了清晨、正午、逆光三组练习", At: "2026-09-12 21:08"},
						{Action: constants.TaskEventApproved, Actor: "林澈", Detail: "三组光线控制到位", At: "2026-09-13 09:30"},
					},
				},
				{
					ID: 2, Title: "人像修图作业", Requirement: "精修 5 张人像并导出修图前后对比图",
					Status: constants.TaskStatusSubmitted,
					Submission: &model.TaskSubmission{
						Note: "已按课程要求完成 5 张精修，对比图放在链接里。",
						Link: "https://works.campus.example/mengye/retouch-5", SubmittedAt: "2026-09-24 22:41",
					},
					History: []model.TaskEvent{
						{Action: constants.TaskEventCreated, Actor: "林澈", Detail: "人像修图作业", At: "2026-09-14 10:00"},
						{Action: constants.TaskEventSubmitted, Actor: "孟野", Detail: "已按课程要求完成 5 张精修", At: "2026-09-24 22:41"},
					},
				},
				{
					ID: 3, Title: "毕业照实战跟拍", Requirement: "独立完成一组毕业照跟拍并交付全部成片",
					Status: constants.TaskStatusRejected,
					Submission: &model.TaskSubmission{
						Note: "跟拍了一组三人毕业照，共交付 40 张。",
						Link: "https://works.campus.example/mengye/graduation-shoot", SubmittedAt: "2026-09-20 18:02",
					},
					RejectReason: "逆光场景成片曝光不稳定，建议补拍逆光机位并统一色温后重新提交。",
					History: []model.TaskEvent{
						{Action: constants.TaskEventCreated, Actor: "林澈", Detail: "毕业照实战跟拍", At: "2026-09-14 10:05"},
						{Action: constants.TaskEventSubmitted, Actor: "孟野", Detail: "跟拍了一组三人毕业照", At: "2026-09-20 18:02"},
						{Action: constants.TaskEventRejected, Actor: "林澈", Detail: "逆光场景成片曝光不稳定", At: "2026-09-21 11:20"},
					},
				},
				{
					ID: 4, Title: "作品集整理", Requirement: "整理 10 张代表作并生成在线作品集页面",
					Status: constants.TaskStatusPending,
					History: []model.TaskEvent{
						{Action: constants.TaskEventCreated, Actor: "林澈", Detail: "作品集整理", At: "2026-09-22 15:00"},
					},
				},
			},
		},
		{
			ID: 2, Mentor: "孟野", Learner: "林澈", Skill: "民谣吉他陪练",
			Time: "周三 19:00", Place: "西校区琴房", Status: constants.CoachingStatusClosed,
			ClosedAt: "2026-09-17 20:30",
			Tasks: []model.PracticeTask{
				{
					ID: 5, Title: "C 调和弦转换", Requirement: "C-G-Am-F 和弦转换连贯，每分钟 60 拍不出错",
					Status: constants.TaskStatusApproved,
					Submission: &model.TaskSubmission{
						Note: "录了一段 1 分钟的转换练习视频。",
						Link: "https://works.campus.example/linche/chord-c", SubmittedAt: "2026-09-15 21:12",
					},
					History: []model.TaskEvent{
						{Action: constants.TaskEventCreated, Actor: "孟野", Detail: "C 调和弦转换", At: "2026-09-10 19:00"},
						{Action: constants.TaskEventSubmitted, Actor: "林澈", Detail: "录了一段 1 分钟的转换练习视频", At: "2026-09-15 21:12"},
						{Action: constants.TaskEventApproved, Actor: "孟野", Detail: "转换连贯，可以进入节奏练习", At: "2026-09-16 08:40"},
					},
				},
				{
					ID: 6, Title: "扫弦节奏型练习", Requirement: "掌握 4/4 拍两种常用扫弦节奏型并弹唱一段",
					Status: constants.TaskStatusApproved,
					Submission: &model.TaskSubmission{
						Note: "弹唱了一段《童年》，两种节奏型各一遍。",
						Link: "https://works.campus.example/linche/strumming", SubmittedAt: "2026-09-17 19:26",
					},
					History: []model.TaskEvent{
						{Action: constants.TaskEventCreated, Actor: "孟野", Detail: "扫弦节奏型练习", At: "2026-09-16 08:45"},
						{Action: constants.TaskEventSubmitted, Actor: "林澈", Detail: "弹唱了一段《童年》", At: "2026-09-17 19:26"},
						{Action: constants.TaskEventApproved, Actor: "孟野", Detail: "节奏稳定，全部任务通过，预约结项", At: "2026-09-17 20:30"},
						{Action: constants.TaskEventClosed, Actor: "系统", Detail: "全部任务已通过，预约自动结项", At: "2026-09-17 20:30"},
					},
				},
			},
		},
		{
			ID: 3, Mentor: "周芮", Learner: "许安", Skill: "Python 数据分析",
			Time: "周二 19:30", Place: "线上会议室", Status: constants.CoachingStatusOpen,
			Tasks: []model.PracticeTask{},
		},
	}

	skillWallEntries = []model.SkillWallEntry{
		{
			ID: 1, User: "孟野", Kind: constants.SkillWallMentor, Skill: "民谣吉他陪练",
			Title: "指导完成「民谣吉他陪练」阶段陪练", Detail: "2 项练习任务全部通过，学员林澈完成和弦与扫弦入门。",
			SessionID: 2, CreatedAt: "2026-09-17 20:30",
		},
		{
			ID: 2, User: "林澈", Kind: constants.SkillWallLearner, Skill: "民谣吉他陪练",
			Title: "完成「民谣吉他陪练」阶段学习", Detail: "通过 2 项练习任务：C 调和弦转换、扫弦节奏型练习。",
			SessionID: 2, CreatedAt: "2026-09-17 20:30",
		},
	}
}
