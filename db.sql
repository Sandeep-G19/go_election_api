CREATE DATABASE IF NOT EXISTS election_db;

USE election_db;

-- Table: elections
CREATE TABLE IF NOT EXISTS elections (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    description TEXT
);

-- Table: political_parties
CREATE TABLE IF NOT EXISTS political_parties (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    leader VARCHAR(255)
);

-- Table: candidates
CREATE TABLE IF NOT EXISTS candidates (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    party_id INT,
    FOREIGN KEY (party_id) REFERENCES political_parties(id)
);

-- Table: voters
CREATE TABLE IF NOT EXISTS voters (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    dob DATE
);

-- Table: districts
CREATE TABLE IF NOT EXISTS districts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

-- Table: campaigns
CREATE TABLE IF NOT EXISTS campaigns (
    id INT AUTO_INCREMENT PRIMARY KEY,
    candidate_id INT NOT NULL,
    district_id INT,
    FOREIGN KEY (candidate_id) REFERENCES candidates(id),
    FOREIGN KEY (district_id) REFERENCES districts(id),
    start_date DATE,
    end_date DATE,
    budget DECIMAL(10, 2)
);

-- Table: results
CREATE TABLE IF NOT EXISTS results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    election_id INT NOT NULL,
    candidate_id INT,
    votes INT NOT NULL,
    FOREIGN KEY (election_id) REFERENCES elections(id),
    FOREIGN KEY (candidate_id) REFERENCES candidates(id)
);

-- Table: users (for authentication)
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert data
INSERT INTO political_parties (name, leader) VALUES ('Party A', 'Leader A'), ('Party B', 'Leader B');
INSERT INTO candidates (name, party_id) VALUES ('Modi', 1), ('Rahul', 2);
INSERT INTO elections (name, date, description) VALUES ('General Elections 2024', '2024-05-01', 'General elections for the country.');
INSERT INTO districts (name, description) VALUES ('District 1', 'Description for District 1'), ('District 2', 'Description for District 2');
INSERT INTO voters (name, address, dob) VALUES ('Voter 1', 'Address 1', '1980-01-01'), ('Voter 2', 'Address 2', '1990-02-02');
INSERT INTO campaigns (candidate_id, district_id, start_date, end_date, budget) VALUES (1, 1, '2024-01-01', '2024-04-30', 50000), (2, 2, '2024-01-01', '2024-04-30', 60000);
INSERT INTO results (election_id, candidate_id, votes) VALUES (1, 1, 1600), (1, 2, 2000);
