FROM python:3.13.3-alpine

WORKDIR /memCache

ENV PYTHONUNBUFFERED=1

RUN pip install --upgrade pip

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 8000

CMD ["uvicorn", "api.main:app", "--reload", "--host", "0.0.0.0", "--port", "8000"]