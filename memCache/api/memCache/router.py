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
    host = "redis"; port = 6379; password = "mypassword"
    PingDb(host=host, port=port, password=password)
    
    r = redis.Redis(host=host, port=port, password=password, decode_responses=True)
    resp = r.get(key)
    
    return resp

@router.post("/set-data/{key}")
def set_data(key: str, cache: Any = Body(...)):
    host = "redis"; port = 6379; password = "mypassword"
    PingDb(host=host, port=port, password=password)

    time = timedelta(minutes=60)
    r = redis.Redis(host=host, port=port, password=password, decode_responses=True)
    r.set(name=f"{key}", value=json.dumps(cache), ex=time)
    
    return "Данные отправились в кэш"

@router.delete("/delete-data/{key}")
def delete_data(key: str):
    host = "redis"; port = 6379; password = "mypassword"
    PingDb(host=host, port=port, password=password)
    r = redis.Redis(host=host, port=port, password=password, decode_responses=True)
    r.delete(key)
    return "Данные удалены из кэша"

@router.get("/get-keys")
def get_keys():
    host = "redis"; port = 6379; password = "mypassword"
    PingDb(host=host, port=port, password=password)
    r = redis.Redis(host=host, port=port, password=password, decode_responses=True)
    keys = r.keys()
    return {"keys": keys}

@router.delete("/delete-all")
def delete_all():
    host = "redis"; port = 6379; password = "mypassword"
    PingDb(host=host, port=port, password=password)
    r = redis.Redis(host=host, port=port, password=password, decode_responses=True)
    r.flushall()
    return "Все данные удалены из кэша"
