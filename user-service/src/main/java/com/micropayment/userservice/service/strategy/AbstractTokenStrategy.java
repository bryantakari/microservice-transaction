package com.micropayment.userservice.service.strategy;

import com.micropayment.userservice.common.exception.ApplicationErrorCode;
import com.micropayment.userservice.common.exception.ServiceException;
import com.micropayment.userservice.config.AppProperties;
import com.micropayment.userservice.config.properties.JwtProperties;
import com.micropayment.userservice.model.dto.JwtValidationDto;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;
import lombok.Getter;

import javax.crypto.SecretKey;
import java.nio.charset.StandardCharsets;

/**
 * Abstract Token Strategy.
 */
@Getter
public abstract class AbstractTokenStrategy<T> implements TokenStrategy<T> {
    private final AppProperties appProperties;
    private JwtProperties jwtProperties;
    private SecretKey key;

    protected AbstractTokenStrategy(AppProperties appProperties) {
        this.appProperties = appProperties;
        this.jwtProperties = appProperties.getJwt();
        this.key = Keys.hmacShaKeyFor(jwtProperties.getSecret().getBytes(StandardCharsets.UTF_8));
    }

    protected Claims validateRaw(String token) {
        try {
            Claims claims = Jwts.parser()
                    .verifyWith(key)
                    .build()
                    .parseSignedClaims(token)
                    .getPayload();

            return claims;
        } catch (Exception e) {
            throw new ServiceException("Token is not valid!", ApplicationErrorCode.UNAUTHORIZED);
        }
    }

    protected abstract T mapToPayload(Claims claims);

    @Override
    public JwtValidationDto<T> validate(String token) {
        T payload = mapToPayload(validateRaw(token));
        return JwtValidationDto.<T>builder().isValid(true).payload(payload).build();
    }
}
