<template>
  <div class="perspective-bar">
    <div class="perspective-bar__field">
      <span>当前身份</span>
      <el-select
        :model-value="currentUser"
        style="width: 140px"
        @update:model-value="(value: string) => emit('update:currentUser', value)"
      >
        <el-option v-for="user in COACHING_USERS" :key="user" :label="user" :value="user" />
      </el-select>
    </div>
    <el-radio-group
      :model-value="perspective"
      @update:model-value="(value: string | number | boolean | undefined) => emit('update:perspective', String(value))"
    >
      <el-radio-button v-for="option in PERSPECTIVE_OPTIONS" :key="option.value" :value="option.value">
        {{ option.label }}
      </el-radio-button>
    </el-radio-group>
    <slot />
  </div>
</template>

<script setup lang="ts">
import { COACHING_USERS, PERSPECTIVE_OPTIONS } from '../../../constants/coaching.constants';

defineProps<{ currentUser: string; perspective: string }>();
const emit = defineEmits<{
  'update:currentUser': [value: string];
  'update:perspective': [value: string];
}>();
</script>
