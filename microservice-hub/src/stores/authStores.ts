import { defineStore } from "pinia";

type JwtPayload = {
  exp?: number;
  [key: string]: unknown;
};

function decodeJwtPayload(token: string): JwtPayload | null {
  try {
    const parts = token.split(".");
    if (parts.length !== 3) return null;

    const payload = parts[1]
      .replace(/-/g, "+")
      .replace(/_/g, "/");

    const decoded = atob(payload);
    return JSON.parse(decodeURIComponent(
      decoded.split("").map((c) => `%${("00" + c.charCodeAt(0).toString(16)).slice(-2)}`).join("")
    ));
  } catch {
    return null;
  }
}

function getTokenExpiration(token: string): number | null {
  const payload = decodeJwtPayload(token);
  if (!payload || typeof payload.exp !== "number") {
    return null;
  }
  return payload.exp;
}

function isTokenValid(token: string): boolean {
  if (!token) return false;

  const exp = getTokenExpiration(token);
  if (exp === null) return false;

  return Date.now() < exp * 1000;
}

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: localStorage.getItem("token") || "",
    refreshToken: localStorage.getItem("refreshToken") || "",
  }),

  getters: {
    isAuthenticated: (state) => isTokenValid(state.token),
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

    validateToken() {
      if (!isTokenValid(this.token)) {
        this.logout();
        return false;
      }
      return true;
    },
  },
});