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
    name TEXT,
);

-- products --
CREATE TABLE public.product
(
    -- on default generetated in db --
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    descrition TEXT NOT NULL,
    price BIGINT,
    currency_id INT,
    rating INT,
    category_id INT NOT NULL,
    specification JSONB,
    image_id UUID,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

-- DATA --

insert into public.currency (name, symbol)
values ('ruble', 'R');

insert into public.currency (name, symbol)
values ('dollar', '$');


COMMIT;