--
-- PostgreSQL database dump
--

-- Dumped from database version 12.1
-- Dumped by pg_dump version 12.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: category; Type: TABLE; Schema: public; Owner:
--

CREATE TABLE public.category (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    enable boolean DEFAULT true NOT NULL
);


--
-- Name: category_product; Type: TABLE; Schema: public;
--

CREATE TABLE public.category_product (
    product_id integer,
    category_id integer
);

--
-- Name: image; Type: TABLE; Schema: public;
--

CREATE TABLE public.image (
    id integer NOT NULL,
    name character varying(255),
    file character varying(50),
    enable boolean DEFAULT true NOT NULL
);



--
-- Name: image_id_seq; Type: SEQUENCE; Schema: public;
--

CREATE SEQUENCE public.image_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: image_id_seq; Type: SEQUENCE OWNED BY; Schema: public;
--

ALTER SEQUENCE public.image_id_seq OWNED BY public.image.id;


--
-- Name: product; Type: TABLE; Schema: public;
--

CREATE TABLE public.product (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    enable boolean DEFAULT true NOT NULL
);


--
-- Name: product_id_seq; Type: SEQUENCE; Schema: public;
--

CREATE SEQUENCE public.product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: product_id_seq; Type: SEQUENCE OWNED BY; Schema: public;
--

ALTER SEQUENCE public.product_id_seq OWNED BY public.product.id;


--
-- Name: product_image; Type: TABLE; Schema: public;
--

CREATE TABLE public.product_image (
    product_id integer,
    image_id integer
);

--
-- Name: table_namecategory_id_seq; Type: SEQUENCE; Schema: public;
--

CREATE SEQUENCE public.table_namecategory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: table_namecategory_id_seq; Type: SEQUENCE OWNED BY; Schema: public;
--

ALTER SEQUENCE public.table_namecategory_id_seq OWNED BY public.category.id;


--
-- Name: category id; Type: DEFAULT; Schema: public;
--

ALTER TABLE ONLY public.category ALTER COLUMN id SET DEFAULT nextval('public.table_namecategory_id_seq'::regclass);


--
-- Name: image id; Type: DEFAULT; Schema: public;
--

ALTER TABLE ONLY public.image ALTER COLUMN id SET DEFAULT nextval('public.image_id_seq'::regclass);


--
-- Name: product id; Type: DEFAULT; Schema: public;
--

ALTER TABLE ONLY public.product ALTER COLUMN id SET DEFAULT nextval('public.product_id_seq'::regclass);


--
-- Data for Name: category; Type: TABLE DATA; Schema: public;
--

INSERT INTO public.category (id, name, enable) VALUES (1, 'Buku', true);
INSERT INTO public.category (id, name, enable) VALUES (2, 'Obat', true);
INSERT INTO public.category (id, name, enable) VALUES (3, 'Snack', true);
INSERT INTO public.category (id, name, enable) VALUES (5, 'Inspirasi', true);
INSERT INTO public.category (id, name, enable) VALUES (4, 'Makanan dan Minuman', true);


--
-- Data for Name: category_product; Type: TABLE DATA; Schema: public;
--

INSERT INTO public.category_product (product_id, category_id) VALUES (1, 1);
INSERT INTO public.category_product (product_id, category_id) VALUES (2, 1);
INSERT INTO public.category_product (product_id, category_id) VALUES (4, 1);
INSERT INTO public.category_product (product_id, category_id) VALUES (4, 5);
INSERT INTO public.category_product (product_id, category_id) VALUES (3, 3);
INSERT INTO public.category_product (product_id, category_id) VALUES (3, 4);


--
-- Data for Name: image; Type: TABLE DATA; Schema: public;
--

INSERT INTO public.image (id, name, file, enable) VALUES (2, 'Biskuit Khong Huan', 'images/konghuan.jpeg', true);


--
-- Data for Name: product; Type: TABLE DATA; Schema: public;
--

INSERT INTO public.product (id, name, description, enable) VALUES (1, 'Mencari Jati Diri', 'Buku tentang mencari jati diri yang sejati', true);
INSERT INTO public.product (id, name, description, enable) VALUES (2, 'Jalan Ninjaku', 'Buku tentang bagaimana cara menjadi ninja seperti naruto', true);
INSERT INTO public.product (id, name, description, enable) VALUES (3, 'Biskuit Macan', 'Biskuit enak bergizi dan menjadi obat laparmu', true);
INSERT INTO public.product (id, name, description, enable) VALUES (4, '100 Merubah Hidup Jadi Lebih Baik', 'Buku bagaiman menjadi web developer selama 100 hari', true);


--
-- Data for Name: product_image; Type: TABLE DATA; Schema: public;
--

INSERT INTO public.product_image (product_id, image_id) VALUES (3, 2);


--
-- Name: image_id_seq; Type: SEQUENCE SET; Schema: public;
--

SELECT pg_catalog.setval('public.image_id_seq', 2, true);


--
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public;
--

SELECT pg_catalog.setval('public.product_id_seq', 4, true);


--
-- Name: table_namecategory_id_seq; Type: SEQUENCE SET; Schema: public;
--

SELECT pg_catalog.setval('public.table_namecategory_id_seq', 5, true);


--
-- Name: image image_pk; Type: CONSTRAINT; Schema: public;
--

ALTER TABLE ONLY public.image
    ADD CONSTRAINT image_pk PRIMARY KEY (id);


--
-- Name: product product_pk; Type: CONSTRAINT; Schema: public;
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pk PRIMARY KEY (id);


--
-- Name: category table_namecategory_pk; Type: CONSTRAINT; Schema: public;
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT table_namecategory_pk PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

