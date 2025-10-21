from locust import task, between
from .base_user import BaseAgoraUser

class CasualBrowser(BaseAgoraUser):
    """Casual Browser - mostly browses, rarely buys"""
    weight = 3 
    wait_time = between(2, 10)
    
    @task(1)
    def get_products_by_category(self):
        super().get_products_by_category()
    
    @task(1)
    def get_product_details(self):
        super().get_product_details()
