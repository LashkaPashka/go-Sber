import shutil
from fastapi import APIRouter, UploadFile
from api.routers.ocr_code import ocr_run
from api.CreateLink.create_link import create_link
from api.PostmanRedis import saveData

from random import randint

router = APIRouter(
    prefix="/images",
    tags=["Загрузка изображения чека"]
)

@router.post("/load-image")
def add_image(file: UploadFile):
    nameID = randint(100000, 999999)
    im_path = f"api/static/images/{nameID}.webp"
    with open(im_path, "wb+") as file_object:
        shutil.copyfileobj(file.file, file_object)
    
    
    ocr_result = ocr_run(im_path)
    
    UrlAndHash = create_link()
    
    saveData.save_data(ocr_result, UrlAndHash[1])
    
    return {
        "status": "success",
        "url": UrlAndHash[0],
        "hash": UrlAndHash[1],
        "message": "Изображение успешно загружено и обработано.",
        "cheque":  ocr_result
    }