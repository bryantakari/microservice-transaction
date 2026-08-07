package com.micropayment.userservice.model.dto;

import lombok.Data;

import java.util.UUID;

/**
 * Jwt Token Dto class.
 */
@Data
public class JwtDto {
    private UUID userId;
    private String username;
    private String fullname;
}
