# Browser Remote Event Viewer in GO

To start client:
- `cd client && harp server -p 8081`

To generate the required certificate:
- `openssl genrsa -out cert.key 2048`
- `openssl ecparam -genkey -name secp384r1 -out cert.key`
- `openssl req -new -x509 -sha256 -key cert.key -out cert.pem -days 3650`
