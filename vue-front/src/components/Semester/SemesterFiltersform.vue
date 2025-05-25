<script setup lang="ts">
import { ref, watch } from 'vue';

interface SemesterFilters {
  name_semester: string;
  date_start_semester: string;
  date_end_semester: string;
}

const emit = defineEmits<{
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const localFilters = ref<SemesterFilters>({
  name_semester: '',
  date_start_semester: '',
  date_end_semester: '',
});

const apply = () => {
  // Only include non-empty filters
  const activeFilters = Object.entries(localFilters.value)
    .reduce((acc, [key, value]) => {
      if (value.trim()) {
        acc[key] = value.trim();
      }
      return acc;
    }, {} as Record<string, string>);

  emit('update-filters', activeFilters);
};

const clearFilters = () => {
  Object.keys(localFilters.value).forEach(key => {
    (localFilters.value as any)[key] = '';
  });
  emit('update-filters', {});
};
</script>

<template>
  <div class="filters-form">
    <div class="filters-grid">
      <div class="filter-group">
        <label>Название</label>
        <input v-model="localFilters.name_semester" placeholder="Введите название" />
      </div>
      <div class="filter-group">
        <label>Дата начала</label>
        <input v-model="localFilters.date_start_semester" type="date" />
      </div>
      <div class="filter-group">
        <label>Дата окончания</label>
        <input v-model="localFilters.date_end_semester" type="date" />
      </div>
    </div>
    <div class="filters-actions">
      <button @click="apply" class="apply-btn">Применить</button>
      <button @click="clearFilters" class="clear-btn">Очистить</button>
    </div>
  </div>
</template>

<style scoped>
@import '../../assets/filter.css';
</style>
