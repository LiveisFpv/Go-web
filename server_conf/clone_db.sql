--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

-- Started on 2024-12-21 15:51:12

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
-- TOC entry 2 (class 3079 OID 16824)
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- TOC entry 4980 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


--
-- TOC entry 240 (class 1255 OID 16834)
-- Name: assign_scholarships(character varying, character varying); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.assign_scholarships(IN current_semester character varying, IN budget_type character varying)
    LANGUAGE plpgsql
    AS $$
DECLARE
    target_budget_id INTEGER;
BEGIN
    SELECT id_budget 
    INTO target_budget_id
    FROM budget
    WHERE name_semester = current_semester AND type_scholarship_budget = budget_type
    LIMIT 1;

    IF target_budget_id IS NULL THEN
        RAISE EXCEPTION 'Бюджет не найден для семестра % и типа %', current_semester, budget_type;
    END IF;

    DELETE FROM scholarship
    WHERE name_semester = current_semester;

    INSERT INTO scholarship (id_num_student, name_semester, size_scholarshp, id_budget)
    SELECT 
        m.id_num_student, 
        current_semester, 
        4500+(AVG(m.score_mark)-80)*100,
        target_budget_id
    FROM 
        mark m
    WHERE 
        m.name_semester = current_semester
    GROUP BY 
        m.id_num_student, m.name_semester
    HAVING 
        AVG(m.score_mark) >= 80;

    RAISE NOTICE 'Стипендии начислены студентам за семестр: % и тип бюджета: %', current_semester, budget_type;
END;
$$;


ALTER PROCEDURE public.assign_scholarships(IN current_semester character varying, IN budget_type character varying) OWNER TO postgres;

--
-- TOC entry 241 (class 1255 OID 16835)
-- Name: check_mark_score(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.check_mark_score() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.score_mark < 0 OR NEW.score_mark > 100 THEN
        RAISE EXCEPTION 'Оценка должна быть в диапазоне от 0 до 100. Некорректное значение: %', NEW.score_mark;
    END IF;
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.check_mark_score() OWNER TO postgres;

--
-- TOC entry 242 (class 1255 OID 16836)
-- Name: set_type_mark(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.set_type_mark() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.score_mark >= 93 THEN
        NEW.type_mark := 'Отлично';
    ELSIF NEW.score_mark >= 80 THEN
        NEW.type_mark := 'Хорошо';
    ELSIF NEW.score_mark >= 53 THEN
        NEW.type_mark := 'Удовлетворительно';
    ELSE
        NEW.type_mark := 'Неудовлетворительно';
    END IF;

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.set_type_mark() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 16837)
-- Name: achievement; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.achievement (
    id_achivment integer NOT NULL,
    id_num_student integer NOT NULL,
    id_category integer NOT NULL,
    name_achivement character varying(200) NOT NULL,
    date_achivment date NOT NULL
);


ALTER TABLE public.achievement OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16840)
-- Name: achievement_id_achivment_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.achievement_id_achivment_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.achievement_id_achivment_seq OWNER TO postgres;

--
-- TOC entry 4981 (class 0 OID 0)
-- Dependencies: 217
-- Name: achievement_id_achivment_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.achievement_id_achivment_seq OWNED BY public.achievement.id_achivment;


--
-- TOC entry 218 (class 1259 OID 16841)
-- Name: budget; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.budget (
    type_scholarship_budget character varying(100) NOT NULL,
    name_semester character varying(5) NOT NULL,
    size_budget numeric NOT NULL,
    id_budget integer NOT NULL
);


ALTER TABLE public.budget OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16844)
-- Name: budget_id_budget_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.budget_id_budget_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.budget_id_budget_seq OWNER TO postgres;

--
-- TOC entry 4982 (class 0 OID 0)
-- Dependencies: 219
-- Name: budget_id_budget_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.budget_id_budget_seq OWNED BY public.budget.id_budget;


--
-- TOC entry 220 (class 1259 OID 16845)
-- Name: category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.category (
    id_category integer NOT NULL,
    achivments_type_category character varying(50) NOT NULL,
    score_category smallint NOT NULL
);


