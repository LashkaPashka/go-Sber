import shutil
from fastapi import APIRouter, UploadFile

router = APIRouter(
    prefix="/images",
    tags=["Загрузка изображения чека"]
)

@router.post("/load-image")
def add_image(name: int, file: UploadFile):
    im_path = f"api/static/images/{name}.webp"
    with open(im_path, "wb+") as file_object:
        shutil.copyfileobj(file.file, file_object)