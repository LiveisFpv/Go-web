PGDMP      .                |            postgres    16.4    16.4 P    q           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            r           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            s           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            t           1262    5    postgres    DATABASE     |   CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';
    DROP DATABASE postgres;
                postgres    false            u           0    0    DATABASE postgres    COMMENT     N   COMMENT ON DATABASE postgres IS 'default administrative connection database';
                   postgres    false    4980                        3079    16384 	   adminpack 	   EXTENSION     A   CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;
    DROP EXTENSION adminpack;
                   false            v           0    0    EXTENSION adminpack    COMMENT     M   COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';
                        false    2            �            1255    16822 9   assign_scholarships(character varying, character varying) 	   PROCEDURE     �  CREATE PROCEDURE public.assign_scholarships(IN current_semester character varying, IN budget_type character varying)
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
 t   DROP PROCEDURE public.assign_scholarships(IN current_semester character varying, IN budget_type character varying);
       public          postgres    false            �            1255    16818    check_mark_score()    FUNCTION     ]  CREATE FUNCTION public.check_mark_score() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.score_mark < 0 OR NEW.score_mark > 100 THEN
        RAISE EXCEPTION 'Оценка должна быть в диапазоне от 0 до 100. Некорректное значение: %', NEW.score_mark;
    END IF;
    RETURN NEW;
