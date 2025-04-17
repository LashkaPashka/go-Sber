import uvicorn
from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from api.routers.ocr import router as router_ocr
from api.routers.images import router as router_images

app = FastAPI()

app.include_router(router_ocr)
app.include_router(router_images)

app.mount("/static", StaticFiles(directory="api/static"), "static")

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)