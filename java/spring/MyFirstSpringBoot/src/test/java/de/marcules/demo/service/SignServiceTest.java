package de.marcules.demo.service;

import de.marcules.demo.model.SignResponse;

import java.io.IOException;

import static de.marcules.demo.Service.Sign.signPayload;
import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

public class SignServiceTest {

        @org.junit.jupiter.api.Test
        void throwsNoException() {

           org.junit.jupiter.api.Assertions.assertDoesNotThrow(() -> signPayload("Test message"));

        }


    @org.junit.jupiter.api.Test
    void responseHasSignatureAndCert() {

        SignResponse signature = signPayload("Test message");

        assertThat(signature.certificate()).isNotBlank();
        assertThat(signature.signedPayload()).isNotBlank();

    }

    @org.junit.jupiter.api.Test
    void responseHasExpectedCert() throws IOException {

        SignResponse signature = signPayload("Test message");

        String expectedCert = new String(java.nio.file.Files.readAllBytes(java.nio.file.Paths.get("src/main/resources/static/sign.crt")));
        assertThat(signature.certificate()).isNotBlank();
        assertThat(signature.certificate().replaceAll("\\r?\\n", "")).isEqualTo(expectedCert.replaceAll("\\r?\\n", ""));

    }
}
