<script setup lang="ts">
import CategoryTable from '@/components/Category/CategoryTable.vue';
import { categoryService } from '@/services/categoryService';
import { useAuthStore } from '@/stores/auth';
import type { CategoryResp } from '@/types/category';
import type { AxiosError } from 'axios';
import { onMounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const authStore = useAuthStore();
const categories = ref<CategoryResp[]>([]);
const loading = ref(true);
const error = ref<string | null>(null);

const page = ref(1);
const limit = ref(10);
const total = ref(0);

const sortOrder = ref<'asc' | 'desc'>('asc');
const sortField = ref('name_semester');
const filters = ref<Record<string, string>>({});

watch(page, () => {
  fetchCategories();
});

const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

const fetchCategories = async () =>{
  if (!checkAuth()) return;

  try {
    loading.value = true;
    const response = await categoryService.getCategories(
      page.value,
      limit.value,
      sortField.value,
      sortOrder.value,
      filters.value
    );
    categories.value = response.data;
    total.value = response.pagination.total;
  } catch (err){
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
  fetchCategories();
};

onMounted(() => {
  // Initialize auth state
  authStore.initialize();
  fetchCategories();
});
</script>
<template>
  <div class="category">
    <CategoryTable
      :categories="categories"
      :sort-field="sortField"
      :sort-order="sortOrder"
      :current-filters="filters"
      @update-filters="handleFiltersUpdate"
      @refresh="fetchCategories"
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
.category {
  padding: 20px;
  animation: fromUp 1s ease-in-out;
}

.error {
  color: red;
}
</style>
