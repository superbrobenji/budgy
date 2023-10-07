--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)

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

--
-- Name: monthly_tracking; Type: SCHEMA; Schema: -; Owner: benji
--

CREATE SCHEMA monthly_tracking;


ALTER SCHEMA monthly_tracking OWNER TO benji;

--
-- Name: owned_resources; Type: SCHEMA; Schema: -; Owner: benji
--

CREATE SCHEMA owned_resources;


ALTER SCHEMA owned_resources OWNER TO benji;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: expense_tracking; Type: TABLE; Schema: monthly_tracking; Owner: benji
--

CREATE TABLE monthly_tracking.expense_tracking (
    id integer NOT NULL,
    name character varying(30) NOT NULL,
    amount double precision NOT NULL,
    expense_id integer NOT NULL
);


ALTER TABLE monthly_tracking.expense_tracking OWNER TO benji;

--
-- Name: expense_types; Type: TABLE; Schema: monthly_tracking; Owner: benji
--

CREATE TABLE monthly_tracking.expense_types (
    id integer NOT NULL,
    type character varying(30) NOT NULL
);


ALTER TABLE monthly_tracking.expense_types OWNER TO benji;

--
-- Name: expenses; Type: TABLE; Schema: monthly_tracking; Owner: benji
--

CREATE TABLE monthly_tracking.expenses (
    id integer NOT NULL,
    name character varying(30) NOT NULL,
    expense_type_id integer NOT NULL,
    amount double precision NOT NULL,
    asset_id integer,
    debit_order_date date,
    liability_id integer,
    user_id integer NOT NULL,
    CONSTRAINT no_assets CHECK ((((asset_id IS NULL) AND (liability_id IS NOT NULL)) OR ((asset_id IS NOT NULL) AND (liability_id IS NULL))))
);


ALTER TABLE monthly_tracking.expenses OWNER TO benji;

--
-- Name: expenses_id_seq; Type: SEQUENCE; Schema: monthly_tracking; Owner: benji
--

CREATE SEQUENCE monthly_tracking.expenses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monthly_tracking.expenses_id_seq OWNER TO benji;

--
-- Name: expenses_id_seq; Type: SEQUENCE OWNED BY; Schema: monthly_tracking; Owner: benji
--

ALTER SEQUENCE monthly_tracking.expenses_id_seq OWNED BY monthly_tracking.expenses.id;


--
-- Name: income; Type: TABLE; Schema: monthly_tracking; Owner: benji
--

CREATE TABLE monthly_tracking.income (
    id integer NOT NULL,
    name character varying(30) NOT NULL,
    amount double precision NOT NULL,
    asset_id integer,
    date_recieved date NOT NULL,
    liability_id integer,
    user_id integer NOT NULL,
    CONSTRAINT no_assets CHECK ((((asset_id IS NULL) AND (liability_id IS NOT NULL)) OR ((asset_id IS NOT NULL) AND (liability_id IS NULL))))
);


ALTER TABLE monthly_tracking.income OWNER TO benji;

--
-- Name: income_id_seq; Type: SEQUENCE; Schema: monthly_tracking; Owner: benji
--

CREATE SEQUENCE monthly_tracking.income_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monthly_tracking.income_id_seq OWNER TO benji;

--
-- Name: income_id_seq; Type: SEQUENCE OWNED BY; Schema: monthly_tracking; Owner: benji
--

ALTER SEQUENCE monthly_tracking.income_id_seq OWNED BY monthly_tracking.income.id;


--
-- Name: types_id_seq; Type: SEQUENCE; Schema: monthly_tracking; Owner: benji
--

CREATE SEQUENCE monthly_tracking.types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monthly_tracking.types_id_seq OWNER TO benji;

--
-- Name: types_id_seq; Type: SEQUENCE OWNED BY; Schema: monthly_tracking; Owner: benji
--

ALTER SEQUENCE monthly_tracking.types_id_seq OWNED BY monthly_tracking.expense_types.id;


--
-- Name: types_id_seq1; Type: SEQUENCE; Schema: monthly_tracking; Owner: benji
--

