<script setup lang="ts">
import Filtersform from '@/components/Student/StudentFiltersform.vue'
import StudentFormModal from '@/components/Student/StudentFormModal.vue'
import type { StudentResp, StudentReq } from '@/types/student';
import { ref } from 'vue';
import { studentService } from '@/services/studentService';

const emit = defineEmits<{
  (e: 'refresh'): void;
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const props = defineProps<{
  students: StudentResp[];
  sortField?: string;
  sortOrder?: 'asc' | 'desc';
  currentFilters?: Record<string, string>;
}>();

const isFiltersOpen = ref(false);
const showModal = ref(false);
const modalMode = ref<'create' | 'edit'>('create');
const selectedStudent = ref<StudentResp | undefined>(undefined);

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

const handleRowDoubleClick = (student: StudentResp) => {
  selectedStudent.value = student;
  modalMode.value = 'edit';
  showModal.value = true;
};

const handleCreateClick = () => {
  selectedStudent.value = undefined;
  modalMode.value = 'create';
  showModal.value = true;
};

const handleModalSubmit = async (student: StudentReq) => {
  try {
    if (modalMode.value === 'create') {
      await studentService.createStudent(student);
    } else if (selectedStudent.value) {
      await studentService.updateStudent( student);
    }
    showModal.value = false;
    emit('refresh');
  } catch (error) {
    console.error('Error saving student:', error);
    // Here you might want to show an error message to the user
  }
};

const handleDelete = async (student: StudentResp) => {
  if (confirm('Вы уверены, что хотите удалить этого студента?')) {
    try {
      await studentService.deleteStudents([student.id_num_student.toString()]);
      emit('refresh');
    } catch (error) {
      console.error('Error deleting student:', error);
      // Here you might want to show an error message to the user
    }
  }
};
</script>

<template>
  <div class="table-window">
    <div class="table-container">
      <div class="table-header">
        <h1>Студенты</h1>
        <div class="filters-wrapper">
          <button class="hamburger" :class="{ rotated: isFiltersOpen }" @click="toggleFilters">☰</button>
          <div class="filters-dropdown" v-show="isFiltersOpen">
            <Filtersform @update-filters="handleFiltersUpdate" />
          </div>
        </div>
      </div>
      <div class="table-scroll">
        <table>
          <thead>
            <tr>
              <th @click="handleSort('id_num_student')" class="sortable">
                Номер билета {{ getSortIcon('id_num_student') }}
              </th>
              <th @click="handleSort('name_group')" class="sortable">
                Группа {{ getSortIcon('name_group') }}
              </th>
              <th @click="handleSort('email_student')" class="sortable">
                Email {{ getSortIcon('email_student') }}
              </th>
              <th @click="handleSort('second_name_student')" class="sortable">
                Фамилия {{ getSortIcon('second_name_student') }}
              </th>
              <th @click="handleSort('first_name_student')" class="sortable">
                Имя {{ getSortIcon('first_name_student') }}
              </th>
              <th @click="handleSort('surname_student')" class="sortable">
                Отчество {{ getSortIcon('surname_student') }}
              </th>
              <th>Действия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="student in students"
                :key="student.id_num_student"
                @dblclick="handleRowDoubleClick(student)"
                class="editable-row">
              <td>{{ student.id_num_student }}</td>
              <td>{{ student.name_group}}</td>
              <td>{{ student.email_student}}</td>
              <td>{{ student.second_name_student }}</td>
              <td>{{ student.first_name_student}}</td>
              <td>{{ student.surname_student }}</td>
              <td>
                <button class="action-button delete" @click="handleDelete(student)">Удалить</button>
              </td>
            </tr>
            <tr class="create-row" @click="handleCreateClick">
              <td colspan="7" class="create-cell">
                <span class="create-icon">+</span> Добавить нового студента
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>

  <StudentFormModal
    :show="showModal"
    :student="selectedStudent"
    :mode="modalMode"
    @close="showModal = false"
    @submit="handleModalSubmit"
  />
</template>

<style scoped>
@import '../../assets/table.css';
</style>

