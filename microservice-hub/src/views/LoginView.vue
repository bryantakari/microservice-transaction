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
  <v-container fluid class="login-page pa-4 pa-sm-8">
    <div class="login-shell">
      <section class="welcome-panel">
        <div class="brand-mark" aria-hidden="true">
          <svg viewBox="0 0 40 40" fill="none">
            <rect x="3" y="3" width="34" height="34" rx="11" fill="currentColor" fill-opacity=".14" />
            <path d="M12 20h16M20 12v16M13.5 13.5l13 13M26.5 13.5l-13 13" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
          </svg>
        </div>
        <p class="panel-eyebrow">MICROSERVICE HUB</p>
        <h1>Everything connected.<br /><span>One place to manage it.</span></h1>
        <p class="panel-copy">Sign in to access your services, monitor activity, and keep your operations moving.</p>
        <div class="status-line"><span class="status-dot"></span> Your workspace, at a glance</div>
      </section>

      <v-card elevation="0" class="login-card">
        <div class="form-heading">
          <p class="form-eyebrow">WELCOME BACK</p>
          <h2>Sign in</h2>
          <p>Enter your account details to continue.</p>
        </div>

        <v-form @submit.prevent="loginUser" class="login-form">
          <v-text-field
            v-model="username"
            label="Username"
            placeholder="Enter your username"
            prepend-inner-icon="mdi-account-outline"
            variant="outlined"
            density="comfortable"
            hide-details="auto"
            rounded="lg"
            class="login-field"
            autocomplete="username"
            required
          />

          <v-text-field
            v-model="password"
            label="Password"
            placeholder="Enter your password"
            prepend-inner-icon="mdi-lock-outline"
            variant="outlined"
            density="comfortable"
            hide-details="auto"
            rounded="lg"
            :type="showPassword ? 'text' : 'password'"
            :append-inner-icon="showPassword ? 'mdi-eye-off-outline' : 'mdi-eye-outline'"
            @click:append-inner="showPassword = !showPassword"
            class="login-field"
            autocomplete="current-password"
            required
          />

          <v-alert v-if="errorMessage" type="error" variant="tonal" density="compact" class="login-alert">
            {{ errorMessage }}
          </v-alert>

          <v-btn
            type="submit"
            color="primary"
            block
            size="large"
            :loading="loading"
            :disabled="loading"
            class="login-button"
          >
            Sign in
            <template #append>
              <v-icon icon="mdi-arrow-right" />
            </template>
          </v-btn>
        </v-form>

        <div class="demo-credentials">
          <div class="demo-icon"><v-icon icon="mdi-flask-outline" size="18" /></div>
          <div>
            <p class="demo-title">Test account</p>
            <p><span>Username</span><strong>root</strong></p>
            <p><span>Password</span><strong>root1234</strong></p>
          </div>
        </div>

        <p class="security-note"><v-icon icon="mdi-shield-check-outline" size="16" /> Secure access to your workspace</p>
      </v-card>
    </div>
  </v-container>
</template>

<style scoped>
.login-page {
  display: grid;
  min-height: 100vh;
  place-items: center;
  background:
    radial-gradient(ellipse at 12% 12%, rgba(59, 130, 246, 0.12), transparent 35%),
    radial-gradient(ellipse at 88% 85%, rgba(20, 184, 166, 0.08), transparent 34%),
    #111827;
  color: #f8fafc;
}

.login-shell {
  display: grid;
  width: min(100%, 1000px);
  grid-template-columns: 1fr 0.88fr;
  overflow: hidden;
  border: 1px solid rgba(148, 163, 184, 0.18);
  border-radius: 26px;
  background: #1f2937;
  box-shadow: 0 28px 80px rgba(0, 0, 0, 0.35);
}

