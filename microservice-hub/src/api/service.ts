import axios from "axios";
import {useAuthStore} from "../stores/authStores"
export const api = axios.create({
  baseURL: "http://localhost:8080/v1",
  timeout: 10000,
});
api.interceptors.request.use((config) => {
  const authStore = useAuthStore();

  if (authStore.token) {
    config.headers.Authorization = `Bearer ${authStore.token}`;
  }

  return config;
});

api.interceptors.response.use((response) => response, (error) => {
  if (error.response && error.response.status === 401) {
    const authStore = useAuthStore();
    authStore.logout();
  }
  return Promise.reject(error);
});
export default api