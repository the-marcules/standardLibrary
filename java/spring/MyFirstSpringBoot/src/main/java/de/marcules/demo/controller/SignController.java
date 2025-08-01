package de.marcules.demo.controller;

import de.marcules.demo.model.SignResponse;
import de.marcules.demo.Service.Sign;
import de.marcules.demo.model.SignRequest;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SignController {

    @PostMapping("/sign")
    public SignResponse sign(@RequestBody SignRequest request) {
        try {
            String payload = request.message();
            return Sign.signPayload(payload);
        } catch (Exception e) {
            throw new org.springframework.web.server.ResponseStatusException(
                    org.springframework.http.HttpStatus.BAD_REQUEST, "Ungültige Anfrage", e);
        }

    }


}
