import { defineStore } from 'pinia';
import {
  addCoachingTask,
  fetchCoachingSessions,
  fetchSkillWall,
  reviewCoachingTask,
  submitCoachingTask,
} from '../services/coaching.service';
import { COACHING_USERS, SESSION_STATUS, TASK_STATUS } from '../constants/coaching.constants';
import type { CoachingSession, Perspective, SkillWallEntry, TodoItem } from '../types/coaching';

interface CoachingState {
  sessions: CoachingSession[];
  skillWall: SkillWallEntry[];
  currentUser: string;
  perspective: Perspective;
  loading: boolean;
}

export const useCoachingStore = defineStore('coaching', {
  state: (): CoachingState => ({
    sessions: [],
    skillWall: [],
    currentUser: COACHING_USERS[0],
    perspective: 'mentor',
    loading: false,
  }),
  getters: {
    visibleSessions(state): CoachingSession[] {
      return state.sessions.filter((session) =>
        state.perspective === 'mentor' ? session.mentor === state.currentUser : session.learner === state.currentUser,
      );
    },
    todos(): TodoItem[] {
      const items: TodoItem[] = [];
      for (const session of this.visibleSessions) {
        if (session.status === SESSION_STATUS.CLOSED) continue;
        if (this.perspective === 'mentor') {
          if (session.tasks.length === 0) {
            items.push({
              key: `empty-${session.id}`,
              text: `为「${session.skill}」（学员 ${session.learner}）布置第一个练习任务`,
              sessionId: session.id,
            });
          }
          for (const task of session.tasks) {
            if (task.status === TASK_STATUS.SUBMITTED) {
              items.push({
                key: `review-${task.id}`,
                text: `审核「${task.title}」— ${session.learner} 已提交，等待处理`,
                sessionId: session.id,
              });
            }
          }
        } else {
          for (const task of session.tasks) {
            if (task.status === TASK_STATUS.PENDING) {
              items.push({ key: `submit-${task.id}`, text: `提交「${task.title}」的作业`, sessionId: session.id });
            }
            if (task.status === TASK_STATUS.REJECTED) {
              items.push({
                key: `resubmit-${task.id}`,
                text: `「${task.title}」被驳回，修改后重新提交`,
                sessionId: session.id,
              });
            }
          }
        }
      }
      return items;
    },
  },
  actions: {
    async loadSessions() {
      this.loading = true;
      try {
        this.sessions = await fetchCoachingSessions();
      } finally {
        this.loading = false;
      }
    },
    async loadSkillWall() {
      this.skillWall = await fetchSkillWall(this.currentUser);
    },
    async reloadAll() {
      await Promise.all([this.loadSessions(), this.loadSkillWall()]);
    },
    async addTask(sessionId: number, title: string, requirement: string) {
      await addCoachingTask(sessionId, { mentor: this.currentUser, title, requirement });
      await this.loadSessions();
    },
    async submitTask(sessionId: number, taskId: number, note: string, link: string) {
      await submitCoachingTask(sessionId, taskId, { learner: this.currentUser, note, link });
      await this.loadSessions();
    },
    async reviewTask(sessionId: number, taskId: number, action: string, reason: string) {
      await reviewCoachingTask(sessionId, taskId, { mentor: this.currentUser, action, reason });
      await this.reloadAll();
    },
  },
});
