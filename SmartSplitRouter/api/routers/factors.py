from fastapi import APIRouter
from api.schemas.factors import SCache
from api.reqToMemcache.save_data import save_data

router = APIRouter(
    prefix="/factors",
    tags=["Учёт факторов"]
)

@router.post("/set")
def factors(cache: SCache):
    res = save_data(cache, "")
    
    return res