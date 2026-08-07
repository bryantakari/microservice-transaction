export interface LoginRequest {
  username: string;
  password: string;
}

export interface DataResponse<T>{
    data: T;
}

export interface LoginResponse {
  token: string;
  refreshToken: string;
}