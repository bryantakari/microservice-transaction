package com.micropayment.userservice.model.dto;

import lombok.Data;

import java.util.UUID;

/**
 * Register Dto class for register response.
 */
@Data
public class RegisterDto {
    private UUID userId;
    private String username;
}
