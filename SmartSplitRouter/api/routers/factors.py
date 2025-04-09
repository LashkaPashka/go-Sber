from fastapi import APIRouter, Request
from api.schemas.factors import SCache
from api.reqToMemcache.save_data import save_data
from api.queues.kafka import Kafka


router = APIRouter(
    prefix="/factors",
    tags=["Учёт факторов"]
)

@router.post("/set")
def factors(request: Request, cache: SCache):
    hash = request.cookies.get("hash")
    res = save_data(cache, hash)
    kafka = Kafka(["localhost:9092"])
    kafka.Publisher("topic-factors", {"hash": hash})
    
    return res