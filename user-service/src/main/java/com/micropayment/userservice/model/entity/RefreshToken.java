package com.micropayment.userservice.model.entity;

import jakarta.persistence.*;
import lombok.Data;

import java.time.LocalDateTime;

/**
 * Refresh Token Entity class.
 */
@Data
@Table(name = "refresh_tokens")
@Entity
public class RefreshToken {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id", nullable = false)
    private long id;
    @Column(name="user_id",nullable = false)
    private int userId;
    @Column(name = "token",nullable = false)
    private String token;
    @Column(name = "is_revoked",nullable = false)
    private boolean is_revoked;
    @Column(name = "expired_at")
    private LocalDateTime expiredAt;
    @Column(name = "created_at")
    private LocalDateTime createdAt;
}
