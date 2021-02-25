const { readFileSync } = require('fs')
const { join } = require('path')
const { JWK, JWT } = require('jose')

const jwk = JSON.parse(readFileSync(join(__dirname, '../privateKey.json')))
const key = JWK.asKey(jwk)

const sub = process.argv[2]

const idToken = JWT.sign(
  {
    iss: 'http://localhost',
    sub,
    aud: 'dummy_client_id',
    exp: Number.MAX_SAFE_INTEGER,
    iat: Date.now() / 1000,
  },
  key
)

console.log(idToken)
