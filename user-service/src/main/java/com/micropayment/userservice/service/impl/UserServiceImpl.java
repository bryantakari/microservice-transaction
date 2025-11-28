package com.micropayment.userservice.service.impl;

import com.micropayment.userservice.common.exception.ApplicationErrorCode;
import com.micropayment.userservice.common.exception.ServiceException;
import com.micropayment.userservice.common.response.BaseValueResponse;
import com.micropayment.userservice.mapper.UserMapper;
import com.micropayment.userservice.model.dto.UserInfoDto;
import com.micropayment.userservice.model.dto.UserTokenPayload;
import com.micropayment.userservice.model.entity.Account;
import com.micropayment.userservice.repository.UserRepository;
import com.micropayment.userservice.service.TokenService;
import com.micropayment.userservice.service.UserService;
import lombok.AllArgsConstructor;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

import java.util.Optional;

@Service
@AllArgsConstructor
public class UserServiceImpl implements UserService {

    private final UserRepository repository;
    private final UserMapper userMapper;

    private final TokenService tokenService;

    @Override
    public Account insertAccount(Account account) {
        return repository.save(account);
    }

    @Override
    public Optional<Account> getAccountByUsername(String username) {
        return repository.findByUsername(username);
    }

    @Override
    public BaseValueResponse<UserInfoDto> userInfo() {
        UserTokenPayload payload = (UserTokenPayload)
                SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        Account acc = repository.findById(payload.getUserId())
                .orElseThrow(() -> new ServiceException("User not valid", ApplicationErrorCode.UNAUTHORIZED));
        return BaseValueResponse.<UserInfoDto>builder().data(userMapper.mappingToUserInfoDto(acc)).build();
    }

}
