<template>
  <v-app>
    <v-app-bar v-if="!route.meta.hideNavbar" density="comfortable">
      <v-app-bar-title ><v-btn href="/">Microservice Hub</v-btn></v-app-bar-title>
      <v-spacer />
      <v-btn v-if="auth.isAuthenticated" variant="text" color="white" @click="logout">
        Logout
      </v-btn>
      <v-btn v-else to="/login" variant="text" color="white">
        Login
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
import { useRoute, useRouter } from "vue-router";
import { useAuthStore } from "@/stores/authStores";

const auth = useAuthStore();
const route = useRoute();
const router = useRouter();
const currentTime = ref("");

function logout() {
  auth.logout();
  void router.replace({ name: "login" });
}

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