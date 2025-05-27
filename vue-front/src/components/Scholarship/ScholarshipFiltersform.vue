<script setup lang="ts">
import router from '@/router';
import { studentService } from '@/services/studentService';
import { useAuthStore } from '@/stores/auth';
import type { StudentResp } from '@/types/student';
import type { AxiosError } from 'axios';
import { onMounted, ref, watch, computed } from 'vue';

const authStore = useAuthStore();

const isStudent = computed(() => {
  return authStore.user_role === 'STUDENT';
});

const emit = defineEmits<{
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

interface Filters {
  id_scholarship: string;
  id_num_student: string;
  name_semester: string;
  size_scholarshp: string;
  id_budget: string;
  surname_student: string;
  first_name_student: string;
  second_name_student: string;
  name_group: string;
  type_scholarship_budget: string;
}

const localFilters = ref<Filters>({
  id_scholarship: '',
  id_num_student: '',
  name_semester: '',
  size_scholarshp: '',
  id_budget: '',
  surname_student: '',
  first_name_student: '',
  second_name_student: '',
  name_group: '',
  type_scholarship_budget: '',
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
    (localFilters.value as any)[key] = key === 'score_mark' || key === 'id_num_student' || key === 'id_mark' ? 0 : '';
  });
  emit('update-filters', {});
};

const students = ref<StudentResp[]>([]);
const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

onMounted(async () => {
  if (!checkAuth()) return;
  if (students.value.length > 0) return;
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
});

</script>

<template>
  <div class="filters-form">
    <div class="filters-grid">
      <div class="hidden filter-item">
        <label for="id_scholarship">ID:</label>
        <input
          type="text"
          id="id_scholarship"
          v-model="localFilters.id_scholarship"
        />
      </div>

      <div v-if="!isStudent" class="filter-item">
        <label for="id_num_student">Номер билета:</label>
        <input
          type="text"
          id="id_num_student"
          v-model="localFilters.id_num_student"
        />
      </div>

      <div class="filter-item">
        <label for="name_semester">Семестр:</label>
        <input
          type="text"
          id="name_semester"
          v-model="localFilters.name_semester"
        />
      </div>

      <div class="filter-item">
        <label for="size_scholarshp">Размер:</label>
        <input
          type="text"
          id="size_scholarshp"
          v-model="localFilters.size_scholarshp"
        />
      </div>

      <div class="hidden filter-item">
        <label for="id_budget">ID бюджета:</label>
        <input
          type="text"
          id="id_budget"
          v-model="localFilters.id_budget"
        />
      </div>

      <div v-if="!isStudent" class="filter-item">
        <label for="surname_student">Фамилия:</label>
        <input
          type="text"
          id="surname_student"
          v-model="localFilters.surname_student"
        />
      </div>

      <div v-if="!isStudent" class="filter-item">
        <label for="first_name_student">Имя:</label>
        <input
          type="text"
          id="first_name_student"
          v-model="localFilters.first_name_student"
        />
      </div>

      <div v-if="!isStudent" class="filter-item">
        <label for="second_name_student">Отчество:</label>
        <input
          type="text"
          id="second_name_student"
          v-model="localFilters.second_name_student"
        />
      </div>

      <div v-if="!isStudent" class="filter-item">
        <label for="name_group">Группа:</label>
        <input
          type="text"
          id="name_group"
          v-model="localFilters.name_group"
        />
      </div>

      <div class="filter-item">
        <label for="type_scholarship_budget">Тип стипендии:</label>
        <input
          type="text"
          id="type_scholarship_budget"
          v-model="localFilters.type_scholarship_budget"
        />
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
