from locust import HttpUser, between
from typing import Optional, Dict

import os
import random
import helpers

class BaseAgoraUser(HttpUser):
    abstract = True
    
    # Dummy assignment to avoid errors, variable not used
    host = "not_used"

    _catalog_host = f"http://{os.getenv('CATALOG_SERVICE', 'localhost')}:8081"
    _cart_host = f"http://{os.getenv('CART_SERVICE', 'localhost')}:8082"

    _products = helpers.get_all_products()
    _categories = helpers.get_all_product_categories()

    _user_counter = 0
    
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        BaseAgoraUser._user_counter += 1
        self.user_id = str(BaseAgoraUser._user_counter)
        print(f"Created user {self.user_id}")

    def _select_random_product(self) -> Optional[Dict]:
        """Get a random product from the fetched products"""
        random_product = random.choice(self._products) 
        print(f"User {self.user_id} selected product: {random_product}")
        return random_product

    def _pick_random_category(self) -> str:
        """Pick a random category using normal distribution to simulate behavior more realistically"""
        length = len(self._categories)
        mean = (length - 1) / 2
        stddev = length / (length / 2)
        index = int(round(random.normalvariate(mean, stddev)))
        clamped_index = max(0, min(length - 1, index))
        return self._categories[clamped_index]

    # Catalog 
    def get_products_by_category(self) -> None: 
        """Get products by category"""
        category = self._pick_random_category()
        self.client.get(f"{self._catalog_host}/products/category/{category}", name="Get products by category")

    def get_product_details(self) -> None:
        """Get details for a specific product"""
        product = self._select_random_product()
        if product and product.get("id"):
            self.client.get(f"{self._catalog_host}/products/{product['id']}", name="Get product details")

    # Cart
    def add_item_to_cart(self):
        """Add a random product to cart"""
        product = self._select_random_product()
        
        # TODO: Randomize quantity with a more realistic distribution
        quantity = random.randint(1, 3) 

        payload = {
            "user_id": self.user_id,
            "item": {
                "product_id": str(product.get("id")),
                "name": product.get("name"),
                "quantity": quantity,
                "price": float(product.get("price"))
            }
        }

        print(f"User {self.user_id} adding to cart: {payload}")

        response = self.client.post(
            f"{self._cart_host}/cart/item/add",
            json=payload,
            headers={"Content-Type": "application/json"},
            name="Add item to cart"
        )
        if response.status_code >= 400:
            print(f"Failed to add item to cart: {response.status_code} - {response.text}")

    def remove_item_from_cart(self):
        """Remove a random product from cart"""
        product = self._select_random_product()

        payload = {
            "user_id": self.user_id,
            "product_id": str(product.get("id")),
        }

        self.client.delete(f"{self._cart_host}/cart/item/delete",
            json=payload,
            headers={"Content-Type": "application/json"},
            name="Remove item from cart"
        )

    def view_cart(self):
        """View the current cart contents"""
        self.client.get(f"{self._cart_host}/cart?userId={self.user_id}", name="View cart")