package com.mycommunity.user.pojo;

import lombok.Data;

@Data
public class User {
    private String username;
    private String alias;
    private String password;
    private String telephone;
    private String email;
    private String school;
    private String communities;
    private String image;
    private Long id;
    private char right;
}

