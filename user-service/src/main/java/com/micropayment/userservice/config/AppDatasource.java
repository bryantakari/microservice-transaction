package com.micropayment.userservice.config;

import com.micropayment.userservice.config.properties.DatasourceProperties;
import lombok.RequiredArgsConstructor;
import org.springframework.boot.flyway.autoconfigure.FlywayDataSource;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;

import javax.sql.DataSource;

@Configuration
@RequiredArgsConstructor
public class AppDatasource {
    private final AppProperties appProperties;

    @Bean
    @Primary // JPA / default app datasource
    public DataSource appDataSource() {
        DatasourceProperties datasourceProperties = appProperties.getDatasource();
        var cfg = new com.zaxxer.hikari.HikariConfig();
        cfg.setDriverClassName("org.postgresql.Driver");
        cfg.setJdbcUrl(datasourceProperties.getUrl());
        cfg.setUsername(datasourceProperties.getUsername());
        cfg.setPassword(datasourceProperties.getPassword());
        return new com.zaxxer.hikari.HikariDataSource(cfg);
    }
}
