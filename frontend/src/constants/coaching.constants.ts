import type { CoachingRole, CoachingTaskStatus } from '../types/domain';

// 演示账号：与后端种子数据中的当前用户一致。
export const CURRENT_USER = '林澈';

export const COACHING_ROLE_OPTIONS: Array<{ value: CoachingRole; label: string }> = [
  { value: 'mentor', label: '导师视角' },
  { value: 'learner', label: '学员视角' },
];

export const TASK_STATUS_META: Record<CoachingTaskStatus, { label: string; tagType: 'info' | 'warning' | 'success' | 'danger' }> = {
  todo: { label: '待提交', tagType: 'info' },
  submitted: { label: '待审核', tagType: 'warning' },
  approved: { label: '已完成', tagType: 'success' },
  rejected: { label: '已驳回', tagType: 'danger' },
};

export const APPOINTMENT_STATUS_META = {
  open: { label: '进行中', tagType: 'warning' as const },
  closed: { label: '已结项', tagType: 'success' as const },
};
