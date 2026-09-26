<template>
  <li class="task-item">
    <div class="task-item__head">
      <strong>{{ task.title }}</strong>
      <el-tag :type="statusMeta.tagType" size="small">{{ statusMeta.label }}</el-tag>
    </div>
    <p class="muted">{{ task.requirement }}</p>

    <div v-if="task.submission" class="task-submission">
      <span v-if="task.submission.note">说明：{{ task.submission.note }}</span>
      <el-link v-if="task.submission.link" :href="task.submission.link" target="_blank" type="primary">
        作品链接
      </el-link>
      <small class="muted">提交于 {{ task.submission.submittedAt }}</small>
    </div>

    <el-alert
      v-if="task.status === 'rejected' && task.rejectReason"
      :title="`导师驳回：${task.rejectReason}`"
      type="error"
      :closable="false"
      show-icon
    />

    <div class="task-actions">
      <template v-if="role === 'learner'">
        <el-button
          v-if="canSubmit"
          size="small"
          type="primary"
          @click="$emit('submit', task)"
        >
          {{ task.status === 'rejected' ? '修改后重新提交' : '提交练习' }}
        </el-button>
        <span v-else-if="task.status === 'submitted'" class="muted">等待导师处理，暂不能重复提交</span>
      </template>
      <template v-else-if="role === 'mentor' && task.status === 'submitted'">
        <el-button size="small" type="success" @click="$emit('review', task, true)">确认通过</el-button>
        <el-button size="small" type="danger" plain @click="$emit('review', task, false)">驳回</el-button>
      </template>
    </div>
  </li>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { TASK_STATUS_META } from '../../constants/coaching.constants';
import type { CoachingRole, CoachingTask } from '../../types/domain';

const props = defineProps<{
  task: CoachingTask;
  role: CoachingRole;
  closed: boolean;
}>();

defineEmits<{
  submit: [task: CoachingTask];
  review: [task: CoachingTask, approve: boolean];
}>();

const statusMeta = computed(() => TASK_STATUS_META[props.task.status]);
// 学员只能在待提交或已驳回时提交；等待审核和已结项都不可改动。
const canSubmit = computed(
  () => !props.closed && (props.task.status === 'todo' || props.task.status === 'rejected'),
);
</script>
