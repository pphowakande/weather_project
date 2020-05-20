-- Table: public.master_city

-- DROP TABLE public.master_city;

CREATE TABLE public.master_city
(
    city_id integer NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    country text COLLATE pg_catalog."default" NOT NULL,
    location geometry NOT NULL,
    downloaded boolean NOT NULL,
    id integer NOT NULL DEFAULT nextval('master_city_id_seq'::regclass),
    CONSTRAINT master_city_pkey PRIMARY KEY (id)
)

-- Table: public.weather_codes

-- DROP TABLE public.weather_codes;

CREATE TABLE public.weather_codes
(
    id integer NOT NULL DEFAULT nextval('weather_codes_id_seq'::regclass),
    code character(50) COLLATE pg_catalog."default" NOT NULL,
    day_condition character(100) COLLATE pg_catalog."default" NOT NULL,
    night_condition character(100) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT weather_codes_pkey PRIMARY KEY (id)
)

