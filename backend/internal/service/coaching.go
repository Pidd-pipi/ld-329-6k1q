package service

import (
	"cyskillswap/internal/model"
	"cyskillswap/internal/repository"
)

func CoachingAppointments(user, role string) []model.CoachingAppointment {
	return repository.ListCoachingAppointments(user, role)
}

func AddCoachingTask(appointmentID int, req model.AddTaskRequest) (model.CoachingAppointment, error) {
	return repository.AddCoachingTask(appointmentID, req)
}

func SubmitCoachingTask(taskID int, req model.SubmitTaskRequest) (model.CoachingAppointment, error) {
	return repository.SubmitCoachingTask(taskID, req)
}

func ReviewCoachingTask(taskID int, req model.ReviewTaskRequest) (model.CoachingAppointment, error) {
	return repository.ReviewCoachingTask(taskID, req)
}

func CoachingSkillWall(user string) []model.SkillWallEntry {
	return repository.ListSkillWall(user)
}
