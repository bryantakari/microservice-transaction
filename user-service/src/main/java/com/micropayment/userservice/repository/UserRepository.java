package com.micropayment.userservice.repository;

import com.micropayment.userservice.model.entity.Account;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;
import java.util.UUID;

/**
 * This is jpa repository of user repository.
 */
@Repository
public interface UserRepository extends JpaRepository<Account, UUID> {
    // Find by username (unique)
    Optional<Account> findByUsername(String username);
}
