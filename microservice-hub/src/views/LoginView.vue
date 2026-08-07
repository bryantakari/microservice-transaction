<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"

import { login } from "@/api/auth"
import { useAuthStore } from "@/stores/authStores"

const router = useRouter();
const auth = useAuthStore();

const showPassword = ref(false);

const username = ref("");
const password = ref("");

const loading = ref(false);
const errorMessage = ref("");


async function loginUser() {
  try {
    loading.value = true;
    errorMessage.value = "";

    const response = await login({
      username: username.value,
      password: password.value,
    })

    auth.setTokens(
      response.data.data.token,
      response.data.data.refreshToken
    );

    router.push("/");
  } catch (error) {
    errorMessage.value = "Invalid username or password.";
    console.error(error);
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <v-container class="fill-height" fluid>
    <v-row justify="center" align="center">
      <v-col cols="12" sm="8" md="5" lg="4">
        <v-card elevation="8" rounded="lg">
          <v-card-title class="text-h5 text-center pt-6">
            Login
          </v-card-title>

          <v-card-text>
            <v-form @submit.prevent="loginUser">
              <v-text-field
                v-model="username"
                label="Username"
                prepend-inner-icon="mdi-account"
                variant="outlined"
                class="mb-4"
                required
              />

              <v-text-field
                v-model="password"
                label="Password"
                prepend-inner-icon="mdi-lock"
                variant="outlined"
                :type="showPassword ? 'text' : 'password'"
                :append-inner-icon="
                  showPassword ? 'mdi-eye-off' : 'mdi-eye'
                "
                @click:append-inner="showPassword = !showPassword"
                class="mb-4"
                required
              />

              <v-alert
                v-if="errorMessage"
                type="error"
                density="compact"
                class="mb-4"
              >
                {{ errorMessage }}
              </v-alert>

              <v-btn
                type="submit"
                color="primary"
                block
                :loading="loading"
                :disabled="loading"
              >
                Login
              </v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
