package com.micropayment.userservice.service;

import com.micropayment.userservice.model.dto.JwtValidationDto;
import com.micropayment.userservice.model.entity.RefreshToken;
import com.micropayment.userservice.service.strategy.TokenType;

public interface TokenService {
       <T> String generateToken(TokenType tokenType,T payload);
       <T> JwtValidationDto<T> validateToken(String token,TokenType tokenType);

       RefreshToken consumeRefreshToken(String refreshToken);


}
