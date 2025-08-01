https://github.com/user-attachments/assets/01694ee6-2bb5-40bb-aab4-9c03d5f24b44

![alt text](https://github.com/fabiose81/watermark-image/blob/master/watermark-image.jpg?raw=true)

### For Golang and AWS/S3
    In golang folder create a file .env and insert:

    TRUSTED_PROXIES="127.0.0.1"
    ALLOW_ORIGINS = "http://localhost:3000"
    PORT = ":9000"
    BUCKET_S3 = {your bucket name}
    S3_REGION = {your region}
    KEY_PROFILE = "golang"
    TMP_FOLDER = "./tmp/"

### For React
    In react folder create a file .env and insert:

    REACT_APP_API_URL=http://localhost:9000
