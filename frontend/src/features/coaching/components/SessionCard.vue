<template>
  <section class="panel session-card">
    <div class="session-card__head">
      <div>
        <h2>{{ session.skill }}</h2>
        <p class="muted">
          导师 {{ session.mentor }} → 学员 {{ session.learner }} · {{ session.time }} · {{ session.place }}
        </p>
      </div>
      <el-tag :type="closed ? 'success' : 'primary'" effect="dark">
        {{ session.status }}<template v-if="closed && session.closedAt"> · {{ session.closedAt }}</template>
      </el-tag>
    </div>

    <div v-if="session.tasks.length" class="session-card__progress">
      <span>任务进度 {{ approvedCount }}/{{ session.tasks.length }} 已通过</span>
      <el-progress :percentage="progressPercent" :status="closed ? 'success' : undefined" />
    </div>

    <el-alert
      v-if="closed"
      title="全部任务已通过，预约已结项。结项成果已写入双方技能墙，历史记录不可改动。"
      type="success"
      :closable="false"
      show-icon
    />
    <el-empty v-else-if="!session.tasks.length" description="导师还没有布置练习任务" :image-size="64" />

    <TaskItem
      v-for="task in session.tasks"
      :key="task.id"
      :task="task"
      :perspective="perspective"
      :session-closed="closed"
      @submit="(t) => emit('submit', session, t)"
      @review="(t) => emit('review', session, t)"
    />

    <div v-if="!closed && perspective === 'mentor'" class="session-card__footer">
      <el-button type="primary" plain @click="emit('addTask', session)">添加练习任务</el-button>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import TaskItem from './TaskItem.vue';
import { SESSION_STATUS, TASK_STATUS } from '../../../constants/coaching.constants';
import type { CoachingSession, PracticeTask } from '../../../types/coaching';

const props = defineProps<{ session: CoachingSession; perspective: string }>();
const emit = defineEmits<{
  addTask: [session: CoachingSession];
  submit: [session: CoachingSession, task: PracticeTask];
  review: [session: CoachingSession, task: PracticeTask];
}>();

const closed = computed(() => props.session.status === SESSION_STATUS.CLOSED);
const approvedCount = computed(
  () => props.session.tasks.filter((task) => task.status === TASK_STATUS.APPROVED).length,
);
const progressPercent = computed(() =>
  props.session.tasks.length ? Math.round((approvedCount.value / props.session.tasks.length) * 100) : 0,
);
</script>
