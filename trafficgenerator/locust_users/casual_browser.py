from locust import task, between
from .base_user import BaseAgoraUser

class CasualBrowser(BaseAgoraUser):
    """Casual Browser - mostly browses, rarely buys"""
    weight = 3 
    wait_time = between(2, 6)
    
    @task(10)
    def get_products_by_category(self):
        super().get_products_by_category()
    
    @task(5)
    def get_product_details(self):
        super().get_product_details()
    
    @task(1)
    def add_item_to_cart(self):
        super().add_item_to_cart()
    
    @task(1)
    def view_cart(self):
        super().view_cart()
