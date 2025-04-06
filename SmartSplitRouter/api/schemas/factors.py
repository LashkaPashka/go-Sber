from pydantic import BaseModel
from typing import List, Optional

class Discounts(BaseModel):
    name: str
    number: int
    
class Tips(BaseModel):
    number: int
    
class SCache(BaseModel):
    discounts: Optional[List[Discounts]] = None
    tips: Optional[List[Tips]] = None