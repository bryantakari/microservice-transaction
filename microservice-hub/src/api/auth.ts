import { userApi } from "./service";
import type { DataResponse, LoginRequest, LoginResponse , RefreshTokenRequest} from "../types/auth";

export function login(request: LoginRequest) {
  return userApi.post<DataResponse<LoginResponse>>("/auth/login", request);
}
export function refreshingToken(request: RefreshTokenRequest) {
  return userApi.post<DataResponse<LoginResponse>>("/auth/refresh-token", request);
}
export function getCurrentUser() {
  return userApi.get("/user/me");
}