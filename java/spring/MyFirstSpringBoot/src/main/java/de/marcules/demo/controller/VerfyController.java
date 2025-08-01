package de.marcules.demo.controller;

import de.marcules.demo.Service.Verify;
import de.marcules.demo.model.SignResponse;
import de.marcules.demo.model.VerifyResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import static de.marcules.demo.Service.Verify.verifySignature;

@RestController
public class VerfyController {
    private static final Logger logger = LoggerFactory.getLogger(VerfyController.class);

    @PostMapping(value = "/verify", consumes = "application/json")
    public VerifyResponse verify(@RequestBody SignResponse request) {

        try {
            boolean isValid = verifySignature(request);
            if (isValid) {
                return new VerifyResponse("success", "Signature is valid");
            } else {
                return new VerifyResponse("failure", "Signature is invalid");
            }

        } catch (Exception e) {
            logger.error("Error while verifying", e);
            throw new org.springframework.web.server.ResponseStatusException(
                    org.springframework.http.HttpStatus.BAD_REQUEST, "Invalid request", e);
        }


    }


}
