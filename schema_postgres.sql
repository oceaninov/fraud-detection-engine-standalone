--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1 (Debian 16.1-1.pgdg120+1)
-- Dumped by pg_dump version 16.8 (Ubuntu 16.8-1.pgdg22.04+1)

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
-- Name: tsel_emoney_fds; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA tsel_emoney_fds;


ALTER SCHEMA tsel_emoney_fds OWNER TO pg_database_owner;

--
-- Name: SCHEMA tsel_emoney_fds; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA tsel_emoney_fds IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: black_list_dttot; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_dttot (
                                                  id character varying(36) DEFAULT NULL::character varying,
                                                  ppatk_id character varying(525) DEFAULT NULL::character varying,
                                                  name character varying(525) DEFAULT NULL::character varying,
                                                  bod character varying(10) DEFAULT NULL::character varying,
                                                  datasource character varying(525) DEFAULT NULL::character varying,
                                                  file_id character varying(525) DEFAULT NULL::character varying,
                                                  file_link text,
                                                  created_at timestamp without time zone,
                                                  created_by character varying(128) DEFAULT ''::character varying,
                                                  updated_at timestamp without time zone,
                                                  updated_by character varying(128) DEFAULT NULL::character varying,
                                                  approved_at timestamp without time zone,
                                                  approved_by character varying(128) DEFAULT NULL::character varying,
                                                  nik character varying(525) DEFAULT ''::character varying
);


ALTER TABLE tsel_emoney_fds.black_list_dttot OWNER TO fds_dev_user;

--
-- Name: black_list_dttot_approval; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_dttot_approval (
                                                           id character varying(36) DEFAULT NULL::character varying,
                                                           black_list_dttot_id character varying(36) DEFAULT NULL::character varying,
                                                           ppatk_id character varying(525) DEFAULT NULL::character varying,
                                                           name character varying(525) DEFAULT NULL::character varying,
                                                           bod character varying(10) DEFAULT NULL::character varying,
                                                           datasource character varying(525) DEFAULT NULL::character varying,
                                                           file_id character varying(525) DEFAULT NULL::character varying,
                                                           file_link text,
                                                           approval_type character varying(10) DEFAULT NULL::character varying,
                                                           note text,
                                                           created_at timestamp without time zone,
                                                           created_by character varying(128) DEFAULT ''::character varying,
                                                           updated_at timestamp without time zone,
                                                           updated_by character varying(128) DEFAULT NULL::character varying,
                                                           approved_at timestamp without time zone,
                                                           approved_by character varying(128) DEFAULT NULL::character varying,
                                                           nik character varying(525) DEFAULT ''::character varying
);


ALTER TABLE tsel_emoney_fds.black_list_dttot_approval OWNER TO fds_dev_user;

--
-- Name: black_list_dttot_file; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_dttot_file (
                                                       id character varying(36) DEFAULT NULL::character varying,
                                                       file_link text,
                                                       file_name character varying(125) DEFAULT NULL::character varying,
                                                       file_path text,
                                                       approval_type character varying(10) DEFAULT NULL::character varying,
                                                       note text,
                                                       created_at timestamp without time zone,
                                                       created_by character varying(128) DEFAULT ''::character varying,
                                                       updated_at timestamp without time zone,
                                                       updated_by character varying(128) DEFAULT NULL::character varying,
                                                       approved_at timestamp without time zone,
                                                       approved_by character varying(128) DEFAULT NULL::character varying,
                                                       active smallint DEFAULT 0,
                                                       status smallint DEFAULT 0
);


ALTER TABLE tsel_emoney_fds.black_list_dttot_file OWNER TO fds_dev_user;

--
-- Name: black_list_history; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_history (
                                                    id character varying(36) NOT NULL,
                                                    phone_number character varying(15) DEFAULT NULL::character varying,
                                                    event character varying(36) DEFAULT NULL::character varying,
                                                    beneficiary_name character varying(1024) DEFAULT ''::character varying NOT NULL,
                                                    created_at timestamp without time zone,
                                                    created_by character varying(128) DEFAULT ''::character varying,
                                                    updated_at timestamp without time zone,
                                                    updated_by character varying(128) DEFAULT NULL::character varying,
                                                    approved_at timestamp without time zone,
                                                    approved_by character varying(128) DEFAULT NULL::character varying,
                                                    transaction_types text
);


