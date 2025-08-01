https://github.com/user-attachments/assets/01694ee6-2bb5-40bb-aab4-9c03d5f24b44

![alt text](https://github.com/fabiose81/watermark-image/blob/master/watermark-image.jpg?raw=true)

### For Golang and AWS/S3
    In golang folder create a file .env and insert:

    TRUSTED_PROXIES="127.0.0.1"
    ALLOW_ORIGINS = "http://localhost:3000"
    PORT = ":9000"
    BUCKET_S3 = {your bucket name}
    S3_REGION = {your region}
    TMP_FOLDER = "./tmp/"
    AWS_PROFILE = "golang" //Profile created in .aws/credentias to set aws_access_key_id and aws_secret_access_key 
    Ex: [golang]
        aws_access_key_id = {your key id}
        aws_secret_access_key = {your access key}

### For React
    In react folder create a file .env and insert:

    REACT_APP_API_URL=http://localhost:9000

### Lambda code for AWS Serveless(Python)
    
    import boto3
    from PIL import Image, ImageDraw, ImageFont
    import os
    import io
    import base64

    s3 = boto3.client('s3')

    source_bucket = "imagenowatermark"
    destination_bucket = "imagewatermarkprocessed"
    watermark_text = "© Fabiose-Watermark"

    def get_image_from_bucket(key):
        return s3.get_object(Bucket=source_bucket, Key=key)

    def create_watermark(image):
        watermark = Image.new("RGBA", image.size)
        return watermark
    
    def add_watermark(image, watermark, font, draw):
        position_getpixel= (image.height/2 , image.height/2)

        r, g, b  = image.getpixel(position_getpixel)
        brightness = (0.299 * r + 0.587 * g + 0.114 * b)
        color = (255, 255, 255, 128)
        if brightness >= 128:
           color = (0, 0, 0, 128)
    
        position_drawtext = (10, image.height/2)
        draw.text(position_drawtext, watermark_text, fill=color, font=font)

        return Image.alpha_composite(image.convert("RGBA"), watermark)

    def save_image_to_bucket(watermarked_image, key):
        buffer = io.BytesIO()
        watermarked_image.convert("RGB").save(buffer, "JPEG")
        buffer.seek(0)

        processed_key = f"processed-{key}"
        s3.upload_fileobj(buffer, destination_bucket, processed_key)

    def delete_image_from_bucket(key):
        s3.delete_object(Bucket=source_bucket, Key=key)

    def lambda_handler(event, context):
        key = event['Records'][0]['s3']['object']['key']
    
        response = get_image_from_bucket(key)
        image_data = response['Body'].read()
        image = Image.open(io.BytesIO(image_data))
    
        watermark = create_watermark(image)
        draw = ImageDraw.Draw(watermark)

        w, h = image.size
        font = ImageFont.load_default(size= round(w/13))
        
        watermarked_image = add_watermark(image, watermark, font, draw)
    
        save_image_to_bucket(watermarked_image, key)

        delete_image_from_bucket(key)

### For lambda layer

    Create a build directory:

        mkdir pillow-layer
        cd pillow-layer
        mkdir python

    Use Docker to build Pillow in Amazon Linux-compatible environment

        Run this in your terminal:

            docker run -v "$PWD"/python:/mnt/output -it public.ecr.aws/lambda/python:3.9 bash

        Inside the container:

            pip install --upgrade pip
            pip install Pillow -t /mnt/output
            exit

    Zip the python folder:

        zip -r pillow_layer.zip python

    
