package com.micropayment.userservice.common.exception;


import org.springframework.security.core.AuthenticationException;

/**
 * Jwt auth exception passed to entry point exception.
 */
public class JwtAuthenticationException extends AuthenticationException {
    public JwtAuthenticationException(String message) {
        super(message);
    }
}