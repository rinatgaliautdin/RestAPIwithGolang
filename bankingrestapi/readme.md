README

##The demo REST API developed using Go(Golang)
The Api uses Redis as a data cache
The Service uses Basic authentication, for the example the Authorization header must be "Basic YWRtaW46cGFzc3dvcmQ="

##Endpoints:

### balance : retrieves the balance of the given wallet specified by the wallet id
GET /api/v1/wallets/{wallet_id}/balance

### credit : credits money to the given wallet specified by the wallet id
POST / api/v1/wallets/{wallet_id}/credit
where the form field name is named "amount"

### debit : debits money from the given wallet specified by the wallet id
POST / api/v1/wallets/{wallet_id}/debit
where the form field name is named "amount"