.welcome-panel {
  display: flex;
  min-height: 570px;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
  padding: clamp(32px, 6vw, 64px);
  background:
    radial-gradient(circle at 85% 18%, rgba(96, 165, 250, 0.15), transparent 28%),
    linear-gradient(145deg, #243449, #172333 80%);
}

.brand-mark {
  width: 48px;
  height: 48px;
  margin-bottom: 48px;
  color: #93c5fd;
}
.brand-mark svg { width: 100%; height: 100%; }
.panel-eyebrow,
.form-eyebrow {
  margin: 0 0 14px;
  color: #93c5fd;
  font-size: 0.72rem;
  font-weight: 700;
  letter-spacing: 0.16em;
}
.welcome-panel h1 {
  margin: 0;
  font-size: clamp(2rem, 4vw, 3.1rem);
  font-weight: 700;
  letter-spacing: -0.055em;
  line-height: 1.12;
}
.welcome-panel h1 span { color: #93c5fd; }
.panel-copy {
  max-width: 370px;
  margin: 22px 0 0;
  color: #cbd5e1;
  font-size: 1rem;
  line-height: 1.7;
}
.status-line {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 48px;
  color: #d1d5db;
  font-size: 0.88rem;
}
.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #34d399;
  box-shadow: 0 0 14px rgba(52, 211, 153, 0.65);
}

.login-card {
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: clamp(30px, 5vw, 56px);
  background: #1f2937;
  color: #f8fafc;
}

.form-heading { margin-bottom: 28px; }
.form-heading h2 {
  margin: 0 0 8px;
  font-size: 2rem;
  font-weight: 700;
  letter-spacing: -0.04em;
}
.form-heading > p:last-child {
  margin: 0;
  color: #9ca3af;
}

.login-form {
  display: grid;
  gap: 18px;
}

.login-field :deep(.v-field) {
  background: #273449;
  color: #f8fafc;
}
.login-field :deep(.v-field__input),
.login-field :deep(.v-label),
.login-field :deep(.v-icon) {
  color: #e5e7eb;
}
.login-field :deep(.v-field__outline) {
  color: #4b5563;
}
.login-field :deep(input::placeholder) {
  color: #9ca3af;
  opacity: 1;
}
.login-field :deep(.v-field:not(.v-field--focused) input::placeholder) {
  opacity: 0;
}
.login-button {
  min-height: 52px;
  margin-top: 4px;
  border-radius: 12px;
  background: #0b1220 !important;
  color: #f8fafc !important;
  font-weight: 600;
  letter-spacing: 0.01em;
  text-transform: none;
}
.login-button:hover { background: #030712 !important; }
.login-alert { margin: 0; }

.demo-credentials {
  display: flex;
  gap: 13px;
  margin-top: 26px;
  padding: 16px;
  border: 1px solid #374151;
  border-radius: 14px;
  background: #182333;
}
.demo-icon {
  display: grid;
  width: 34px;
  height: 34px;
  flex: 0 0 auto;
  place-items: center;
  border-radius: 10px;
  background: rgba(96, 165, 250, 0.13);
  color: #93c5fd;
}
.demo-credentials p { margin: 0; }
.demo-title {
  margin-bottom: 8px !important;
  color: #e5e7eb;
  font-size: 0.86rem;
  font-weight: 700;
}
.demo-credentials p:not(.demo-title) {
  display: flex;
  justify-content: space-between;
  gap: 20px;
  color: #9ca3af;
  font-size: 0.82rem;
  line-height: 1.7;
}
.demo-credentials strong {
  color: #f3f4f6;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-weight: 600;
}
.security-note {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 7px;
  margin: 22px 0 0;
  color: #9ca3af;
  font-size: 0.78rem;
}

@media (max-width: 760px) {
  .login-page { min-height: calc(100vh - 56px); }
  .login-shell { max-width: 480px; grid-template-columns: 1fr; }
  .welcome-panel { min-height: auto; padding: 30px 30px 26px; }
  .brand-mark { width: 40px; height: 40px; margin-bottom: 24px; }
  .welcome-panel h1 { font-size: 2rem; }
  .panel-copy { margin-top: 14px; font-size: 0.92rem; }
  .status-line { margin-top: 20px; }
  .login-card { padding: 32px 30px; }
}

@media (max-width: 420px) {
  .login-page { padding: 12px !important; }
  .login-shell { border-radius: 20px; }
  .welcome-panel,
  .login-card { padding-right: 22px; padding-left: 22px; }
}
</style>
