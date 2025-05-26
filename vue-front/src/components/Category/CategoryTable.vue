<script setup lang="ts">
import { categoryService } from '@/services/categoryService';
import type { CategoryReq, CategoryResp } from '@/types/category';
import { ref } from 'vue';
import CategoryFilterform from './CategoryFilterform.vue';
import CategoryFormModal from './CategoryFormModal.vue';

const emit = defineEmits<{
  (e: 'refresh'): void;
  (e: 'update-filters', filters: Record<string, string>): void;
}>();

const props = defineProps<{
  categories: CategoryResp[];
  sortField?: string;
  sortOrder?: 'asc' | 'desc';
  currentFilters: Record<string, string>;
}>();

const isFiltersOpen = ref(false);
const showModal = ref(false);
const modalMode = ref<'create' | 'edit'>('create');
const selectedCategory = ref<CategoryResp | undefined>(undefined);

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

const handleRowDoubleClick = (category: CategoryResp) => {
  selectedCategory.value = category;
  modalMode.value = 'edit';
  showModal.value = true;
};

const handleCreateClick = () => {
  selectedCategory.value = undefined;
  modalMode.value = 'create';
  showModal.value = true;
};

const handleModalSubmit = async (category: CategoryReq) =>{
  try{
    if (modalMode.value === 'create'){
      await categoryService.createCategory(category);
    } else if (selectedCategory.value){
      await categoryService.updateCategory(category);
    }
    showModal.value=false;
    emit('refresh');
  } catch(error) {
    console.error('Error saving category:', error);
  }
};

const handleDelete = async (category: CategoryReq)=>{
  if (confirm('Вы уверены, что хотите удалить этот семестр?')) {
    try {
      await categoryService.deleteCategories([category.id_category]);
      emit('refresh');
    } catch (error){
      console.error('Error deleting category:', error);
    }
  }
};
</script>

<template>
  <div class="table-window">
    <div class="table-container">
      <div class="table-header">
        <h1>Категории</h1>
        <div class="filters-wrapper">
          <button class="hamburger" :class="{ rotated: isFiltersOpen }" @click="toggleFilters">☰</button>
          <div class="filters-dropdown" v-show="isFiltersOpen">
            <CategoryFilterform @update-filters="handleFiltersUpdate" />
          </div>
        </div>
      </div>
      <div class="table-scroll">
        <table>
          <thead>
            <tr>
              <th @click="handleSort('id_category')" class="sortable" hidden>
                ID {{ getSortIcon('id_category') }}
              </th>
              <th @click="handleSort('achivments_type_category')" class="sortable">
                Тип достижения {{ getSortIcon('achivments_type_category') }}
              </th>
              <th @click="handleSort('score_category')" class="sortable">
                Баллы {{ getSortIcon('score_category') }}
              </th>
              <th>Действия</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="category in categories"
            :key="category.id_category"
            @dblclick="handleRowDoubleClick(category)"
            class="editable-row">
              <td hidden>{{ category.id_category }}</td>
              <td>{{ category.achivments_type_category }}</td>
              <td>{{ category.score_category }}</td>
              <td>
                <button class="action-button delete" @click="handleDelete(category)">Удалить</button>
              </td>
            </tr>
            <tr class="create-row" @click="handleCreateClick">
              <td colspan="3" class="create-cell">
                <span class="create-icon">+</span> Добавить новую категорию
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
    <CategoryFormModal
      :show="showModal"
      :mode="modalMode"
      :category="selectedCategory"
      @submit="handleModalSubmit"
      @close="showModal = false"
    />
</template>

<style scoped>
@import '../../assets/table.css';
</style>
