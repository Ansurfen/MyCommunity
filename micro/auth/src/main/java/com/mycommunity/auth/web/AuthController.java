package com.mycommunity.auth.web;

import com.mycommunity.auth.conf.ProxyConfig;
import com.mycommunity.auth.service.AuthService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class AuthController {
    @Autowired
    private AuthService authService;

    @GetMapping("/addr")
    public String getAddr() {
        return authService.GetAddr();
    }

//    @PostMapping("/jwt/release")
//    public String releaseJWT() {
//
//    }
//
//    @PostMapping("/jwt/parser")
//    public String parserJWT() {
//
//    }

    @PostMapping("/crypto/encode/aes")
    public String encodeWithAES(@RequestParam("data") String data, @RequestParam("key") String key) {
        return authService.encodeWithAES(data, key);
    }

    @PostMapping("/crypto/encode/md5")
    public String encodeWithMD5(@RequestParam("data") String data) {
        return authService.encodeWithMD5(data);
    }

    @PostMapping("/crypto/encode/sha256")
    public String encodeWithSHA256(@RequestParam("data") String data) {
        return authService.encodeWithSHA256(data);
    }

    @PostMapping("/crypto/decode/aes")
    public String decodeWithAES(@RequestParam("data") String data, @RequestParam("key") String key) {
        return authService.decodeWithAES(data, key);
    }
}
