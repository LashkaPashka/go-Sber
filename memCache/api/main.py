from fastapi import FastAPI
from api.memCache.router import router as router_cManipulation

app = FastAPI()

app.include_router(router_cManipulation)