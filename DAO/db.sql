CREATE TABLE users (
    id integer NOT NULL,
    login character varying(255),
    password character varying(255)
);


ALTER TABLE users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Name: users_login_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_login_key UNIQUE (login);


--
-- Name: insertdata; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER insertdata AFTER INSERT ON users FOR EACH ROW EXECUTE PROCEDURE createuserdata();

CREATE TABLE usersdata (
    userid integer NOT NULL,
    rating integer,
    countgames integer
);


ALTER TABLE usersdata OWNER TO postgres;

--
-- Name: usersdata_userid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY usersdata
    ADD CONSTRAINT usersdata_userid_key UNIQUE (userid);



CREATE TABLE sessions (
    userid integer,
    sessionid character varying(255) NOT NULL,
    createdate timestamp without time zone DEFAULT now()
);


ALTER TABLE sessions OWNER TO postgres;

--
-- Name: sessions_sessionid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY sessions
    ADD CONSTRAINT sessions_sessionid_key UNIQUE (sessionid);


--
-- Name: sessions_userid_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY sessions
    ADD CONSTRAINT sessions_userid_key UNIQUE (userid);


