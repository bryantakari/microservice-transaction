<template>
  <v-app>
    <v-app-bar density="comfortable">
      <v-app-bar-title>Microservice Hub</v-app-bar-title>
      <v-spacer />
      <v-btn
        v-for="item in items"
        :key="item.title"
        :to="item.to"
        variant="text"
        color="white"
      >
        {{ item.title }}
      </v-btn>
      <span class="ml-3 text-caption">{{ currentTime }}</span>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";
const currentTime = ref("");
const items = [
  { title: "Home", to: "/" },
  { title: "Login", to: "/login" },
];

let timer: number | undefined;

function updateTime() {
  const now = new Date();

  currentTime.value = now.toLocaleTimeString("en-GB", {
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
  });
}
onMounted(() => {
  updateTime();
  timer = window.setInterval(updateTime, 1000);
});

onUnmounted(() => {
  if (timer) {
    clearInterval(timer);
  }
});

</script>