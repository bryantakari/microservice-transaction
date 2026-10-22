<script setup lang="ts">
import type { ItemsDTO } from "@/types/inventory";

defineProps<{
  item: ItemsDTO;
}>();

function formatPrice(price: number) {
  return new Intl.NumberFormat(undefined, {
    style: "currency",
    currency: "USD",
  }).format(price);
}
</script>

<template>
  <div class="item-details">
    <header class="item-details__heading">
      <p class="eyebrow">CATALOG MANAGEMENT</p>
      <h2 class="item-details__title">Item details</h2>
      <p class="item-details__name">{{ item.item_name || "Unnamed item" }}</p>
    </header>
    <p class="item-details__description">
      {{ item.item_description || "No description provided." }}
    </p>
    <dl class="item-details__grid">
      <div>
        <dt>Item code</dt>
        <dd>{{ item.item_code || "—" }}</dd>
      </div>
      <div>
        <dt>Price</dt>
        <dd>{{ formatPrice(item.price) }}</dd>
      </div>
      <div>
        <dt>Quantity</dt>
        <dd>{{ item.quantity }}</dd>
      </div>
    </dl>
  </div>
</template>

<style scoped>
.item-details {
  padding: 30px;
  color: #f8fafc;
}

.item-details__heading {
  margin-bottom: 24px;
}

.eyebrow {
  margin: 0 0 7px;
  color: #93c5fd;
  font-size: 0.72rem;
  font-weight: 700;
  letter-spacing: 0.14em;
}

.item-details__title {
  margin: 0 0 6px;
  font-size: 1.35rem;
}

.item-details__name {
  margin: 0;
  color: #cbd5e1;
}

.item-details__description {
  margin: 0;
  padding: 16px;
  border: 1px solid #374151;
  border-radius: 12px;
  background: #182333;
  color: #cbd5e1;
  line-height: 1.6;
}

.item-details__grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 16px;
  margin: 24px 0 0;
}

.item-details__grid dt {
  margin-bottom: 5px;
  color: #9ca3af;
  font-size: 0.8rem;
}

.item-details__grid dd {
  margin: 0;
  font-weight: 600;
}
</style>