ALTER TABLE public.category OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16848)
-- Name: category_id_category_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.category_id_category_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.category_id_category_seq OWNER TO postgres;

--
-- TOC entry 4983 (class 0 OID 0)
-- Dependencies: 221
-- Name: category_id_category_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.category_id_category_seq OWNED BY public.category.id_category;


--
-- TOC entry 222 (class 1259 OID 16849)
-- Name: group; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."group" (
    name_group character varying(10) NOT NULL,
    studies_direction_group character varying(8) NOT NULL,
    studies_profile_group character varying(50) NOT NULL,
    start_date_group date NOT NULL,
    studies_period_group smallint NOT NULL
);


ALTER TABLE public."group" OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16852)
-- Name: mark; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mark (
    id_mark integer NOT NULL,
    id_num_student integer NOT NULL,
    name_semester character varying(5) NOT NULL,
    lesson_name_mark character varying(100) NOT NULL,
    score_mark smallint NOT NULL,
    type_mark character varying(20)
);


ALTER TABLE public.mark OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 16855)
-- Name: mark_id_mark_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mark_id_mark_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.mark_id_mark_seq OWNER TO postgres;

--
-- TOC entry 4984 (class 0 OID 0)
-- Dependencies: 224
-- Name: mark_id_mark_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mark_id_mark_seq OWNED BY public.mark.id_mark;


--
-- TOC entry 225 (class 1259 OID 16856)
-- Name: scholarship; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scholarship (
    id_scholarship integer NOT NULL,
    id_num_student integer NOT NULL,
    name_semester character varying(5) NOT NULL,
    size_scholarshp numeric NOT NULL,
    id_budget integer NOT NULL
);


ALTER TABLE public.scholarship OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 16859)
-- Name: scholarship_id_scholarship_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scholarship_id_scholarship_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.scholarship_id_scholarship_seq OWNER TO postgres;

--
-- TOC entry 4985 (class 0 OID 0)
-- Dependencies: 226
-- Name: scholarship_id_scholarship_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scholarship_id_scholarship_seq OWNED BY public.scholarship.id_scholarship;


--
-- TOC entry 227 (class 1259 OID 16860)
-- Name: semester; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.semester (
    name_semester character varying(5) NOT NULL,
    date_start_semester date NOT NULL,
    date_end_semester date NOT NULL
);


ALTER TABLE public.semester OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 16863)
-- Name: student; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.student (
    id_num_student integer NOT NULL,
    name_group character varying(10) NOT NULL,
    email_student character varying(40) NOT NULL,
    second_name_student character varying(40) NOT NULL,
    first_name_student character varying(40) NOT NULL,
    surname_student character varying(40)
);

CREATE TABLE public.user (
    user_id serial NOT NULL,
    user_login character varying(40) NOT NULL UNIQUE,
    user_email character varying(40) NOT NULL UNIQUE,
    user_password text NOT NULL,
    PRIMARY KEY (user_id)
);


ALTER TABLE public.student OWNER TO postgres;

--
-- TOC entry 4771 (class 2604 OID 16866)
-- Name: achievement id_achivment; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.achievement ALTER COLUMN id_achivment SET DEFAULT nextval('public.achievement_id_achivment_seq'::regclass);


--
-- TOC entry 4772 (class 2604 OID 16867)
-- Name: budget id_budget; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.budget ALTER COLUMN id_budget SET DEFAULT nextval('public.budget_id_budget_seq'::regclass);


--
-- TOC entry 4773 (class 2604 OID 16868)
-- Name: category id_category; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category ALTER COLUMN id_category SET DEFAULT nextval('public.category_id_category_seq'::regclass);


--
-- TOC entry 4774 (class 2604 OID 16869)
-- Name: mark id_mark; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mark ALTER COLUMN id_mark SET DEFAULT nextval('public.mark_id_mark_seq'::regclass);


--
-- TOC entry 4775 (class 2604 OID 16870)
-- Name: scholarship id_scholarship; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scholarship ALTER COLUMN id_scholarship SET DEFAULT nextval('public.scholarship_id_scholarship_seq'::regclass);


