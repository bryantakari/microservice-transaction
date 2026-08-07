package com.micropayment.userservice.config.properties;

import lombok.Getter;
import lombok.Setter;
import org.springframework.beans.factory.annotation.Value;

/**
 * JWT configuration properties.
 */
@Getter
@Setter
public class JwtProperties {
    private String secret;
    private Long accessTokenExpiryMillis;
    private Long refreshTokenExpiryMillis;
    private String publicKeyPath;
    private String privateKeyPath;
}
