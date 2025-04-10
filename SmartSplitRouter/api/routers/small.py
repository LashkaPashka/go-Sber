from fastapi import APIRouter, Response, Request
from datetime import timedelta
from api.schemas.PositionPrice import PositionPrice
from api.queues.rabbit import RabbitMQ


router = APIRouter(
    prefix="/small-router",
    tags=["Умный роутер"]
)

@router.get("/{hash}")
def routing(response: Response, hash: str):
    response.set_cookie("hash", hash, httponly=True, expires=timedelta(minutes=200))
    
    print(hash)

@router.post("/get-price-position")
def get_price_position(request: Request, pos: PositionPrice):
    hash = request.cookies.get("hash")
    
    rabbit = RabbitMQ()
    
    rabbit.Publish("topic-divide", {"hash": hash, "numClients": str(pos.num_clients), "useClients": str(pos.use_clients)})
    #kafka.Subscriber(["topic-factors"])