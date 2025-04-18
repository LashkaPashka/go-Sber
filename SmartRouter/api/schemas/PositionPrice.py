from pydantic import BaseModel

class SPosition(BaseModel):
    position: str
    num_clients: int
    use_clients: int
    
class SAccount(BaseModel):
    num_clients: int
    use_clients: int