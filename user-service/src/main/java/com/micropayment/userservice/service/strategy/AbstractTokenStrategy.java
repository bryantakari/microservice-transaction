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
import org.springframework.beans.factory.annotation.Autowired;

import javax.crypto.SecretKey;
import java.nio.charset.StandardCharsets;
import java.security.PrivateKey;
import java.security.PublicKey;

/**
 * Abstract Token Strategy.
 */
public abstract class AbstractTokenStrategy<T> implements TokenStrategy<T> {
    @Autowired
    private AppProperties appProperties;
    @Autowired
    protected PrivateKey privateKey;
    @Autowired
    protected PublicKey publicKey;
    protected final String TYPE_CONSTANT = "type";

    protected Claims validateRaw(String token) {
        try {
            Claims claims = Jwts.parser()
                    .verifyWith(publicKey)
                    .build()
                    .parseSignedClaims(token)
                    .getPayload();
            Object type = claims.get(TYPE_CONSTANT);
            if( null == type || !claims.get(TYPE_CONSTANT).equals(expectedType())){
                throw new ServiceException("Token is not valid!", ApplicationErrorCode.UNAUTHORIZED);
            }
            return claims;
        } catch (Exception e) {
            throw new ServiceException("Token is not valid!", ApplicationErrorCode.UNAUTHORIZED);
        }
    }
    protected abstract String expectedType();
    protected abstract T mapToPayload(Claims claims);
    protected JwtProperties getJwtProperties(){
        return appProperties.getJwt();
    }
    @Override
    public JwtValidationDto<T> validate(String token) {
        T payload = mapToPayload(validateRaw(token));
        return JwtValidationDto.<T>builder().isValid(true).payload(payload).build();
    }
}
