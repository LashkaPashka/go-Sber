import uvicorn
from fastapi import FastAPI
from api.routers.ocr import router as router_ocr

app = FastAPI()

app.include_router(router_ocr)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)