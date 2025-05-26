<script setup lang="ts">
import router from '@/router';
import { studentService } from '@/services/studentService';
import { categoryService } from '@/services/categoryService';
import { useAuthStore } from '@/stores/auth';
import type { StudentResp } from '@/types/student';
import type { AxiosError } from 'axios';
import { onMounted, ref } from 'vue';
import type { CategoryResp } from '@/types/category';

const emit = defineEmits<{
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

interface Filters {
  id_num_student: string;
  name_achivement: string;
  date_achivment: string;
  surname_student: string;
  first_name_student: string;
  second_name_student: string;
  achivments_type_category: string;
  name_group: string;
}

const localFilters = ref<Filters>({
  id_num_student: '',
  name_achivement: '',
  date_achivment: '',
  surname_student: '',
  first_name_student: '',
  second_name_student: '',
  achivments_type_category: '',
  name_group: ''
});

const apply = () => {
  // Only include non-empty filters
  const activeFilters = Object.entries(localFilters.value)
    .reduce((acc, [key, value]) => {
      if (value !== '') {
        acc[key] = String(value);
      }
      return acc;
    }, {} as Record<string, string>);

  emit('update-filters', activeFilters);
};

const clearFilters = () => {
  Object.keys(localFilters.value).forEach(key => {
    (localFilters.value as any)[key] = key === 'id_achivment' || key === 'id_num_student' || key === 'id_category' ? 0 : '';
  });
  emit('update-filters', {});
};

const students = ref<StudentResp[]>([]);
const categories = ref<CategoryResp[]>([]);
const authStore = useAuthStore();
const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

onMounted(async () => {
  if (!checkAuth()) return;

  // Загрузка студентов
  if (students.value.length === 0) {
    try {
      const response = await studentService.getStudents(1, 1000);
      students.value = response.data;
    } catch (err) {
      const axiosError = err as AxiosError;
      if (axiosError.response?.status === 401) {
        authStore.logout();
        router.push('/auth');
      } else {
        console.error('Failed to fetch students:', axiosError);
      }
    }
  }
  if (categories.value.length > 0) return;
  // Загрузка категорий
  try {
    const response = await categoryService.getCategories(1, 100);
    categories.value = response.data; // Сохраняем категории
  } catch (err) {
    console.error('Failed to fetch categories:', err);
  }
});

</script>

<template>
  <div class="filters-form">
    <div class="filters-grid">
      <div class="filter-item">
        <label for="id_num_student">Номер билета:</label>
        <input
          type="text"
          id="id_num_student"
          v-model="localFilters.id_num_student"
        />
      </div>

      <div class="filter-item">
        <label for="name_achivement">Название достижения:</label>
        <input
          type="text"
          id="name_achivement"
          v-model="localFilters.name_achivement"
        />
      </div>

      <div class="filter-item">
        <label for="date_achivment">Дата достижения:</label>
        <input
          type="date"
          id="date_achivment"
          v-model="localFilters.date_achivment"
        />
      </div>

      <div class="filter-item">
        <label for="second_name_student">Фамилия:</label>
        <input
          type="text"
          id="second_name_student"
          v-model="localFilters.second_name_student"
        />
      </div>

      <div class="filter-item">
        <label for="first_name_student">Имя:</label>
        <input
          type="text"
          id="first_name_student"
          v-model="localFilters.first_name_student"
        />
      </div>

      <div class="filter-item">
        <label for="surname_student">Отчество:</label>
        <input
          type="text"
          id="surname_student"
          v-model="localFilters.surname_student"
        />
      </div>

      <div class="filter-item">
        <label for="name_group">Группа:</label>
        <input
          type="text"
          id="name_group"
          v-model="localFilters.name_group"
        />
      </div>

      <div class="filter-item">
        <label for="achivments_type_category">Тип достижения:</label>
        <select
          id="achivments_type_category"
          v-model="localFilters.achivments_type_category"
        >
          <option value="">Выберите тип</option>
          <option
            v-for="type in categories"
            :key="type.id_category"
            :value="type.achivments_type_category"
          >
            {{ type.achivments_type_category }}
          </option>
        </select>
      </div>
    </div>

    <div class="filters-actions">
      <button @click="apply" class="apply-btn">Применить</button>
      <button @click="clearFilters" class="clear-button">Очистить фильтры</button>
    </div>
  </div>
</template>

<style scoped>
@import '../../assets/filter.css';
</style>
