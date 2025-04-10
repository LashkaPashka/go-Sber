import requests
import json
from fastapi import status, HTTPException
from api.schemas.factors import SCache

def save_data(model: SCache, hash: str) -> str:
    url = f"http://localhost:8000/cache/set-data/factors:{hash}"
    
    payload = model.dict()
    resonse = requests.post(url, json=payload)
    
    if resonse.status_code != status.HTTP_200_OK:
        raise HTTPException(status_code=status.HTTP_409_CONFLICT, detail="не удалось сохранить данные в Redis")
    
    return "Данные отправлены в микросервис Redis успешно!"