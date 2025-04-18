import requests
import json
from fastapi import status, HTTPException

def get_data(hash: str):
    url = f"http://memcache:8000/cache/get-data/cheque:{hash}"
    
    response = requests.get(url)
    
    if response.status_code != status.HTTP_200_OK:
        raise HTTPException(status_code=status.HTTP_409_CONFLICT, detail="не удалось получить данные из Redis")
    
    return json.loads(response.json())
