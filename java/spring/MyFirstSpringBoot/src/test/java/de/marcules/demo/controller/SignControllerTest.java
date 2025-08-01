package de.marcules.demo.controller;

import de.marcules.demo.model.SignRequest;
import de.marcules.demo.model.SignResponse;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import java.util.Map;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class SignControllerTest {
    @Autowired
    private TestRestTemplate restTemplate;

    @Test
    public void TestSignEndpointShouldReturn2xx() throws Exception {

        SignRequest request = new SignRequest("When the shit hits the fan");

        ResponseEntity<SignResponse> response = this.restTemplate.postForEntity("/sign", request, SignResponse.class);

        assertThat(response.getBody().payload()).isEqualTo(request.message());
        assertThat(response.getStatusCode().is2xxSuccessful()).isTrue();

    }

    @Test
    public void TestSignEndpointWithNoBodyShouldReturn400() throws Exception {

        ResponseEntity<SignResponse> response = this.restTemplate.postForEntity("/sign", null, SignResponse.class);

        assertThat(response.getStatusCode()).isEqualTo( HttpStatus.BAD_REQUEST );
        Assertions.assertNotNull(response.getBody());
        assertThat(response.getBody().signedPayload()).isNull();
        assertThat(response.getBody().payload()).isNull();
        assertThat(response.getBody().certificate()).isNull();

    }

    @Test
    public void TestSignEndpointWithMalformedBodyShouldReturn400() throws Exception {
        Map<String, String> shittyRequest = Map.of("this", "is all messed up");

        ResponseEntity<SignResponse> response = this.restTemplate.postForEntity("/sign", shittyRequest, SignResponse.class);

        assertThat(response.getStatusCode()).isEqualTo( HttpStatus.BAD_REQUEST );
        Assertions.assertNotNull(response.getBody());
        assertThat(response.getBody().signedPayload()).isNull();
        assertThat(response.getBody().payload()).isNull();
        assertThat(response.getBody().certificate()).isNull();

    }
}
