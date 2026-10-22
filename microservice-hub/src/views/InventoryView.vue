<template>
  <v-container fluid class="inventory-page py-8">
    <div class="page-header">
      <div>
        <p class="eyebrow">CATALOG MANAGEMENT</p>
        <h1>Inventory</h1>
        <p class="subtitle">Browse and manage the items available in your inventory.</p>
      </div>
      <v-btn color="primary" class="create-button" @click="openCreateModal">
        <svg class="button-icon" viewBox="0 0 24 24" aria-hidden="true">
          <path d="M12 5v14M5 12h14" />
        </svg>
        Create item
      </v-btn>
    </div>

    <v-card rounded="xl" elevation="0" class="inventory-card">
      <div class="filters">
        <div class="filter-title">
          <h2>Items</h2>
          <span>{{ items.length }} on this page</span>
        </div>
        <div class="filter-controls">
          <v-select
            v-model="sortBy"
            :items="sortOptions"
            item-title="label"
            item-value="value"
            label="Sort by"
            density="comfortable"
            hide-details
            class="sort-select"
          />
          <v-select
            v-model="orderBy"
            :items="orderOptions"
            item-title="label"
            item-value="value"
            label="Order"
            density="comfortable"
            hide-details
            class="order-select"
          />
          <v-select
            v-model="limit"
            :items="pageSizeOptions"
            label="Rows"
            density="comfortable"
            hide-details
            class="rows-select"
          />
        </div>
      </div>

      <v-alert v-if="errorMessage" type="error" variant="tonal" class="mx-5 mt-2">
        {{ errorMessage }}
      </v-alert>

      <v-progress-linear v-if="loading" indeterminate color="primary" />

      <v-table class="inventory-table">
        <thead>
          <tr>
            <th scope="col">Item</th>
            <th scope="col">Item code</th>
            <th scope="col">Price</th>
            <th scope="col">Quantity</th>
            <th scope="col" class="actions-heading">Details</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in items" :key="`${item.item_code}-${index}`">
            <td>
              <div class="item-name">{{ item.item_name || 'Unnamed item' }}</div>
              <div class="item-description">{{ item.item_description || 'No description' }}</div>
            </td>
            <td><span class="code-badge">{{ item.item_code || '—' }}</span></td>
            <td>{{ formatPrice(item.price) }}</td>
            <td>
              <span class="quantity-badge" :class="{ 'quantity-low': item.quantity <= 5 }">
                {{ item.quantity }}
              </span>
            </td>
            <td class="actions-cell">
              <v-btn
                icon
                variant="tonal"
                color="primary"
                size="small"
                :aria-label="`View details for ${item.item_name || 'item'}`"
                @click="openDetails(item)"
              >
                <svg class="eye-icon" viewBox="0 0 24 24" aria-hidden="true">
                  <path d="M2 12s3.6-7 10-7 10 7 10 7-3.6 7-10 7S2 12 2 12Z" />
                  <circle cx="12" cy="12" r="3" />
                </svg>
              </v-btn>
            </td>
          </tr>
          <tr v-if="!loading && items.length === 0">
            <td colspan="5" class="empty-state">
              {{ errorMessage ? 'Inventory could not be loaded.' : 'No inventory items found.' }}
            </td>
          </tr>
        </tbody>
      </v-table>

      <div class="table-footer">
        <span>Page {{ page }}</span>
        <div class="pagination-controls">
          <v-btn variant="text" :disabled="page <= 1 || loading" @click="page--">Previous</v-btn>
          <v-btn variant="text" :disabled="items.length < limit || loading" @click="page++">Next</v-btn>
        </div>
      </div>
    </v-card>

    <BaseModal
      v-model="showModal"
      :width="700"
      :max-width="900"
      color="#1f2937"
      :hide-close-button="false"
      :on-closed="handleModalClosed"
      @close-request="handleModalCloseRequest"
    >
      <InventoryItemDetails
        v-if="modalContent === 'details' && selectedItem"
        :item="selectedItem"
      />
      <CreateInventoryItem
        v-if="modalContent === 'create'"
        ref="createItemRef"
        @created="handleItemCreated"
        @discarded="showModal = false"
      />
    </BaseModal>
  </v-container>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { listItems } from '@/api/inventory';