ALTER TABLE tsel_emoney_fds.black_list_history OWNER TO fds_dev_user;

--
-- Name: black_list_merchant; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_merchant (
                                                     id character varying(36) DEFAULT NULL::character varying,
                                                     nmid character varying(525) DEFAULT NULL::character varying,
                                                     merchant_name character varying(525) DEFAULT NULL::character varying,
                                                     datasource character varying(525) DEFAULT NULL::character varying,
                                                     file_id character varying(525) DEFAULT NULL::character varying,
                                                     file_link text,
                                                     created_at timestamp without time zone,
                                                     created_by character varying(128) DEFAULT ''::character varying,
                                                     updated_at timestamp without time zone,
                                                     updated_by character varying(128) DEFAULT NULL::character varying,
                                                     approved_at timestamp without time zone,
                                                     approved_by character varying(128) DEFAULT NULL::character varying
);


ALTER TABLE tsel_emoney_fds.black_list_merchant OWNER TO fds_dev_user;

--
-- Name: black_list_merchant_approval; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_merchant_approval (
                                                              id character varying(36) DEFAULT NULL::character varying,
                                                              black_list_merchant_id character varying(36) DEFAULT NULL::character varying,
                                                              nmid character varying(525) DEFAULT NULL::character varying,
                                                              merchant_name character varying(525) DEFAULT NULL::character varying,
                                                              datasource character varying(525) DEFAULT NULL::character varying,
                                                              file_id character varying(525) DEFAULT NULL::character varying,
                                                              file_link text,
                                                              approval_type character varying(10) DEFAULT NULL::character varying,
                                                              note text,
                                                              created_at timestamp without time zone,
                                                              created_by character varying(128) DEFAULT ''::character varying,
                                                              updated_at timestamp without time zone,
                                                              updated_by character varying(128) DEFAULT NULL::character varying,
                                                              approved_at timestamp without time zone,
                                                              approved_by character varying(128) DEFAULT NULL::character varying
);


ALTER TABLE tsel_emoney_fds.black_list_merchant_approval OWNER TO fds_dev_user;

--
-- Name: black_list_merchant_file; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_merchant_file (
                                                          id character varying(36) DEFAULT NULL::character varying,
                                                          file_link text,
                                                          file_name character varying(125) DEFAULT NULL::character varying,
                                                          file_path text,
                                                          approval_type character varying(10) DEFAULT NULL::character varying,
                                                          note text,
                                                          created_at timestamp without time zone,
                                                          created_by character varying(128) DEFAULT ''::character varying,
                                                          updated_at timestamp without time zone,
                                                          updated_by character varying(128) DEFAULT NULL::character varying,
                                                          approved_at timestamp without time zone,
                                                          approved_by character varying(128) DEFAULT NULL::character varying,
                                                          active smallint DEFAULT 0,
                                                          status smallint DEFAULT 0
);


ALTER TABLE tsel_emoney_fds.black_list_merchant_file OWNER TO fds_dev_user;

--
-- Name: black_list_receiver; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_receiver (
                                                     id character varying(36) NOT NULL,
                                                     phone_number character varying(15) DEFAULT NULL::character varying,
                                                     status smallint DEFAULT 1,
                                                     beneficiary_name character varying(1024) DEFAULT ''::character varying NOT NULL,
                                                     created_at timestamp without time zone,
                                                     created_by character varying(128) DEFAULT ''::character varying,
                                                     updated_at timestamp without time zone,
                                                     updated_by character varying(128) DEFAULT NULL::character varying,
                                                     approved_at timestamp without time zone,
                                                     approved_by character varying(128) DEFAULT NULL::character varying,
                                                     transaction_types text
);


ALTER TABLE tsel_emoney_fds.black_list_receiver OWNER TO fds_dev_user;

