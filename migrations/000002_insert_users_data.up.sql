-- Filename: migrations/000002_insert_users_data.up.sql
INSERT INTO users(first_name, last_name, email, age, user_address, phone_number, roles_id, user_password, activated)
VALUES 
('Jane', 'Doe', 'sweetgirl@gmail.com', '20', '13 Mahogany Street', '666-9999', 2, 'iloveUB2023', TRUE),
('Starla', 'Donde', 'sweetstar@gmail.com', '21', '10 Gill Street', '615-9876', 1, 'UBisthebest23', TRUE),
('John', 'Smith', 'johnsmith@gmail.com', '22', '5 Oak Lane', '555-1234', 2, 'password123', TRUE),
('Emily', 'Johnson', 'emilyjohnson@gmail.com', '19', '20 Maple Street', '555-5678', 2, 'mypassword', TRUE),
('Michael', 'Brown', 'michaelbrown@gmail.com', '20', '15 Elm Street', '555-4321', 2, '12345678', TRUE),
('Jessica', 'Davis', 'jessicadavis@gmail.com', '18', '3 Pine Road', '555-8765', 2, 'password1', TRUE),
('David', 'Miller', 'davidmiller@gmail.com', '21', '12 Birch Lane', '555-9876', 2, 'abcdefg', TRUE),
('Jennifer', 'Wilson', 'jenniferwilson@gmail.com', '19', '8 Cedar Street', '555-2468', 2, 'mypassword2', TRUE),
('William', 'Anderson', 'williamanderson@gmail.com', '20', '19 Maple Road', '555-1357', 2, 'password1234', TRUE),
('Ava', 'Jackson', 'avajackson@gmail.com', '18', '7 Elm Lane', '555-8642', 2, 'qwerty', TRUE),
('James', 'Thomas', 'jamesthomas@gmail.com', '21', '17 Oak Street', '555-3698', 2, '12345', TRUE),
('Sophia', 'Robinson', 'sophiarobinson@gmail.com', '19', '22 Pine Lane', '555-7856', 2, 'mypassword3', TRUE),
('Daniel', 'White', 'danielwhite@gmail.com', '20', '14 Cedar Road', '555-4682', 2, 'password12345', TRUE),
('Olivia', 'Harris', 'oliviaharris@gmail.com', '18', '6 Maple Lane', '555-9753', 2, 'abcdef', TRUE),
('Benjamin', 'Young', 'benjaminyoung@gmail.com', '21', '11 Elm Street', '555-3579', 2, 'mypassword4', TRUE),
('Elizabeth', 'Allen', 'elizabethallen@gmail.com', '19', '16 Oak Road', '555-6812', 2, 'password123456', TRUE),
('Christopher', 'King', 'christopherking@gmail.com', '20', '9 Pine Lane', '555-2589', 2, 'qwertyuiop', TRUE);