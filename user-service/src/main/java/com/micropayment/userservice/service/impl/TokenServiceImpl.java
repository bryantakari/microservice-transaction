package com.micropayment.userservice.service.impl;

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
    public <T> String generateToken(TokenType tokenType,T payload) {
        TokenStrategy<T> strategy = getStrategy(tokenType);
        return strategy.generate(payload);
    }

    @Override
    public <T> JwtValidationDto<T> validateToken(String token,TokenType tokenType) {
        TokenStrategy<T> strategy = getStrategy(tokenType);
        return strategy.validate(token);
    }

    @Override
    public RefreshToken getRefreshToken(String refreshToken) {
        return null;
    }

    @Override
    public boolean invalidRefreshToken(String refreshToken) {
        return false;
    }

    @SuppressWarnings("unchecked")
    private <T> TokenStrategy<T> getStrategy(TokenType type) {
        return (TokenStrategy<T>) strategies.stream()
                .filter(s -> s.support(type))
                .findFirst()
                .orElseThrow(() -> new RuntimeException("Unsupported token type"));
    }


}