--
-- Name: black_list_receiver_approval; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_receiver_approval (
                                                              id character varying(36) NOT NULL,
                                                              phone_number character varying(15) NOT NULL,
                                                              approval_type character varying(10) DEFAULT NULL::character varying,
                                                              status smallint DEFAULT 0,
                                                              note text,
                                                              event character varying(36) DEFAULT NULL::character varying,
                                                              blacklist_id character varying(36) DEFAULT NULL::character varying,
                                                              beneficiary_name character varying(1024) DEFAULT ''::character varying NOT NULL,
                                                              created_at timestamp without time zone,
                                                              created_by character varying(128) DEFAULT ''::character varying,
                                                              updated_at timestamp without time zone,
                                                              updated_by character varying(128) DEFAULT NULL::character varying,
                                                              approved_at timestamp without time zone,
                                                              approved_by character varying(128) DEFAULT NULL::character varying,
                                                              transaction_types text
);


ALTER TABLE tsel_emoney_fds.black_list_receiver_approval OWNER TO fds_dev_user;

--
-- Name: black_list_sender; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_sender (
                                                   id character varying(36) NOT NULL,
                                                   phone_number character varying(15) DEFAULT NULL::character varying,
                                                   created_at timestamp without time zone,
                                                   created_by character varying(128) DEFAULT ''::character varying,
                                                   updated_at timestamp without time zone,
                                                   updated_by character varying(128) DEFAULT NULL::character varying,
                                                   approved_at timestamp without time zone,
                                                   approved_by character varying(128) DEFAULT NULL::character varying,
                                                   status smallint DEFAULT 1,
                                                   beneficiary_name character varying(1024) DEFAULT ''::character varying NOT NULL,
                                                   transaction_types text
);


ALTER TABLE tsel_emoney_fds.black_list_sender OWNER TO fds_dev_user;

--
-- Name: black_list_sender_approval; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.black_list_sender_approval (
                                                            id character varying(36) NOT NULL,
                                                            phone_number character varying(15) NOT NULL,
                                                            approval_type character varying(10) DEFAULT NULL::character varying,
                                                            status smallint DEFAULT 0,
                                                            note text,
                                                            event character varying(36) DEFAULT NULL::character varying,
                                                            blacklist_id character varying(36) DEFAULT NULL::character varying,
                                                            beneficiary_name character varying(1024) DEFAULT ''::character varying NOT NULL,
                                                            created_at timestamp without time zone,
                                                            created_by character varying(128) DEFAULT ''::character varying,
                                                            updated_at timestamp without time zone,
                                                            updated_by character varying(128) DEFAULT NULL::character varying,
                                                            approved_at timestamp without time zone,
                                                            approved_by character varying(128) DEFAULT NULL::character varying,
                                                            transaction_types text
);


ALTER TABLE tsel_emoney_fds.black_list_sender_approval OWNER TO fds_dev_user;

--
-- Name: flag; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.flag (
                                      id character varying(36) NOT NULL,
                                      title character varying(255) DEFAULT NULL::character varying,
                                      created_at timestamp without time zone,
                                      updated_at timestamp without time zone
);


ALTER TABLE tsel_emoney_fds.flag OWNER TO fds_dev_user;

--
-- Name: keyword; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.keyword (
                                         id character varying(36) DEFAULT NULL::character varying,
                                         keyword text,
                                         action character varying(36) DEFAULT NULL::character varying,
                                         channel character varying(255) DEFAULT NULL::character varying,
                                         created_at timestamp without time zone,
                                         created_by character varying(128) DEFAULT ''::character varying,
                                         updated_at timestamp without time zone,
                                         updated_by character varying(128) DEFAULT NULL::character varying,
                                         approved_at timestamp without time zone,
                                         approved_by character varying(128) DEFAULT NULL::character varying
);


ALTER TABLE tsel_emoney_fds.keyword OWNER TO fds_dev_user;

--
-- Name: keyword_approval; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.keyword_approval (
                                                  id character varying(36) DEFAULT NULL::character varying,
                                                  keyword_id character varying(36) DEFAULT NULL::character varying,
                                                  keyword text,
                                                  action character varying(36) DEFAULT NULL::character varying,
                                                  channel character varying(255) DEFAULT NULL::character varying,
                                                  note text,
                                                  status smallint,
                                                  approval_type character varying(10) DEFAULT NULL::character varying,
                                                  created_at timestamp without time zone,
                                                  created_by character varying(128) DEFAULT ''::character varying,
                                                  updated_at timestamp without time zone,
                                                  updated_by character varying(128) DEFAULT NULL::character varying,
                                                  approved_at timestamp without time zone,
                                                  approved_by character varying(128) DEFAULT NULL::character varying
);


