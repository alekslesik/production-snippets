BEGIN;

SET statement_timeout = 0;

SET client_encoding = 'UTF8';

SET standard_conforming_strings = ON;

SET check_function_bodies = FALSE;

SET client_min_messages = WARNING;

SET search_path = public, extensions;

SET default_tablespace = '';

SET default_with_oids = FALSE;

-- EXTENSIONS --
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- TABLES --

-- currencies --
CREATE TABLE public.currency 
(
    id SERIAL PRIMARY KEY,
    name TEXT,
    symbol TEXT
);

-- categories --
CREATE TABLE public.category
(
    id SERIAL PRIMARY KEY,
    name TEXT
);

-- products --
CREATE TABLE public.product
(
    -- on default generetated in db --
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    descrition TEXT NOT NULL,
    price BIGINT,
currency_id INT REFERENCES public.currency(id),
    rating INT,
    category_id INT REFERENCES public.category(id),
    specification JSONB,
    image_id UUID,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT positive_price CHECK (price > 0),
    CONSTRAINT valid_rating CHECK (rating <= 5)
);

-- DATA --

INSERT INTO public.currency (name, symbol)
VALUES ('ruble', 'R');

INSERT INTO public.currency (name, symbol)
VALUES ('dollar', '$');

INSERT INTO public.category (name)
VALUES ('coupons');

INSERT INTO public.category (name)
VALUES ('tickets');



COMMIT;