const { writeFileSync } = require('fs')
const { join } = require('path')
const { JWK } = require('jose')

const jwk = JWK.generateSync('RSA', 2048, { alg: 'RS256', use: 'sig' })
const privateKey = jwk.toJWK(true)
const jwks = { keys: [jwk.toJWK()] }

writeFileSync(
  join(__dirname, '../privateKey.json'),
  JSON.stringify(privateKey, null, 2)
)
writeFileSync(
  join(__dirname, '../dummy-idp/.well-known/jwks.json'),
  JSON.stringify(jwks, null, 2)
)
