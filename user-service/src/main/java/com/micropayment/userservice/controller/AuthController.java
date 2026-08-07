package com.micropayment.userservice.controller;

import com.micropayment.userservice.common.response.BaseValueResponse;
import com.micropayment.userservice.model.dto.LoginDto;
import com.micropayment.userservice.model.dto.RegisterDto;
import com.micropayment.userservice.model.request.LoginRequest;
import com.micropayment.userservice.model.request.RefreshTokenRequest;
import com.micropayment.userservice.model.request.RegisterRequest;
import com.micropayment.userservice.service.AuthService;
import jakarta.validation.Valid;
import lombok.AllArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/v1/auth")
@AllArgsConstructor
public class AuthController {

    private final AuthService authService;

    @PostMapping("/register")
    public ResponseEntity<BaseValueResponse<RegisterDto>> registerAccount(@Valid @RequestBody
                                                                          RegisterRequest request) {
        return new ResponseEntity<>(authService.registerAccount(request), HttpStatus.OK);
    }

    @PostMapping("/login")
    public ResponseEntity<BaseValueResponse<LoginDto>> loginAccount(@Valid @RequestBody
                                                                    LoginRequest request) {
        return new ResponseEntity<>(authService.loginAccount(request), HttpStatus.OK);
    }
    @PostMapping("/refresh-token")
    public ResponseEntity<BaseValueResponse<LoginDto>> refreshToken(
            @Valid @RequestBody RefreshTokenRequest request) {
        return new ResponseEntity<>(authService.refreshToken(request), HttpStatus.OK);
    }
}
