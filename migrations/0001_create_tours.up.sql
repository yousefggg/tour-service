CREATE TABLE IF NOT EXISTS tours (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    title TEXT NOT NULL,
    location TEXT NOT NULL,
    country TEXT NOT NULL,
    season TEXT NOT NULL,

    price BIGINT NOT NULL,

    image_url TEXT NOT NULL,

    description TEXT NOT NULL,
    includes TEXT NOT NULL,

    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_tours_country ON tours(country);
CREATE INDEX IF NOT EXISTS idx_tours_season ON tours(season);
CREATE INDEX IF NOT EXISTS idx_tours_price ON tours(price);