ALTER TABLE tsel_emoney_fds.keyword_approval OWNER TO fds_dev_user;

--
-- Name: log; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.log (
                                     id character varying(36) NOT NULL,
                                     user_id character varying(128) DEFAULT NULL::character varying,
                                     amount double precision,
                                     start_date timestamp without time zone,
                                     body_req text,
                                     channel text,
                                     transaction_type text,
                                     end_date timestamp without time zone,
                                     destination_id character varying(128) DEFAULT NULL::character varying
);


ALTER TABLE tsel_emoney_fds.log OWNER TO fds_dev_user;

--
-- Name: registered_channel; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.registered_channel (
                                                    id character varying(36) NOT NULL,
                                                    channel_name character varying(255) NOT NULL,
                                                    channel_status character varying(50) NOT NULL,
                                                    created_at timestamp without time zone,
                                                    created_by character varying(128) DEFAULT ''::character varying,
                                                    updated_at timestamp without time zone,
                                                    updated_by character varying(128) DEFAULT NULL::character varying,
                                                    approved_at timestamp without time zone,
                                                    approved_by character varying(128) DEFAULT NULL::character varying
);


ALTER TABLE tsel_emoney_fds.registered_channel OWNER TO fds_dev_user;

--
-- Name: reset_password; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.reset_password (
                                                id bigint NOT NULL,
                                                user_id uuid,
                                                token character varying(50) DEFAULT NULL::character varying,
                                                expired_at timestamp without time zone,
                                                created_at timestamp without time zone
);


ALTER TABLE tsel_emoney_fds.reset_password OWNER TO fds_dev_user;

--
-- Name: reset_password_id_seq; Type: SEQUENCE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE SEQUENCE tsel_emoney_fds.reset_password_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE tsel_emoney_fds.reset_password_id_seq OWNER TO fds_dev_user;

--
-- Name: reset_password_id_seq; Type: SEQUENCE OWNED BY; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER SEQUENCE tsel_emoney_fds.reset_password_id_seq OWNED BY tsel_emoney_fds.reset_password.id;


--
-- Name: roles; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.roles (
                                       id bigint NOT NULL,
                                       title character varying(25) NOT NULL,
                                       created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                       updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                       description text
);


ALTER TABLE tsel_emoney_fds.roles OWNER TO fds_dev_user;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE SEQUENCE tsel_emoney_fds.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE tsel_emoney_fds.roles_id_seq OWNER TO fds_dev_user;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER SEQUENCE tsel_emoney_fds.roles_id_seq OWNED BY tsel_emoney_fds.roles.id;


--
-- Name: rules; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.rules (
                                       id character varying(36) DEFAULT ''::character varying NOT NULL,
                                       rule_name character varying(100) DEFAULT NULL::character varying,
                                       type character varying(100) DEFAULT NULL::character varying,
                                       transaction_type text,
                                       "interval" text,
                                       amount double precision,
                                       actions text,
                                       status smallint DEFAULT 1,
                                       channel character varying(255) DEFAULT NULL::character varying,
                                       time_range_type character varying(10) DEFAULT 'NONE'::character varying,
                                       start_time_range character varying(19) DEFAULT 'NONE'::character varying,
                                       end_time_range character varying(19) DEFAULT 'NONE'::character varying,
                                       created_at timestamp without time zone,
                                       created_by character varying(128) DEFAULT ''::character varying,
                                       updated_at timestamp without time zone,
                                       updated_by character varying(128) DEFAULT NULL::character varying,
                                       approved_at timestamp without time zone,
                                       approved_by character varying(128) DEFAULT NULL::character varying,
                                       sofs text
);


ALTER TABLE tsel_emoney_fds.rules OWNER TO fds_dev_user;

