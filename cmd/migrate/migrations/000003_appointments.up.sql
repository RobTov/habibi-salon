CREATE TABLE IF NOT EXISTS appointments(
    id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    client_id INT NOT NULL,
    service_id INT NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',

    FOREIGN KEY (client_id) REFERENCES clients(id),
    FOREIGN KEY (service_id) REFERENCES services(id)
);