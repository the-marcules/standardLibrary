package de.marcules.demo.controller;


import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.http.ResponseEntity;

import java.util.Map;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@SpringBootTest (webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class HelloControllerITest {

    @Autowired
    private TestRestTemplate restTemplate;

    @Test
    public void testGetHello() throws Exception {
        String expectedMessage = "Greetings from Spring Boot!";

        ResponseEntity<Map> response = restTemplate.getForEntity("/", Map.class);
        assertThat(response.getBody().get("message")).isEqualTo(expectedMessage);
    }
}