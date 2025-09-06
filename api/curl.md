1. Health Check
curl -X GET http://localhost:8080/


2. Fetch Templates
curl -X GET http://localhost:8080/templates/330090043524228


3. Create Template
curl -X POST http://localhost:8080/templates \
  -H "Content-Type: application/json" \
  -d '{
    "waba_id": "318794741325676",
    "template_name": "example_template1",
    "language": "en_US",
    "category": "MARKETING",
    "header_type": "headerImage",
    "header_content": "4::YXBwbGljYXRpb24vb2N0ZXQtc3RyZWFt:ARbp3NxfKCTnotO_6ZojyIsDak1YDrcmaku_jJl4cJvKk-nrt6lKBgNDKTf-G9kElPUR-74ab9bwgGOqSunXE_6HkQz1ltZL3_23pnKYcg0Zxg:e:1720928506:1002275394751227:61560603673003:ARbQbHWnFfKRtUWd7_g",
    "body_text": "This is the body of the template.",
    "footer_text": "This is the footer.",
    "call_button_text": "Call Us",
    "phone_number": "+917905968734",
    "url_button_text": "Visit Website",
    "website_url": "https://example.com"
}'


4. Upload Media (generate media ID)
curl -X POST http://localhost:8080/upload \
  -F "phone_number_id=1234567890" \
  -F "media_file=@path/to/image.jpg"

5. Upload Header Media (generate header_handle)
curl -X POST http://localhost:8080/upload/header \
  -F "header_file=@path/to/image.jpg"

6. Send Message
curl -X POST http://localhost:8080/send \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number_id": "1234567890",
    "template_name": "example_template1",
    "language": "en_US",
    "media_type": "IMAGE",
    "media_id": "123456_media_id",
    "contact_list": [
      "919000000000",
      "919111111111"
    ]
}'

1. Exchange auth_code for Access Token

curl -X GET "https://graph.facebook.com/v20.0/oauth/access_token \
  ?client_id=<APP_ID> \
  &client_secret=<APP_SECRET> \
  &code=<AUTH_CODE> \
  &redirect_uri=<REDIRECT_URI>"

{
  "access_token": "EAABsbCS1...",
  "token_type": "bearer",
  "expires_in": 5183944
}

2. Get business_id and waba_id from Access Token

curl -X GET "https://graph.facebook.com/v20.0/me?fields=id,name,whatsapp_business_account \
  &access_token=<ACCESS_TOKEN>"

{
  "id": "123456789",
  "name": "Test Business",
  "whatsapp_business_account": {
    "id": "987654321"
  }
}
3. Get Phone Number ID

curl -X GET "https://graph.facebook.com/v20.0/<WABA_ID>/phone_numbers?access_token=<ACCESS_TOKEN>"