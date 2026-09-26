<template>
  <el-dialog
    :model-value="visible"
    :title="task ? `驳回任务：${task.title}` : '驳回任务'"
    width="480px"
    @update:model-value="$emit('update:visible', $event)"
    @open="reason = ''"
  >
    <el-form label-position="top">
      <el-form-item label="驳回原因" required>
        <el-input
          v-model="reason"
          type="textarea"
          :rows="3"
          placeholder="说明哪里不达标、建议如何修改，学员改完后可重新提交"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="$emit('update:visible', false)">取消</el-button>
      <el-button type="danger" :loading="submitting" @click="confirm">确认驳回</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import type { CoachingTask } from '../../types/domain';

defineProps<{
  visible: boolean;
  task: CoachingTask | null;
  submitting: boolean;
}>();

const emit = defineEmits<{
  'update:visible': [value: boolean];
  confirm: [reason: string];
}>();

const reason = ref('');

function confirm() {
  if (!reason.value.trim()) {
    ElMessage.warning('驳回必须说明原因');
    return;
  }
  emit('confirm', reason.value.trim());
}
</script>
