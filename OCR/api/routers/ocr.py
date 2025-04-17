from fastapi import APIRouter
from api.CreateLink import create_link
from api.PostmanRedis import saveData
from api.PostmanRedis.schemas import SCache

router = APIRouter(
    prefix="/ocr-work",
    tags=["OCR"]
)

@router.post("/create")
def work(modelCheque: SCache):
    ## Данные получаем из чека
    
    ## Формируем модель json для cheque
    
    ## Создаём ссылку 
    url = create_link.create_link()
    hash = url.split("http://localhost:8090/small-router/")
    
    ##Сохраняем в redis
    saveData.save_data(modelCheque, hash[1])
    
    return url