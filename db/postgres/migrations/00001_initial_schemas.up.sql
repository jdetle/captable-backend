CREATE TABLE public.shareholders (
    id serial NOT NULL,
    ownership_chunk_ids INTEGER[] NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    PRIMARY KEY (id)
);


CREATE TABLE public.cap_tables (
    id serial NOT NULl,
    total_shares int NOT NULL, 
    shareholder_ids INTEGER NOT NULL,
    company_name text NOT NULL, 
    share_price float NOT NULL,
    updated_at timestamp NOT NULL DEFAULT NOW(),
    created_at timestamp NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)    
);

CREATE TABLE public.ownership_chunks (
    id serial NOT NULL,
    shares_owned int NOT NULL,
    share_price float NOT NULL,
    captable_id int NOT NULL,
    shareholder_id int NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (captable_id) REFERENCES cap_tables(id),
    FOREIGN KEY (shareholder_id) REFERENCES shareholders(id)
);
