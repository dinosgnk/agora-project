from locust import task, between
from .base_user import BaseAgoraUser

class ActiveShopper(BaseAgoraUser):
    """Active Shopper - browses, then creates order"""
    weight = 2  
    wait_time = between(1, 3)
    
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.order_created = False
    
    @task(1)
    def get_products_by_category(self):
        if self.order_created:
            return
        super().get_products_by_category()
    
    @task(5)
    def add_item_to_cart(self):
        if self.order_created:
            return
        super().add_item_to_cart()
    
    @task(2)
    def view_cart(self):
        if self.order_created:
            return
        super().view_cart()
    
    @task(2)
    def get_product_details(self):
        if self.order_created:
            return
        super().get_product_details()

    @task(1)
    def try_create_order(self):
        if self.order_created:
            return

        if len(self.cart_items) > 0:
            super().create_order()
            self.order_created = True
            self.stop()
