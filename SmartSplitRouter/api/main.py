import uvicorn
from fastapi import FastAPI
from api.routers.small import router as small_router

app = FastAPI()
app.include_router(small_router)


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8090)