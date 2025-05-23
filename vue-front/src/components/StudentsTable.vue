<script setup lang="ts">
import Filtersform from '@/components/StudentFiltersform.vue'
import StudentFormModal from '@/components/StudentFormModal.vue'
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
.table-window {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.table-container {
  width: 100%;
  max-width: 1200px;
  border-radius: 15px;
  background-color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px 20px;
  background-color: var(--color-accent);
  width: 100%;
  box-sizing: border-box;
  border-bottom: 1px solid var(--color-background);
}

.table-header h1 {
  margin: 0;
  color: white;
  font-size: 1.5rem;
  font-weight: 600;
}

.table-scroll {
  overflow-x: auto;
  width: 100%;
}

table {
  width: 100%;
  border-collapse: collapse;
  min-width: 800px;
}

th, td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
  white-space: nowrap;
}

th {
  background-color: var(--color-accent);
  color: white;
  position: sticky;
  top: 0;
}

td {
  background-color: #f9f9f9;
}

td:hover {
  background-color: #f1f1f1;
}

.sortable {
  cursor: pointer;
  user-select: none;
}

.sortable:hover {
  background-color: var(--color-accent-hover);
}

.filters-wrapper {
  position: relative;
}

.hamburger {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  padding: 0.5rem;
  transition: transform 0.3s ease;
  text-align: center;
  vertical-align: middle;
  color: white;
  z-index: 2;
}

.hamburger.rotated {
  transform: rotate(90deg);
}

.filters-dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  background-color: var(--color-background);
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 1000;
  min-width: 300px;
  margin-top: 5px;
  animation: slideDown 0.3s ease-out;
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.editable-row {
  cursor: pointer;
}

.editable-row:hover {
  background-color: #f0f0f0;
}

.create-row {
  cursor: pointer;
  background-color: #f8f9fa;
}

.create-cell {
  text-align: center;
  color: var(--color-accent);
  font-weight: 500;
  padding: 15px;
}

.create-icon {
  font-size: 1.2em;
  margin-right: 5px;
}

.action-button {
  padding: 4px 8px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9em;
  transition: background-color 0.2s;
}

.action-button.delete {
  background-color: #dc3545;
  color: white;
}

.action-button.delete:hover {
  background-color: #c82333;
}
</style>

