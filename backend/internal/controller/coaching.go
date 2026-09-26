package controller

import (
	"net/http"
	"strconv"

	bizerrors "cyskillswap/internal/errors"
	"cyskillswap/internal/model"
	"cyskillswap/internal/service"
	"github.com/gin-gonic/gin"
)

// respondError 统一把业务错误转成 400 响应，未知错误按 500 处理。
func respondError(c *gin.Context, err error) {
	if be, ok := err.(bizerrors.BusinessError); ok {
		c.JSON(http.StatusBadRequest, be)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL", "message": "服务器内部错误"})
}

// CoachingAppointments 按角色（mentor/learner）返回当前用户的阶段陪练预约。
func CoachingAppointments(c *gin.Context) {
	user := c.Query("user")
	role := c.DefaultQuery("role", "learner")
	c.JSON(http.StatusOK, service.CoachingAppointments(user, role))
}

// AddCoachingTask 导师为未结项预约追加练习任务。
func AddCoachingTask(c *gin.Context) {
	appointmentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondError(c, bizerrors.ErrAppointmentNotFound)
		return
	}
	var req model.AddTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, bizerrors.ErrEmptyTaskTitle)
		return
	}
	appointment, err := service.AddCoachingTask(appointmentID, req)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, appointment)
}

// SubmitCoachingTask 学员提交练习说明或作品链接。
func SubmitCoachingTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondError(c, bizerrors.ErrTaskNotFound)
		return
	}
	var req model.SubmitTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, bizerrors.ErrEmptySubmission)
		return
	}
	appointment, err := service.SubmitCoachingTask(taskID, req)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, appointment)
}

// ReviewCoachingTask 导师确认通过或驳回（驳回必须说明原因）。
func ReviewCoachingTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondError(c, bizerrors.ErrTaskNotFound)
		return
	}
	var req model.ReviewTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, bizerrors.ErrTaskNotAwaiting)
		return
	}
	appointment, err := service.ReviewCoachingTask(taskID, req)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, appointment)
}

// CoachingSkillWall 返回用户技能墙上由结项预约写入的记录。
func CoachingSkillWall(c *gin.Context) {
	c.JSON(http.StatusOK, service.CoachingSkillWall(c.Query("user")))
}
