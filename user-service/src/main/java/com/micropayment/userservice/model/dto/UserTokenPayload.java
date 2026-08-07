package com.micropayment.userservice.model.dto;

import lombok.Data;

import java.util.UUID;

/**
 * User Token Payload dto.
 */
@Data
public class UserTokenPayload {
    private UUID userId;
}
