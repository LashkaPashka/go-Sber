import uvicorn
from fastapi import FastAPI
from api.memCache.router import router as router_cManipulation

app = FastAPI()

app.include_router(router_cManipulation)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)