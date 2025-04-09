import uvicorn
from fastapi import FastAPI
from api.routers.small import router as small_router
from api.routers.factors import router as factors_router

app = FastAPI()
app.include_router(small_router)
app.include_router(factors_router)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8090)