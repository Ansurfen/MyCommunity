package com.mycommunity.user.web;

import com.mycommunity.user.pojo.User;
import com.mycommunity.user.service.AuthService;
import com.mycommunity.user.service.UserService;
import lombok.extern.slf4j.Slf4j;
import org.apache.ibatis.annotations.Param;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

@Slf4j
@RestController
public class UserController {
    @Autowired
    private UserService userService;

    @Autowired
    private AuthService authService;

    @PostMapping("/login")
    public String login(@RequestBody Map<String, String> params) {
        return userService.login(params);
    }

    @PostMapping("/md5")
    public String encodeMD5(@Param("data") String data) {
        return authService.encodeMD5(data);
    }
}
