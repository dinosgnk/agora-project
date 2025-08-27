from locust import task, between
from .base_user import BaseAgoraUser

class ActiveShopper(BaseAgoraUser):
    """Active Shopper - frequently adds items and checks cart"""
    weight = 2  
    wait_time = between(1, 3)
    
    @task(3)
    def get_products_by_category(self):
        super().get_products_by_category()
    
    @task(5)
    def add_item_to_cart(self):
        super().add_item_to_cart()
    
    @task(3)
    def view_cart(self):
        super().view_cart()
    
    @task(2)
    def get_product_details(self):
        super().get_product_details()
