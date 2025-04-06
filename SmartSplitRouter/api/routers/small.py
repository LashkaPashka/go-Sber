from fastapi import APIRouter

router = APIRouter(
    prefix="/small-router",
    tags=["Умный роутер"]
)

@router.get("/{hash}")
def routing(hash: str):
    print(hash)