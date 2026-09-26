<template>
  <el-dialog
    :model-value="modelValue"
    :title="task ? `提交作业：${task.title}` : '提交作业'"
    width="480px"
    @update:model-value="(value: boolean) => emit('update:modelValue', value)"
    @open="reset"
  >
    <el-alert
      v-if="task?.status === TASK_STATUS.REJECTED && task.rejectReason"
      :title="`导师驳回原因：${task.rejectReason}`"
      type="error"
      :closable="false"
      show-icon
      class="submit-dialog__reject"
    />
    <el-form label-position="top">
      <el-form-item label="提交说明">
        <el-input v-model="note" type="textarea" :rows="3" placeholder="说明完成情况、遇到的问题等" maxlength="300" />
      </el-form-item>
      <el-form-item label="作品链接">
        <el-input v-model="link" placeholder="https://…（作品集、视频或仓库地址）" maxlength="200" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="emit('update:modelValue', false)">取消</el-button>
      <el-button type="primary" :disabled="!note.trim() && !link.trim()" @click="confirm">
        {{ task?.status === TASK_STATUS.REJECTED ? '重新提交' : '提交' }}
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { TASK_STATUS } from '../../../constants/coaching.constants';
import type { PracticeTask } from '../../../types/coaching';

defineProps<{ modelValue: boolean; task: PracticeTask | null }>();
const emit = defineEmits<{
  'update:modelValue': [value: boolean];
  confirm: [payload: { note: string; link: string }];
}>();

const note = ref('');
const link = ref('');

function reset() {
  note.value = '';
  link.value = '';
}

function confirm() {
  emit('confirm', { note: note.value.trim(), link: link.value.trim() });
}
</script>
