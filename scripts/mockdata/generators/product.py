import random
from typing import List, Dict, Any
from .base import BaseDataGenerator

class ProductGenerator(BaseDataGenerator):
    """Generator for product mock data"""
    
    def __init__(self):
        super().__init__()
        self.categories = [
            "Movies", "Home", "Music", "Grocery", "Computers", "Baby", "Sports", "Jewelry", "Toys", "Health",
            "Clothing", "Games", "Kids", "Shoes", "Beauty", "Automotive", "Outdoors", "Electronics", "Industrial",
            "Garden", "Books", "Tools"
        ]
        self.used_upcs = set()
    
    def _generate_unique_upc(self) -> str:
        """Generate a unique UPC code"""
        while True:
            upc = self.fake.ean13()[:12]  # Use first 12 digits of EAN13
            if upc not in self.used_upcs:
                self.used_upcs.add(upc)
                return upc
    
    def _generate_price(self) -> float:
        """Generate price with realistic distribution"""
        price_tier = random.choices(
            [1, 2, 3, 4],  # price tiers
            weights=[50, 30, 15, 5],  # probability weights
            k=1
        )[0]
        
        if price_tier == 1:  
            return round(random.uniform(0.99, 49.99), 2)
        elif price_tier == 2:
            return round(random.uniform(50.00, 199.99), 2)
        elif price_tier == 3:
            return round(random.uniform(200.00, 499.99), 2)
        else:
            return round(random.uniform(500.00, 1999.99), 2)
    
    def generate(self, num_products: int) -> List[Dict[str, Any]]:
        """Generate specified number of products"""
        products = []
        
        for i in range(num_products):
            # Generate description without newlines
            description = self.fake.text(max_nb_chars=200).replace('\n', ' ').replace('\r', ' ')
            
            product = {
                "name": self.fake.catch_phrase(),
                "category": random.choice(self.categories),
                "price": self._generate_price(),
                "description": description,
                "product_code": self._generate_unique_upc(),
                "created_at": self.fake.date_time_this_year().isoformat()
            }
            products.append(product)
        
        return products
    
    def generate_sql(self, products: List[Dict[str, Any]]) -> str:
        """Generate SQL INSERT statements for products"""
        sql_lines = []
        
        for product in products:
            # Escape single quotes in text fields
            name = product['name'].replace("'", "''")
            description = product['description'].replace("'", "''")
            category = product['category'].replace("'", "''")
            
            # Match the database schema: Name, Category, Description, Price
            sql = (
                f"INSERT INTO products.t_product (Name, Category, Description, Price) "
                f"VALUES ('{name}', '{category}', '{description}', {product['price']});"
            )
            sql_lines.append(sql)
        
        return '\n'.join(sql_lines)
