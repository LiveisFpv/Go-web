<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import SemesterTable from '@/components/Semester/SemesterTable.vue';
import { semesterService } from '@/services/semesterService';
import type { SemesterResp } from '@/types/semester';
import type { AxiosError } from 'axios';
import { watch } from 'vue';
import PaginationBlock from '@/components/PaginationBlock.vue';

const router = useRouter();
const authStore = useAuthStore();
const semesters = ref<SemesterResp[]>([]);
const loading = ref(true);
const error = ref<string | null>(null);

const page = ref(1);
const limit = ref(10);
const total = ref(0);

const sortOrder = ref<'asc' | 'desc'>('asc');
const sortField = ref('name_semester');
const filters = ref<Record<string, string>>({});

watch(page, () => {
  fetchSemesters();
});

// Check authentication
const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

// Fetch semesters data
const fetchSemesters = async () => {
  if (!checkAuth()) return;

  try {
    loading.value = true;
    const response = await semesterService.getSemesters(
      page.value,
      limit.value,
      sortField.value,
      sortOrder.value,
      filters.value
    );
    semesters.value = response.data;
    total.value = response.pagination.total;
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      error.value = 'Failed to fetch semesters';
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
  fetchSemesters();
};

onMounted(() => {
  // Initialize auth state
  authStore.initialize();
  fetchSemesters();
});
</script>

<template>
  <div class="semester">
    <SemesterTable
      :semesters="semesters"
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
.semester {
  padding: 20px;
  animation: fromUp 1s ease-in-out;
}

.error {
  color: red;
}
</style>
