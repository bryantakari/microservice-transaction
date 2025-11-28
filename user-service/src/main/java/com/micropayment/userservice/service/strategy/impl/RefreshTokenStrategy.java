package com.micropayment.userservice.service.strategy.impl;

import com.micropayment.userservice.config.AppProperties;
import com.micropayment.userservice.model.dto.UserTokenPayload;
import com.micropayment.userservice.service.strategy.AbstractTokenStrategy;
import com.micropayment.userservice.service.strategy.TokenType;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import org.springframework.stereotype.Service;

import java.util.Date;

/**
 * refresh strategy pattern.
 */
@Service("REFRESH")
public class RefreshTokenStrategy extends AbstractTokenStrategy<UserTokenPayload> {

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

    @Override
    protected UserTokenPayload mapToPayload(Claims claims) {
        UserTokenPayload payload = new UserTokenPayload();
        payload.setUserId(Integer.parseInt(claims.getSubject()));
        return payload;
    }

}
