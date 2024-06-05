CREATE TABLE IF NOT EXISTS category(
    id uuid PRIMARY KEY,
    slug VARCHAR(25) DEFAULT '',
    name_uz VARCHAR(25) default '', 
    name_ru VARCHAR(25) default '',
    name_en VARCHAR(25) default '',
    active BOOLEAN DEFAULT TRUE,
    order_no INTEGER DEFAULT 0,
    parent_id uuid REFERENCES category (id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product(
    id uuid PRIMARY KEY,
    slug VARCHAR(25),
    name_uz VARCHAR(25) default '', 
    name_ru VARCHAR(25) default '',
    name_en VARCHAR(25) default '',
    description_uz VARCHAR(25) default '', 
    description_ru VARCHAR(25) default '',
    description_en VARCHAR(25) default '',
    active BOOLEAN DEFAULT TRUE,
    order_no INTEGER DEFAULT 0,
    in_price FLOAT,
    out_price FLOAT,
    left_count INTEGER,
    discount_percent FLOAT default 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_categories (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    category_id UUID NOT NULL,
    CONSTRAINT unique_product_category UNIQUE (product_id, category_id)
);

CREATE TABLE IF NOT EXISTS product_reviews (
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL,
    product_id UUID NOT NULL,
    text VARCHAR(500),
    rating FLOAT,
    order_id UUID,
    created_at TIMESTAMP DEFAULT NOW()
);
