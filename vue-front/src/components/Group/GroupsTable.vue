<script setup lang="ts">
import GroupFilterform from './GroupFilterform.vue';
import GroupFormModal from './GroupFormModal.vue';
import {ref} from 'vue';
import { groupService } from '@/services/groupService';
import type { GroupResp } from '@/types/group';
import { pdfService } from '@/services/pdfService';

const emit = defineEmits<{
  (e: 'refresh'): void;
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const props = defineProps<{
  groups: GroupResp[];
  sortField?: string;
  sortOrder?: 'asc' | 'desc';
  currentFilters?: Record<string, string>;
}>();

const isFiltersOpen = ref(false);
const showModal = ref(false);
const modalMode = ref<'create' | 'edit'>('create');
const selectedGroup = ref<GroupResp | undefined>(undefined);


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

const handleRowDoubleClick = (group: GroupResp) => {
  selectedGroup.value = group;
  modalMode.value = 'edit';
  showModal.value = true;
};

const handleCreateClick = () => {
  selectedGroup.value = undefined;
  modalMode.value = 'create';
  showModal.value = true;
};

const handleModalSubmit = async (group: GroupResp) => {
  try {
    if (modalMode.value === 'create') {
      await groupService.createGroup(group);
    } else if (selectedGroup.value) {
      await groupService.updateGroup(group);
    }
    showModal.value = false;
    emit('refresh');
  } catch (error) {
    console.error('Error saving group:', error);
  }
};

const handleDelete = async (group: GroupResp) => {
  if (confirm(`Вы уверены, что хотите удалить группу "${group.name_group}"?`)) {
    try {
      await groupService.deleteGroups([group.name_group]);
      emit('refresh');
    } catch (error) {
      console.error('Error deleting group:', error);
      alert('Не удалось удалить группу. Пожалуйста, попробуйте позже.');
    }
  }
}

const handleExport = async () => {
  try {
    // Получаем все группы с учетом фильтров и сортировки
    const response = await groupService.getGroups(
      1, // Первая страница
      1000, // Большой лимит, чтобы получить все записи
      props.sortField,
      props.sortOrder,
      props.currentFilters
    );
    
    await pdfService.generateReport(
      response.data,
      'groups',
      'Отчет по группам'
    );
  } catch (error) {
    console.error('Error exporting groups:', error);
  }
};
</script>

<template>
  <div class="table-window">
    <div class="table-container">
      <div class="table-header">
        <h1>Группы</h1>
        <div class="header-actions">
          <button class="export-button" @click="handleExport">📄 Экспорт</button>
          <div class="filters-wrapper">
            <button class="hamburger" :class="{ rotated: isFiltersOpen }" @click="toggleFilters">☰</button>
            <div class="filters-dropdown" v-show="isFiltersOpen">
              <GroupFilterform @update-filters="handleFiltersUpdate" />
            </div>
          </div>
        </div>
      </div>
      <div class="table-scroll">
        <table>
          <thead>
            <tr>
              <th @click="handleSort('name_group')" class="sortable">
                Название группы {{ getSortIcon('name_group') }}
              </th>
              <th @click="handleSort('studies_direction_group')" class="sortable">
                Направление обучения {{ getSortIcon('studies_direction_group') }}
              </th>
              <th @click="handleSort('studies_profile_group')" class="sortable">
                Профиль обучения {{ getSortIcon('studies_profile_group') }}
              </th>
              <th @click="handleSort('start_date_group')" class="sortable">
                Дата начала обучения {{ getSortIcon('start_date_group') }}
              </th>
              <th @click="handleSort('studies_period_group')" class="sortable">
                Срок обучения {{ getSortIcon('studies_period_group') }}
              </th>
              <th>Действия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="group in groups"
                :key="group.name_group"
                @dblclick="handleRowDoubleClick(group)"
                class="editable-row">
              <td>{{ group.name_group }}</td>
              <td>{{ group.studies_direction_group }}</td>
              <td>{{ group.studies_profile_group }}</td>
              <td>{{ group.start_date_group }}</td>
              <td>{{ group.studies_period_group }}</td>
              <td>
                <button class="action-button delete" @click="handleDelete(group)">Удалить</button>
              </td>
            </tr>
            <tr class="create-row" @click="handleCreateClick">
              <td colspan="6" class="create-cell">
                <span class="create-icon">+</span> Cоздать новую группу
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>

  <GroupFormModal
    v-if="showModal"
    :show="showModal"
    :mode="modalMode"
    :group="selectedGroup"
    @close="showModal = false"
    @submit="handleModalSubmit"
  />
</template>

<style scoped>
@import '../../assets/table.css';
</style>
