    CREATE TABLE IF NOT EXISTS stock(
        id SERIAL PRIMARY KEY,
        product_id INT NOT NULL,
        service_id INT NOT NULL,
        quantity INT NOT NULL,

        FOREIGN KEY (product_id) REFERENCES products(id),
        FOREIGN KEY (service_id) REFERENCES services(id)
    );