import json
import redis
from fastapi import APIRouter
from api.connDb import PingDb
from api.memCache.schemas import SCache
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

@router.post("/set-data")
def set_data(cache: SCache):
    host = "localhost"; port = 6379; password = "mypassword"
    PingDb(host=host, port=port, password=password)

    time = timedelta(minutes=2)
    r = redis.Redis(host=host, port=port, password=password, decode_responses=True)
    r.set(name="key", value=cache.model_dump_json(), ex=time)
    
    return "Данные отправились в кэш"