CREATE SEQUENCE monthly_tracking.types_id_seq1
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monthly_tracking.types_id_seq1 OWNER TO benji;

--
-- Name: types_id_seq1; Type: SEQUENCE OWNED BY; Schema: monthly_tracking; Owner: benji
--

ALTER SEQUENCE monthly_tracking.types_id_seq1 OWNED BY monthly_tracking.expense_tracking.id;


--
-- Name: assets; Type: TABLE; Schema: owned_resources; Owner: benji
--

CREATE TABLE owned_resources.assets (
    id integer NOT NULL,
    name character varying(30) NOT NULL,
    value double precision NOT NULL,
    user_id integer
);


ALTER TABLE owned_resources.assets OWNER TO benji;

--
-- Name: assets_id_seq; Type: SEQUENCE; Schema: owned_resources; Owner: benji
--

CREATE SEQUENCE owned_resources.assets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE owned_resources.assets_id_seq OWNER TO benji;

--
-- Name: assets_id_seq; Type: SEQUENCE OWNED BY; Schema: owned_resources; Owner: benji
--

ALTER SEQUENCE owned_resources.assets_id_seq OWNED BY owned_resources.assets.id;


--
-- Name: investments; Type: TABLE; Schema: owned_resources; Owner: benji
--

CREATE TABLE owned_resources.investments (
    id integer NOT NULL,
    name character varying(30) NOT NULL,
    total_invested double precision NOT NULL,
    asset_id integer NOT NULL
);


ALTER TABLE owned_resources.investments OWNER TO benji;

--
-- Name: investments_id_seq; Type: SEQUENCE; Schema: owned_resources; Owner: benji
--

CREATE SEQUENCE owned_resources.investments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE owned_resources.investments_id_seq OWNER TO benji;

--
-- Name: investments_id_seq; Type: SEQUENCE OWNED BY; Schema: owned_resources; Owner: benji
--

ALTER SEQUENCE owned_resources.investments_id_seq OWNED BY owned_resources.investments.id;


--
-- Name: liabilities; Type: TABLE; Schema: owned_resources; Owner: benji
--

CREATE TABLE owned_resources.liabilities (
    id integer NOT NULL,
    name character varying(30) NOT NULL,
    debt_amount double precision NOT NULL,
    user_id integer
);


ALTER TABLE owned_resources.liabilities OWNER TO benji;

--
-- Name: liabilities_id_seq; Type: SEQUENCE; Schema: owned_resources; Owner: benji
--

CREATE SEQUENCE owned_resources.liabilities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE owned_resources.liabilities_id_seq OWNER TO benji;

--
-- Name: liabilities_id_seq; Type: SEQUENCE OWNED BY; Schema: owned_resources; Owner: benji
--

ALTER SEQUENCE owned_resources.liabilities_id_seq OWNED BY owned_resources.liabilities.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: benji
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name text NOT NULL,
    surname text NOT NULL,
    email text NOT NULL,
    hash text,
    salt text
);


ALTER TABLE public.users OWNER TO benji;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: benji
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO benji;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: benji
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: expense_tracking id; Type: DEFAULT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expense_tracking ALTER COLUMN id SET DEFAULT nextval('monthly_tracking.types_id_seq1'::regclass);


--
-- Name: expense_types id; Type: DEFAULT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expense_types ALTER COLUMN id SET DEFAULT nextval('monthly_tracking.types_id_seq'::regclass);


--
-- Name: expenses id; Type: DEFAULT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expenses ALTER COLUMN id SET DEFAULT nextval('monthly_tracking.expenses_id_seq'::regclass);


--
-- Name: income id; Type: DEFAULT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.income ALTER COLUMN id SET DEFAULT nextval('monthly_tracking.income_id_seq'::regclass);


--
-- Name: assets id; Type: DEFAULT; Schema: owned_resources; Owner: benji
--

ALTER TABLE ONLY owned_resources.assets ALTER COLUMN id SET DEFAULT nextval('owned_resources.assets_id_seq'::regclass);


--
-- Name: investments id; Type: DEFAULT; Schema: owned_resources; Owner: benji
--

