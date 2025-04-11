import json
import requests
from typing import Any
from fastapi import status, HTTPException

def fetchFromService(path: str, body: Any):
    url = f"http://localhost:8085/{path}"
    
    response = requests.post(url=url, json=body, headers={"Content-Type": "application/json"})
    
    if response.status_code != status.HTTP_200_OK:
        raise HTTPException(status_code=status.HTTP_409_CONFLICT, detail="не удалось получить API от Go")
    

    return response.json()

