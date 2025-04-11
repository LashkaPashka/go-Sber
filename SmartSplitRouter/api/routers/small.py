from fastapi import APIRouter, Response, Request
from datetime import timedelta
from api.schemas.PositionPrice import SPosition, SAccount
from api.requestsToSerivce.requestToGo import fetchFromService

router = APIRouter(
    prefix="/small-router",
    tags=["Умный роутер"]
)

@router.get("/{hash}")
def routing(response: Response, hash: str):
    response.set_cookie("hash", hash, httponly=True, expires=timedelta(minutes=100))
    
    print(hash)

@router.post("/get-position")
def receive_split_data(request: Request, pos: SPosition):
    hash = request.cookies.get("hash")
    data = fetchFromService(
        "split-position",
        {"hash": hash, "position": pos.position, "numClients": str(pos.num_clients), "useClients": str(pos.use_clients)}
    )

    return data

@router.post("/get-account")
def receive_split_data(request: Request, pos: SAccount):
    hash = request.cookies.get("hash")
    data = fetchFromService("split-account", {"hash": hash, "numClients": str(pos.num_clients), "useClients": str(pos.use_clients)})

    return data