--
-- TOC entry 4962 (class 0 OID 16837)
-- Dependencies: 216
-- Data for Name: achievement; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.achievement (id_achivment, id_num_student, id_category, name_achivement, date_achivment) FROM stdin;
1	1220060502	3	Олимпиада	2023-05-12
2	1220060408	2	Олимпиада	2023-05-09
3	1220060407	4	Грант	2023-04-11
5	1220060502	3	Олимпиада	2023-07-07
6	1220060506	1	Олимпиада	2023-06-06
7	1220060505	2	Олимпиада	2023-05-05
8	1220060503	2	Грант	2023-04-04
9	1220060508	4	Грант	2023-03-03
10	1220060504	4	Грант	2022-09-10
11	1220060504	3	Грант	2023-07-09
12	1220060504	3	Публикация	2022-11-11
13	1220060404	2	Публикация	2023-11-10
14	1220060407	4	Публикация	2022-10-05
15	1220060405	3	Публикация	2023-11-10
16	1220060407	3	Олимпиада	2023-12-19
18	1220060404	2	Олимпиада	2022-08-04
19	1220060402	3	Публикация	2022-10-14
20	1220060402	3	Публикация	2022-10-04
21	1220060402	4	Публикация	2023-11-04
22	1220060402	1	Грант	2023-10-04
23	1220060401	2	Олимпиада	2023-10-14
24	1220060401	3	Олимпиада	2023-10-11
25	1220060401	4	Олимпиада	2023-10-04
\.


--
-- TOC entry 4964 (class 0 OID 16841)
-- Dependencies: 218
-- Data for Name: budget; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.budget (type_scholarship_budget, name_semester, size_budget, id_budget) FROM stdin;
ПГАС	2023В	500000.00	1
ГАС	2023О	500000.00	2
ПГАС	2023В	500000.00	3
ГАС	2023О	500000.00	4
\.


--
-- TOC entry 4966 (class 0 OID 16845)
-- Dependencies: 220
-- Data for Name: category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category (id_category, achivments_type_category, score_category) FROM stdin;
1	Международное	10
2	Всероссийское	8
3	Региональная	6
4	Городская	4
\.


--
-- TOC entry 4968 (class 0 OID 16849)
-- Dependencies: 222
-- Data for Name: group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."group" (name_group, studies_direction_group, studies_profile_group, start_date_group, studies_period_group) FROM stdin;
ПИ-22-2	09.03.04	Разработка программных систем	2022-09-01	4
ПИ-22-1	09.03.04	Разработка программных систем	2022-09-01	4
ПИ-21-2	09.03.04	Разработка программных систем	2021-09-01	4
ПИ-21-1	09.03.04	Разработка программных систем	2021-09-01	4
ПИ-20-2	09.03.04	Разработка программных систем	2020-09-01	4
ПИ-20-1	09.03.04	Разработка программных систем	2020-09-01	4
АС-23-2	09.03.01	Информатика и вычислительная техника	2023-09-01	4
АС-23-1	09.03.01	Информатика и вычислительная техника	2023-09-01	4
АС-22-2	09.03.01	Информатика и вычислительная техника	2022-09-01	4
АС-22-1	09.03.01	Информатика и вычислительная техника	2022-09-01	4
АС-21-2	09.03.01	Информатика и вычислительная техника	2021-09-01	4
АС-21-1	09.03.01	Информатика и вычислительная техника	2021-09-01	4
АС-20-1	09.03.01	Информатика и вычислительная техника	2020-09-01	4
АИ-23	02.03.03	Разработка и автоматизация	2023-09-01	4
АИ-22	02.03.03	Разработка и автоматизация	2022-09-01	4
АИ-21	02.03.03	Разработка и автоматизация	2021-09-01	4
АИ-20	02.03.03	Разработка и автоматизация	2020-09-01	4
ПИ-23-2	09.03.04	Разработка программных систем	2023-09-01	6
ПИ-23-1	09.03.04	Разработка программных систем	2024-12-19	4
\.


