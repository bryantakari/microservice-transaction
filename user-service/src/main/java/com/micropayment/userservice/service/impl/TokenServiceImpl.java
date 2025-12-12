package com.micropayment.userservice.service.impl;

import com.micropayment.userservice.common.exception.ApplicationErrorCode;
import com.micropayment.userservice.common.exception.ServiceException;
import com.micropayment.userservice.config.AppProperties;
import com.micropayment.userservice.config.properties.JwtProperties;
import com.micropayment.userservice.model.dto.JwtValidationDto;
import com.micropayment.userservice.model.entity.RefreshToken;
import com.micropayment.userservice.repository.RefreshTokenRepository;
import com.micropayment.userservice.service.TokenService;
import com.micropayment.userservice.service.strategy.TokenStrategy;
import com.micropayment.userservice.service.strategy.TokenType;
import io.jsonwebtoken.security.Keys;
import jakarta.annotation.PostConstruct;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import javax.crypto.SecretKey;
import java.nio.charset.StandardCharsets;
import java.util.List;
import java.util.Optional;

/**
 * Token Service implementation.
 */
@Service
@RequiredArgsConstructor
public class TokenServiceImpl implements TokenService {
    private final List<TokenStrategy<?>> strategies;
    private final RefreshTokenRepository refreshTokenRepository;
    private final AppProperties appProperties;
    private JwtProperties jwtProperties;
    private SecretKey key;

    @PostConstruct
    public void init() {
        this.jwtProperties = appProperties.getJwt();
        this.key = Keys.hmacShaKeyFor(jwtProperties.getSecret().getBytes(StandardCharsets.UTF_8));
    }

    @Override
    public <T> String generateToken(TokenType tokenType, T payload) {
        try {
            TokenStrategy<T> strategy = getStrategy(tokenType);
            return strategy.generate(payload);
        } catch (Exception ex) {
            throw new ServiceException("Generate Token failed!", ApplicationErrorCode.INTERNAL_SERVER_ERROR);
        }
    }

    @Override
    public <T> JwtValidationDto<T> validateToken(String token, TokenType tokenType) {
        TokenStrategy<T> strategy = getStrategy(tokenType);
        return strategy.validate(token);
    }

    @Override
    public RefreshToken consumeRefreshToken(String refreshToken) {
        RefreshToken token = refreshTokenRepository
                .findValidToken(refreshToken)
                .orElseThrow(() -> new ServiceException("Refresh token invalid",ApplicationErrorCode.ITEM_NOT_FOUND));
        token.setRevoked(true);
        refreshTokenRepository.save(token);

        return token;
    }

    @SuppressWarnings("unchecked")
    private <T> TokenStrategy<T> getStrategy(TokenType type) {
        return (TokenStrategy<T>) strategies.stream()
                .filter(s -> s.support(type))
                .findFirst()
                .orElseThrow(() -> new RuntimeException("Unsupported token type"));
    }


}
