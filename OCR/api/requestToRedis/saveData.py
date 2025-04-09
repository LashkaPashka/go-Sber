import requests
import json
from api.requestToRedis.schemas import SCache
from fastapi import status, HTTPException

def save_data(model: SCache, hash: str) -> str:
    url = f"http://localhost:8000/cache/set-data/cheque:{hash}"
    
    payload = model.dict()
    resonse = requests.post(url, json=payload)
    
    if resonse.status_code != status.HTTP_200_OK:
        raise HTTPException(status_code=status.HTTP_409_CONFLICT, detail="не удалось сохранить данные в Redis")
    
    return "Данные отправлены в микросервис Redis успешно!"