# User Contact to QR Converter

This is a REST web service written in order to learn Go. 

## Description

This is a REST web service written in Go that converts any user's contact information into a QR code so that when scanned from a mobile device, the user can be added as a contact with the given details. 

The REST web service accepts a user's contact information in JSON format.

```
{
    "Name":"Random User",
    "Mobile": "+8112345678",
    "Email": "randomuser@gmail.com",
    "Address": "Tokyo,Japan"
}
```
This then generates a QR code that will be saved locally to the server (/images) running this web service.

## Dependency
* Install Go on the machine that will run this REST web service: https://go.dev/dl/ 

## Running the REST web service

To run the web service, type the following on the terminal/cmd:
```
cd contact-qr-converter
go run main.go
```
## Calling the REST web service

Use Postman to make a request to http://localhost:8080/generateQR and pass in the contact details in JSON format.

### Buit With
* Go (https://go.dev)

### Acknoledgements
* Gin Web Framework (https://gin-gonic.com/docs/)
* go-qrcode Package (https://github.com/skip2/go-qrcode)
* Tutorial: Developing a RESTful API with Go and Gin (https://go.dev/doc/tutorial/web-service-gin)
