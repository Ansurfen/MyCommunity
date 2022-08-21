package com.mycommunity.user.mapper;

import com.mycommunity.user.pojo.User;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;

@Mapper
public interface UserMapper {
    @Select("select * from users where id = #{id}")
    User findById(@Param("id") Long id);
}
