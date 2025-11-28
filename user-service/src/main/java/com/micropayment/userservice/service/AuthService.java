package com.micropayment.userservice.service;

import com.micropayment.userservice.common.response.BaseValueResponse;
import com.micropayment.userservice.model.dto.LoginDto;
import com.micropayment.userservice.model.dto.RegisterDto;
import com.micropayment.userservice.model.request.LoginRequest;
import com.micropayment.userservice.model.request.RegisterRequest;

public interface AuthService {
    BaseValueResponse<RegisterDto> registerAccount(RegisterRequest request);

    BaseValueResponse<LoginDto> loginAccount(LoginRequest request);

}
