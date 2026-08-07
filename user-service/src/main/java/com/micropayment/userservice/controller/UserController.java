package com.micropayment.userservice.controller;

import com.micropayment.userservice.common.response.BaseValueResponse;
import com.micropayment.userservice.model.dto.UserInfoDto;
import com.micropayment.userservice.service.UserService;
import lombok.AllArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * User Controller class.
 */
@RestController
@RequestMapping("/v1/user")
@AllArgsConstructor
public class UserController {

    private final UserService userService;

    @PostMapping("/me")
    public ResponseEntity<BaseValueResponse<UserInfoDto>> getMe() {
        return new ResponseEntity<>(userService.userInfo(), HttpStatus.OK);
    }

    @PostMapping("/list-menu")
    public ResponseEntity<BaseValueResponse<UserInfoDto>> getListMenu() {
        return new ResponseEntity<>(userService.userInfo(), HttpStatus.OK);
    }
}