--
-- TOC entry 4969 (class 0 OID 16852)
-- Dependencies: 223
-- Data for Name: mark; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mark (id_mark, id_num_student, name_semester, lesson_name_mark, score_mark, type_mark) FROM stdin;
28	1220060405	2023В	Физика	82	Хорошо
3	1220060405	2023В	Базы данных	89	Хорошо
4	1220060405	2023В	Базы данных	88	Хорошо
5	1220060407	2023О	Базы данных	87	Хорошо
6	1220060407	2023О	Базы данных	76	Удовлетворительно
9	1220060405	2023О	Базы данных	65	Удовлетворительно
10	1220060405	2023О	Базы данных	61	Удовлетворительно
11	1220060404	2023В	Программирование	75	Удовлетворительно
12	1220060404	2023В	Программирование	100	Отлично
16	1220060404	2023О	Программирование	56	Удовлетворительно
19	1220060402	2023В	Математика	66	Удовлетворительно
20	1220060402	2023В	Математика	56	Удовлетворительно
21	1220060401	2023В	Математика	76	Удовлетворительно
22	1220060401	2023В	Математика	99	Отлично
23	1220060402	2023О	Математика	89	Хорошо
24	1220060402	2023О	Математика	77	Удовлетворительно
25	1220060401	2023О	Математика	67	Удовлетворительно
26	1220060401	2023О	Математика	80	Хорошо
15	1220060404	2023О	Программирование	79	Удовлетворительно
\.


--
-- TOC entry 4971 (class 0 OID 16856)
-- Dependencies: 225
-- Data for Name: scholarship; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scholarship (id_scholarship, id_num_student, name_semester, size_scholarshp, id_budget) FROM stdin;
11	1220060503	2023В	5700.00	1
12	1220060402	2023В	5300.00	1
13	1220060401	2023В	5400.00	1
14	1220060401	2023В	15500.00	2
15	1220060402	2023В	15600.00	2
17	1220060404	2023В	5200.00	1
18	1220060405	2023В	5500.00	1
20	1220060407	2023В	5300.00	1
21	1220060402	2023О	4800.00	2
22	1220060407	2023О	4650.00	2
\.


--
-- TOC entry 4973 (class 0 OID 16860)
-- Dependencies: 227
-- Data for Name: semester; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.semester (name_semester, date_start_semester, date_end_semester) FROM stdin;
2022В	2022-02-01	2022-08-31
2022О	2022-09-01	2023-01-31
2023В	2023-02-01	2023-08-31
2023О	2023-09-01	2024-01-31
\.


--
-- TOC entry 4974 (class 0 OID 16863)
-- Dependencies: 228
-- Data for Name: student; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.student (id_num_student, name_group, email_student, second_name_student, first_name_student, surname_student) FROM stdin;
1220060508	АС-22-1	cvg@gmail.com	Обыденных	Дмитрий	Александрович
1220060506	АС-22-1	cvh@gmail.com	Коновалов	Александр	Сергеевич
1220060505	АС-22-1	rgy@gmail.com	Ефремов	Дмитрий	Сергеевич
1220060504	АС-22-1	fgh@mail.ru	Бурлов	Егор	Павлович
1220060503	АС-22-1	fth@mail.ru	Бобенко	Максим	Геннадьевич
1220060502	АС-22-1	rth@mail.ru	Безрукавников	Дмитрий	Алексеевич
1220060501	АС-22-1	rty@mail.ru	Балалыкин	Данила	Андреевич
1220060412	ПИ-22-1	asc@mail.ru	Клименко	Никита	Дмитриевич
1220060410	ПИ-22-1	zsc@mail.ru	Чупахин	Алексей	Сергеевич
1220060409	ПИ-22-1	zxf@mail.ru	Сазонов	Данила	Андреевич
1220060408	ПИ-22-1	zxc@mail.ru	Первеева	Елизавета	Юрьевна
1220060407	ПИ-22-1	ase@gmail.com	Пахомов	Александр	Александрович
1220060405	ПИ-22-1	qse@gmail.com	Кочетков	Артем	Игоревич
1220060404	ПИ-22-1	qwd@gmail.com	Кистирёв	Виктор	Денисович
1220060402	ПИ-22-1	asd@gmail.com	Боярчук	Григорий	Сергеевич
1220060401	ПИ-22-1	qwe@gmail.com	Баженов	Кирилл	Александрович
1220060603	АИ-22	IVAN@gmail.com	Иван	Синюков	Евгеньевич
1220060420	ПИ-22-1	ivantest@mail.ru	Иванов	Иван	Иванович
1220060605	АИ-22	cvb@gmail.com	Данила	Абрамов	Юрьевич
1220060604	АИ-22	dfg@gmail.com	Денис	Ульянкин	Сергеевич
1220060507	АС-22-1	vlad@gmail.com	Владимир	Коростелев	Александрович
1220060602	АИ-22	maks@gmail.com	Синицкий	Тимофей	Максимович
1220060601	АИ-22	ert	Свиридов	Александр	Сергеевич
1220060411	ПИ-22-1	teplovte@mail.ru	Владимир	Теплов	Сергеевич
1220060999	ПИ-22-1	asdad@mail.ru	ada	dad	ada
\.


