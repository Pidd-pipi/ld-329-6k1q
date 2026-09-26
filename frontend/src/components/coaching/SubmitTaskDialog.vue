<template>
  <el-dialog
    :model-value="visible"
    :title="task ? `提交练习：${task.title}` : '提交练习'"
    width="480px"
    @update:model-value="$emit('update:visible', $event)"
    @open="reset"
  >
    <el-form label-position="top">
      <el-form-item label="练习说明">
        <el-input v-model="note" type="textarea" :rows="3" placeholder="说明练习过程、遇到的问题或成果" />
      </el-form-item>
      <el-form-item label="作品链接">
        <el-input v-model="link" placeholder="视频 / 音频 / 相册 / 仓库链接（可和说明二选一）" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="$emit('update:visible', false)">取消</el-button>
      <el-button type="primary" :loading="submitting" @click="confirm">提交</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import type { CoachingTask } from '../../types/domain';

const props = defineProps<{
  visible: boolean;
  task: CoachingTask | null;
  submitting: boolean;
}>();

const emit = defineEmits<{
  'update:visible': [value: boolean];
  confirm: [payload: { note: string; link: string }];
}>();

const note = ref('');
const link = ref('');

// 重新提交时预填上次内容，方便学员在驳回后修改。
function reset() {
  note.value = props.task?.submission?.note ?? '';
  link.value = props.task?.submission?.link ?? '';
}

function confirm() {
  if (!note.value.trim() && !link.value.trim()) {
    ElMessage.warning('请填写练习说明或作品链接');
    return;
  }
  emit('confirm', { note: note.value.trim(), link: link.value.trim() });
}
</script>
