from locust import HttpUser, between
from typing import Optional, Dict

import os
import random
import helpers

class BaseAgoraUser(HttpUser):
    abstract = True
    
    # Dummy assignment to avoid errors, variable not used
    host = "not_used"

    _catalog_host = f"http://{os.getenv('CATALOG_SERVICE')}"
    _cart_host = f"http://{os.getenv('CART_SERVICE')}"
    _order_host = f"http://{os.getenv('ORDER_SERVICE')}"

    _products = helpers.get_all_products()
    _categories = helpers.get_all_product_categories()

    _user_counter = 0
    
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        BaseAgoraUser._user_counter += 1
        self.user_id = str(BaseAgoraUser._user_counter)
        self.cart_items = {}

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
        if product and product.get("ProductCode"):
            self.client.get(f"{self._catalog_host}/products/{product['ProductCode']}", name="Get product details")

    # Cart
    def add_item_to_cart(self):
        """Add a random product to cart"""
        product = self._select_random_product()
        
        # TODO: Randomize quantity with a more realistic distribution
        quantity = random.randint(1, 3) 
        product_code = str(product.get("ProductCode"))

        payload = {
            "user_id": self.user_id,
            "item": {
                "product_code": product_code,
                "name": product.get("Name"),
                "quantity": quantity,
                "price": float(product.get("Price"))
            }
        }

        print(f"User {self.user_id} adding to cart: {payload}")

        response = self.client.post(
            f"{self._cart_host}/cart/item/add/{self.user_id}",
            json=payload,
            headers={"Content-Type": "application/json"},
            name="Add item to cart"
        )
        
        if response.status_code >= 400:
            print(f"Failed to add item to cart: {response.status_code} - {response.text}")
        else:
            if product_code in self.cart_items:
                # Update quantity if item already exists
                self.cart_items[product_code]["quantity"] += quantity
            else:
                self.cart_items[product_code] = {
                    "name": product.get("Name"),
                    "quantity": quantity,
                    "price": float(product.get("Price"))
                }
            print(f"User {self.user_id} cart now has {len(self.cart_items)} unique items")

    def remove_item_from_cart(self):
        """Remove a random product from cart"""
        product = self._select_random_product()
        product_code = str(product.get("ProductCode"))

        payload = {
            "user_id": self.user_id,
            "product_code": product_code,
        }

        response = self.client.delete(f"{self._cart_host}/cart/item/delete",
            json=payload,
            headers={"Content-Type": "application/json"},
            name="Remove item from cart"
        )
        
        if response.status_code < 400 and product_code in self.cart_items:
            del self.cart_items[product_code]
            print(f"User {self.user_id} removed {product_code} from cart")

    def view_cart(self):
        """View the current cart contents"""
        self.client.get(f"{self._cart_host}/cart/{self.user_id}", name="View cart")

    # Order
    def create_order(self):
        """Create an order from the current cart"""

        if not self.cart_items:
            print(f"User {self.user_id} has empty cart, cannot create order")
            return
        
        products = []
        for product_code, item_data in self.cart_items.items():
            products.append({
                "code": product_code,
                "product_name": item_data["name"],
                "quantity": item_data["quantity"],
                "price": item_data["price"]
            })

        payload = {
            "user_id": self.user_id,
            "products": products,
            "shipping_address": f"{random.randint(100, 9999)} Main St, City, State {random.randint(10000, 99999)}",
            "payment_method": "crypto"
        }
        
        response = self.client.post(
            f"{self._order_host}/orders",
            json=payload,
            headers={"Content-Type": "application/json"},
            name="Create order"
        )
        
        if response.status_code >= 400:
            print(f"Failed to create order for user {self.user_id}: {response.status_code} - {response.text}")
        else:
            print(f"User {self.user_id} successfully created order")
            self.cart_items = {}