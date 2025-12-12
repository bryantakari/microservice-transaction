package com.micropayment.userservice.service.strategy.impl;

import com.micropayment.userservice.config.AppProperties;
import com.micropayment.userservice.model.dto.UserTokenPayload;
import com.micropayment.userservice.model.entity.RefreshToken;
import com.micropayment.userservice.repository.RefreshTokenRepository;
import com.micropayment.userservice.service.strategy.AbstractTokenStrategy;
import com.micropayment.userservice.service.strategy.TokenType;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.Instant;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.Date;

/**
 * refresh strategy pattern.
 */
@Service("REFRESH")
public class RefreshTokenStrategy extends AbstractTokenStrategy<UserTokenPayload> {
    @Autowired
    private RefreshTokenRepository refreshTokenRepository;


    protected RefreshTokenStrategy(AppProperties appProperties) {
        super(appProperties);
    }

    @Override
    public boolean support(TokenType type) {
        return TokenType.REFRESH == type;
    }
    @Override
    protected String expectedType() {
        return TokenType.REFRESH.name();
    }
    @Override
    public String generate(Object subject) {
        long expiryMillis = this.getJwtProperties().getRefreshTokenExpiryMillis();
        Instant issuedAt = Instant.now();
        Instant expirationInstant = issuedAt.plusMillis(expiryMillis);

        String token = Jwts.builder()
                .subject(String.valueOf(subject))
                .claim(TYPE_CONSTANT,TokenType.REFRESH)
                .issuedAt(Date.from(issuedAt))
                .expiration(Date.from(expirationInstant))
                .signWith(this.getKey())
                .compact();

        createRefreshToken(
                token,
                Integer.parseInt(subject.toString()),
                expirationInstant
        );

        return token;
    }

    private void createRefreshToken(String token, int userId, Instant expirationInstant) {
        LocalDateTime expirationLocalDateTime = LocalDateTime.ofInstant(
                expirationInstant,
                ZoneId.systemDefault()
        );
        RefreshToken refreshToken = new RefreshToken();
        refreshToken.setToken(token);
        refreshToken.setUserId(userId);
        refreshToken.setCreatedAt(LocalDateTime.now());
        refreshToken.setExpiredAt(expirationLocalDateTime);
        refreshTokenRepository.save(refreshToken);
    }

    @Override
    protected UserTokenPayload mapToPayload(Claims claims) {
        UserTokenPayload payload = new UserTokenPayload();
        payload.setUserId(Integer.parseInt(claims.getSubject()));
        return payload;
    }

}
