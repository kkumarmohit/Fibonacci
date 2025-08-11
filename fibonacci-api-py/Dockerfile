# Copyright: 2024 Arm Ltd. All Rights Reserved.
FROM python:slim
WORKDIR /app
COPY app.py .
RUN pip install flask
RUN adduser --disabled-password --gecos '' fibuser
USER fibuser
EXPOSE 8080
CMD ["python", "app.py"]
