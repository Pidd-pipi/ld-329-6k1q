<template>
  <section class="panel">
    <h2>{{ user }} 的技能墙</h2>
    <el-empty v-if="!entries.length" description="结项后的学习/指导成果会写入这里" :image-size="64" />
    <article v-for="entry in entries" :key="entry.id" class="feature-card">
      <div class="feature-card__title">
        <strong>{{ entry.title }}</strong>
        <el-tag :type="WALL_KIND_TAG[entry.kind] ?? 'info'" size="small">{{ entry.kind }}</el-tag>
      </div>
      <p>{{ entry.detail }}</p>
      <small>{{ entry.skill }} · 结项于 {{ entry.createdAt }}</small>
    </article>
  </section>
</template>

<script setup lang="ts">
import { WALL_KIND_TAG } from '../../../constants/coaching.constants';
import type { SkillWallEntry } from '../../../types/coaching';

defineProps<{ user: string; entries: SkillWallEntry[] }>();
</script>
