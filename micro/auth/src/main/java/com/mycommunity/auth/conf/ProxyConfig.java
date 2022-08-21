package com.mycommunity.auth.conf;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

import java.util.List;

@Data
@Component
@ConfigurationProperties("proxy")
public class ProxyConfig {
    private List<String> addr;
    private boolean balanced;
}
