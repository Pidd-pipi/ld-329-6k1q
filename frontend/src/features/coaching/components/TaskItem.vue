<template>
  <article class="task-item">
    <div class="task-item__head">
      <strong>{{ task.title }}</strong>
      <el-tag :type="TASK_STATUS_TAG[task.status] ?? 'info'" size="small">{{ task.status }}</el-tag>
    </div>
    <p class="muted">{{ task.requirement }}</p>

    <div v-if="task.submission" class="task-item__submission">
      <span>提交说明：{{ task.submission.note || '（未填写）' }}</span>
      <el-link v-if="task.submission.link" :href="task.submission.link" target="_blank" type="primary">
        作品链接
      </el-link>
      <small>{{ task.submission.submittedAt }}</small>
    </div>

    <el-alert
      v-if="task.status === TASK_STATUS.REJECTED && task.rejectReason"
      :title="`驳回原因：${task.rejectReason}`"
      type="error"
      :closable="false"
      show-icon
    />

    <el-collapse class="task-item__history">
      <el-collapse-item title="任务历史" name="history">
        <el-timeline>
          <el-timeline-item v-for="(event, index) in task.history" :key="index" :timestamp="event.at">
            {{ event.actor }} · {{ event.action }}<span v-if="event.detail">：{{ event.detail }}</span>
          </el-timeline-item>
        </el-timeline>
      </el-collapse-item>
    </el-collapse>

    <div class="task-item__actions">
      <template v-if="!sessionClosed && perspective === 'learner'">
        <el-button v-if="task.status === TASK_STATUS.PENDING" type="primary" size="small" @click="emit('submit', task)">
          提交作业
        </el-button>
        <el-button v-else-if="task.status === TASK_STATUS.REJECTED" type="warning" size="small" @click="emit('submit', task)">
          修改后重新提交
        </el-button>
        <span v-else-if="task.status === TASK_STATUS.SUBMITTED" class="muted">等待导师处理，暂不能重复提交</span>
      </template>
      <el-button
        v-if="!sessionClosed && perspective === 'mentor' && task.status === TASK_STATUS.SUBMITTED"
        type="primary"
        size="small"
        @click="emit('review', task)"
      >
        审核作业
      </el-button>
    </div>
  </article>
</template>

<script setup lang="ts">
import { TASK_STATUS, TASK_STATUS_TAG } from '../../../constants/coaching.constants';
import type { PracticeTask } from '../../../types/coaching';

defineProps<{ task: PracticeTask; perspective: string; sessionClosed: boolean }>();
const emit = defineEmits<{
  submit: [task: PracticeTask];
  review: [task: PracticeTask];
}>();
</script>
