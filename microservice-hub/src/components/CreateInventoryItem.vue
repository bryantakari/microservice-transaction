<script setup lang="ts">
import { ref } from "vue";
import { insertItems } from "@/api/inventory";
import type { ItemsDTO } from "@/types/inventory";
import BaseModal from "@/components/BaseModal.vue";

const emit = defineEmits<{
  created: [];
  discarded: [];
}>();

const form = ref<{ validate: () => Promise<{ valid: boolean }> } | null>(null);
const item = ref<ItemsDTO>({
  item_name: "",
  item_description: "",
  price: 0,
  item_code: "",
  quantity: 0,
});
const loading = ref(false);
const errorMessage = ref("");
const showCreateConfirmation = ref(false);
const showDiscardConfirmation = ref(false);

const requiredRule = (value: string) => !!value || "This field is required.";
const nonNegativeRule = (value: number) =>
  Number.isFinite(value) && value >= 0 || "Enter a non-negative number.";

async function createItem() {
  const result = await form.value?.validate();
  if (!result?.valid) {
    return;
  }

  showCreateConfirmation.value = true;
}

async function confirmCreate() {
  loading.value = true;
  errorMessage.value = "";
  showCreateConfirmation.value = false;

  try {
    await insertItems({ items: [item.value] });
    emit("created");
  } catch (error) {
    console.error("Failed to create inventory item:", error);
    errorMessage.value = "Unable to create item. Please try again.";
  } finally {
    loading.value = false;
  }
}

function requestDiscard() {
  if (!loading.value) {
    showDiscardConfirmation.value = true;
  }
}

function confirmDiscard() {
  showDiscardConfirmation.value = false;
  emit("discarded");
}

defineExpose({ requestDiscard });
</script>

<template>
  <section class="create-item-form">
    <header class="form-heading">
      <p class="eyebrow">CATALOG MANAGEMENT</p>
      <h2>Create new item</h2>
      <p>Enter the item details to add it to your inventory.</p>
    </header>

    <v-form ref="form" @submit.prevent="createItem">
      <div class="form-fields">
        <v-text-field
          v-model="item.item_name"
          label="Item name"
          :rules="[requiredRule]"
          required
        />
        <v-text-field
          v-model="item.item_code"
          label="Item code"
          :rules="[requiredRule]"
          required
        />
        <v-textarea
          v-model="item.item_description"
          label="Description"
          rows="3"
        />
        <v-text-field
          v-model.number="item.price"
          label="Price"
          type="number"
          min="0"
          step="1"
          :rules="[nonNegativeRule]"
          required
        />
        <v-text-field
          v-model.number="item.quantity"
          label="Quantity"
          type="number"
          min="0"
          step="1"
          :rules="[nonNegativeRule]"
          required
        />
      </div>

      <v-alert v-if="errorMessage" type="error" variant="tonal" class="mb-4">
        {{ errorMessage }}
      </v-alert>

      <div class="form-actions">
        <v-btn variant="text" :disabled="loading" @click="requestDiscard">
          Cancel
        </v-btn>
        <v-btn type="submit" color="primary" :loading="loading" :disabled="loading">
          Create item
        </v-btn>
      </div>
    </v-form>
  </section>

  <BaseModal
    v-model="showCreateConfirmation"
    :width="440"
    :max-width="440"
    color="#1f2937"
    @close-request="showCreateConfirmation = false"
  >
    <section class="confirmation-content">
      <p class="eyebrow">CONFIRM CREATION</p>
      <h2>Create this inventory item?</h2>
      <p>
        <strong>{{ item.item_name }}</strong> will be added to your inventory.
      </p>
      <div class="confirmation-actions">
        <v-btn variant="text" @click="showCreateConfirmation = false">Go back</v-btn>
        <v-btn color="primary" :loading="loading" @click="confirmCreate">
          Confirm create
        </v-btn>
      </div>
    </section>
  </BaseModal>

  <BaseModal
    v-model="showDiscardConfirmation"
    :width="440"
    :max-width="440"
    color="#1f2937"
    @close-request="showDiscardConfirmation = false"
  >
    <section class="confirmation-content">
      <p class="eyebrow">DISCARD CHANGES</p>
      <h2>Discard this item?</h2>
      <p>Your entered details will be lost if you discard them.</p>
      <div class="confirmation-actions">
        <v-btn variant="text" @click="showDiscardConfirmation = false">
          Keep editing
        </v-btn>
        <v-btn color="error" @click="confirmDiscard">Discard</v-btn>
      </div>
    </section>
  </BaseModal>
</template>

<style scoped>
.create-item-form {
  padding: 30px;
  color: #f8fafc;
}

.form-heading {
  margin-bottom: 26px;
}

.eyebrow {
  margin: 0 0 7px;
  color: #93c5fd;
  font-size: 0.72rem;
  font-weight: 700;
  letter-spacing: 0.14em;
}

.form-heading h2,
.confirmation-content h2 {
  margin: 0 0 8px;
  font-size: 1.45rem;
  font-weight: 700;
}

.form-heading > p:last-child,
.confirmation-content > p:last-of-type {
  margin: 0;
  color: #cbd5e1;
  line-height: 1.6;
}

.form-fields {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 8px 16px;
}

.form-fields > :nth-child(3) {
  grid-column: 1 / -1;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 8px;
}

.confirmation-content {
  padding: 28px;
  color: #f8fafc;
}

.confirmation-content > p:last-of-type {
  margin-top: 12px;
}

.confirmation-actions {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 26px;
}

@media (max-width: 560px) {
  .create-item-form,
  .confirmation-content {
    padding: 22px;
  }

  .form-fields {
    grid-template-columns: 1fr;
  }

  .form-fields > :nth-child(3) {
    grid-column: auto;
  }
}
</style>
