<script setup lang="ts">
import router from '@/router';
import { groupService } from '@/services/groupService';
import { studentService } from '@/services/studentService';
import { useAuthStore } from '@/stores/auth';
import type { GroupResp } from '@/types/group';
import type { StudentResp } from '@/types/student';
import type { AxiosError } from 'axios';
import { onMounted, ref, watch, computed } from 'vue';

interface MarkFilters {
  id_num_student: number;
  name_semester: string;
  lesson_name_mark: string;
  type_mark: string;
  name_group?: string; // Optional field for group name
}

const emit = defineEmits<{
  (e: 'update-filters', filters: Record<string, string>): void;
}>()

const localFilters = ref<MarkFilters>({
  id_num_student: 0,
  name_semester: '',
  lesson_name_mark: '',
  type_mark: '',
  name_group: '',
});

const apply = () => {
  // Only include non-empty filters
  const activeFilters = Object.entries(localFilters.value)
    .reduce((acc, [key, value]) => {
      if (value || value === 0) { // Include zero values
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
const groups = ref<string[]>([]);

const onclick = async () => {
  if (!checkAuth()) return;
  // Fetch groups from a service or prop
  if (groups.value.length > 0) return;
  try{
    const response = await groupService.getGroups(1, 1000);
    groups.value = response.data.map((group: GroupResp) => group.name_group);
    groups.value.sort((a, b) => a.localeCompare(b));
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      console.error(err);
    }
  }
};

const isStudent = computed(() => {
  return authStore.user_role === 'STUDENT';
});

const props = defineProps<{
  filters: {
    studentId?: number;
    subjectId?: number;
    semester?: number;
    mark?: number;
  };
}>();

</script>

<template>
  <div class="filters-form">
    <div class="filters-grid">
      <div v-if="!isStudent" class="filter-group">
        <label>Номер студента</label>
        <select v-model="localFilters.id_num_student">
          <option v-for="student in students" :key="student.id_num_student" :value="student.id_num_student">{{ student.second_name_student+" "+student.first_name_student+" "+student.surname_student+" "+student.id_num_student }}</option>
        </select>
      </div>
      <div v-if="!isStudent" class="filter-group">
        <label>Группа</label>
        <select @click="onclick" v-model="localFilters.name_group" placeholder="Выберите группу">
          <option v-for="group in groups" :key="group" :value="group">{{ group }}</option>
        </select>
      </div>
      <div class="filter-group">
        <label>Семестр</label>
        <input v-model="localFilters.name_semester" placeholder="Введите семестр" />
      </div>
      <div class="filter-group">
        <label>Название предмета</label>
        <input v-model="localFilters.lesson_name_mark" placeholder="Введите название предмета" />
      </div>
      <div class="filter-group">
        <label>Тип оценки</label>
        <input v-model="localFilters.type_mark" placeholder="Введите тип оценки" />
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
