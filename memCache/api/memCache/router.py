import json
import redis
from typing import Any
from fastapi import APIRouter, Body
from api.connDb import PingDb
from datetime import timedelta

router = APIRouter(
    prefix="/cache",
    tags=["Кэширование"]
)

@router.get("/get-data/{key}")
def get_data(key: str):
    host = "localhost"; port = 6379; password = "mypassword"
    PingDb(host=host, port=port, password=password)
    
    r = redis.Redis(host=host, port=port, password=password, decode_responses=True)
    resp = r.get(key)
    
    return resp

@router.post("/set-data/{key}")
def set_data(key: str, cache: Any = Body(...)):
    host = "localhost"; port = 6379; password = "mypassword"
    PingDb(host=host, port=port, password=password)

    time = timedelta(minutes=120)
    r = redis.Redis(host=host, port=port, password=password, decode_responses=True)
    r.set(name=f"{key}", value=json.dumps(cache), ex=time)
    
    return "Данные отправились в кэш"