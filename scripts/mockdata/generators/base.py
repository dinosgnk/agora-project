from faker import Faker
from typing import List, Dict, Any


class BaseDataGenerator:
    """Base class for all data generators"""
    
    def __init__(self):
        self.fake = Faker()
    
    def generate(self, num_items: int) -> List[Dict[str, Any]]:
        """Generate specified number of items"""
        raise NotImplementedError("Subclasses must implement generate method")
