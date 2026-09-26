<template>
  <div class="panel panel--wide">
    <div class="coaching-header">
      <h2>阶段陪练</h2>
      <el-radio-group v-model="role" size="small">
        <el-radio-button v-for="option in COACHING_ROLE_OPTIONS" :key="option.value" :value="option.value">
          {{ option.label }}
        </el-radio-button>
      </el-radio-group>
    </div>

    <el-alert v-if="error" :title="error" type="error" show-icon :closable="false" />

    <div v-loading="loading">
      <el-empty v-if="!loading && appointments.length === 0" description="当前视角下暂无已确认的陪练预约" />
      <AppointmentCard
        v-for="appointment in appointments"
        :key="appointment.id"
        :appointment="appointment"
        :role="role"
        @add-task="openAddTask"
        @submit="openSubmit"
        @review="onReview"
      />
    </div>

    <section class="skill-wall">
      <h3>我的技能墙（结项写入）</h3>
      <el-empty v-if="skillWall.length === 0" description="预约结项后，成果会自动写入技能墙" :image-size="60" />
      <ul v-else class="skill-wall__list">
        <li v-for="entry in skillWall" :key="`${entry.skill}-${entry.closedAt}-${entry.partner}`">
          <el-tag type="success" effect="plain">{{ entry.skill }}</el-tag>
          <span>{{ entry.note }}</span>
          <small class="muted">与 {{ entry.partner }} · {{ entry.closedAt }}</small>
        </li>
      </ul>
    </section>

    <SubmitTaskDialog
      v-model:visible="submitVisible"
      :task="activeTask"
      :submitting="submitting"
      @confirm="onSubmitConfirm"
    />
    <AddTaskDialog v-model:visible="addVisible" :submitting="submitting" @confirm="onAddConfirm" />
    <RejectTaskDialog
      v-model:visible="rejectVisible"
      :task="activeTask"
      :submitting="submitting"
      @confirm="onRejectConfirm"
    />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import { COACHING_ROLE_OPTIONS, CURRENT_USER } from '../../constants/coaching.constants';
import {
  addCoachingTask,
  fetchCoachingAppointments,
  fetchSkillWall,
  reviewCoachingTask,
  submitCoachingTask,
} from '../../services/coaching.service';
import type { CoachingAppointment, CoachingRole, CoachingTask, SkillWallEntry } from '../../types/domain';
import AddTaskDialog from './AddTaskDialog.vue';
import AppointmentCard from './AppointmentCard.vue';
import RejectTaskDialog from './RejectTaskDialog.vue';
import SubmitTaskDialog from './SubmitTaskDialog.vue';

const role = ref<CoachingRole>('learner');
const appointments = ref<CoachingAppointment[]>([]);
const skillWall = ref<SkillWallEntry[]>([]);
const loading = ref(false);
const submitting = ref(false);
const error = ref('');

const activeTask = ref<CoachingTask | null>(null);
const activeAppointment = ref<CoachingAppointment | null>(null);
const submitVisible = ref(false);
const addVisible = ref(false);
const rejectVisible = ref(false);

async function reload() {
  loading.value = true;
  error.value = '';
  try {
    const [appointmentList, wall] = await Promise.all([
      fetchCoachingAppointments(CURRENT_USER, role.value),
      fetchSkillWall(CURRENT_USER),
    ]);
    appointments.value = appointmentList;
    skillWall.value = wall;
  } catch (err) {
    error.value = err instanceof Error ? err.message : '加载阶段陪练数据失败';
  } finally {
    loading.value = false;
  }
}

onMounted(reload);
watch(role, reload);

function openSubmit(task: CoachingTask) {
  activeTask.value = task;
  submitVisible.value = true;
}

function openAddTask(appointment: CoachingAppointment) {
  activeAppointment.value = appointment;
  addVisible.value = true;
}

async function run(action: () => Promise<unknown>, successMessage: string) {
  submitting.value = true;
  try {
    await action();
    ElMessage.success(successMessage);
    return true;
  } catch (err) {
    ElMessage.error(err instanceof Error ? err.message : '操作失败');
    return false;
  } finally {
    submitting.value = false;
  }
}

async function onSubmitConfirm(payload: { note: string; link: string }) {
  if (!activeTask.value) return;
  const ok = await run(
    () => submitCoachingTask(activeTask.value!.id, { user: CURRENT_USER, ...payload }),
    '已提交，等待导师确认',
  );
  if (ok) {
    submitVisible.value = false;
    await reload();
  }
}

async function onAddConfirm(payload: { title: string; requirement: string }) {
  if (!activeAppointment.value) return;
  const ok = await run(
    () => addCoachingTask(activeAppointment.value!.id, { user: CURRENT_USER, ...payload }),
    '已追加练习任务',
  );
  if (ok) {
    addVisible.value = false;
    await reload();
  }
}

async function onReview(task: CoachingTask, approve: boolean) {
  if (approve) {
    try {
      await ElMessageBox.confirm(`确认通过「${task.title}」吗？全部任务通过后预约将结项。`, '确认通过', {
        type: 'success',
        confirmButtonText: '通过',
        cancelButtonText: '再想想',
      });
    } catch {
      return;
    }
    const ok = await run(
      () => reviewCoachingTask(task.id, { user: CURRENT_USER, approve: true, reason: '' }),
      '已确认通过',
    );
    if (ok) await reload();
    return;
  }
  activeTask.value = task;
  rejectVisible.value = true;
}

async function onRejectConfirm(reason: string) {
  if (!activeTask.value) return;
  const ok = await run(
    () => reviewCoachingTask(activeTask.value!.id, { user: CURRENT_USER, approve: false, reason }),
    '已驳回并通知学员',
  );
  if (ok) {
    rejectVisible.value = false;
    await reload();
  }
}
</script>
