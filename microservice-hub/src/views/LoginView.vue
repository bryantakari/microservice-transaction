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
    if(response.status === 200){
      auth.setTokens(
        response.data.data.token,
        response.data.data.refreshToken
      );
      router.push("/");
    }else{
      errorMessage.value = "Error Server Please try again later.";
    }
    
  } catch (error) {
    errorMessage.value = "Invalid username or password.";
    console.error(error);
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <v-container class="fill-height login-page-bg" fluid>
    <v-row justify="center" align="center" class="fill-height">
      <v-col cols="12" sm="8" md="5" lg="4">
        <v-card elevation="10" class="login-card">
          <v-card-title class="text-h5 text-center login-card-title">
            Login
          </v-card-title>

          <v-card-text>
            <v-form @submit.prevent="loginUser" class="login-form">
              <v-text-field
                v-model="username"
                label="Username"
                prepend-inner-icon="mdi-account"
                variant="outlined"
                density="comfortable"
                hide-details="auto"
                rounded="lg"
                class="login-field"
                required
              />

              <v-text-field
                v-model="password"
                label="Password"
                prepend-inner-icon="mdi-lock"
                variant="outlined"
                density="comfortable"
                hide-details="auto"
                rounded="lg"
                :type="showPassword ? 'text' : 'password'"
                :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
                @click:append-inner="showPassword = !showPassword"
                class="login-field"
                required
              />

              <v-alert
                v-if="errorMessage"
                type="error"
                density="compact"
                class="login-alert"
              >
                {{ errorMessage }}
              </v-alert>

              <v-btn
                type="submit"
                color="primary"
                block
                :loading="loading"
                :disabled="loading"
                class="login-button"
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

<style scoped>
.login-page-bg {
  min-height: 100vh;
  background: radial-gradient(circle at top, #1f2937 0%, #111827 55%, #0f172a 100%);
}

.login-card {
  border-radius: 32px;
  padding: 2rem 1.25rem;
  background: rgba(31, 41, 55, 0.96);
  color: #f8fafc;
  box-shadow: 0 20px 60px rgba(15, 23, 42, 0.35);
}

.login-card-title {
  padding-top: 0.5rem;
  padding-bottom: 0.75rem;
  color: #f8fafc;
}

.login-form {
  display: grid;
  gap: 1rem;
}

.login-field {
  width: 100%;
}

.login-field .v-field__input {
  color: #f8fafc;
}

.login-field .v-text-field__slot {
  background: rgba(255, 255, 255, 0.04);
}

.login-alert {
  margin-top: 0.5rem;
}

.login-button {
  margin-top: 0.5rem;
}
</style>
