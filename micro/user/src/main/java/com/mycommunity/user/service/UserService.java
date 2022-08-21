package com.mycommunity.user.service;

import com.alibaba.fastjson.JSONObject;
import com.mycommunity.user.mapper.UserMapper;
import com.mycommunity.user.pojo.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.Map;

@Service
public class UserService {
    @Autowired
    private UserMapper userMapper;

    @Autowired
    private RestTemplate restTemplate;

    @Autowired
    private AuthService authService;

    public User queryById(Long id) {
        return userMapper.findById(id);
    }

    public String login(Map<String,String> data) {
        JSONObject json = restTemplate.postForEntity("http://localhost:9090/user/login",data, JSONObject.class).getBody();
        assert json != null;
        return json.getString("msg");
    }

    public String encodeMD5(String data) {
        return authService.encodeMD5(data);
    }
}