import type { ItemsDTO, ListItemsQuery } from '@/types/inventory';
import BaseModal from '@/components/BaseModal.vue';
import CreateInventoryItem from '@/components/CreateInventoryItem.vue';
import InventoryItemDetails from '@/components/InventoryItemDetails.vue';

const items = ref<ItemsDTO[]>([]);
const loading = ref(false);
const errorMessage = ref('');
const page = ref(1);
const limit = ref(10);
const sortBy = ref('created_at');
const orderBy = ref<ListItemsQuery['order_by']>('desc');
const selectedItem = ref<ItemsDTO | null>(null);
const showModal = ref(false);
const modalContent = ref<'create' | 'details' | null>(null);
const createItemRef = ref<{ requestDiscard: () => void } | null>(null);

const sortOptions = [
  { label: 'Date added', value: 'created_at' },
  { label: 'Item name', value: 'item_name' },
  { label: 'Item code', value: 'item_code' },
  { label: 'Price', value: 'price' },
  { label: 'Quantity', value: 'quantity' },
];
const orderOptions = [  { label: 'Newest / highest first', value: 'desc' },
  { label: 'Oldest / lowest first', value: 'asc' },
];
const pageSizeOptions = [10, 25, 50];

async function loadItems() {
  loading.value = true;
  errorMessage.value = '';

  const query: ListItemsQuery = {
    page: page.value,
    limit: limit.value,
    sort_by: sortBy.value,
    order_by: orderBy.value,
  };

  try {
    const response = await listItems(query);
    items.value = response.data.data ?? [];
  } catch {
    items.value = [];
    errorMessage.value = 'Unable to load inventory. Please try again.';
  } finally {
    loading.value = false;
  }
}

function openDetails(item: ItemsDTO) {
  selectedItem.value = item;
  modalContent.value = 'details';
  showModal.value = true;
}

function openCreateModal() {
  selectedItem.value = null;
  modalContent.value = 'create';
  showModal.value = true;
}

function handleModalClosed() {
  modalContent.value = null;
  selectedItem.value = null;
}

function handleModalCloseRequest() {
  if (modalContent.value === 'create') {
    createItemRef.value?.requestDiscard();
    return;
  }

  showModal.value = false;
}

function handleItemCreated() {
  showModal.value = false;
  void loadItems();
}

function formatPrice(price: number) {
  return new Intl.NumberFormat(undefined, { style: 'currency', currency: 'USD' }).format(price);
}

watch([limit, sortBy, orderBy], () => {
  if (page.value !== 1) {
    page.value = 1;
  } else {
    void loadItems();
  }
});
watch(page, loadItems, { immediate: true });
</script>

<style scoped>
.inventory-page {
  display: flex;
  min-height: calc(100vh - 64px);
  flex-direction: column;
  background: #111827;
  color: #f8fafc;
}

.page-header,
.filters,
.table-footer,
.pagination-controls {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 20px;
}

.page-header {
  margin-bottom: 28px;
}

.eyebrow {
  margin: 0 0 8px;
  color: #9ca3af;
  font-size: 0.72rem;
  font-weight: 700;
  letter-spacing: 0.12em;
}

h1,
h2,
p {
  margin-top: 0;
}

h1 {
  margin-bottom: 6px;
  font-size: clamp(1.8rem, 3vw, 2.3rem);
  font-weight: 700;
  letter-spacing: -0.04em;
}

.subtitle {
  margin-bottom: 0;
  color: #cbd5e1;
}

.create-button {
  display: inline-flex;
  gap: 8px;
  border-radius: 10px;
  background: #0b1220 !important;
  color: #f8fafc !important;
  text-transform: none;
  font-weight: 600;
}

