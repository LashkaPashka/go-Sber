from fastapi import APIRouter
from api.createLink import createLink
from api.requestToRedis import saveData
from api.requestToRedis.schemas import SCache, Product
from api import data

router = APIRouter(
    prefix="/ocr-work",
    tags=["OCR"]
)

@router.get("/")
def work():
    ## Данные получаем из чека
    
    ## Формируем модель json для products
    modelProducts = SCache()
    modelProducts.numberClients = 4
    modelProducts.numberClients = 4070
    modelProducts.products = data.dataProducts
    
    ## Создаём ссылку 
    url = createLink.create_link()
    hash = url.split("http://localhost:8000/")
    
    ## Сохраняем в redis
    saveData.save_data(modelProducts, hash[0])
    
    print(modelProducts, hash)
    return "Что-то получилось"