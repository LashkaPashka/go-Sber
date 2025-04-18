import requests
from typing import Any
from fastapi import status, HTTPException

def save_data(model: dict[str, Any], hash: str) -> str:
    url = f"http://memcache:8000/cache/set-data/cheque:{hash}"
    
    resonse = requests.post(url, json=model)
    
    if resonse.status_code != status.HTTP_200_OK:
        raise HTTPException(status_code=status.HTTP_409_CONFLICT, detail="не удалось сохранить данные в Redis")
    
    return "Данные отправлены в микросервис Redis успешно!"