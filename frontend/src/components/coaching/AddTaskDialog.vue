<template>
  <el-dialog
    :model-value="visible"
    title="追加练习任务"
    width="480px"
    @update:model-value="$emit('update:visible', $event)"
    @open="reset"
  >
    <el-form label-position="top">
      <el-form-item label="任务标题" required>
        <el-input v-model="title" placeholder="例如：音阶每日练习" />
      </el-form-item>
      <el-form-item label="任务要求">
        <el-input v-model="requirement" type="textarea" :rows="3" placeholder="说明练习内容、验收标准和提交形式" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="$emit('update:visible', false)">取消</el-button>
      <el-button type="primary" :loading="submitting" @click="confirm">添加任务</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { ElMessage } from 'element-plus';

defineProps<{
  visible: boolean;
  submitting: boolean;
}>();

const emit = defineEmits<{
  'update:visible': [value: boolean];
  confirm: [payload: { title: string; requirement: string }];
}>();

const title = ref('');
const requirement = ref('');

function reset() {
  title.value = '';
  requirement.value = '';
}

function confirm() {
  if (!title.value.trim()) {
    ElMessage.warning('任务标题不能为空');
    return;
  }
  emit('confirm', { title: title.value.trim(), requirement: requirement.value.trim() });
}
</script>
