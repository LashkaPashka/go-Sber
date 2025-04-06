from fastapi import APIRouter
from api.createLink import createLink
from api.requestToRedis import saveData
from api.requestToRedis.schemas import SCache
from api import data

router = APIRouter(
    prefix="/ocr-work",
    tags=["OCR"]
)

@router.post("/create")
def work(modelCheque: SCache):
    ## Данные получаем из чека
    
    ## Формируем модель json для cheque
    
    ## Создаём ссылку 
    url = createLink.create_link()
    hash = url.split("http://localhost:8000/")
    
    ##Сохраняем в redis
    saveData.save_data(modelCheque, hash[1])
    
    print(modelCheque.products, hash[1])
    return url