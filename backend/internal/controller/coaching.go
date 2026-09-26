package controller

import (
	"net/http"
	"strconv"

	berrors "cyskillswap/internal/errors"
	"cyskillswap/internal/service"
	"cyskillswap/internal/validator"
	"github.com/gin-gonic/gin"
)

// CoachingSessions 列出全部阶段陪练预约
func CoachingSessions(c *gin.Context) {
	c.JSON(http.StatusOK, service.CoachingSessions())
}

// CoachingSkillWall 按用户查询技能墙
func CoachingSkillWall(c *gin.Context) {
	c.JSON(http.StatusOK, service.SkillWall(c.Query("user")))
}

// AddCoachingTask 导师追加练习任务
func AddCoachingTask(c *gin.Context) {
	sessionID, ok := pathID(c, "id")
	if !ok {
		return
	}
	var req validator.AddTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, berrors.New(berrors.CodeInvalidRequest, "请求体格式不正确"))
		return
	}
	if err := validator.CheckAddTask(req); err != nil {
		respondError(c, err)
		return
	}
	session, err := service.AddCoachingTask(sessionID, req.Mentor, req.Title, req.Requirement)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, session)
}

// SubmitCoachingTask 学员提交作业
func SubmitCoachingTask(c *gin.Context) {
	sessionID, ok := pathID(c, "id")
	if !ok {
		return
	}
	taskID, ok := pathID(c, "taskId")
	if !ok {
		return
	}
	var req validator.SubmitTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, berrors.New(berrors.CodeInvalidRequest, "请求体格式不正确"))
		return
	}
	if err := validator.CheckSubmitTask(req); err != nil {
		respondError(c, err)
		return
	}
	session, err := service.SubmitCoachingTask(sessionID, taskID, req.Learner, req.Note, req.Link)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, session)
}

// ReviewCoachingTask 导师审核作业（通过 / 驳回）
func ReviewCoachingTask(c *gin.Context) {
	sessionID, ok := pathID(c, "id")
	if !ok {
		return
	}
	taskID, ok := pathID(c, "taskId")
	if !ok {
		return
	}
	var req validator.ReviewTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, berrors.New(berrors.CodeInvalidRequest, "请求体格式不正确"))
		return
	}
	if err := validator.CheckReviewTask(req); err != nil {
		respondError(c, err)
		return
	}
	session, err := service.ReviewCoachingTask(sessionID, taskID, req.Mentor, req.Action, req.Reason)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, session)
}

func pathID(c *gin.Context, name string) (int, bool) {
	id, err := strconv.Atoi(c.Param(name))
	if err != nil || id <= 0 {
		respondError(c, berrors.New(berrors.CodeInvalidRequest, "路径参数 ID 不合法"))
		return 0, false
	}
	return id, true
}

func respondError(c *gin.Context, err error) {
	if be, isBusiness := err.(berrors.BusinessError); isBusiness {
		c.JSON(berrors.HTTPStatus(be), be)
		return
	}
	c.JSON(http.StatusInternalServerError, berrors.New(berrors.CodeInternal, "服务器内部错误"))
}
