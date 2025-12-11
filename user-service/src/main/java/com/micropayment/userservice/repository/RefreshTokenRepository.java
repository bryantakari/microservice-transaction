package com.micropayment.userservice.repository;

import com.micropayment.userservice.model.entity.RefreshToken;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

/**
 * RefreshToken Repository class.
 */
@Repository
public interface RefreshTokenRepository extends JpaRepository<RefreshToken,Long> {
}
