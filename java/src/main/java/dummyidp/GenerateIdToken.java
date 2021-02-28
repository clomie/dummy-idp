package dummyidp;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.time.Instant;
import java.util.HashMap;
import java.util.Map;

import com.nimbusds.jose.JWSAlgorithm;
import com.nimbusds.jose.JWSHeader;
import com.nimbusds.jose.JWSSigner;
import com.nimbusds.jose.crypto.factories.DefaultJWSSignerFactory;
import com.nimbusds.jose.jwk.JWK;
import com.nimbusds.jwt.JWT;
import com.nimbusds.jwt.JWTClaimsSet;
import com.nimbusds.jwt.SignedJWT;

public class GenerateIdToken {
    public static void main(String[] args) throws Exception {
        JWK privateKey = JWK.parse(Files.readString(Paths.get("../privateKey.json")));

        String sub = args.length >= 1 && args[1].isBlank() ? args[1] : "dummy-user-id";
        long now = Instant.now().getEpochSecond();

        Map<String, Object> claims = new HashMap<>();
        claims.put("iss", "http://localhost");
        claims.put("sub", sub);
        claims.put("aud", "dummy-audience");
        claims.put("exp", Integer.MAX_VALUE);
        claims.put("iat", now);

        JWSHeader header = new JWSHeader.Builder(JWSAlgorithm.RS256).keyID(privateKey.getKeyID()).build();
        JWTClaimsSet claimsSet = JWTClaimsSet.parse(claims);
        SignedJWT jwt = new SignedJWT(header, claimsSet);

        JWSSigner signer = new DefaultJWSSignerFactory().createJWSSigner(privateKey, JWSAlgorithm.RS256);
        jwt.sign(signer);

        System.out.println(jwt.serialize());
    }
}
