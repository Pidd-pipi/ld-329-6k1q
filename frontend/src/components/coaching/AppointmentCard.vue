<template>
  <article class="feature-card coaching-card">
    <div class="feature-card__title">
      <strong>{{ appointment.skill }}</strong>
      <el-tag :type="statusMeta.tagType" size="small">{{ statusMeta.label }}</el-tag>
    </div>
    <small class="muted">
      {{ partnerLabel }} · {{ appointment.time }} · {{ appointment.place }}
      <template v-if="appointment.closed"> · 结项于 {{ appointment.closedAt }}</template>
    </small>

    <div class="coaching-progress">
      <el-progress :percentage="progressPercent" :status="appointment.closed ? 'success' : undefined" />
      <span class="muted">待办 {{ pendingCount }} 项 · 已完成 {{ doneCount }}/{{ appointment.tasks.length }}</span>
    </div>

    <ul class="task-list">
      <TaskItem
        v-for="task in appointment.tasks"
        :key="task.id"
        :task="task"
        :role="role"
        :closed="appointment.closed"
        @submit="$emit('submit', $event)"
        @review="(task, approve) => $emit('review', task, approve)"
      />
    </ul>

    <el-button
      v-if="role === 'mentor'"
      size="small"
      :disabled="appointment.closed"
      @click="$emit('add-task', appointment)"
    >
      {{ appointment.closed ? '已结项，不能追加任务' : '追加练习任务' }}
    </el-button>
  </article>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { APPOINTMENT_STATUS_META } from '../../constants/coaching.constants';
import type { CoachingAppointment, CoachingRole, CoachingTask } from '../../types/domain';
import TaskItem from './TaskItem.vue';

const props = defineProps<{
  appointment: CoachingAppointment;
  role: CoachingRole;
}>();

defineEmits<{
  'add-task': [appointment: CoachingAppointment];
  submit: [task: CoachingTask];
  review: [task: CoachingTask, approve: boolean];
}>();

const statusMeta = computed(() =>
  props.appointment.closed ? APPOINTMENT_STATUS_META.closed : APPOINTMENT_STATUS_META.open,
);

const partnerLabel = computed(() =>
  props.role === 'mentor' ? `学员：${props.appointment.learner}` : `导师：${props.appointment.mentor}`,
);

const doneCount = computed(
  () => props.appointment.tasks.filter((task) => task.status === 'approved').length,
);

const pendingCount = computed(
  () => props.appointment.tasks.filter((task) => task.status === 'todo' || task.status === 'rejected').length,
);

const progressPercent = computed(() => {
  const total = props.appointment.tasks.length;
  return total === 0 ? 0 : Math.round((doneCount.value / total) * 100);
});
</script>
