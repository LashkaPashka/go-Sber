from pydantic import BaseModel

class PositionPrice(BaseModel):
    use_clients: int
    num_clients: int
    