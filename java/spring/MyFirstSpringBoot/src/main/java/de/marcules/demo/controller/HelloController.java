package de.marcules.demo.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;

@RestController
public class HelloController {

    @GetMapping("/")
    public Map<String,String> index() {
        String msg = "Greetings from Spring Boot!";

        Map<String, String> response = new HashMap<>();
        response.put("message", msg);





        return response;
    }
}
