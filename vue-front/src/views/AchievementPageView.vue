<script setup lang="ts">
import AchievementTable from '@/components/Achievement/AchievementTable.vue';
import PaginationBlock from '@/components/PaginationBlock.vue';
import { achievementService } from '@/services/achievementService';
import { useAuthStore } from '@/stores/auth';
import type { AchivementResp } from '@/types/achievement';
import type { AxiosError } from 'axios';
import { onMounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const authStore = useAuthStore();
const achievements = ref<AchivementResp[]>([]);
const loading = ref(true);
const error = ref<string | null>(null);

const page = ref(1);
const limit = ref(10);
const total = ref(0);

const sortOrder = ref<'asc' | 'desc'>('asc');
const sortField = ref('id_achivment');
const filters = ref<Record<string, string>>({});

watch(page, () => {
  fetchAchievements();
});

const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

const fetchAchievements = async () => {
  if (!checkAuth()) return;

  try {
    loading.value = true;
    const response = await achievementService.getAchievements(
      page.value,
      limit.value,
      sortField.value,
      sortOrder.value,
      filters.value
    );
    achievements.value = response.data;
    total.value = response.pagination.total;
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      console.error('Failed to fetch achievements:', axiosError);
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
  fetchAchievements();
};

onMounted(() => {
  // Initialize auth state
  authStore.initialize();
  fetchAchievements();
});
</script>

<template>
  <div class="achievement">
    <AchievementTable
      :achievements="achievements"
      :sort-field="sortField"
      :sort-order="sortOrder"
      :current-filters="filters"
      @update-filters="handleFiltersUpdate"
      @refresh="fetchAchievements"
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
.achievement {
  padding: 20px;
  animation: fromUp 1s ease-in-out;
}

.error {
  color: red;
}
</style>
