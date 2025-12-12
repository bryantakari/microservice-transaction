package com.micropayment.userservice.security;

import com.micropayment.userservice.common.exception.ApplicationErrorCode;
import com.micropayment.userservice.common.exception.ServiceException;
import com.micropayment.userservice.common.exception.dto.BaseErrorResponse;
import com.micropayment.userservice.common.exception.dto.ErrorMessage;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.MediaTypeEditor;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.AuthenticationEntryPoint;
import org.springframework.stereotype.Component;
import org.springframework.web.context.request.ServletWebRequest;
import tools.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.util.Date;

@Component
public class JwtAuthenticationEntryPoint implements AuthenticationEntryPoint {

    @Override
    public void commence(HttpServletRequest request,
                         HttpServletResponse response,
                         AuthenticationException authException) throws IOException {

        ErrorMessage errorMessage = ErrorMessage.builder().timestamp(new Date())
                .code(ApplicationErrorCode.UNAUTHORIZED.getErrorCode())
                .error(ApplicationErrorCode.UNAUTHORIZED)
                .message(authException.getMessage())
                .path(request.getRequestURI()).build();
        BaseErrorResponse authResponse =
                BaseErrorResponse.builder().error(errorMessage).build();
        response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);
        response.setContentType(MediaType.APPLICATION_JSON_VALUE);
        response.getWriter().write(new ObjectMapper().writeValueAsString(authResponse));
    }
}
