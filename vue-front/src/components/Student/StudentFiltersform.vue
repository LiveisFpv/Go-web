<script setup lang="ts">
import type { Filter } from '@/types/meta';
import { ref } from 'vue';

interface StudentFilters {
  id_num_student: string;
  name_group: string;
  email_student: string;
  second_name_student: string;
  first_name_student: string;
  surname_student: string;
}

const emit = defineEmits<{
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const localFilters = ref<StudentFilters>({
  id_num_student: '',
  name_group: '',
  email_student: '',
  second_name_student: '',
  first_name_student: '',
  surname_student: '',
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
        <label>Номер билета</label>
        <input v-model="localFilters.id_num_student" placeholder="Введите номер" />
      </div>
      <div class="filter-group">
        <label>Группа</label>
        <input v-model="localFilters.name_group" placeholder="Введите группу" />
      </div>
      <div class="filter-group">
        <label>Почта</label>
        <input v-model="localFilters.email_student" placeholder="Введите почту" />
      </div>
      <div class="filter-group">
        <label>Фамилия</label>
        <input v-model="localFilters.second_name_student" placeholder="Введите фамилию" />
      </div>
      <div class="filter-group">
        <label>Имя</label>
        <input v-model="localFilters.first_name_student" placeholder="Введите имя" />
      </div>
      <div class="filter-group">
        <label>Отчество</label>
        <input v-model="localFilters.surname_student" placeholder="Введите отчество" />
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