.create-button:hover {
  background: #030712 !important;
}

.button-icon {
  width: 18px;
  height: 18px;
  fill: none;
  stroke: currentColor;
  stroke-linecap: round;
  stroke-width: 2;
}

.inventory-card {
  display: flex;
  min-height: 580px;
  flex: 1;
  flex-direction: column;
  overflow: hidden;
  border: 1px solid #374151;
  background: #1f2937;
  color: #f8fafc;
}

.filters {
  min-height: 100px;
  padding: 22px 24px;
  border-bottom: 1px solid #374151;
}

.filter-title h2 {
  margin-bottom: 4px;
  font-size: 1.1rem;
}

.filter-title span {
  color: #9ca3af;
  font-size: 0.85rem;
}

.filter-controls {
  display: flex;
  align-items: center;
  gap: 12px;
}

.sort-select { width: 175px; }
.order-select { width: 205px; }
.rows-select { width: 105px; }

.filter-controls :deep(.v-field) {
  background: #273449;
  color: #f8fafc;
}

.filter-controls :deep(.v-field__input),
.filter-controls :deep(.v-label),
.filter-controls :deep(.v-field__append-inner) {
  color: #e5e7eb;
}

.inventory-table {
  min-height: 360px;
  flex: 1;
  background: #1f2937;
  color: #f8fafc;
}

.inventory-table :deep(.v-table__wrapper) {
  height: 100%;
}

.inventory-table :deep(table) {
  height: 100%;
}

.inventory-table :deep(th) {
  height: 52px !important;
  color: #cbd5e1 !important;
  background: #111827;
  font-size: 0.75rem !important;
  font-weight: 700 !important;
  letter-spacing: 0.04em;
  text-transform: uppercase;
  white-space: nowrap;
}

.inventory-table :deep(td) {
  height: 72px !important;
  border-bottom: 1px solid #374151 !important;
  color: #f3f4f6;
}

.inventory-table :deep(tbody tr) {
  background: #1f2937;
  transition: background-color 0.15s ease;
}

.inventory-table :deep(tbody tr:hover) {
  background: #273449;
}

.item-name { font-weight: 600; }
.item-description {
  max-width: 440px;
  overflow: hidden;
  color: #9ca3af;
  font-size: 0.82rem;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.code-badge {
  padding: 5px 9px;
  border-radius: 6px;
  background: #374151;
  color: #e5e7eb;
  font-size: 0.82rem;
}

.quantity-badge {
  display: inline-block;
  min-width: 34px;
  padding: 5px 8px;
  border-radius: 20px;
  background: #14532d;
  color: #dcfce7;
  text-align: center;
  font-weight: 600;
}

.quantity-low {
  background: #78350f;
  color: #ffedd5;
}

.actions-heading,
.actions-cell { text-align: center !important; }
.eye-icon {
  width: 19px;
  height: 19px;
  fill: none;
  stroke: currentColor;
  stroke-linecap: round;
  stroke-linejoin: round;
  stroke-width: 1.8;
}

.empty-state {
  height: 140px !important;
  color: #9ca3af;
  text-align: center;
}

.table-footer {
  padding: 12px 24px;
  color: #cbd5e1;
  font-size: 0.875rem;
}

.pagination-controls { gap: 4px; }
.pagination-controls :deep(.v-btn) { text-transform: none; }

@media (max-width: 700px) {
  .page-header,
  .filters { align-items: stretch; flex-direction: column; }
  .page-header { gap: 16px; }
  .create-button { align-self: flex-start; }
  .filter-controls { flex-wrap: wrap; }
  .sort-select { flex: 1 1 160px; }
  .order-select { flex: 1 1 180px; }
  .rows-select { flex: 1 1 90px; }
  .inventory-table :deep(table) { min-width: 650px; }
  .inventory-card { overflow-x: auto; }
  .inventory-card { min-height: 580px; }
}
</style>
