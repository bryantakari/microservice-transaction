import axios from "axios";
import { useAuthStore } from "../stores/authStores";

const userServiceBaseURL =
  import.meta.env.VITE_USER_SERVICE_URL || "http://localhost:8080/v1";
let refreshPromise: Promise<string> | null = null;

type TokenResponse = {
  data: {
    token: string;
    refreshToken: string;
  };
};

export function refreshSession(): Promise<string> {
  const authStore = useAuthStore();
  if (!authStore.refreshToken) {
    return Promise.reject(new Error("No refresh token is available"));
  }

  if (!refreshPromise) {
    const refreshToken = authStore.refreshToken;
    refreshPromise = axios
      .post<TokenResponse>(
        `${userServiceBaseURL.replace(/\/$/, "")}/auth/refresh-token`,
        { refreshToken },
        { timeout: 10000 },
      )
      .then(({ data }) => {
        authStore.setTokens(data.data.token, data.data.refreshToken);
        return data.data.token;
      })
      .finally(() => {
        refreshPromise = null;
      });
  }

  return refreshPromise;
}

function redirectToLogin() {
  const loginUrl = `${import.meta.env.BASE_URL.replace(/\/$/, "")}/login`;
  if (window.location.pathname !== loginUrl) {
    window.location.assign(loginUrl);
  }
}

function createServiceClient(baseURL: string) {
  const client = axios.create({
    baseURL,
    timeout: 10000,
  });

  client.interceptors.request.use((config) => {
    const authStore = useAuthStore();

    if (authStore.token) {
      config.headers.Authorization = `Bearer ${authStore.token}`;
    }

    return config;
  });

  client.interceptors.response.use(
    (response) => response,
    async (error) => {
      if (error.response?.status !== 401) {
        return Promise.reject(error);
      }

      const request = error.config as (typeof error.config & { _retry?: boolean }) | undefined;
      if (!request || request.url?.includes("/auth/")) {
        return Promise.reject(error);
      }

      const authStore = useAuthStore();
      if (request._retry) {
        authStore.logout();
        redirectToLogin();
        return Promise.reject(error);
      }

      request._retry = true;
      try {
        const token = await refreshSession();
        request.headers.Authorization = `Bearer ${token}`;
        return client.request(request);
      } catch (refreshError) {
        authStore.logout();
        redirectToLogin();
        return Promise.reject(refreshError);
      }
    },
  );

  return client;
}

export const userApi = createServiceClient(
  userServiceBaseURL,
);
export const orderApi = createServiceClient(
  import.meta.env.VITE_ORDER_SERVICE_URL || "http://localhost:3001",
);
export const inventoryApi = createServiceClient(
  import.meta.env.VITE_INVENTORY_SERVICE_URL || "http://localhost:3000",
);