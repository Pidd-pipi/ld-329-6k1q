<template>
  <el-dialog
    :model-value="modelValue"
    :title="task ? `审核作业：${task.title}` : '审核作业'"
    width="480px"
    @update:model-value="(value: boolean) => emit('update:modelValue', value)"
    @open="reset"
  >
    <div v-if="task?.submission" class="review-dialog__submission">
      <p>提交说明：{{ task.submission.note || '（未填写）' }}</p>
      <el-link v-if="task.submission.link" :href="task.submission.link" target="_blank" type="primary">
        查看作品链接
      </el-link>
      <small class="muted">提交于 {{ task.submission.submittedAt }}</small>
    </div>
    <el-form label-position="top">
      <el-form-item label="审核结果" required>
        <el-radio-group v-model="action">
          <el-radio-button :value="REVIEW_ACTION.APPROVE">通过</el-radio-button>
          <el-radio-button :value="REVIEW_ACTION.REJECT">驳回</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item v-if="action === REVIEW_ACTION.REJECT" label="驳回原因（必填，学员修改后可重新提交）" required>
        <el-input v-model="reason" type="textarea" :rows="3" placeholder="说明不通过的原因和修改方向" maxlength="200" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="emit('update:modelValue', false)">取消</el-button>
      <el-button type="primary" :disabled="action === REVIEW_ACTION.REJECT && !reason.trim()" @click="confirm">
        提交审核
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { REVIEW_ACTION } from '../../../constants/coaching.constants';
import type { PracticeTask } from '../../../types/coaching';

defineProps<{ modelValue: boolean; task: PracticeTask | null }>();
const emit = defineEmits<{
  'update:modelValue': [value: boolean];
  confirm: [payload: { action: string; reason: string }];
}>();

const action = ref<string>(REVIEW_ACTION.APPROVE);
const reason = ref('');

function reset() {
  action.value = REVIEW_ACTION.APPROVE;
  reason.value = '';
}

function confirm() {
  emit('confirm', { action: action.value, reason: reason.value.trim() });
}
</script>
