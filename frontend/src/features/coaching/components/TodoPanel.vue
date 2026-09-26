<template>
  <section class="panel todo-panel">
    <h2>
      待办事项
      <el-badge v-if="todos.length" :value="todos.length" type="danger" />
    </h2>
    <el-empty v-if="!todos.length" description="当前视角没有待办事项" :image-size="64" />
    <ul v-else class="todo-list">
      <li v-for="todo in todos" :key="todo.key">
        <el-tag size="small" effect="plain">待办</el-tag>
        <span>{{ todo.text }}</span>
      </li>
    </ul>
  </section>
</template>

<script setup lang="ts">
import type { TodoItem } from '../../../types/coaching';

defineProps<{ todos: TodoItem[] }>();
</script>
