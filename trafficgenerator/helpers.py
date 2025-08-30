from typing import List, Dict

import requests
import os

CATALOG_SERVICE = f"http://{os.getenv('CATALOG_SERVICE', 'localhost')}:8081"  

def get_all_product_categories() -> List[str]:
    """Fetch all unique product categories from the catalog service"""
    try:
        response = requests.get(f"{CATALOG_SERVICE}/products")
        if response.status_code == 200:
            products = response.json()
            categories = list(set(product.get("category", "") for product in products))
            print(f"Loaded {len(categories)} unique categories from catalog")
            return categories
        else:
            print(f"Failed to fetch products: {response.status_code}")
            return []
    except Exception as e:
        print(f"Error fetching products: {e}")
        return []

def get_all_products() -> List[Dict]:
    """Fetch products from catalog service"""
    try:
        response = requests.get(f"{CATALOG_SERVICE}/products")
        if response.status_code == 200:
            products = response.json()
            print(f"Loaded {len(products)} products from catalog")
            return products
        else:
            print(f"Failed to fetch products: {response.status_code}")
            return []
    except Exception as e:
        print(f"Error fetching products: {e}")
        return []
    

print(get_all_product_categories())