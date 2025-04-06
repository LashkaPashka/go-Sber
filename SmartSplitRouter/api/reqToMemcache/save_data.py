import requests
import json
from fastapi import status, HTTPException

def save_data(model: any, hash: str) -> str:
    url = f"http://localhost:8000/cache/set-data/{hash}"
    
    payload = json.dumps(model)
    resonse = requests.post(url, json=payload)
    
    if resonse.status_code != status.HTTP_200_OK:
        raise HTTPException(status_code=status.HTTP_409_CONFLICT, detail="не удалось сохранить данные в Redis")
    
    return "Данные отправлены в микросервис Redis успешно!"