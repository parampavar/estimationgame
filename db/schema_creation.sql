
-- Table: drop tables

DROP TABLE IF EXISTS estimate;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS toys;
DROP TABLE IF EXISTS tools;


-- Table: users

CREATE TABLE IF NOT EXISTS users
(
  id serial NOT NULL,
  idp_user_id character varying(100),
  name character varying(100),
  last_updated date,
  status character varying(100),
  user_attributes json,
  CONSTRAINT pk_users PRIMARY KEY (id),
  CONSTRAINT unq_users_01 UNIQUE (name)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE users
  OWNER TO postgres;

-- Table: toys

CREATE TABLE IF NOT EXISTS toys
(
  id serial NOT NULL,
  name character varying(100),
  "isActive" boolean,
  CONSTRAINT pk_toys_01 PRIMARY KEY (id),
  CONSTRAINT unq_toys_01 UNIQUE (name)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE toys
  OWNER TO postgres;

-- Table: tools

CREATE TABLE IF NOT EXISTS tools
(
  id serial NOT NULL,
  name character varying(100),
  "isActive" boolean,
  CONSTRAINT pk_tools_01 PRIMARY KEY (id),
  CONSTRAINT unq_tools_01 UNIQUE (name)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE tools
  OWNER TO postgres;

-- Table: create estimate

CREATE TABLE IF NOT EXISTS estimate
(
  id serial NOT NULL,
  value integer NOT NULL DEFAULT 0,
  "createdDate" date NOT NULL DEFAULT ('now'::text)::date,
  userid bigint NOT NULL,
  toyid bigint NOT NULL,
  CONSTRAINT pk_estimate PRIMARY KEY (id),
  CONSTRAINT fk_toys_estimate_id FOREIGN KEY (toyid)
      REFERENCES toys (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT fk_users_estimate_id FOREIGN KEY (userid)
      REFERENCES users (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
ALTER TABLE estimate
  OWNER TO postgres;

-- Index: fki_toys_estimate_id

DROP INDEX IF EXISTS fki_toys_estimate_id;

CREATE INDEX fki_toys_estimate_id
  ON estimate
  USING btree
  (toyid);

-- Index: fki_users_estimate_id

DROP INDEX IF EXISTS fki_users_estimate_id;

CREATE INDEX fki_users_estimate_id
  ON estimate
  USING btree
  (userid);

