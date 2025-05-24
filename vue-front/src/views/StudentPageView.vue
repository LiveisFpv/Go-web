<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import StudentsTable from '@/components/Student/StudentsTable.vue';
import { studentService } from '@/services/studentService';
import type { StudentResp } from '@/types/student';
import type { AxiosError } from 'axios';
import { watch } from 'vue';
import PaginationBlock from '@/components/PaginationBlock.vue';

const router = useRouter();
const authStore = useAuthStore();
const students = ref<StudentResp[]>([]);
const loading = ref(true);
const error = ref<string | null>(null);

const page = ref(1);
const limit = ref(10);
const total = ref(0);

const sortOrder = ref<'asc' | 'desc'>('asc');
const sortField = ref('id_num_student');
const filters = ref<Record<string, string>>({});

watch(page, () => {
  fetchStudents();
});


// Check authentication
const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

// Fetch students data
const fetchStudents = async () => {
  if (!checkAuth()) return;

  try {
    loading.value = true;
    const response = await studentService.getStudents(
      page.value,
      limit.value,
      sortField.value,
      sortOrder.value,
      filters.value
    );
    students.value = response.data;
    total.value = response.pagination.total;
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      error.value = 'Failed to fetch students';
      console.error(err);
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
  filters.value = filterParams;

  page.value = 1; // Reset to first page when filters change
  fetchStudents();
};

onMounted(() => {
  // Initialize auth state
  authStore.initialize();
  fetchStudents();
});
</script>

<template>
  <div class="student">
    <StudentsTable
      :students="students"
      :sort-field="sortField"
      :sort-order="sortOrder"
      :current-filters="filters"
      @update-filters="handleFiltersUpdate"
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
.student {
  padding: 20px;
  animation: fromUp 1s ease-in-out;
}

.error {
  color: red;
}
</style>

