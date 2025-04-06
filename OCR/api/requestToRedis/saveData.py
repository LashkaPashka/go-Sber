import requests
from api.requestToRedis.schemas import SCache
from fastapi import status, HTTPException


def save_data(model: SCache, hash: str) -> str:
    url = f"http://localhost:8000/set-data/{hash}"
    
    mJson = model.model_dump_json()
    x = requests.post(url, json=mJson)
    
    if x.status_code != status.HTTP_200_OK:
        raise HTTPException(status_code=status.HTTP_409_CONFLICT, detail="не удалось сохранить данные в Redis")
    
    return "Данные отправлены в микросервис Redis успешно!"