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

import java.time.Duration;
import java.time.LocalDateTime;
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
    public String generate(Object subject) {
        return Jwts.builder()
                .subject(subject.toString())
                .issuedAt(new Date())
                .expiration(
                        new Date((new Date()).getTime() + this.getJwtProperties().getRefreshTokenExpiryMillis()))
                .signWith(this.getKey())
                .compact();
    }

    private RefreshToken createRefreshToken(String token, int subject) {
        LocalDateTime expired = LocalDateTime.now()
                .plus(Duration.ofMillis(this.getJwtProperties().getRefreshTokenExpiryMillis()));
        RefreshToken refreshToken = new RefreshToken();
        refreshToken.setToken(token);
        refreshToken.setUserId(subject);
        refreshToken.setCreatedAt(LocalDateTime.now());
        refreshToken.setExpiredAt(expired);
        refreshTokenRepository.save(refreshToken);
        return refreshToken;
    }

    @Override
    protected UserTokenPayload mapToPayload(Claims claims) {
        UserTokenPayload payload = new UserTokenPayload();
        payload.setUserId(Integer.parseInt(claims.getSubject()));
        return payload;
    }

}
