<script setup lang="ts">
import { markService } from '@/services/markService';
import type { MarkReq, MarkResp } from '@/types/mark';
import { ref } from 'vue';
import MarkFiltersform from './MarkFiltersform.vue';
import MarkFormModal from './MarkFormModal.vue';


const emit = defineEmits<{
  (e: 'refresh'): void;
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const props = defineProps<{
  marks: MarkResp[];
  sortField?: string;
  sortOrder?: 'asc' | 'desc';
  currentFilters?: Record<string, string>;
}>();

const isFiltersOpen = ref(false);
const showModal = ref(false);
const modalMode = ref<'create' | 'edit'>('create');
const selectedMark = ref<MarkResp | undefined>(undefined);

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

const handleRowDoubleClick = (mark: MarkResp) => {
  selectedMark.value = mark;
  modalMode.value = 'edit';
  showModal.value = true;
};

const handleCreateClick = () => {
  selectedMark.value = undefined;
  modalMode.value = 'create';
  showModal.value = true;
};

const handleModalSubmit = async (mark: MarkReq) => {
  try {
    if (modalMode.value === 'create') {
      await markService.createMark(mark);
    } else if (selectedMark.value) {
      await markService.updateMark(mark);
    }
    showModal.value = false;
    emit('refresh');
  } catch (error) {
    console.error('Error saving mark:', error);
  }
};

const handleDelete = async (mark: MarkResp) => {
  if (confirm('Вы уверены, что хотите удалить оценку?')) {
    try {
      await markService.deleteMarks([mark.id_mark.toString()]);
      emit('refresh');
    } catch (error) {
      console.error('Error deleting mark:', error);
      // Here you might want to show an error message to the user
    }
  }
};
</script>

<template>
  <div class="table-window">
    <div class="table-container">
      <div class="table-header">
        <h1>Оценки</h1>
        <div class="filters-wrapper">
          <button class="hamburger" :class="{ rotated: isFiltersOpen }" @click="toggleFilters">☰</button>
          <div class="filters-dropdown" v-show="isFiltersOpen">
            <MarkFiltersform @update-filters="handleFiltersUpdate" />
          </div>
        </div>
      </div>
      <div class="table-scroll">
        <table>
          <thead>
            <tr>
              <th @click="handleSort('id_mark')" class="sortable" hidden>
                ID оценки {{ getSortIcon('id_mark') }}
              </th>
              <th @click="handleSort('id_num_student')" class="sortable">
                Cтудент {{ getSortIcon('id_num_student') }}
              </th>
              <th @click="handleSort('second_name_student')" class="sortable">
                  ФИО {{ getSortIcon('second_name_student') }}
                </th>
              <th @click="handleSort('name_group')" class="sortable">
                Группа {{ getSortIcon('name_group') }}
              </th>
              <th @click="handleSort('name_semester')" class="sortable">
                Семестр {{ getSortIcon('name_semester') }}
              </th>
              <th @click="handleSort('lesson_name_mark')" class="sortable">
                Название предмета {{ getSortIcon('lesson_name_mark') }}
              </th>
              <th @click="handleSort('score_mark')" class="sortable">
                Оценка {{ getSortIcon('score_mark') }}
              </th>
              <th @click="handleSort('type_mark')" class="sortable">
                Тип оценки {{ getSortIcon('type_mark') }}
              </th>
              <th>Действия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="mark in props.marks"
              :key="mark.id_mark"
              @dblclick="handleRowDoubleClick(mark)"
              class="editable-row">
              <td hidden>{{ mark.id_mark }}</td>
              <td>{{ mark.id_num_student }}</td>
              <td>{{ mark.second_name_student+" "+mark.first_name_student+" "+mark.surname_student }}</td>
              <td>{{ mark.name_group }}</td>
              <td>{{ mark.name_semester }}</td>
              <td>{{ mark.lesson_name_mark }}</td>
              <td>{{ mark.score_mark }}</td>
              <td>{{ mark.type_mark }}</td>
              <td>
                <button class="action-button delete" @click="handleDelete(mark)">Удалить</button>
              </td>
            </tr>
            <tr class="create-row" @click="handleCreateClick">
              <td colspan="9" class="create-cell">
                <span class="create-icon">+</span> Добавить новую оценку
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
  <MarkFormModal
    v-if="showModal"
    :show="showModal"
    :mode="modalMode"
    :mark="selectedMark"
    @close="showModal = false"
    @submit="handleModalSubmit"
  />
</template>

<style scoped>
@import '../../assets/table.css';
</style>
