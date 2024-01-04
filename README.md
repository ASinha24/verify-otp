# verify-otp

## A complete golang project using gin framework for API to verify the otp through TWILIO sms service, can be customized to veriy email or whatsapp

### setup

1. create a repository in github and clone it
   
2. create Twilio account in console.twilio.com

3. get the twilio id and authtoken

4. go the explore services and find verify service create a service with sms option

5. get the twilio account id and authoken along with service id.

6. add it in your. env file

```bash
TWILIO_ACCOUNT_SID=<ACCOUNT SID>
TWILIO_AUTHTOKEN=<AUTH TOKEN>
TWILIO_SERVICE_ID=<SERVICE ID>
```

7. install dependencied using `go mod tidy`

8. run the server using `go run cmd/main.go`

### curl command to test the service

---
### Send OTP

```bash
curl -H "Content-Type: application/json" -X POST -d '{"phoneNumber": "<number with country code>"}' http://localhost:8000/otp
```

Response

```json
{
  "status": 202,
  "message": "success",
  "data": "OTP sent successfully"
}
```
---
### Verify OTP

POST /verifyOTP

Request Body

```bash
curl -H "Content-Type: application/json" -X POST -d '{"user": {"phoneNumber": <phone-number-with-country-code>}, "code":"795279"}' http://localhost:8000/verifyOTP
```

Response

```json
{
  "status": 202,
  "message": "success",
  "data": "OTP verified successfully"
}
```
---


