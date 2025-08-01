package de.marcules.demo.Service;

import de.marcules.demo.model.SignResponse;

public class Sign {
    public static SignResponse signPayload(String payload) {
        try {
            // Lade privaten Schlüssel aus Datei
            java.nio.file.Path keyPath = java.nio.file.Paths.get("src/main/resources/static/sign.key");
            String keyPem = new String(java.nio.file.Files.readAllBytes(keyPath));
            String privateKeyPEM = keyPem
                    .replace("-----BEGIN PRIVATE KEY-----", "")
                    .replace("-----END PRIVATE KEY-----", "")
                    .replaceAll("\\s", "");
            byte[] encoded = java.util.Base64.getDecoder().decode(privateKeyPEM);

            java.security.spec.PKCS8EncodedKeySpec keySpec = new java.security.spec.PKCS8EncodedKeySpec(encoded);
            java.security.PrivateKey privateKey = java.security.KeyFactory.getInstance("RSA").generatePrivate(keySpec);

            java.security.Signature signature = java.security.Signature.getInstance("SHA256withRSA");
            signature.initSign(privateKey);
            signature.update(payload.getBytes());
            byte[] signed = signature.sign();
            return new SignResponse(java.util.Base64.getEncoder().encodeToString(signed), getCert(), payload);
        } catch (Exception e) {
            throw new RuntimeException("Fehler beim Signieren des Payloads", e);
        }
    }

    private static String getCert() {
        try {
            java.nio.file.Path certPath = java.nio.file.Paths.get("src/main/resources/static/sign.crt");
            return new String(java.nio.file.Files.readAllBytes(certPath));
        } catch (Exception e) {
            throw new RuntimeException("Fehler beim Lesen des Zertifikats", e);
        }
    }

}