END;
$$;
 )   DROP FUNCTION public.check_mark_score();
       public          postgres    false            �            1255    16820    set_type_mark()    FUNCTION     �  CREATE FUNCTION public.set_type_mark() RETURNS trigger
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
 &   DROP FUNCTION public.set_type_mark();
       public          postgres    false            �            1259    16399    achievement    TABLE     �   CREATE TABLE public.achievement (
    id_achivment integer NOT NULL,
    id_num_student integer NOT NULL,
    id_category integer NOT NULL,
    name_achivement character varying(200) NOT NULL,
    date_achivment date NOT NULL
);
    DROP TABLE public.achievement;
       public         heap    postgres    false            �            1259    16398    achievement_id_achivment_seq    SEQUENCE     �   CREATE SEQUENCE public.achievement_id_achivment_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 3   DROP SEQUENCE public.achievement_id_achivment_seq;
       public          postgres    false    217            w           0    0    achievement_id_achivment_seq    SEQUENCE OWNED BY     ]   ALTER SEQUENCE public.achievement_id_achivment_seq OWNED BY public.achievement.id_achivment;
          public          postgres    false    216            �            1259    16527    budget    TABLE     �   CREATE TABLE public.budget (
    type_scholarship_budget character varying(100) NOT NULL,
    name_semester character varying(5) NOT NULL,
    size_budget money NOT NULL,
    id_budget integer NOT NULL
);
    DROP TABLE public.budget;
       public         heap    postgres    false            �            1259    16526    budget_id_budget_seq    SEQUENCE     �   CREATE SEQUENCE public.budget_id_budget_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.budget_id_budget_seq;
       public          postgres    false    228            x           0    0    budget_id_budget_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.budget_id_budget_seq OWNED BY public.budget.id_budget;
          public          postgres    false    227            �            1259    16416    category    TABLE     �   CREATE TABLE public.category (
    id_category integer NOT NULL,
    achivments_type_category character varying(50) NOT NULL,
    score_category smallint NOT NULL
);
    DROP TABLE public.category;
       public         heap    postgres    false            �            1259    16415    category_id_category_seq    SEQUENCE     �   CREATE SEQUENCE public.category_id_category_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.category_id_category_seq;
       public          postgres    false    219            y           0    0    category_id_category_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.category_id_category_seq OWNED BY public.category.id_category;
          public          postgres    false    218            �            1259    16423    group    TABLE       CREATE TABLE public."group" (
    name_group character varying(10) NOT NULL,
    studies_direction_group character varying(8) NOT NULL,
    studies_profile_group character varying(50) NOT NULL,
    start_date_group date NOT NULL,
    "studies_period-group" smallint NOT NULL
);
    DROP TABLE public."group";
       public         heap    postgres    false            �            1259    16456    mark    TABLE       CREATE TABLE public.mark (
    id_mark integer NOT NULL,
    id_num_student integer NOT NULL,
    name_semester character varying(5) NOT NULL,
    lesson_name_mark character varying(100) NOT NULL,
    score_mark smallint NOT NULL,
    type_mark character varying(20)
);
    DROP TABLE public.mark;
       public         heap    postgres    false            �            1259    16455    mark_id_mark_seq    SEQUENCE     �   CREATE SEQUENCE public.mark_id_mark_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.mark_id_mark_seq;
       public          postgres    false    226            z           0    0    mark_id_mark_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.mark_id_mark_seq OWNED BY public.mark.id_mark;
          public          postgres    false    225            �            1259    16430    scholarship    TABLE     �   CREATE TABLE public.scholarship (
    id_scholarship integer NOT NULL,
    id_num_student integer NOT NULL,
    name_semester character varying(5) NOT NULL,
    size_scholarshp money NOT NULL,
    id_budget integer NOT NULL
);
    DROP TABLE public.scholarship;
       public         heap    postgres    false            �            1259    16429    scholarship_id_scholarship_seq    SEQUENCE     �   CREATE SEQUENCE public.scholarship_id_scholarship_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 5   DROP SEQUENCE public.scholarship_id_scholarship_seq;
       public          postgres    false    222            {           0    0    scholarship_id_scholarship_seq    SEQUENCE OWNED BY     a   ALTER SEQUENCE public.scholarship_id_scholarship_seq OWNED BY public.scholarship.id_scholarship;
          public          postgres    false    221            �            1259    16440    semester    TABLE     �   CREATE TABLE public.semester (
    name_semester character varying(5) NOT NULL,
    date_start_semester date NOT NULL,
    date_end_semester date NOT NULL
);
    DROP TABLE public.semester;
       public         heap    postgres    false            �            1259    16446    student    TABLE     @  CREATE TABLE public.student (
    id_num_student integer NOT NULL,
    name_group character varying(10) NOT NULL,
    email_student character varying(40) NOT NULL,
    second_name_student character varying(40) NOT NULL,
    first_name_student character varying(40) NOT NULL,
    surname_student character varying(40)
);
    DROP TABLE public.student;
       public         heap    postgres    false            �           2604    16402    achievement id_achivment    DEFAULT     �   ALTER TABLE ONLY public.achievement ALTER COLUMN id_achivment SET DEFAULT nextval('public.achievement_id_achivment_seq'::regclass);
 G   ALTER TABLE public.achievement ALTER COLUMN id_achivment DROP DEFAULT;
       public          postgres    false    216    217    217            �           2604    16530    budget id_budget    DEFAULT     t   ALTER TABLE ONLY public.budget ALTER COLUMN id_budget SET DEFAULT nextval('public.budget_id_budget_seq'::regclass);
 ?   ALTER TABLE public.budget ALTER COLUMN id_budget DROP DEFAULT;
       public          postgres    false    227    228    228            �           2604    16419    category id_category    DEFAULT     |   ALTER TABLE ONLY public.category ALTER COLUMN id_category SET DEFAULT nextval('public.category_id_category_seq'::regclass);
 C   ALTER TABLE public.category ALTER COLUMN id_category DROP DEFAULT;
       public          postgres    false    218    219    219            �           2604    16459    mark id_mark    DEFAULT     l   ALTER TABLE ONLY public.mark ALTER COLUMN id_mark SET DEFAULT nextval('public.mark_id_mark_seq'::regclass);
 ;   ALTER TABLE public.mark ALTER COLUMN id_mark DROP DEFAULT;
       public          postgres    false    225    226    226            �           2604    16433    scholarship id_scholarship    DEFAULT     �   ALTER TABLE ONLY public.scholarship ALTER COLUMN id_scholarship SET DEFAULT nextval('public.scholarship_id_scholarship_seq'::regclass);
 I   ALTER TABLE public.scholarship ALTER COLUMN id_scholarship DROP DEFAULT;
       public          postgres    false    222    221    222            c          0    16399    achievement 
   TABLE DATA           q   COPY public.achievement (id_achivment, id_num_student, id_category, name_achivement, date_achivment) FROM stdin;
    public          postgres    false    217   �d       n          0    16527    budget 
   TABLE DATA           `   COPY public.budget (type_scholarship_budget, name_semester, size_budget, id_budget) FROM stdin;
    public          postgres    false    228   �e       e          0    16416    category 
   TABLE DATA           Y   COPY public.category (id_category, achivments_type_category, score_category) FROM stdin;
    public          postgres    false    219   0f       f          0    16423    group 
   TABLE DATA           �   COPY public."group" (name_group, studies_direction_group, studies_profile_group, start_date_group, "studies_period-group") FROM stdin;
    public          postgres    false    220   �f       l          0    16456    mark 
   TABLE DATA           o   COPY public.mark (id_mark, id_num_student, name_semester, lesson_name_mark, score_mark, type_mark) FROM stdin;
    public          postgres    false    226   �g       h          0    16430    scholarship 
   TABLE DATA           p   COPY public.scholarship (id_scholarship, id_num_student, name_semester, size_scholarshp, id_budget) FROM stdin;
    public          postgres    false    222   �h       i          0    16440    semester 
   TABLE DATA           Y   COPY public.semester (name_semester, date_start_semester, date_end_semester) FROM stdin;
    public          postgres    false    223   Yi       j          0    16446    student 
   TABLE DATA           �   COPY public.student (id_num_student, name_group, email_student, second_name_student, first_name_student, surname_student) FROM stdin;
    public          postgres    false    224   �i       |           0    0    achievement_id_achivment_seq    SEQUENCE SET     K   SELECT pg_catalog.setval('public.achievement_id_achivment_seq', 25, true);
          public          postgres    false    216            }           0    0    budget_id_budget_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.budget_id_budget_seq', 4, true);
          public          postgres    false    227            ~           0    0    category_id_category_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.category_id_category_seq', 4, true);
          public          postgres    false    218                       0    0    mark_id_mark_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.mark_id_mark_seq', 28, true);
          public          postgres    false    225            �           0    0    scholarship_id_scholarship_seq    SEQUENCE SET     M   SELECT pg_catalog.setval('public.scholarship_id_scholarship_seq', 22, true);
          public          postgres    false    221            �           2606    16452     student ak_email_student_student 
   CONSTRAINT     d   ALTER TABLE ONLY public.student
    ADD CONSTRAINT ak_email_student_student UNIQUE (email_student);
 J   ALTER TABLE ONLY public.student DROP CONSTRAINT ak_email_student_student;
       public            postgres    false    224            �           2606    16404    achievement pk_achievement 
   CONSTRAINT     b   ALTER TABLE ONLY public.achievement
    ADD CONSTRAINT pk_achievement PRIMARY KEY (id_achivment);
 D   ALTER TABLE ONLY public.achievement DROP CONSTRAINT pk_achievement;
       public            postgres    false    217            �           2606    16532    budget pk_budget 
   CONSTRAINT     U   ALTER TABLE ONLY public.budget
    ADD CONSTRAINT pk_budget PRIMARY KEY (id_budget);
 :   ALTER TABLE ONLY public.budget DROP CONSTRAINT pk_budget;
       public            postgres    false    228            �           2606    16421    category pk_category 
   CONSTRAINT     [   ALTER TABLE ONLY public.category
    ADD CONSTRAINT pk_category PRIMARY KEY (id_category);
 >   ALTER TABLE ONLY public.category DROP CONSTRAINT pk_category;
       public            postgres    false    219            �           2606    16427    group pk_group 
   CONSTRAINT     V   ALTER TABLE ONLY public."group"
    ADD CONSTRAINT pk_group PRIMARY KEY (name_group);
 :   ALTER TABLE ONLY public."group" DROP CONSTRAINT pk_group;
       public            postgres    false    220            �           2606    16461    mark pk_mark 
   CONSTRAINT     O   ALTER TABLE ONLY public.mark
    ADD CONSTRAINT pk_mark PRIMARY KEY (id_mark);
 6   ALTER TABLE ONLY public.mark DROP CONSTRAINT pk_mark;
       public            postgres    false    226            �           2606    16435    scholarship pk_scholarship 
   CONSTRAINT     d   ALTER TABLE ONLY public.scholarship
    ADD CONSTRAINT pk_scholarship PRIMARY KEY (id_scholarship);
 D   ALTER TABLE ONLY public.scholarship DROP CONSTRAINT pk_scholarship;
       public            postgres    false    222            �           2606    16444    semester pk_semester 
   CONSTRAINT     ]   ALTER TABLE ONLY public.semester
    ADD CONSTRAINT pk_semester PRIMARY KEY (name_semester);
 >   ALTER TABLE ONLY public.semester DROP CONSTRAINT pk_semester;
       public            postgres    false    223            �           2606    16450    student pk_student 
   CONSTRAINT     \   ALTER TABLE ONLY public.student
    ADD CONSTRAINT pk_student PRIMARY KEY (id_num_student);
 <   ALTER TABLE ONLY public.student DROP CONSTRAINT pk_student;
       public            postgres    false    224            �           1259    16405    achievement_pk    INDEX     U   CREATE UNIQUE INDEX achievement_pk ON public.achievement USING btree (id_achivment);
 "   DROP INDEX public.achievement_pk;
       public            postgres    false    217            �           1259    16439 	   assign_fk    INDEX     J   CREATE INDEX assign_fk ON public.scholarship USING btree (name_semester);
    DROP INDEX public.assign_fk;
       public            postgres    false    222            �           1259    16533 	   budget_pk    INDEX     H   CREATE UNIQUE INDEX budget_pk ON public.budget USING btree (id_budget);
    DROP INDEX public.budget_pk;
       public            postgres    false    228            �           1259    16422    category_pk    INDEX     N   CREATE UNIQUE INDEX category_pk ON public.category USING btree (id_category);
    DROP INDEX public.category_pk;
       public            postgres    false    219            �           1259    16463    get_fk    INDEX     A   CREATE INDEX get_fk ON public.mark USING btree (id_num_student);
    DROP INDEX public.get_fk;
       public            postgres    false    226            �           1259    16464    give_fk    INDEX     A   CREATE INDEX give_fk ON public.mark USING btree (name_semester);
    DROP INDEX public.give_fk;
       public            postgres    false    226            �           1259    16428    group_pk    INDEX     I   CREATE UNIQUE INDEX group_pk ON public."group" USING btree (name_group);
    DROP INDEX public.group_pk;
       public            postgres    false    220            �           1259    16406    have_fk    INDEX     I   CREATE INDEX have_fk ON public.achievement USING btree (id_num_student);
    DROP INDEX public.have_fk;
       public            postgres    false    217            �           1259    16462    mark_pk    INDEX     B   CREATE UNIQUE INDEX mark_pk ON public.mark USING btree (id_mark);
    DROP INDEX public.mark_pk;
       public            postgres    false    226            �           1259    16437 
   receive_fk    INDEX     L   CREATE INDEX receive_fk ON public.scholarship USING btree (id_num_student);
    DROP INDEX public.receive_fk;
       public            postgres    false    222            �           1259    16407    refer_fk    INDEX     G   CREATE INDEX refer_fk ON public.achievement USING btree (id_category);
    DROP INDEX public.refer_fk;
       public            postgres    false    217            �           1259    16436    scholarship_pk    INDEX     W   CREATE UNIQUE INDEX scholarship_pk ON public.scholarship USING btree (id_scholarship);
 "   DROP INDEX public.scholarship_pk;
       public            postgres    false    222            �           1259    16445    semester_pk    INDEX     P   CREATE UNIQUE INDEX semester_pk ON public.semester USING btree (name_semester);
    DROP INDEX public.semester_pk;
       public            postgres    false    223            �           1259    16453 
   student_pk    INDEX     O   CREATE UNIQUE INDEX student_pk ON public.student USING btree (id_num_student);
    DROP INDEX public.student_pk;
       public            postgres    false    224            �           1259    16454    study_fk    INDEX     B   CREATE INDEX study_fk ON public.student USING btree (name_group);
    DROP INDEX public.study_fk;
       public            postgres    false    224            �           2620    16819    mark trg_check_mark_score    TRIGGER     �   CREATE TRIGGER trg_check_mark_score BEFORE INSERT OR UPDATE ON public.mark FOR EACH ROW EXECUTE FUNCTION public.check_mark_score();
 2   DROP TRIGGER trg_check_mark_score ON public.mark;
       public          postgres    false    226    229            �           2620    16821    mark trigger_set_type_mark    TRIGGER     x   CREATE TRIGGER trigger_set_type_mark BEFORE INSERT ON public.mark FOR EACH ROW EXECUTE FUNCTION public.set_type_mark();
 3   DROP TRIGGER trigger_set_type_mark ON public.mark;
       public          postgres    false    230    226            �           2606    16466 $   achievement fk_achievem_have_student    FK CONSTRAINT     �   ALTER TABLE ONLY public.achievement
    ADD CONSTRAINT fk_achievem_have_student FOREIGN KEY (id_num_student) REFERENCES public.student(id_num_student) ON UPDATE RESTRICT ON DELETE CASCADE;
 N   ALTER TABLE ONLY public.achievement DROP CONSTRAINT fk_achievem_have_student;
       public          postgres    false    4798    217    224            �           2606    16471 &   achievement fk_achievem_refer_category    FK CONSTRAINT     �   ALTER TABLE ONLY public.achievement
    ADD CONSTRAINT fk_achievem_refer_category FOREIGN KEY (id_category) REFERENCES public.category(id_category) ON UPDATE RESTRICT ON DELETE RESTRICT;
 P   ALTER TABLE ONLY public.achievement DROP CONSTRAINT fk_achievem_refer_category;
       public          postgres    false    4783    217    219            �           2606    16659    mark fk_mark_get_student    FK CONSTRAINT     �   ALTER TABLE ONLY public.mark
    ADD CONSTRAINT fk_mark_get_student FOREIGN KEY (id_num_student) REFERENCES public.student(id_num_student) ON UPDATE RESTRICT ON DELETE CASCADE;
 B   ALTER TABLE ONLY public.mark DROP CONSTRAINT fk_mark_get_student;
       public          postgres    false    226    224    4798            �           2606    16506    mark fk_mark_give_semester    FK CONSTRAINT     �   ALTER TABLE ONLY public.mark
    ADD CONSTRAINT fk_mark_give_semester FOREIGN KEY (name_semester) REFERENCES public.semester(name_semester) ON UPDATE RESTRICT ON DELETE RESTRICT;
 D   ALTER TABLE ONLY public.mark DROP CONSTRAINT fk_mark_give_semester;
       public          postgres    false    4793    226    223            �           2606    16491 '   scholarship fk_scholars_assign_semester    FK CONSTRAINT     �   ALTER TABLE ONLY public.scholarship
    ADD CONSTRAINT fk_scholars_assign_semester FOREIGN KEY (name_semester) REFERENCES public.semester(name_semester) ON UPDATE RESTRICT ON DELETE RESTRICT;
 Q   ALTER TABLE ONLY public.scholarship DROP CONSTRAINT fk_scholars_assign_semester;
       public          postgres    false    222    223    4793            �           2606    16534 (   scholarship fk_scholars_calculate_budget    FK CONSTRAINT     �   ALTER TABLE ONLY public.scholarship
    ADD CONSTRAINT fk_scholars_calculate_budget FOREIGN KEY (id_budget) REFERENCES public.budget(id_budget) ON UPDATE RESTRICT ON DELETE RESTRICT;
 R   ALTER TABLE ONLY public.scholarship DROP CONSTRAINT fk_scholars_calculate_budget;
       public          postgres    false    4808    222    228            �           2606    16486 '   scholarship fk_scholars_receive_student    FK CONSTRAINT     �   ALTER TABLE ONLY public.scholarship
    ADD CONSTRAINT fk_scholars_receive_student FOREIGN KEY (id_num_student) REFERENCES public.student(id_num_student) ON UPDATE RESTRICT ON DELETE CASCADE;
 Q   ALTER TABLE ONLY public.scholarship DROP CONSTRAINT fk_scholars_receive_student;
       public          postgres    false    4798    222    224            �           2606    16496    student fk_student_study_group    FK CONSTRAINT     �   ALTER TABLE ONLY public.student
    ADD CONSTRAINT fk_student_study_group FOREIGN KEY (name_group) REFERENCES public."group"(name_group) ON UPDATE RESTRICT ON DELETE RESTRICT;
 H   ALTER TABLE ONLY public.student DROP CONSTRAINT fk_student_study_group;
       public          postgres    false    224    220    4786            c     x����M�0�s�K��8m�� ��r�+��'~�
�F8����@��9��#f���N�UO���]�Y��~�0c�踇q_����}�2���7$�?^?ف�`�m��L�_z�{E���xn�|��(9(�`B ��?֚r��;��n�i��E�t��t�U�(�IU��hx�z�Bv+j�ǜ��C����c]��]�lڸ����c^�Es?H����3վ��:k�\{:lt�Zm�:�r�Yj����9�Zz�:�� �Ok�)      n   >   x��0���.,�4202�0������{NC.d�yh�F\��6ƫۄ+F��� L;,9      e   g   x���	�0Dϻň� c1���W� ��L���e���V��2��.="���@��0�CH\������k¦��8&���f��-aQ+?f'a�]��?�\H      f     x���Mn�0F��)r�D3��� *�J�d��i+DT����F�@[]Ů"9�5~o��XW����q-M-#�7�غ	���hK��p�W���wp7/���2f���B��+Q3*pk�M�����2��*c���2��*c���2�^�����`�]�DN�����J|���Q�q���{��Q��Oݜ�><轋E��b��b#K�ZB�f�E#K�ZB�d��u�8�����qAhh��\�[���ZF��ߵ�@ u(P����.���G��      l   �   x���Kj�0�t����Ɩd_�'�aJ�P�*Yw��M�S��׍*��~��[��R���QMV1q���=a��]�9�V�M5�Z���7�۫��ws�'wɷ�z漳�*	�k��|�Gxň�x�q�� �ߒ�W>/Əd���$�,,]7�������-���>��u�/z��3���~�i����	3��Vv(W�z��>�h�K���vRόwٻR��ܹ���}Fs8;V���1�ų���G��      h   {   x�m���@��] 4��ws��TE(��J�![��?����@G�	��zH�� ���.�ƃ8XH+�%�B����8�bR��^�H�%ŒZ�D!��H2�?�Nc������t�Gz��zS�7��P�      i   A   x�3202�0��H� �!�i�kl�b^���Ji��1D�1B�1B�1D�1B�	Tg� =�      j   �  x���Mo�@���ie�N��.%_E�B�ZNI?ȡM�HQDiܐj��a���ڱgc"�C��γ�̼�븮mo㓷�E�׵
o��V�^��,�:��;U'��4��E_�sF#��o�OUW�ЀF��s,/f�ɺ���TW ���1-`�jXԧ ���B�������4���@��m�`�,K��J4����=Xt�/��� �߯��U#\\�:�Z3������Pe�4�+y��+r��O8��z��3�r	g�B�s=�?��=d��;o�~��w�:"TC���$��ĝǡ#}R}�3�����@�\7n�?��v�r,�mu�� q�Q�LO0˕�]�����Z�"&��F��� t���-ɩI�������,�����vo>t�"_���x
&��!gf�8#w��K��[?����wA+�~��9.ר2�	��P6R\�oZ4�2Q�N1�<��}\�(�ֿv�slA|/U�@M�W�aֲ|�|��O%��5\8�:�] xBaOkD1>{�;�F�j&w5��0w�ޔ���t����d��?g������M>�/M-/�i����U��+]O��M�H�Uh���7aAOi�Ԥ5��4��Tu^��
é�J[}�x�c(-�T�~k=:�n�@Lh�NM]��Z鰖!z��C��A�6s����Pj     