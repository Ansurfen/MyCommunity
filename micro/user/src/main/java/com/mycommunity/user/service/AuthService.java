package com.mycommunity.user.service;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.stereotype.Service;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;

@Service
@FeignClient(value = "auth-service")
public interface AuthService {
    @PostMapping("/crypto/encode/md5")
    String encodeMD5(@RequestParam("data") String data);
}