--
-- TOC entry 4986 (class 0 OID 0)
-- Dependencies: 217
-- Name: achievement_id_achivment_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.achievement_id_achivment_seq', 25, true);


--
-- TOC entry 4987 (class 0 OID 0)
-- Dependencies: 219
-- Name: budget_id_budget_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.budget_id_budget_seq', 4, true);


--
-- TOC entry 4988 (class 0 OID 0)
-- Dependencies: 221
-- Name: category_id_category_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.category_id_category_seq', 4, true);


--
-- TOC entry 4989 (class 0 OID 0)
-- Dependencies: 224
-- Name: mark_id_mark_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mark_id_mark_seq', 29, true);


--
-- TOC entry 4990 (class 0 OID 0)
-- Dependencies: 226
-- Name: scholarship_id_scholarship_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scholarship_id_scholarship_seq', 22, true);


--
-- TOC entry 4804 (class 2606 OID 16872)
-- Name: student ak_email_student_student; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.student
    ADD CONSTRAINT ak_email_student_student UNIQUE (email_student);


--
-- TOC entry 4779 (class 2606 OID 16874)
-- Name: achievement pk_achievement; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.achievement
    ADD CONSTRAINT pk_achievement PRIMARY KEY (id_achivment);


--
-- TOC entry 4783 (class 2606 OID 16876)
-- Name: budget pk_budget; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.budget
    ADD CONSTRAINT pk_budget PRIMARY KEY (id_budget);


--
-- TOC entry 4786 (class 2606 OID 16878)
-- Name: category pk_category; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT pk_category PRIMARY KEY (id_category);


--
-- TOC entry 4789 (class 2606 OID 16880)
-- Name: group pk_group; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."group"
    ADD CONSTRAINT pk_group PRIMARY KEY (name_group);


--
-- TOC entry 4794 (class 2606 OID 16882)
-- Name: mark pk_mark; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mark
    ADD CONSTRAINT pk_mark PRIMARY KEY (id_mark);


--
-- TOC entry 4797 (class 2606 OID 16884)
-- Name: scholarship pk_scholarship; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scholarship
    ADD CONSTRAINT pk_scholarship PRIMARY KEY (id_scholarship);


--
-- TOC entry 4801 (class 2606 OID 16886)
-- Name: semester pk_semester; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.semester
    ADD CONSTRAINT pk_semester PRIMARY KEY (name_semester);


--
-- TOC entry 4806 (class 2606 OID 16888)
-- Name: student pk_student; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.student
    ADD CONSTRAINT pk_student PRIMARY KEY (id_num_student);


--
-- TOC entry 4776 (class 1259 OID 16889)
-- Name: achievement_pk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX achievement_pk ON public.achievement USING btree (id_achivment);


--
-- TOC entry 4795 (class 1259 OID 16890)
-- Name: assign_fk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX assign_fk ON public.scholarship USING btree (name_semester);


--
-- TOC entry 4781 (class 1259 OID 16891)
-- Name: budget_pk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX budget_pk ON public.budget USING btree (id_budget);


--
-- TOC entry 4784 (class 1259 OID 16892)
-- Name: category_pk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX category_pk ON public.category USING btree (id_category);


--
-- TOC entry 4790 (class 1259 OID 16893)
-- Name: get_fk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX get_fk ON public.mark USING btree (id_num_student);


--
-- TOC entry 4791 (class 1259 OID 16894)
-- Name: give_fk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX give_fk ON public.mark USING btree (name_semester);


