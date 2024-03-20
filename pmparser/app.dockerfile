FROM python:3.12-alpine
# System Dependencies
RUN apk add --no-cache chromium
# App Dependencies
COPY ./requirements.txt ./application/requirements.txt
WORKDIR /application
RUN pip install -r requirements.txt
# App
COPY ./app.py /application/app.py
COPY ./app /application/app
CMD ["python", "app.py"]