<script setup lang="ts">
import BudgetFormModal from '@/components/Budget/BudgetFormModal.vue'
import type { BudgetResp, BudgetReq } from '@/types/budget';
import { ref } from 'vue';
import { budgetService } from '@/services/budgetService';
import BudgetFiltersform from '@/components/Budget/BudgetFiltersform.vue';

const emit = defineEmits<{
  (e: 'refresh'): void;
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const props = defineProps<{
  budgets: BudgetResp[];
  sortField?: string;
  sortOrder?: 'asc' | 'desc';
  currentFilters?: Record<string, string>;
}>();

const isFiltersOpen = ref(false);
const showModal = ref(false);
const modalMode = ref<'create' | 'edit'>('create');
const selectedBudget = ref<BudgetResp | undefined>(undefined);

const handleFiltersUpdate = (filters: Record<string, string>) => {
  emit('update-filters', filters);
};

const handleSort = (field: string) => {
  const currentOrder = props.sortOrder || 'asc';
  const newOrder = currentOrder === 'asc' ? 'desc' : 'asc';

  const updatedFilters = {
    ...props.currentFilters,
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

const handleRowDoubleClick = (budget: BudgetResp) => {
  selectedBudget.value = budget;
  modalMode.value = 'edit';
  showModal.value = true;
};

const handleCreateClick = () => {
  selectedBudget.value = undefined;
  modalMode.value = 'create';
  showModal.value = true;
};

const handleModalSubmit = async (budget: BudgetReq) => {
  try {
    if (modalMode.value === 'create') {
      await budgetService.createBudget(budget);
    } else if (selectedBudget.value) {
      await budgetService.updateBudget(budget);
    }
    showModal.value = false;
    emit('refresh');
  } catch (error) {
    console.error('Error saving budget:', error);
  }
};

const handleDelete = async (budget: BudgetResp) => {
  if (confirm('Вы уверены, что хотите удалить этот бюджет?')) {
    try {
      await budgetService.deleteBudgets([budget.id_budget]);
      emit('refresh');
    } catch (error) {
      console.error('Error deleting budget:', error);
    }
  }
};
</script>

<template>
  <div class="table-window">
    <div class="table-container">
      <div class="table-header">
        <h1>Бюджеты</h1>
        <div class="filters-wrapper">
          <button class="hamburger" :class="{ rotated: isFiltersOpen }" @click="toggleFilters">☰</button>
          <div class="filters-dropdown" v-show="isFiltersOpen">
            <BudgetFiltersform @update-filters="handleFiltersUpdate" />
          </div>
        </div>
      </div>
      <div class="table-scroll">
        <table>
          <thead>
            <tr>
              <th @click="handleSort('id_budget')" class="sortable" hidden>
                ID {{ getSortIcon('id_budget') }}
              </th>
              <th @click="handleSort('size_budget')" class="sortable">
                Размер {{ getSortIcon('size_budget') }}
              </th>
              <th @click="handleSort('type_scholarship_budget')" class="sortable">
                Тип стипендии {{ getSortIcon('type_scholarship_budget') }}
              </th>
              <th @click="handleSort('name_semester')" class="sortable">
                Семестр {{ getSortIcon('name_semester') }}
              </th>
              <th>Действия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="budget in budgets"
                :key="budget.id_budget"
                @dblclick="handleRowDoubleClick(budget)"
                class="editable-row">
              <td hidden>{{ budget.id_budget }}</td>
              <td>{{ budget.size_budget }}</td>
              <td>{{ budget.type_scholarship_budget }}</td>
              <td>{{ budget.name_semester}}</td>
              <td>
                <button class="action-button delete" @click="handleDelete(budget)">Удалить</button>
              </td>
            </tr>
            <tr class="create-row" @click="handleCreateClick">
              <td colspan="4" class="create-cell">
                <span class="create-icon">+</span> Добавить новый бюджет
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>

  <BudgetFormModal
    :show="showModal"
    :budget="selectedBudget"
    :mode="modalMode"
    @close="showModal = false"
    @submit="handleModalSubmit"
  />
</template>

<style scoped>
@import '../../assets/table.css';
</style>
