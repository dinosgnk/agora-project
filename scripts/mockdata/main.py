from generators.manager import MockDataManager

def main():
    manager = MockDataManager()
    
    products = manager.generate_data('products', 3000)

    manager.save_to_json_file(products, 'products.json')
    manager.save_to_sql_file('products', products, 'product.sql')


if __name__ == "__main__":
    main()