export const TASK_STATUS = {
  PENDING: '待提交',
  SUBMITTED: '待审核',
  APPROVED: '已通过',
  REJECTED: '已驳回',
} as const;

export const SESSION_STATUS = {
  OPEN: '进行中',
  CLOSED: '已结项',
} as const;

export const REVIEW_ACTION = {
  APPROVE: 'approve',
  REJECT: 'reject',
} as const;

export const TASK_STATUS_TAG: Record<string, 'info' | 'warning' | 'success' | 'danger'> = {
  [TASK_STATUS.PENDING]: 'info',
  [TASK_STATUS.SUBMITTED]: 'warning',
  [TASK_STATUS.APPROVED]: 'success',
  [TASK_STATUS.REJECTED]: 'danger',
};

export const PERSPECTIVE_OPTIONS = [
  { value: 'mentor', label: '导师视角' },
  { value: 'learner', label: '学员视角' },
] as const;

export const COACHING_USERS = ['林澈', '孟野', '周芮', '许安'] as const;

export const WALL_KIND_TAG: Record<string, 'warning' | 'success'> = {
  指导成果: 'warning',
  学习成果: 'success',
};
