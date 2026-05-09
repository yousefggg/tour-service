ALTER TABLE bookings
ADD COLUMN phone_number TEXT DEFAULT '',
ADD COLUMN people_count INT DEFAULT 1,
ADD COLUMN notes TEXT DEFAULT '',
ADD COLUMN medical_info TEXT DEFAULT '',
ADD COLUMN payment_method TEXT DEFAULT 'card';