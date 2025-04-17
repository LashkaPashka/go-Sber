from fastapi import APIRouter, Request
from api.schemas.factors import SCache
from api.Postman.save_data_in_redis import save_data
from api.queues.kafka import Kafka

router = APIRouter(
    prefix="/factors",
    tags=["Учёт факторов"]
)

@router.post("/set")
def apply_discounts_and_tips(request: Request, cache: SCache):
    hash = request.cookies.get("hash")
    _ = save_data(cache, hash)
    kafka = Kafka(["localhost:9092"])
    kafka.Publisher("topic-factors", {"hash": hash})
    
    return "Факторы применены"