from pydantic import BaseModel

class PositionPrice(BaseModel):
    position: str
    use_clients: int
    num_clients: int
    