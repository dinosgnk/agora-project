import json
from typing import List, Dict, Any
from .product import ProductGenerator


class MockDataManager:
    """Main manager class for generating and saving mock data"""
    
    def __init__(self):
        self.generators = {
            'products': ProductGenerator()
        }
    
    def generate_data(self, data_type: str, num_items: int) -> List[Dict[str, Any]]:
        """Generate data of specified type"""
        if data_type not in self.generators:
            raise ValueError(f"Unknown data type: {data_type}")
        
        return self.generators[data_type].generate(num_items)
    
    def save_to_json_file(self, data: List[Dict[str, Any]], filename: str) -> None:
        """Save data to JSON file"""
        with open(filename, 'w') as f:
            json.dump(data, f, indent=2)
        print(f"Generated {len(data)} items and saved to '{filename}'")
    
    def save_to_sql_file(self, data_type: str, data: List[Dict[str, Any]], filename: str) -> None:
        """Save data to SQL file"""
        if data_type not in self.generators:
            raise ValueError(f"Unknown data type: {data_type}")
        
        generator = self.generators[data_type]
        if hasattr(generator, 'generate_sql'):
            sql_content = generator.generate_sql(data)
            with open(filename, 'w') as f:
                f.write(sql_content)
            print(f"Generated SQL INSERT statements for {len(data)} items and saved to '{filename}'")
        else:
            raise ValueError(f"Generator for {data_type} does not support SQL generation")
