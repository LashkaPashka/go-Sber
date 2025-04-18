import uvicorn
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from fastapi.staticfiles import StaticFiles
from api.routers.main_page import router as router_images

app = FastAPI()

app.include_router(router_images)

app.mount("/static", StaticFiles(directory="api/static"), "static")

origins = [
    "http://localhost:2173",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["GET", "POST", "OPTIONS", "DELETE", "PATCH", "PUT"],
    allow_headers=["Content-Type", "Set-Cookie", "Access-Control-Allow-Headers", 
                   "Access-Control-Allow-Origin", "Authorization"],
)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)