CREATE TABLE public.owner (
	id serial NOT NULL,
	tid integer NOT NULL,
    ownership_type string NOT NULL,
	email text NOT NULL
);
