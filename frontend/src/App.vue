<template>
  <nav class="app-nav">
    <el-radio-group v-model="view">
      <el-radio-button value="dashboard">平台总览</el-radio-button>
      <el-radio-button value="coaching">阶段陪练</el-radio-button>
    </el-radio-group>
  </nav>
  <DashboardView v-if="view === 'dashboard'" />
  <CoachingView v-else />
</template>

<script setup lang="ts">
import { ref } from 'vue';
import DashboardView from './features/DashboardView.vue';
import CoachingView from './features/coaching/CoachingView.vue';

const view = ref<'dashboard' | 'coaching'>('coaching');
</script>
