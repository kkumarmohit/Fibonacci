# Copyright: 2024 Arm Ltd. All Rights Reserved.
FROM python:slim
WORKDIR /app
COPY app.py .
RUN pip install flask requests
RUN adduser --disabled-password --gecos '' fibuser
USER fibuser
EXPOSE 8081
CMD ["python", "app.py"]
