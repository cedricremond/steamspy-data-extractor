-- Table: public.games

-- DROP TABLE public.games;

CREATE TABLE public.games
(
    appid integer NOT NULL,
    name character varying COLLATE pg_catalog."default",
    negative integer,
    owners integer,
    developer character varying COLLATE pg_catalog."default",
    playersforevervariance integer,
    players2weeksvariance integer,
    average2weeks integer,
    ownersvariance integer,
    positive integer,
    publisher character varying COLLATE pg_catalog."default",
    playersforever integer,
    scorerank integer,
    players2weeks integer,
    userscore integer,
    averageforever integer,
    medianforever integer,
    median2weeks integer,
    price integer,
    CONSTRAINT games_pkey PRIMARY KEY (appid)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.games
    OWNER to "steam-spy";