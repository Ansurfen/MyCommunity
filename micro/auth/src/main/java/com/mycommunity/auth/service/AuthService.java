package com.mycommunity.auth.service;

import com.alibaba.fastjson2.JSONObject;
import com.mycommunity.auth.conf.ProxyConfig;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;
import org.springframework.web.client.RestTemplate;

import java.util.*;

@Service
public class AuthService {

    @Autowired
    private ProxyConfig proxyConfig;

    @Autowired
    private RestTemplate restTemplate;

    private HashMap<String, Integer> rule = new HashMap<>();

    private int factor = 0;

    public String GetAddr() {
        List<String> address = proxyConfig.getAddr();
        int n = new Random().nextInt(address.size());
        String ret = address.get(n);
        if (proxyConfig.isBalanced()) {
            if (rule == null) address.forEach(addr -> rule.put(addr, factor));
//            int count = rule.get(ret);
//            if (count == factor) {
//                factor++;
//                rule.put(ret, factor);
//            } else if (rule.get(ret) > factor) {
// total = 4 , factor=1; total / cur == factor ;
//            }
            System.out.println("触发负载均衡");
        }
        return ret;
    }

    public String CallService(Map<String, String> data) {
        JSONObject json = restTemplate.postForEntity(this.GetAddr(), data, JSONObject.class).getBody();

        return this.GetAddr();
    }

    public String encodeWithAES(String data, String key) {
        MultiValueMap<Object,Object> params = new LinkedMultiValueMap<>();
        params.add("data",data);
        params.add("key",key);
        return restTemplate.postForEntity(this.GetAddr()+"/crypto/encode/aes",params,String.class).getBody();
    }

    public String encodeWithMD5(String data) {
        MultiValueMap<Object,Object> params = new LinkedMultiValueMap<>();
        params.add("data",data);
        return restTemplate.postForEntity(this.GetAddr()+"/crypto/encode/md5",params,String.class).getBody();
    }

    public String encodeWithSHA256(String data) {
        MultiValueMap<Object,Object> params = new LinkedMultiValueMap<>();
        params.add("data",data);
        return restTemplate.postForEntity(this.GetAddr()+"/crypto/encode/sha256",params,String.class).getBody();
    }

    public String decodeWithAES(String data,String  key) {
        MultiValueMap<Object,Object> params = new LinkedMultiValueMap<>();
        params.add("data",data);
        params.add("key",key);
        return restTemplate.postForEntity(this.GetAddr()+"/crypto/decode/aes",params,String.class).getBody();
    }
}
