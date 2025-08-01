package de.marcules.demo.Service;

import de.marcules.demo.model.SignResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.ByteArrayInputStream;
import java.nio.charset.StandardCharsets;
import java.security.Signature;
import java.security.cert.CertificateFactory;
import java.security.cert.X509Certificate;
import java.util.Base64;


public class Verify {
    private static final Logger logger = LoggerFactory.getLogger(Verify.class);
    private static final String algorithm = "SHA256withRSA";

    public static boolean verifySignature(SignResponse signature) throws Exception{

        try {
            logger.info("Verifying signedPayload with algorithm: {}", algorithm);
            logger.info("Verifying signedPayload {}", signature.signedPayload());
            X509Certificate certificate = (X509Certificate) CertificateFactory.getInstance("X.509")
                    .generateCertificate(new ByteArrayInputStream(signature.certificate().getBytes(StandardCharsets.UTF_8)));
            Signature sig = Signature.getInstance(algorithm);
            sig.initVerify(certificate.getPublicKey());
            sig.update(signature.payload().getBytes(StandardCharsets.UTF_8));
            return sig.verify(Base64.getDecoder().decode(signature.signedPayload()));
        } catch (Exception e) {
            System.out.println("Error during signedPayload verification: " + e.getMessage());
            return false;
        }
    }

}
