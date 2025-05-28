<script setup lang="ts">
import SemesterFormModal from '@/components/Semester/SemesterFormModal.vue'
import type { SemesterResp, SemesterReq } from '@/types/semester';
import { ref } from 'vue';
import { semesterService } from '@/services/semesterService';
import SemesterFiltersform from '@/components/Semester/SemesterFiltersform.vue';
import { pdfService } from '@/services/pdfService';

const emit = defineEmits<{
  (e: 'refresh'): void;
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const props = defineProps<{
  semesters: SemesterResp[];
  sortField?: string;
  sortOrder?: 'asc' | 'desc';
  currentFilters?: Record<string, string>;
}>();

const isFiltersOpen = ref(false);
const showModal = ref(false);
const modalMode = ref<'create' | 'edit'>('create');
const selectedSemester = ref<SemesterResp | undefined>(undefined);

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

const handleRowDoubleClick = (semester: SemesterResp) => {
  selectedSemester.value = semester;
  modalMode.value = 'edit';
  showModal.value = true;
};

const handleCreateClick = () => {
  selectedSemester.value = undefined;
  modalMode.value = 'create';
  showModal.value = true;
};

const handleModalSubmit = async (semester: SemesterReq) => {
  try {
    if (modalMode.value === 'create') {
      await semesterService.createSemester(semester);
    } else if (selectedSemester.value) {
      await semesterService.updateSemester(semester);
    }
    showModal.value = false;
    emit('refresh');
  } catch (error) {
    console.error('Error saving semester:', error);
  }
};

const handleDelete = async (semester: SemesterResp) => {
  if (confirm('Вы уверены, что хотите удалить этот семестр?')) {
    try {
      await semesterService.deleteSemesters([semester.name_semester]);
      emit('refresh');
    } catch (error) {
      console.error('Error deleting semester:', error);
    }
  }
};

const handleExport = async () => {
  try {
    // Получаем все семестры с учетом фильтров и сортировки
    const response = await semesterService.getSemesters(
      1, // Первая страница
      1000, // Большой лимит, чтобы получить все записи
      props.sortField,
      props.sortOrder,
      props.currentFilters
    );
    
    await pdfService.generateReport(
      response.data,
      'semesters',
      'Отчет по семестрам'
    );
  } catch (error) {
    console.error('Error exporting semesters:', error);
  }
};
</script>

<template>
  <div class="table-window">
    <div class="table-container">
      <div class="table-header">
        <h1>Семестры</h1>
        <div class="header-actions">
          <button class="export-button" @click="handleExport">📄 Экспорт</button>
          <div class="filters-wrapper">
            <button class="hamburger" :class="{ rotated: isFiltersOpen }" @click="toggleFilters">☰</button>
            <div class="filters-dropdown" v-show="isFiltersOpen">
              <SemesterFiltersform @update-filters="handleFiltersUpdate" />
            </div>
          </div>
        </div>
      </div>
      <div class="table-scroll">
        <table>
          <thead>
            <tr>
              <th @click="handleSort('name_semester')" class="sortable">
                Название {{ getSortIcon('name_semester') }}
              </th>
              <th @click="handleSort('date_start_semester')" class="sortable">
                Дата начала {{ getSortIcon('date_start_semester') }}
              </th>
              <th @click="handleSort('date_end_semester')" class="sortable">
                Дата окончания {{ getSortIcon('date_end_semester') }}
              </th>
              <th>Действия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="semester in semesters"
                :key="semester.name_semester"
                @dblclick="handleRowDoubleClick(semester)"
                class="editable-row">
              <td>{{ semester.name_semester }}</td>
              <td>{{ semester.date_start_semester }}</td>
              <td>{{ semester.date_end_semester }}</td>
              <td>
                <button class="action-button delete" @click="handleDelete(semester)">Удалить</button>
              </td>
            </tr>
            <tr class="create-row" @click="handleCreateClick">
              <td colspan="4" class="create-cell">
                <span class="create-icon">+</span> Добавить новый семестр
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>

  <SemesterFormModal
    :show="showModal"
    :semester="selectedSemester"
    :mode="modalMode"
    @close="showModal = false"
    @submit="handleModalSubmit"
  />
</template>

<style scoped>
@import '../../assets/table.css';
</style>
