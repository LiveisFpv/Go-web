<script setup lang="ts">
import { ref } from 'vue';

interface SemesterFilters {
  size_budget: '',
  type_scholarship_budget: '',
  name_semester: ''
};

const emit = defineEmits<{
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const localFilters = ref<SemesterFilters>({
  size_budget: '',
  type_scholarship_budget: '',
  name_semester: ''
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
        <label for="size_budget">Размер бюджета:</label>
        <input
          type="number"
          id="size_budget"
          v-model="localFilters.size_budget"
          min="0"
          step="0.01"
        />
      </div>

      <div class="filter-group">
        <label for="type_scholarship_budget">Тип стипендии:</label>
        <input
          type="text"
          id="type_scholarship_budget"
          v-model="localFilters.type_scholarship_budget"
        />
      </div>

      <div class="filter-group">
        <label for="name_semester">Семестр:</label>
        <input
          type="text"
          id="name_semester"
          v-model="localFilters.name_semester"
        />
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
