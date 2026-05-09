CREATE TABLE IF NOT EXISTS bookings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,
    tour_id UUID NOT NULL,

    status TEXT NOT NULL DEFAULT 'pending',

    created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_bookings_user_id ON bookings(user_id);
CREATE INDEX IF NOT EXISTS idx_bookings_tour_id ON bookings(tour_id);

ALTER TABLE bookings
ADD CONSTRAINT fk_bookings_tour
FOREIGN KEY (tour_id) REFERENCES tours(id)
ON DELETE CASCADE;