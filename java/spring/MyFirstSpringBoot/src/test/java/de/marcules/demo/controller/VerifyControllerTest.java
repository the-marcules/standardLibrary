package de.marcules.demo.controller;


import de.marcules.demo.model.SignResponse;
import de.marcules.demo.model.VerifyResponse;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.http.ResponseEntity;

import java.util.HashMap;
import java.util.Map;

import static de.marcules.demo.Service.Sign.signPayload;
import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class VerifyControllerTest {

    @Autowired
    private TestRestTemplate restTemplate;

    @Test
    public void testVerifyEndpointShouldReturnSuccess() throws Exception {

        String payload = "when the shit hits the fan";

        SignResponse signature = signPayload(payload);

        ResponseEntity<VerifyResponse> response = restTemplate.postForEntity("/verify", signature, VerifyResponse.class);

        Assertions.assertNotNull(response.getBody());
        assertThat(response.getBody().result).isEqualTo("success");
        assertThat(response.getBody().message).isEqualTo("Signature is valid");

    }

    @Test
    public void testVerifyEndpointShouldReturnFailure() throws Exception {

        String payload = "when the shit hits the fan";
        SignResponse signature = signPayload(payload);
        SignResponse signatureFalsy = new SignResponse("falseSignature", signature.certificate(), payload);

        ResponseEntity<VerifyResponse> response = restTemplate.postForEntity("/verify", signatureFalsy, VerifyResponse.class);

        Assertions.assertNotNull(response.getBody());
        assertThat(response.getBody().result).isEqualTo("failure");
        assertThat(response.getBody().message).isEqualTo("Signature is invalid");

    }

    @Test
    public void testVerifyEndpointWithShittyInput() {

        Map<String, String> shittyPayload = Map.of("someKey", "value");

        ResponseEntity<VerifyResponse> response = restTemplate.postForEntity("/verify", shittyPayload, VerifyResponse.class);

        Assertions.assertNotNull(response.getBody());
        assertThat(response.getBody().result).isEqualTo("failure");
        assertThat(response.getBody().message).isEqualTo("Signature is invalid");
    }


}