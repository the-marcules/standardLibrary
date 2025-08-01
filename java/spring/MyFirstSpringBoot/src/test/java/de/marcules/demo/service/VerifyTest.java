package de.marcules.demo.service;

import de.marcules.demo.model.SignResponse;
import org.junit.jupiter.api.Test;

import static de.marcules.demo.Service.Sign.signPayload;
import static de.marcules.demo.Service.Verify.verifySignature;
import static org.junit.jupiter.api.Assertions.*;

class VerifyTest {

    @Test
    void testVerifyReturnsTrueOnValidSignature() throws Exception {
        String data = "testdata";
        SignResponse validSignature = signPayload(data); // Annahme: sign erzeugt eine gültige Signatur
        assertTrue(verifySignature(validSignature));
    }


}