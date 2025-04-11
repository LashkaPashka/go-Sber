import json
import requests
from fastapi import status, HTTPException



def fetchFromService(body: any):
    url = "http://localhost:8050/divide-bill"
    
    response = requests.post(url=url, data=json.dumps(body), headers={"Content-Type": "application/json"})
    
    if response.status_code != status.HTTP_200_OK:
        raise HTTPException(status_code=status.HTTP_409_CONFLICT, detail="не удалось получить API от Go")
    
    return response.json()

