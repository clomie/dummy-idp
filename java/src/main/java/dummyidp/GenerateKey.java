package dummyidp;

import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import com.nimbusds.jose.JWSAlgorithm;
import com.nimbusds.jose.jwk.JWK;
import com.nimbusds.jose.jwk.JWKSet;
import com.nimbusds.jose.jwk.KeyUse;
import com.nimbusds.jose.jwk.RSAKey;
import com.nimbusds.jose.jwk.gen.RSAKeyGenerator;

public class GenerateKey {
    public static void main(String... args) throws Exception {
        RSAKey generatedPrivateKey = new RSAKeyGenerator(2048).keyUse(KeyUse.SIGNATURE).algorithm(JWSAlgorithm.RS256)
                .keyIDFromThumbprint(true).generate();

        Path pkPath = Paths.get("../privateKey.json");
        Files.writeString(pkPath, generatedPrivateKey.toString());
        System.out.println("Output: " + pkPath.toRealPath());
        
        JWK privateKey = JWK.parse(Files.readString(pkPath));
        JWKSet jwks = new JWKSet(privateKey).toPublicJWKSet();
        
        Path jwksPath = Paths.get("../dummy-idp/.well-known/jwks.json");
        Files.writeString(jwksPath, jwks.toString());
        System.out.println("Output: " + jwksPath.toRealPath());
    }
}