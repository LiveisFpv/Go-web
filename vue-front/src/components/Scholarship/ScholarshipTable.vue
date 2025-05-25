<script setup lang="ts">
import { ref, watch } from 'vue';
import type { ScholarshipReq, ScholarshipResp } from '@/types/scholarship';
import ScholarshipFiltersform from '@/components/Scholarship/ScholarshipFiltersform.vue';
import ScholarshipFormModal from '@/components/Scholarship/ScholarshipFormModal.vue';
import { scholarshipService } from '@/services/scholarshipService';
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const props = defineProps<{
  scholarships: ScholarshipResp[];
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
const selectedScholarship = ref<ScholarshipResp | undefined>(undefined);

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

const handleRowDoubleClick = (scholarship: ScholarshipResp) => {
  selectedScholarship.value = scholarship;
  modalMode.value = 'edit';
  showModal.value = true;
};

const handleCreateClick = () => {
  selectedScholarship.value = undefined;
  modalMode.value = 'create';
  showModal.value = true;
};

const handleModalSubmit = async (scholarship: ScholarshipReq) => {
  try {
    if (modalMode.value === 'create') {
      await scholarshipService.createScholarship(scholarship);
    } else if (selectedScholarship.value) {
      await scholarshipService.updateScholarship(scholarship);
    }
    showModal.value = false;
    emit('refresh');
  } catch (error) {
    console.error('Error saving scholarship:', error);
  }
};
const handleDelete = async (scholarship: ScholarshipResp) => {
  if (confirm('Вы уверены, что хотите удалить стипендию?')) {
    try {
      await scholarshipService.deleteScholarships([scholarship.id_scholarship.toString()]);
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
        <h1>Стипендии</h1>
        <div class="filters-wrapper">
            <button class="hamburger" :class="{ rotated: isFiltersOpen }" @click="toggleFilters">☰</button>
            <div class="filters-dropdown" v-show="isFiltersOpen">
              <ScholarshipFiltersform
                @update-filters="handleFiltersUpdate"
              />
            </div>
        </div>
      </div>
      <div class="table-scroll">
        <table>
          <thead>
            <tr>
                <th @click="handleSort('id_scholarship')"class="sortable" hidden>
                  ID Стипендии {{ getSortIcon('id_scholarship') }}
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
                <th @click="handleSort('size_scholarshp')">
                  Размер {{ getSortIcon('size_scholarshp') }}
                </th>
                <th @click="handleSort('type_scholarship_budget')">
                  Тип стипендии {{ getSortIcon('type_scholarship_budget') }}
                </th>
                <th>Действия</th>
              </tr>
          </thead>
          <tbody>
            <tr v-for="scholarship in scholarships"
              :key="scholarship.id_scholarship"
              @dblclick="handleRowDoubleClick(scholarship)"
              class="editable-row">
              <td hidden>{{ scholarship.id_scholarship }}</td>
              <td>{{ scholarship.id_num_student }}</td>
              <td>{{ scholarship.second_name_student+" "+scholarship.first_name_student+" "+scholarship.surname_student }}</td>
              <td>{{ scholarship.name_group }}</td>
              <td>{{ scholarship.name_semester }}</td>
              <td>{{ scholarship.size_scholarshp }}</td>
              <td>{{ scholarship.type_scholarship_budget }}</td>
              <td>
                <button class="action-button delete" @click="handleDelete(scholarship)">Удалить</button>
              </td>
            </tr>
            <tr class="create-row" @click="handleCreateClick">
              <td colspan="7" class="create-cell">
                <span class="create-icon">+</span> Добавить стипендию
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>

  <ScholarshipFormModal
    v-if="showModal"
    :show="showModal"
    :mode="modalMode"
    :mark="selectedScholarship"
    @close="showModal = false"
    @submit="handleModalSubmit"
  />
</template>

<style scoped>
@import '../../assets/table.css';
</style>
