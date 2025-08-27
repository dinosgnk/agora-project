FROM python:3.11-slim

ENV CATALOG_SERVICE=host.docker.internal:8081
ENV CART_SERVICE=host.docker.internal:8082

WORKDIR /app
RUN pip install locust requests
COPY . .
EXPOSE 8089
CMD ["python", "./main.py"]
