import { defineStore } from "pinia";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: localStorage.getItem("token") || "",
    refreshToken: localStorage.getItem("refreshToken") || "",
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
  },

  actions: {
    setTokens(token: string, refreshToken: string) {
      this.token = token;
      this.refreshToken = refreshToken;

      localStorage.setItem("token", token);
      localStorage.setItem("refreshToken", refreshToken);
    },

    logout() {
      this.token = "";
      this.refreshToken = "";

      localStorage.removeItem("token");
      localStorage.removeItem("refreshToken");
    },
  },
});