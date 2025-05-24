<script setup lang="ts">
import { ref } from 'vue';

interface GroupFilters{
  name_group: string
  studies_direction_group: string
  studies_profile_group: string
  start_date_group: string
  studies_period_group: string
}

const emit = defineEmits<{
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const localFilters = ref<GroupFilters>({
  name_group: '',
  studies_direction_group: '',
  studies_profile_group: '',
  start_date_group: '',
  studies_period_group: '',
});

const applyFilter = () =>{
  const activeFilters = Object.entries(localFilters.value)
  .reduce((acc, [key, value]) => {
    if (value.trim()) {
      acc[key] = value.trim();
    }
    return acc;
  }, {} as Record<string, string>);
  emit('update-filters', activeFilters);
}
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
        <label>Название группы</label>
        <input v-model="localFilters.name_group" placeholder="Введите название группы" />
      </div>
      <div class="filter-group">
        <label>Направление обучения</label>
        <input v-model="localFilters.studies_direction_group" placeholder="Введите направление обучения" />
      </div>
      <div class="filter-group">
        <label>Профиль обучения</label>
        <input v-model="localFilters.studies_profile_group" placeholder="Введите профиль обучения" />
      </div>
      <div class="filter-group">
        <label>Дата начала обучения</label>
        <input v-model="localFilters.start_date_group" type="date" placeholder="Выберите дату начала" />
      </div>
      <div class="filter-group">
        <label>Срок обучения</label>
        <input v-model="localFilters.studies_period_group" placeholder="Введите срок обучения" />
      </div>
    </div>
    <div class="filters-actions">
      <button @click="applyFilter" class="apply-btn">Применить</button>
      <button @click="clearFilters" class="clear-btn">Сбросить</button>
    </div>
  </div>
</template>

<style scoped>
@import '../../assets/filter.css';
</style>
