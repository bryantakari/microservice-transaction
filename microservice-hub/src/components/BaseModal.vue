<script setup lang="ts">
import { ref, watch } from "vue";

const props = withDefaults(
  defineProps<{
    modelValue: boolean;
    width?: string | number;
    maxWidth?: string | number;
    color?: string;
    hideCloseButton?: boolean;
    onClosed?: () => void;
  }>(),
  {
    width: "auto",
    maxWidth: 600,
    color: "surface",
    hideCloseButton: false,
    onClosed: undefined,
  },
);

const emit = defineEmits<{
  "update:modelValue": [value: boolean];
  "close-request": [];
}>();

const isOpen = ref(props.modelValue);

watch(
  () => props.modelValue,
  (value) => {
    isOpen.value = value;
  },
);

watch(isOpen, (value, previousValue) => {
  emit("update:modelValue", value);
  if (previousValue && !value) {
    props.onClosed?.();
  }
});

function requestClose() {
  emit("close-request");
}
</script>

<template>
  <v-dialog
    v-model="isOpen"
    :width="width"
    :max-width="maxWidth"
    persistent
    @click:outside="requestClose"
  >
    <v-card :color="color" class="base-modal">
      <v-btn
        v-if="!hideCloseButton"
        icon="mdi-close"
        variant="text"
        aria-label="Close dialog"
        class="base-modal__close"
        @click="requestClose"
      />
      <slot />
    </v-card>
  </v-dialog>
</template>

<style scoped>
.base-modal {
  position: relative;
}

.base-modal__close {
  position: absolute;
  top: 8px;
  right: 8px;
  z-index: 1;
}
</style>