--
-- TOC entry 4787 (class 1259 OID 16895)
-- Name: group_pk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX group_pk ON public."group" USING btree (name_group);


--
-- TOC entry 4777 (class 1259 OID 16896)
-- Name: have_fk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX have_fk ON public.achievement USING btree (id_num_student);


--
-- TOC entry 4792 (class 1259 OID 16897)
-- Name: mark_pk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX mark_pk ON public.mark USING btree (id_mark);


--
-- TOC entry 4798 (class 1259 OID 16898)
-- Name: receive_fk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX receive_fk ON public.scholarship USING btree (id_num_student);


--
-- TOC entry 4780 (class 1259 OID 16899)
-- Name: refer_fk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX refer_fk ON public.achievement USING btree (id_category);


--
-- TOC entry 4799 (class 1259 OID 16900)
-- Name: scholarship_pk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX scholarship_pk ON public.scholarship USING btree (id_scholarship);


--
-- TOC entry 4802 (class 1259 OID 16901)
-- Name: semester_pk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX semester_pk ON public.semester USING btree (name_semester);


--
-- TOC entry 4807 (class 1259 OID 16902)
-- Name: student_pk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX student_pk ON public.student USING btree (id_num_student);


--
-- TOC entry 4808 (class 1259 OID 16903)
-- Name: study_fk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX study_fk ON public.student USING btree (name_group);


--
-- TOC entry 4817 (class 2620 OID 16904)
-- Name: mark trg_check_mark_score; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trg_check_mark_score BEFORE INSERT OR UPDATE ON public.mark FOR EACH ROW EXECUTE FUNCTION public.check_mark_score();


--
-- TOC entry 4818 (class 2620 OID 16905)
-- Name: mark trigger_set_type_mark; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trigger_set_type_mark BEFORE INSERT ON public.mark FOR EACH ROW EXECUTE FUNCTION public.set_type_mark();


--
-- TOC entry 4809 (class 2606 OID 16906)
-- Name: achievement fk_achievem_have_student; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.achievement
    ADD CONSTRAINT fk_achievem_have_student FOREIGN KEY (id_num_student) REFERENCES public.student(id_num_student) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- TOC entry 4810 (class 2606 OID 16911)
-- Name: achievement fk_achievem_refer_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.achievement
    ADD CONSTRAINT fk_achievem_refer_category FOREIGN KEY (id_category) REFERENCES public.category(id_category) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 4811 (class 2606 OID 16916)
-- Name: mark fk_mark_get_student; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mark
    ADD CONSTRAINT fk_mark_get_student FOREIGN KEY (id_num_student) REFERENCES public.student(id_num_student) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- TOC entry 4812 (class 2606 OID 16921)
-- Name: mark fk_mark_give_semester; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mark
    ADD CONSTRAINT fk_mark_give_semester FOREIGN KEY (name_semester) REFERENCES public.semester(name_semester) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 4813 (class 2606 OID 16926)
-- Name: scholarship fk_scholars_assign_semester; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scholarship
    ADD CONSTRAINT fk_scholars_assign_semester FOREIGN KEY (name_semester) REFERENCES public.semester(name_semester) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 4814 (class 2606 OID 16931)
-- Name: scholarship fk_scholars_calculate_budget; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scholarship
    ADD CONSTRAINT fk_scholars_calculate_budget FOREIGN KEY (id_budget) REFERENCES public.budget(id_budget) ON UPDATE RESTRICT ON DELETE RESTRICT;


--
-- TOC entry 4815 (class 2606 OID 16936)
-- Name: scholarship fk_scholars_receive_student; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scholarship
    ADD CONSTRAINT fk_scholars_receive_student FOREIGN KEY (id_num_student) REFERENCES public.student(id_num_student) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- TOC entry 4816 (class 2606 OID 16941)
-- Name: student fk_student_study_group; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.student
    ADD CONSTRAINT fk_student_study_group FOREIGN KEY (name_group) REFERENCES public."group"(name_group) ON UPDATE RESTRICT ON DELETE RESTRICT;


-- Completed on 2024-12-21 15:51:13

--
-- PostgreSQL database dump complete
--

