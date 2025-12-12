package com.micropayment.userservice.repository;

import com.micropayment.userservice.model.entity.RefreshToken;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.Optional;

/**
 * RefreshToken Repository class.
 */
@Repository
public interface RefreshTokenRepository extends JpaRepository<RefreshToken, Long> {

    @Query("SELECT r FROM RefreshToken r WHERE r.token = :token AND r.isRevoked = false")
    Optional<RefreshToken> findValidToken(@Param("token") String token);
}