ALTER TABLE ONLY owned_resources.investments ALTER COLUMN id SET DEFAULT nextval('owned_resources.investments_id_seq'::regclass);


--
-- Name: liabilities id; Type: DEFAULT; Schema: owned_resources; Owner: benji
--

ALTER TABLE ONLY owned_resources.liabilities ALTER COLUMN id SET DEFAULT nextval('owned_resources.liabilities_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: benji
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: expense_tracking; Type: TABLE DATA; Schema: monthly_tracking; Owner: benji
--

COPY monthly_tracking.expense_tracking (id, name, amount, expense_id) FROM stdin;
2	home loan	2000	2
\.


--
-- Data for Name: expense_types; Type: TABLE DATA; Schema: monthly_tracking; Owner: benji
--

COPY monthly_tracking.expense_types (id, type) FROM stdin;
1	non-essential
\.


--
-- Data for Name: expenses; Type: TABLE DATA; Schema: monthly_tracking; Owner: benji
--

COPY monthly_tracking.expenses (id, name, expense_type_id, amount, asset_id, debit_order_date, liability_id, user_id) FROM stdin;
2	tfsa	1	2000	1	2023-06-26	\N	1
\.


--
-- Data for Name: income; Type: TABLE DATA; Schema: monthly_tracking; Owner: benji
--

COPY monthly_tracking.income (id, name, amount, asset_id, date_recieved, liability_id, user_id) FROM stdin;
1	Shyft	300	1	2023-06-23	\N	1
4	tfsa	300	\N	2023-06-23	2	1
6	tfsa	300	1	2023-06-23	\N	1
\.


--
-- Data for Name: assets; Type: TABLE DATA; Schema: owned_resources; Owner: benji
--

COPY owned_resources.assets (id, name, value, user_id) FROM stdin;
1	TFSA	93000	2
\.


--
-- Data for Name: investments; Type: TABLE DATA; Schema: owned_resources; Owner: benji
--

COPY owned_resources.investments (id, name, total_invested, asset_id) FROM stdin;
1	TFSA	43000	1
\.


--
-- Data for Name: liabilities; Type: TABLE DATA; Schema: owned_resources; Owner: benji
--

COPY owned_resources.liabilities (id, name, debt_amount, user_id) FROM stdin;
2	home loan	3000	1
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: benji
--

COPY public.users (id, name, surname, email, hash, salt) FROM stdin;
1	Jerry	Pringle	jerry@example.com	\N	\N
2	George	Smith	george@example.com	\N	\N
4	Benjamin	Swanepoel	bennie.h.swanepoel@gmail.com	1f98ad9dcb1cf55c110ae4f7bc345dec61752a84e66198dcfb23800303d5a17c005e36a28db80f50451f7b7593712e228dd0d884b3f7c44db206004fd0de0f31	57a94c6b278a1024b7848f9712a86a27
\.


--
-- Name: expenses_id_seq; Type: SEQUENCE SET; Schema: monthly_tracking; Owner: benji
--

SELECT pg_catalog.setval('monthly_tracking.expenses_id_seq', 3, true);


--
-- Name: income_id_seq; Type: SEQUENCE SET; Schema: monthly_tracking; Owner: benji
--

SELECT pg_catalog.setval('monthly_tracking.income_id_seq', 6, true);


--
-- Name: types_id_seq; Type: SEQUENCE SET; Schema: monthly_tracking; Owner: benji
--

SELECT pg_catalog.setval('monthly_tracking.types_id_seq', 1, true);


--
-- Name: types_id_seq1; Type: SEQUENCE SET; Schema: monthly_tracking; Owner: benji
--

SELECT pg_catalog.setval('monthly_tracking.types_id_seq1', 2, true);


--
-- Name: assets_id_seq; Type: SEQUENCE SET; Schema: owned_resources; Owner: benji
--

SELECT pg_catalog.setval('owned_resources.assets_id_seq', 1, true);


--
-- Name: investments_id_seq; Type: SEQUENCE SET; Schema: owned_resources; Owner: benji
--

SELECT pg_catalog.setval('owned_resources.investments_id_seq', 2, true);


--
-- Name: liabilities_id_seq; Type: SEQUENCE SET; Schema: owned_resources; Owner: benji
--

SELECT pg_catalog.setval('owned_resources.liabilities_id_seq', 2, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: benji
--

SELECT pg_catalog.setval('public.users_id_seq', 4, true);


--
-- Name: expenses expenses_pkey; Type: CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expenses
    ADD CONSTRAINT expenses_pkey PRIMARY KEY (id);


--
-- Name: income income_pkey; Type: CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.income
    ADD CONSTRAINT income_pkey PRIMARY KEY (id);


--
-- Name: expense_types types_pkey; Type: CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expense_types
    ADD CONSTRAINT types_pkey PRIMARY KEY (id);


--
-- Name: expense_tracking types_pkey1; Type: CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expense_tracking
    ADD CONSTRAINT types_pkey1 PRIMARY KEY (id);


--
-- Name: assets assets_pkey; Type: CONSTRAINT; Schema: owned_resources; Owner: benji
--

ALTER TABLE ONLY owned_resources.assets
    ADD CONSTRAINT assets_pkey PRIMARY KEY (id);


--
-- Name: investments investments_asset_id_key; Type: CONSTRAINT; Schema: owned_resources; Owner: benji
--

ALTER TABLE ONLY owned_resources.investments
    ADD CONSTRAINT investments_asset_id_key UNIQUE (asset_id);


--
-- Name: investments investments_pkey; Type: CONSTRAINT; Schema: owned_resources; Owner: benji
--

ALTER TABLE ONLY owned_resources.investments
    ADD CONSTRAINT investments_pkey PRIMARY KEY (id);


--
-- Name: liabilities liabilities_pkey; Type: CONSTRAINT; Schema: owned_resources; Owner: benji
--

ALTER TABLE ONLY owned_resources.liabilities
    ADD CONSTRAINT liabilities_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: benji
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: benji
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: income fk_asset; Type: FK CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.income
    ADD CONSTRAINT fk_asset FOREIGN KEY (asset_id) REFERENCES owned_resources.assets(id) ON DELETE CASCADE;


--
-- Name: expenses fk_asset; Type: FK CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expenses
    ADD CONSTRAINT fk_asset FOREIGN KEY (asset_id) REFERENCES owned_resources.assets(id) ON DELETE CASCADE;


--
-- Name: expense_tracking fk_expense; Type: FK CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expense_tracking
    ADD CONSTRAINT fk_expense FOREIGN KEY (expense_id) REFERENCES monthly_tracking.expenses(id) ON DELETE CASCADE;


--
-- Name: expenses fk_expense_type; Type: FK CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expenses
    ADD CONSTRAINT fk_expense_type FOREIGN KEY (expense_type_id) REFERENCES monthly_tracking.expense_types(id) ON DELETE CASCADE;


--
-- Name: income fk_liability; Type: FK CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.income
    ADD CONSTRAINT fk_liability FOREIGN KEY (liability_id) REFERENCES owned_resources.liabilities(id) ON DELETE CASCADE;


--
-- Name: expenses fk_liability; Type: FK CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expenses
    ADD CONSTRAINT fk_liability FOREIGN KEY (liability_id) REFERENCES owned_resources.liabilities(id) ON DELETE CASCADE;


--
-- Name: income fk_user; Type: FK CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.income
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: expenses fk_user; Type: FK CONSTRAINT; Schema: monthly_tracking; Owner: benji
--

ALTER TABLE ONLY monthly_tracking.expenses
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: investments fk_asset; Type: FK CONSTRAINT; Schema: owned_resources; Owner: benji
--

ALTER TABLE ONLY owned_resources.investments
    ADD CONSTRAINT fk_asset FOREIGN KEY (asset_id) REFERENCES owned_resources.assets(id) ON DELETE CASCADE;


--
-- Name: liabilities fk_user; Type: FK CONSTRAINT; Schema: owned_resources; Owner: benji
--

ALTER TABLE ONLY owned_resources.liabilities
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: assets fk_user; Type: FK CONSTRAINT; Schema: owned_resources; Owner: benji
--

ALTER TABLE ONLY owned_resources.assets
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

