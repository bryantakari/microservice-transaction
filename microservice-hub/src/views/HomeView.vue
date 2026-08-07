<script setup lang="ts">
import { useRouter } from 'vue-router'
import { menuItems } from '@/constants/menuItems'

const router = useRouter()
</script>
<template>
  <v-container fluid class="py-8 dashboard-bg">
    <v-row>
      <v-col cols="12" class="mb-6">
        <h1 class="text-h4 mb-2">Dashboard</h1>
        <p class="text-subtitle-1">Quick access to your main sections.</p>
      </v-col>
    </v-row>

    <v-row dense>
      <v-col v-for="item in menuItems" :key="item.route" cols="12" sm="6" md="3">
        <v-card
          elevation="2"
          class="pa-6 dashboard-card"
          @click="router.push(item.route)"
          style="cursor: pointer"
        >
          <div class="d-flex align-center justify-space-between mb-4">
            <div>
              <div class="text-h6 mb-2">{{ item.title }}</div>
              <div class="text-body-2 text--secondary">{{ item.subtitle }}</div>
            </div>
            <v-icon size="40">{{ item.icon }}</v-icon>
          </div>

          <v-btn color="primary" variant="tonal" @click.stop="router.push(item.route)">
            Open
          </v-btn>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.dashboard-bg {
  min-height: 100vh;
  background-color: #111827;
  color: #f8fafc;
}

.dashboard-card {
  background-color: #1f2937;
  color: #f8fafc;
  transition: transform 0.15s ease, box-shadow 0.15s ease;
}

.dashboard-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 16px 40px rgba(15, 23, 42, 0.25);
}

.dashboard-card .text-body-2,
.dashboard-card .text-h6 {
  color: #e5e7eb;
}
</style>