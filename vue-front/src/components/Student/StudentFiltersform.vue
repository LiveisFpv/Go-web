<script setup lang="ts">
import type { Filter } from '@/types/meta';
import { groupService } from '@/services/groupService';

import { onMounted, ref } from 'vue';
import type { GroupResp } from '@/types/group';
import { useAuthStore } from '@/stores/auth';
import router from '@/router';
import type { AxiosError } from 'axios';

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

const groups = ref<string[]>([]);
const authStore = useAuthStore();
const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

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
        <select @click="onclick" v-model="localFilters.name_group" placeholder="Выберите группу">
          <option v-for="group in groups" :key="group" :value="group">{{ group }}</option>
        </select>
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
@import '../../assets/filter.css';
</style>
