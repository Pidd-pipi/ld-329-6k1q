<template>
  <main class="page-shell" v-loading="store.loading">
    <PerspectiveBar
      :current-user="store.currentUser"
      :perspective="store.perspective"
      @update:current-user="switchUser"
      @update:perspective="(value) => (store.perspective = value as Perspective)"
    >
      <el-button :loading="store.loading" @click="reload">刷新</el-button>
    </PerspectiveBar>

    <el-alert v-if="error" :title="error" type="error" show-icon class="coaching-error" />

    <div class="coaching-grid">
      <div class="coaching-grid__main">
        <TodoPanel :todos="store.todos" />
        <el-empty v-if="!store.visibleSessions.length" description="当前视角下没有进行中的陪练预约" />
        <SessionCard
          v-for="session in store.visibleSessions"
          :key="session.id"
          :session="session"
          :perspective="store.perspective"
          @add-task="openAddTask"
          @submit="openSubmit"
          @review="openReview"
        />
      </div>
      <SkillWallPanel :user="store.currentUser" :entries="store.skillWall" />
    </div>

    <TaskFormDialog v-model="taskDialogVisible" @confirm="handleAddTask" />
    <SubmitDialog v-model="submitDialogVisible" :task="activeTask" @confirm="handleSubmit" />
    <ReviewDialog v-model="reviewDialogVisible" :task="activeTask" @confirm="handleReview" />
  </main>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { ElMessage } from 'element-plus';
import PerspectiveBar from './components/PerspectiveBar.vue';
import TodoPanel from './components/TodoPanel.vue';
import SessionCard from './components/SessionCard.vue';
import SkillWallPanel from './components/SkillWallPanel.vue';
import TaskFormDialog from './components/TaskFormDialog.vue';
import SubmitDialog from './components/SubmitDialog.vue';
import ReviewDialog from './components/ReviewDialog.vue';
import { useCoachingStore } from '../../stores/coaching.store';
import { REVIEW_ACTION, SESSION_STATUS } from '../../constants/coaching.constants';
import type { CoachingSession, Perspective, PracticeTask } from '../../types/coaching';

const store = useCoachingStore();
const error = ref('');

const taskDialogVisible = ref(false);
const submitDialogVisible = ref(false);
const reviewDialogVisible = ref(false);
const activeSession = ref<CoachingSession | null>(null);
const activeTask = ref<PracticeTask | null>(null);

onMounted(load);

async function load() {
  error.value = '';
  try {
    await store.reloadAll();
  } catch (err) {
    error.value = err instanceof Error ? err.message : '加载失败';
  }
}

async function reload() {
  await load();
}

async function switchUser(user: string) {
  store.currentUser = user;
  await load();
}

function openAddTask(session: CoachingSession) {
  activeSession.value = session;
  taskDialogVisible.value = true;
}

function openSubmit(session: CoachingSession, task: PracticeTask) {
  activeSession.value = session;
  activeTask.value = task;
  submitDialogVisible.value = true;
}

function openReview(session: CoachingSession, task: PracticeTask) {
  activeSession.value = session;
  activeTask.value = task;
  reviewDialogVisible.value = true;
}

async function runAction(action: () => Promise<void>, successText: string, close: () => void) {
  try {
    await action();
    close();
    ElMessage.success(successText);
    if (activeSession.value?.status === SESSION_STATUS.OPEN) {
      const closed = store.sessions.find((s) => s.id === activeSession.value?.id && s.status === SESSION_STATUS.CLOSED);
      if (closed) ElMessage.success('全部任务已通过，预约自动结项并写入双方技能墙');
    }
  } catch (err) {
    ElMessage.error(err instanceof Error ? err.message : '操作失败');
  }
}

async function handleAddTask(payload: { title: string; requirement: string }) {
  const session = activeSession.value;
  if (!session) return;
  await runAction(
    () => store.addTask(session.id, payload.title, payload.requirement),
    '练习任务已布置',
    () => (taskDialogVisible.value = false),
  );
}

async function handleSubmit(payload: { note: string; link: string }) {
  const session = activeSession.value;
  const task = activeTask.value;
  if (!session || !task) return;
  await runAction(
    () => store.submitTask(session.id, task.id, payload.note, payload.link),
    '作业已提交，等待导师处理',
    () => (submitDialogVisible.value = false),
  );
}

async function handleReview(payload: { action: string; reason: string }) {
  const session = activeSession.value;
  const task = activeTask.value;
  if (!session || !task) return;
  await runAction(
    () => store.reviewTask(session.id, task.id, payload.action, payload.reason),
    payload.action === REVIEW_ACTION.REJECT ? '已驳回并通知学员修改' : '已通过该任务',
    () => (reviewDialogVisible.value = false),
  );
}
</script>
