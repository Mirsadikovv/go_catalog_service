ALTER TABLE product_categories ADD FOREIGN KEY (product_id) REFERENCES product (id);
ALTER TABLE product_categories ADD FOREIGN KEY (category_id) REFERENCES category (id);
