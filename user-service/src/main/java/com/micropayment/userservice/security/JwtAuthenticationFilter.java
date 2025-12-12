package com.micropayment.userservice.security;

import com.micropayment.userservice.common.exception.JwtAuthenticationException;
import com.micropayment.userservice.common.exception.ServiceException;
import com.micropayment.userservice.model.dto.JwtValidationDto;
import com.micropayment.userservice.model.dto.UserTokenPayload;
import com.micropayment.userservice.service.TokenService;
import com.micropayment.userservice.service.strategy.TokenType;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.util.List;

/**
 * User Jwt Authentication Filter.
 */
@Component
public class JwtAuthenticationFilter extends OncePerRequestFilter {

    private final TokenService tokenService;

    public JwtAuthenticationFilter(TokenService tokenService) {
        this.tokenService = tokenService;
    }

    @Override
    protected void doFilterInternal(HttpServletRequest request,
                                    HttpServletResponse response,
                                    FilterChain filterChain) throws ServletException, IOException {

        String header = request.getHeader("Authorization");

        try{
            if (header != null && header.startsWith("Bearer ")) {
                String token = header.substring(7);

                JwtValidationDto<UserTokenPayload> dto =
                        tokenService.validateToken(token, TokenType.ACCESS);

                // put user info into security context
                UsernamePasswordAuthenticationToken auth =
                        new UsernamePasswordAuthenticationToken(
                                dto.getPayload(), null, List.of() // or roles
                        );

                SecurityContextHolder.getContext().setAuthentication(auth);
            }
        }catch (Exception ex){
            throw new JwtAuthenticationException(ex.getMessage());
        }

        filterChain.doFilter(request, response);
    }
}
