<script setup lang="ts">
import { ref } from 'vue';


interface CategoryFilter{
  achivments_type_category: string;
  score_category: number;
}

const emit = defineEmits<{
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const localFilters = ref<CategoryFilter>({
  achivments_type_category: '',
  score_category: 0,
});

const apply = () => {
  // Only include non-empty filters
  const activeFilters = Object.entries(localFilters.value)
    .reduce((acc, [key, value]) => {
      if (typeof value === 'string' && value.trim()) {
        acc[key] = value.trim();
      } else if (typeof value === 'number' && value !== 0) {
        acc[key] = value.toString();
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
        <label>Тип достижения</label>
        <input v-model="localFilters.achivments_type_category" placeholder="Введите тип достижения" />
      </div>
      <div class="filter-group">
        <label>Балл</label>
        <input v-model.number="localFilters.score_category" type="number" placeholder="Введите балл" />
      </div>
    </div>
  </div>
</template>
<style scoped>
@import '../../assets/filter.css';
</style>
