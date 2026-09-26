<template>
  <el-dialog
    :model-value="modelValue"
    title="添加练习任务"
    width="480px"
    @update:model-value="(value: boolean) => emit('update:modelValue', value)"
    @open="reset"
  >
    <el-form label-position="top">
      <el-form-item label="任务标题" required>
        <el-input v-model="title" placeholder="例如：人像修图作业" maxlength="40" />
      </el-form-item>
      <el-form-item label="任务要求">
        <el-input v-model="requirement" type="textarea" :rows="3" placeholder="说明完成标准和交付形式" maxlength="200" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="emit('update:modelValue', false)">取消</el-button>
      <el-button type="primary" :disabled="!title.trim()" @click="confirm">布置任务</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';

defineProps<{ modelValue: boolean }>();
const emit = defineEmits<{
  'update:modelValue': [value: boolean];
  confirm: [payload: { title: string; requirement: string }];
}>();

const title = ref('');
const requirement = ref('');

function reset() {
  title.value = '';
  requirement.value = '';
}

function confirm() {
  emit('confirm', { title: title.value.trim(), requirement: requirement.value.trim() });
}
</script>
