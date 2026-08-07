package com.micropayment.userservice.service.strategy.impl;

import com.micropayment.userservice.config.AppProperties;
import com.micropayment.userservice.model.dto.UserTokenPayload;
import com.micropayment.userservice.service.strategy.AbstractTokenStrategy;
import com.micropayment.userservice.service.strategy.TokenType;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import org.springframework.stereotype.Service;

import java.security.PrivateKey;
import java.util.Date;
import java.util.UUID;

@Service("ACCESS")
public class AccessTokenStrategy extends AbstractTokenStrategy<UserTokenPayload> {

    @Override
    protected String expectedType() {
        return TokenType.ACCESS.name();
    }

    @Override
    protected UserTokenPayload mapToPayload(Claims claims) {
        UserTokenPayload payload = new UserTokenPayload();
        payload.setUserId(UUID.fromString(claims.getSubject()));
        return payload;
    }

    @Override
    public boolean support(TokenType type) {
        return TokenType.ACCESS == type;
    }

    @Override
    public String generate(Object subject) {
        return Jwts.builder()
                .subject(subject.toString())
                .claim(TYPE_CONSTANT,TokenType.ACCESS)
                .issuedAt(new Date())
                .expiration(
                        new Date((new Date()).getTime() + this.getJwtProperties().getAccessTokenExpiryMillis()))
                .signWith(privateKey)
                .compact();
    }



}
