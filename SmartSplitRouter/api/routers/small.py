from fastapi import APIRouter, Response
from datetime import timedelta

router = APIRouter(
    prefix="/small-router",
    tags=["Умный роутер"]
)

@router.get("/{hash}")
def routing(response: Response, hash: str):
    response.set_cookie("hash", hash, httponly=True, expires=timedelta(minutes=200))
    
    print(hash)