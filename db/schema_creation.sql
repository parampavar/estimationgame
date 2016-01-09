
-- Table: drop tables

DROP TABLE IF EXISTS tblestimate;
DROP TABLE IF EXISTS tbluser;
DROP TABLE IF EXISTS tbltoy;
DROP TABLE IF EXISTS tbltool;


-- Table: tbluser

CREATE TABLE IF NOT EXISTS tbluser
(
  id serial NOT NULL,
  idp_tbluser_id character varying(100),
  name character varying(100),
  last_updated date,
  status character varying(100),
  tbluser_attributes json,
  CONSTRAINT pk_tbluser PRIMARY KEY (id),
  CONSTRAINT unq_tbluser_01 UNIQUE (name)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE tbluser
  OWNER TO postgres;

-- Table: tbltoy

CREATE TABLE IF NOT EXISTS tbltoy
(
  id serial NOT NULL,
  name character varying(100),
  "isActive" boolean,
  CONSTRAINT pk_tbltoy_01 PRIMARY KEY (id),
  CONSTRAINT unq_tbltoy_01 UNIQUE (name)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE tbltoy
  OWNER TO postgres;

-- Table: tbltool

CREATE TABLE IF NOT EXISTS tbltool
(
  id serial NOT NULL,
  name character varying(100),
  "isActive" boolean,
  CONSTRAINT pk_tbltool_01 PRIMARY KEY (id),
  CONSTRAINT unq_tbltool_01 UNIQUE (name)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE tbltool
  OWNER TO postgres;

-- Table: create tblestimate

CREATE TABLE IF NOT EXISTS tblestimate
(
  id serial NOT NULL,
  tbluserid bigint NOT NULL,
  tbltoyid bigint NOT NULL,
  value integer NOT NULL DEFAULT 0,
  "createdDate" date NOT NULL DEFAULT ('now'::text)::date,
  CONSTRAINT pk_tblestimate PRIMARY KEY (id),
  CONSTRAINT fk_tbltoy_tblestimate_id FOREIGN KEY (tbltoyid)
      REFERENCES tbltoy (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT fk_tbluser_tblestimate_id FOREIGN KEY (tbluserid)
      REFERENCES tbluser (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE tblestimate
  OWNER TO postgres;

-- Index: fki_tbltoy_tblestimate_id

DROP INDEX IF EXISTS fki_tbltoy_tblestimate_id;

CREATE INDEX fki_tbltoy_tblestimate_id
  ON tblestimate
  USING btree
  (tbltoyid);

-- Index: fki_tbluser_tblestimate_id

DROP INDEX IF EXISTS fki_tbluser_tblestimate_id;

CREATE INDEX fki_tbluser_tblestimate_id
  ON tblestimate
  USING btree
  (tbluserid);

