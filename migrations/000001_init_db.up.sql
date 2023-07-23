CREATE TABLE users (
          id        SERIAL PRIMARY KEY,
          login      VARCHAR(255) NOT NULL UNIQUE,
          password   VARCHAR(255) NOT NULL,
          created_at TIMESTAMP    NOT NULL DEFAULT now()
);


CREATE TABLE login_credentials (
          id SERIAL PRIMARY KEY,
          user_id INT NOT NULL,
          login VARCHAR(255) NOT NULL,
          password VARCHAR(255) NOT NULL,
          created_at TIMESTAMP DEFAULT NOW(),
          FOREIGN KEY (user_id) REFERENCES users (id)
);


CREATE TABLE text_data (
          id SERIAL PRIMARY KEY,
          user_id INT NOT NULL,
          data TEXT NOT NULL,
          created_at TIMESTAMP DEFAULT NOW(),
          FOREIGN KEY (user_id) REFERENCES users (id)
);


CREATE TABLE binary_data (
          id SERIAL PRIMARY KEY,
          user_id INT NOT NULL,
          data BYTEA NOT NULL,
          created_at TIMESTAMP DEFAULT NOW(),
          FOREIGN KEY (user_id) REFERENCES users (id)
);


CREATE TABLE bank_cards (
          id SERIAL PRIMARY KEY,
          user_id INT NOT NULL,
          card_number VARCHAR(16) NOT NULL,
          expiration_date DATE NOT NULL,
          cvv VARCHAR(3) NOT NULL,
          created_at TIMESTAMP DEFAULT NOW(),
          FOREIGN KEY (user_id) REFERENCES users (id)
);

