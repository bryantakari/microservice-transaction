import api from "./service";
import type { DataResponse, LoginRequest, LoginResponse , RefreshTokenRequest} from "../types/auth";

export function login(request: LoginRequest) {
  return api.post<DataResponse<LoginResponse>>("/auth/login", request);
}
export function refreshingToken(request: RefreshTokenRequest) {
  return api.post<DataResponse<LoginResponse>>("/auth/refresh-token", request);
}
export function getCurrentUser() {
  return api.get("/user/me");
}