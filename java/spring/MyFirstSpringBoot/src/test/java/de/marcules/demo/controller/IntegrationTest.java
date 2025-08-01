package de.marcules.demo.controller;

import de.marcules.demo.model.SignRequest;
import de.marcules.demo.model.SignResponse;
import de.marcules.demo.model.VerifyResponse;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.http.ResponseEntity;

import static de.marcules.demo.Service.Sign.signPayload;
import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class IntegrationTest {

    @Autowired
    public TestRestTemplate restTemplate;

    @Test
    public void shouldSignAndVerifyPayloadAsValid() throws Exception {
        SignRequest request = new SignRequest("When the shit hits the fan");

        ResponseEntity<SignResponse> signResponse = this.restTemplate.postForEntity("/sign", request, SignResponse.class);

        Assertions.assertNotNull(signResponse.getBody());
        assertThat(signResponse.getStatusCode().is2xxSuccessful()).isTrue();


        ResponseEntity<VerifyResponse> response = restTemplate.postForEntity("/verify", signResponse, VerifyResponse.class);

        Assertions.assertNotNull(response.getBody());
        assertThat(response.getBody().result).isEqualTo("success");
        assertThat(response.getBody().message).isEqualTo("Signature is valid");

    }
}
