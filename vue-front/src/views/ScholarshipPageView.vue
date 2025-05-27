<script setup lang="ts">
import ScholarshipTable from '@/components/Scholarship/ScholarshipTable.vue';
import PaginationBlock from '@/components/PaginationBlock.vue';
import { scholarshipService } from '@/services/scholarshipService';
import { useAuthStore } from '@/stores/auth';
import type { ScholarshipResp } from '@/types/scholarship';
import type { AxiosError } from 'axios';
import { onMounted, ref, watch, computed } from 'vue';
import { useRouter } from 'vue-router';
import { userService } from '@/services/userService';

const router = useRouter();
const authStore = useAuthStore();
const scholarships = ref<ScholarshipResp[]>([]);
const loading = ref(true);
const error = ref<string | null>(null);

const isStudent = computed(() => {
  return authStore.user_role === 'STUDENT';
});

const page = ref(1);
const limit = ref(10);
const total = ref(0);

const sortOrder = ref<'asc' | 'desc'>('asc');
const sortField = ref('id_scholarship');
const filters = ref<Record<string, string>>({});

watch(page, () => {
  fetchScholarships();
});

const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

const getStudentId = async () => {
  if (!authStore.email) return;
  try {
    const userData = await userService.getUserbyEmail(authStore.email);
    if (userData.user_student_id) {
      // Устанавливаем фильтр по ID студента
      filters.value.id_num_student = userData.user_student_id.toString();
    }
  } catch (err) {
    console.error('Failed to fetch user data:', err);
  }
};

const fetchScholarships = async () => {
  if (!checkAuth()) return;

  try {
    loading.value = true;
    // Если пользователь студент, получаем его ID и устанавливаем фильтр
    if (isStudent.value) {
      await getStudentId();
    }
    const response = await scholarshipService.getScholarships(
      page.value,
      limit.value,
      sortField.value,
      sortOrder.value,
      filters.value
    );
    scholarships.value = response.data;
    total.value = response.pagination.total;
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      console.error('Failed to fetch scholarships:', axiosError);
    }
  } finally {
    loading.value = false;
  }
};

const handleFiltersUpdate = (newFilters: Record<string, string>) => {
  // Extract sort parameters
  if (newFilters.sort) {
    sortField.value = newFilters.sort;
    sortOrder.value = newFilters.order as 'asc' | 'desc';
  }

  // Remove sort parameters from filters
  const { sort, order, ...filterParams } = newFilters;
  
  // Если пользователь студент, сохраняем фильтр по его ID
  if (isStudent.value && filters.value.id_num_student) {
    filterParams.id_num_student = filters.value.id_num_student;
  }
  
  filters.value = filterParams;

  page.value = 1; // Reset to first page when filters change
  fetchScholarships();
};

onMounted(async () => {
  // Initialize auth state
  authStore.initialize();
  // Если пользователь студент, получаем его ID перед загрузкой данных
  if (isStudent.value) {
    await getStudentId();
  }
  fetchScholarships();
});
</script>

<template>
  <div class="scholarship">
    <ScholarshipTable
      :scholarships="scholarships"
      :sort-field="sortField"
      :sort-order="sortOrder"
      :current-filters="filters"
      @update-filters="handleFiltersUpdate"
      @refresh="fetchScholarships"
    />
    <PaginationBlock
      :page="page"
      :limit="limit"
      :total="total"
      @update:page="page = $event"
    />
  </div>
</template>

<style scoped>
.scholarship {
  padding: 20px;
  animation: fromUp 1s ease-in-out;
}

.error {
  color: red;
}
</style>
