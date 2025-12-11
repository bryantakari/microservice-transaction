package com.micropayment.userservice.service.impl;

import com.micropayment.userservice.common.JwtHelper;
import com.micropayment.userservice.common.exception.ApplicationErrorCode;
import com.micropayment.userservice.common.exception.ServiceException;
import com.micropayment.userservice.common.response.BaseValueResponse;
import com.micropayment.userservice.mapper.UserMapper;
import com.micropayment.userservice.model.dto.*;
import com.micropayment.userservice.model.entity.Account;
import com.micropayment.userservice.model.request.LoginRequest;
import com.micropayment.userservice.model.request.RefreshTokenRequest;
import com.micropayment.userservice.model.request.RegisterRequest;
import com.micropayment.userservice.service.AuthService;
import com.micropayment.userservice.service.TokenService;
import com.micropayment.userservice.service.UserService;
import com.micropayment.userservice.service.strategy.TokenType;
import lombok.AllArgsConstructor;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

/**
 * Auth Service Implementation.
 */
@Service
@AllArgsConstructor
public class AuthServiceImpl implements AuthService {

    private final UserService userService;
    private final UserMapper userMapper;
    private final JwtHelper jwtHelper;
    private final PasswordEncoder passwordEncoder;
    private final TokenService tokenService;

    @Override
    public BaseValueResponse<RegisterDto> registerAccount(RegisterRequest request) {
        Account account = userMapper.mappingRegisterToAccount(request);
        account.setPassword(passwordEncoder.encode(account.getPassword()));
        return BaseValueResponse.<RegisterDto>builder()
                .data(userMapper.mappingAccountToDto(userService.insertAccount(account))).build();
    }

    @Override
    public BaseValueResponse<LoginDto> loginAccount(LoginRequest request) {
        Account user = userService.getAccountByUsername(request.getUsername())
                .orElseThrow(() -> new ServiceException("Invalid username or password",
                        ApplicationErrorCode.BAD_REQUEST));

        if (!passwordEncoder.matches(request.getPassword(), user.getPassword())) {
            throw new ServiceException("Invalid username or password",
                    ApplicationErrorCode.BAD_REQUEST);
        }

        LoginDto loginDto = new LoginDto();
        loginDto.setToken(tokenService.generateToken(TokenType.ACCESS, user.getId()));
        loginDto.setRefreshToken(tokenService.generateToken(TokenType.REFRESH, user.getId()));

        return BaseValueResponse.<LoginDto>builder()
                .data(loginDto)
                .build();
    }

    @Override
    public BaseValueResponse<LoginDto> refreshToken(RefreshTokenRequest request) {
        JwtValidationDto<UserTokenPayload> validationDto = tokenService.validateToken(request.getRefreshToken(), TokenType.REFRESH);
        if (!validationDto.isValid()) {
            throw new ServiceException("Invalid refresh token",
                    ApplicationErrorCode.UNAUTHORIZED);
        }

        // Invalidate old refresh token in DB here

        String newAccess = tokenService.generateToken(TokenType.ACCESS, validationDto.getPayload().getUserId());
        String newRefresh = tokenService.generateToken(TokenType.REFRESH, validationDto.getPayload().getUserId());

        LoginDto dto = new LoginDto();
        dto.setToken(newAccess);
        dto.setRefreshToken(newRefresh);

        return BaseValueResponse.<LoginDto>builder()
                .data(dto)
                .build();
    }

}