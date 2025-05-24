<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import GroupsTable from '@/components/Group/GroupsTable.vue';
import { groupService } from '@/services/groupService';
import type { GroupResp } from '@/types/group';
import type { AxiosError } from 'axios';
import { watch } from 'vue';
import PaginationBlock from '@/components/PaginationBlock.vue';

const router = useRouter();
const authStore = useAuthStore();
const groups = ref<GroupResp[]>([]);
const loading = ref(true);
const error = ref<string | null>(null);

const page = ref(1);
const limit = ref(10);
const total = ref(0);

const sortOrder = ref<'asc' | 'desc'>('asc');
const sortField = ref('name_group');
const filters = ref<Record<string, string>>({});

watch(page, () => {
  fetchGroups();
});

const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

const fetchGroups = async () => {
  if (!checkAuth()) return;

  try {
    loading.value = true;
    const response = await groupService.getGroups(
      page.value,
      limit.value,
      sortField.value,
      sortOrder.value,
      filters.value
    );
    groups.value = response.data;
    total.value = response.pagination.total;
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      error.value = axiosError.message;
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
  fetchGroups();
};

onMounted(() => {
  // Initialize auth state
  authStore.initialize();
  fetchGroups();
});
</script>
<template>
  <div class="group">
    <GroupsTable
      :groups="groups"
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
.group {
  padding: 20px;
  animation: fromUp 1s ease-in-out;
}

.error {
  color: red;
}
</style>