--
-- Name: rules_approval; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.rules_approval (
                                                id character varying(36) NOT NULL,
                                                rule_id character varying(36) DEFAULT NULL::character varying,
                                                rule_name character varying(100) DEFAULT NULL::character varying,
                                                type character varying(100) DEFAULT NULL::character varying,
                                                transaction_type text,
                                                "interval" text,
                                                amount double precision,
                                                actions text,
                                                status smallint,
                                                approval_type character varying(10) DEFAULT NULL::character varying,
                                                note text,
                                                channel character varying(255) DEFAULT NULL::character varying,
                                                time_range_type character varying(10) DEFAULT 'NONE'::character varying NOT NULL,
                                                start_time_range character varying(19) DEFAULT 'NONE'::character varying NOT NULL,
                                                end_time_range character varying(19) DEFAULT 'NONE'::character varying NOT NULL,
                                                created_at timestamp without time zone,
                                                created_by character varying(128) DEFAULT ''::character varying,
                                                updated_at timestamp without time zone,
                                                updated_by character varying(128) DEFAULT NULL::character varying,
                                                approved_at timestamp without time zone,
                                                approved_by character varying(128) DEFAULT NULL::character varying,
                                                sofs text
);


ALTER TABLE tsel_emoney_fds.rules_approval OWNER TO fds_dev_user;

--
-- Name: sofs; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.sofs (
                                      id character varying(36) NOT NULL,
                                      sof_name character varying(255) NOT NULL,
                                      sof_status character varying(50) NOT NULL,
                                      created_at timestamp without time zone,
                                      created_by character varying(128) DEFAULT ''::character varying,
                                      updated_at timestamp without time zone,
                                      updated_by character varying(128) DEFAULT NULL::character varying,
                                      approved_at timestamp without time zone,
                                      approved_by character varying(128) DEFAULT NULL::character varying
);


ALTER TABLE tsel_emoney_fds.sofs OWNER TO fds_dev_user;

--
-- Name: transaction; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.transaction (
                                             id character varying(36) NOT NULL,
                                             transaction_id character varying(128) DEFAULT NULL::character varying,
                                             transaction_type character varying(128) DEFAULT NULL::character varying,
                                             rules text,
                                             title character varying(255) DEFAULT NULL::character varying,
                                             channel character varying(255) DEFAULT NULL::character varying,
                                             body_req text,
                                             flag_id character varying(36) DEFAULT NULL::character varying,
                                             created_at timestamp without time zone,
                                             user_id character varying(128) DEFAULT NULL::character varying,
                                             amount character varying(100) DEFAULT NULL::character varying,
                                             destination_id character varying(128) DEFAULT ''::character varying
);


ALTER TABLE tsel_emoney_fds.transaction OWNER TO fds_dev_user;

--
-- Name: transaction_type; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.transaction_type (
                                                  id character varying(36) NOT NULL,
                                                  name character varying(128) DEFAULT NULL::character varying,
                                                  created_at timestamp without time zone,
                                                  created_by character varying(128) DEFAULT ''::character varying,
                                                  updated_at timestamp without time zone,
                                                  updated_by character varying(128) DEFAULT NULL::character varying,
                                                  approved_at timestamp without time zone,
                                                  approved_by character varying(128) DEFAULT NULL::character varying
);


ALTER TABLE tsel_emoney_fds.transaction_type OWNER TO fds_dev_user;

--
-- Name: transaction_type_approval; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.transaction_type_approval (
                                                           id character varying(36) NOT NULL,
                                                           transaction_type_id character varying(36) NOT NULL,
                                                           name character varying(128) DEFAULT NULL::character varying,
                                                           note text,
                                                           status smallint DEFAULT 0,
                                                           created_at timestamp without time zone,
                                                           approval_type character varying(10) DEFAULT NULL::character varying,
                                                           created_by character varying(128) DEFAULT ''::character varying,
                                                           updated_at timestamp without time zone,
                                                           updated_by character varying(128) DEFAULT NULL::character varying,
                                                           approved_at timestamp without time zone,
                                                           approved_by character varying(128) DEFAULT NULL::character varying
);


ALTER TABLE tsel_emoney_fds.transaction_type_approval OWNER TO fds_dev_user;

--
-- Name: users; Type: TABLE; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

