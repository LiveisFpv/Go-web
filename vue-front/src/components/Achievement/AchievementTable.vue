<script setup lang="ts">
import { ref, computed } from 'vue';
import type { AchivementReq, AchivementResp } from '@/types/achievement';
import AchievementFiltersform from '@/components/Achievement/AchievementFilterForm.vue';
import AchievementFormModal from '@/components/Achievement/AchievementFormModal.vue';
import { achievementService } from '@/services/achievementService';
import { useAuthStore } from '@/stores/auth';
import { pdfService } from '@/services/pdfService';

const authStore = useAuthStore();

const props = defineProps<{
  achievements: AchivementResp[];
  sortField?: string;
  sortOrder?: 'asc' | 'desc';
  currentFilters: Record<string, string>;
}>();

const emit = defineEmits<{
  (e: 'refresh'): void;
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const isFiltersOpen = ref(false);
const showModal = ref(false);
const modalMode = ref<'create' | 'edit'>('create');
const selectedAchievement = ref<AchivementResp | undefined>(undefined);

const handleFiltersUpdate = (filters: Record<string, string>) => {
  emit('update-filters', filters);
};

const handleSort = (field: string) => {
  const currentOrder = props.sortOrder || 'asc';
  const newOrder = currentOrder === 'asc' ? 'desc' : 'asc';

  // Create new filters object with existing filters and new sort parameters
  const updatedFilters = {
    ...props.currentFilters, // Preserve existing filters
    sort: field,
    order: newOrder
  };

  emit('update-filters', updatedFilters);
};

const getSortIcon = (field: string) => {
  if (props.sortField !== field) return '↕';
  return props.sortOrder === 'asc' ? '↑' : '↓';
};

const toggleFilters = () => {
  isFiltersOpen.value = !isFiltersOpen.value;
};

const handleRowDoubleClick = (achievement: AchivementResp) => {
  selectedAchievement.value = achievement;
  modalMode.value = 'edit';
  showModal.value = true;
};

const handleCreateClick = () => {
  selectedAchievement.value = undefined;
  modalMode.value = 'create';
  showModal.value = true;
};

const handleModalSubmit = async (achievement: AchivementReq) => {
  try {
    if (modalMode.value === 'create') {
      await achievementService.createAchievement(achievement);
    } else if (selectedAchievement.value) {
      await achievementService.updateAchievement(achievement);
    }
    showModal.value = false;
    emit('refresh');
  } catch (error) {
    console.error('Error saving achievement:', error);
  }
};

const handleDelete = async (achievement: AchivementResp) => {
  if (confirm('Вы уверены, что хотите удалить достижение?')) {
    try {
      await achievementService.deleteAchievements([achievement.id_achivment.toString()]);
      emit('refresh');
    } catch (error) {
      console.error('Error deleting achievement:', error);
    }
  }
};

const handleExport = async () => {
  try {
    // Получаем все достижения с учетом фильтров и сортировки
    const response = await achievementService.getAchievements(
      1, // Первая страница
      1000, // Большой лимит, чтобы получить все записи
      props.sortField,
      props.sortOrder,
      props.currentFilters
    );
    
    await pdfService.generateReport(
      response.data,
      'achievements',
      'Отчет по достижениям'
    );
  } catch (error) {
    console.error('Error exporting achievements:', error);
  }
};

const isStudent = computed(() => {
  return authStore.user_role === 'STUDENT';
});

const canEdit = computed(() => {
  return authStore.user_role !== 'STUDENT';
});
</script>

<template>
  <div class="table-window">
    <div class="table-container">
      <div class="table-header">
        <h1>Достижения</h1>
        <div class="header-actions">
          <button class="export-button" @click="handleExport">📄 Экспорт</button>
          <div class="filters-wrapper">
            <button class="hamburger" :class="{ rotated: isFiltersOpen }" @click="toggleFilters">☰</button>
            <div class="filters-dropdown" v-show="isFiltersOpen">
              <AchievementFiltersform
                @update-filters="handleFiltersUpdate"
              />
            </div>
          </div>
        </div>
      </div>
      <div class="table-scroll">
        <table>
          <thead>
            <tr>
              <th @click="handleSort('id_achivment')" class="sortable" hidden>
                ID Достижения {{ getSortIcon('id_achivment') }}
              </th>
              <th v-if="!isStudent" @click="handleSort('id_num_student')" class="sortable">
                Студент {{ getSortIcon('id_num_student') }}
              </th>
              <th v-if="!isStudent" @click="handleSort('second_name_student')" class="sortable">
                ФИО {{ getSortIcon('second_name_student') }}
              </th>
              <th v-if="!isStudent" @click="handleSort('name_group')" class="sortable">
                Группа {{ getSortIcon('name_group') }}
              </th>
              <th @click="handleSort('name_achivement')" class="sortable">
                Название достижения {{ getSortIcon('name_achivement') }}
              </th>
              <th @click="handleSort('date_achivment')" class="sortable">
                Дата достижения {{ getSortIcon('date_achivment') }}
              </th>
              <th @click="handleSort('achivments_type_category')" class="sortable">
                Тип достижения {{ getSortIcon('achivments_type_category') }}
              </th>
              <th >Действия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="achievement in achievements"
              :key="achievement.id_achivment"
              @dblclick="handleRowDoubleClick(achievement)"
              class="editable-row">
              <td hidden>{{ achievement.id_achivment }}</td>
              <td v-if="!isStudent">{{ achievement.id_num_student }}</td>
              <td v-if="!isStudent">{{ achievement.second_name_student+" "+achievement.first_name_student+" "+achievement.surname_student }}</td>
              <td v-if="!isStudent">{{ achievement.name_group }}</td>
              <td>{{ achievement.name_achivement }}</td>
              <td>{{ achievement.date_achivement }}</td>
              <td>{{ achievement.achivments_type_category }}</td>
              <td >
                <button class="action-button delete" @click="handleDelete(achievement)">Удалить</button>
              </td>
            </tr>
            <tr class="create-row" @click="handleCreateClick">
              <td :colspan="isStudent ? 4 : 7" class="create-cell">
                <span class="create-icon">+</span> Добавить достижение
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>

  <AchievementFormModal
    v-if="showModal"
    :show="showModal"
    :mode="modalMode"
    :achievement="selectedAchievement"
    @close="showModal = false"
    @submit="handleModalSubmit"
  />
</template>

<style scoped>
@import '../../assets/table.css';
</style>
