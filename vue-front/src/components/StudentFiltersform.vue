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
.filters-form {
  display: flex;
  flex-direction: column;
  gap: 15px;
  padding: 15px;
}

.filters-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
}

.filter-group {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.filter-group label {
  font-size: 0.85rem;
  color: var(--color-text-light);
  font-weight: 500;
}

.filters-form input {
  padding: 8px;
  border-radius: 4px;
  border: 1px solid #ddd;
  font-size: 0.9rem;
  transition: all 0.2s ease;
}

.filters-form input:focus {
  border-color: var(--color-accent);
  outline: none;
  box-shadow: 0 0 0 2px rgba(var(--color-accent-rgb), 0.2);
}

.filters-actions {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
  margin-top: 5px;
}

.filters-form button {
  padding: 8px 16px;
  border-radius: 4px;
  border: none;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.2s ease;
}

.apply-btn {
  background-color: var(--color-accent);
  color: white;
}

.clear-btn {
  background-color: #f0f0f0;
  color: #666;
}

.filters-form button:hover {
  transform: translateY(-1px);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.filters-form button:active {
  transform: translateY(0);
}
</style>
