--
-- TOC entry 211 (class 1259 OID 32951)
-- Name: audit_trails; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.audit_trails (
    id serial,
    action character varying(45),
    url character varying(191),
    client character varying(255),
    ip_address character varying(15),
    data text,
    user_id integer,
    description character varying(191),
    method character varying(10),
    resvd7 character varying(1),
    resvd6 character varying(1),
    resvd5 character varying(1),
    resvd4 character varying(1),
    resvd3 character varying(1),
    resvd2 character varying(1),
    resvd1 character varying(1),
    created_at timestamp(0) WITH TIME ZONE  DEFAULT (current_timestamp at time zone 'EAT') NOT NULL,
    deleted_at timestamp(0) WITH TIME ZONE,
    updated_at timestamp(0) WITH TIME ZONE,
    primary key (id)
    );
    
ALTER TABLE public.audit_trails OWNER TO postgres;