CREATE TABLE tsel_emoney_fds.users (
                                       id character varying(36) NOT NULL,
                                       role_id bigint DEFAULT 3 NOT NULL,
                                       email character varying(255) DEFAULT NULL::character varying,
                                       password text,
                                       full_name character varying(255) DEFAULT 'no name'::character varying,
                                       avatar_url text,
                                       gender character varying(1) DEFAULT 'U'::character varying,
                                       activated smallint DEFAULT 0,
                                       created_at timestamp without time zone,
                                       created_by character varying(128) DEFAULT ''::character varying,
                                       updated_at timestamp without time zone,
                                       updated_by character varying(128) DEFAULT NULL::character varying,
                                       approved_at timestamp without time zone,
                                       approved_by character varying(128) DEFAULT NULL::character varying
);


ALTER TABLE tsel_emoney_fds.users OWNER TO fds_dev_user;

--
-- Name: reset_password id; Type: DEFAULT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.reset_password ALTER COLUMN id SET DEFAULT nextval('tsel_emoney_fds.reset_password_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.roles ALTER COLUMN id SET DEFAULT nextval('tsel_emoney_fds.roles_id_seq'::regclass);

--
-- Name: reset_password_id_seq; Type: SEQUENCE SET; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

SELECT pg_catalog.setval('tsel_emoney_fds.reset_password_id_seq', 11, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

SELECT pg_catalog.setval('tsel_emoney_fds.roles_id_seq', 1, false);


--
-- Name: black_list_history black_list_history_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.black_list_history
    ADD CONSTRAINT black_list_history_pkey PRIMARY KEY (id);


--
-- Name: black_list_receiver_approval black_list_receiver_approval_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.black_list_receiver_approval
    ADD CONSTRAINT black_list_receiver_approval_pkey PRIMARY KEY (id);


--
-- Name: black_list_receiver black_list_receiver_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.black_list_receiver
    ADD CONSTRAINT black_list_receiver_pkey PRIMARY KEY (id);


--
-- Name: black_list_sender_approval black_list_sender_approval_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.black_list_sender_approval
    ADD CONSTRAINT black_list_sender_approval_pkey PRIMARY KEY (id);


--
-- Name: black_list_sender black_list_sender_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.black_list_sender
    ADD CONSTRAINT black_list_sender_pkey PRIMARY KEY (id);


--
-- Name: flag flag_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.flag
    ADD CONSTRAINT flag_pkey PRIMARY KEY (id);


--
-- Name: log log_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.log
    ADD CONSTRAINT log_pkey PRIMARY KEY (id);


--
-- Name: registered_channel registered_channel_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.registered_channel
    ADD CONSTRAINT registered_channel_pkey PRIMARY KEY (id);


--
-- Name: reset_password reset_password_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.reset_password
    ADD CONSTRAINT reset_password_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: rules_approval rules_approval_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.rules_approval
    ADD CONSTRAINT rules_approval_pkey PRIMARY KEY (id);


--
-- Name: rules rules_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.rules
    ADD CONSTRAINT rules_pkey PRIMARY KEY (id);


--
-- Name: sofs sofs_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.sofs
    ADD CONSTRAINT sofs_pkey PRIMARY KEY (id);


--
-- Name: transaction transaction_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);


--
-- Name: transaction_type_approval transaction_type_approval_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.transaction_type_approval
    ADD CONSTRAINT transaction_type_approval_pkey PRIMARY KEY (id);


--
-- Name: transaction_type transaction_type_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.transaction_type
    ADD CONSTRAINT transaction_type_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users fk_roleuser; Type: FK CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.users
    ADD CONSTRAINT fk_roleuser FOREIGN KEY (role_id) REFERENCES tsel_emoney_fds.roles(id);


--
-- Name: transaction transaction_ibfk_1; Type: FK CONSTRAINT; Schema: tsel_emoney_fds; Owner: fds_dev_user
--

ALTER TABLE ONLY tsel_emoney_fds.transaction
    ADD CONSTRAINT transaction_ibfk_1 FOREIGN KEY (flag_id) REFERENCES tsel_emoney_fds.flag(id);


--
-- Name: SCHEMA tsel_emoney_fds; Type: ACL; Schema: -; Owner: pg_database_owner
--

GRANT ALL ON SCHEMA tsel_emoney_fds TO fds_dev_user;
GRANT ALL ON SCHEMA tsel_emoney_fds TO fds_app_user;


--
-- PostgreSQL database dump complete
--
