-- -------------------------------------------------------------
-- TablePlus 6.3.2(586)
--
-- https://tableplus.com/
--
-- Database: kreditplus-test
-- Generation Time: 2025-03-21 2:59:53.6240 PM
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS credit_limits_credit_id_seq;

-- Table Definition
CREATE TABLE "public"."credit_limits" (
    "credit_id" int8 NOT NULL DEFAULT nextval('credit_limits_credit_id_seq'::regclass),
    "user_id" int8 NOT NULL,
    "tenor" int8 NOT NULL,
    "initial_limit_amount" numeric,
    "used_limit_amount" numeric,
    "remaining_limit_amount" numeric,
    "created_by" text,
    "created_at" timestamptz,
    "updated_by" text,
    "updated_at" timestamptz,
    PRIMARY KEY ("credit_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS payments_id_seq;

-- Table Definition
CREATE TABLE "public"."payments" (
    "id" int8 NOT NULL DEFAULT nextval('payments_id_seq'::regclass),
    "transaction_id" int8 NOT NULL,
    "payment_amount" numeric,
    "payment_date" text,
    "created_by" text,
    "created_at" timestamptz,
    "updated_by" text,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS transactions_id_seq;

-- Table Definition
CREATE TABLE "public"."transactions" (
    "id" int8 NOT NULL DEFAULT nextval('transactions_id_seq'::regclass),
    "contract_number" text NOT NULL,
    "user_id" int8 NOT NULL,
    "otr" numeric,
    "tenor" int8,
    "admin_fee" numeric,
    "installment_amount" numeric,
    "interest" numeric,
    "asset_name" text,
    "created_by" text,
    "created_at" timestamptz,
    "updated_by" text,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_user_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "user_id" int8 NOT NULL DEFAULT nextval('users_user_id_seq'::regclass),
    "full_name" text NOT NULL,
    "legal_name" text,
    "email" text NOT NULL,
    "password" text NOT NULL,
    "nik" text NOT NULL,
    "birth_place" text,
    "birth_date" text,
    "salary" numeric,
    "ktp_photo" text,
    "selfie_photo" text,
    "created_by" text,
    "created_at" timestamptz,
    "updated_by" text,
    "updated_at" timestamptz,
    PRIMARY KEY ("user_id")
);

INSERT INTO "public"."credit_limits" ("credit_id", "user_id", "tenor", "initial_limit_amount", "used_limit_amount", "remaining_limit_amount", "created_by", "created_at", "updated_by", "updated_at") VALUES
(1, 1, 1, 100000, 0, 100000, '', '2025-03-20 14:07:21.446178+07', '', '2025-03-20 14:07:21.446178+07'),
(2, 2, 1, 1000000, 0, 1000000, '', '2025-03-20 14:07:21.446178+07', '', '2025-03-20 14:07:21.446178+07'),
(3, 1, 2, 200000, 0, 200000, '', '2025-03-20 14:07:21.446178+07', '', '2025-03-20 14:07:21.446178+07'),
(4, 2, 2, 1200000, 0, 1200000, '', '2025-03-20 14:07:21.446178+07', '', '2025-03-20 14:07:21.446178+07'),
(5, 1, 3, 500000, 0, 500000, '', '2025-03-20 14:07:21.446178+07', '', '2025-03-20 14:07:21.446178+07'),
(6, 2, 3, 1500000, 0, 1500000, '', '2025-03-20 14:07:21.446178+07', '', '2025-03-20 14:07:21.446178+07'),
(7, 1, 6, 700000, 0, 700000, '', '2025-03-20 14:07:21.446178+07', '', '2025-03-20 14:07:21.446178+07'),
(8, 2, 6, 2000000, 0, 2000000, '', '2025-03-20 14:07:21.446178+07', '', '2025-03-20 14:07:21.446178+07'),
(9, 3, 1, 1000000, 0, 1000000, 'Saipul Muiz', '2025-03-20 14:46:38.863594+07', 'Saipul Muiz', '2025-03-20 14:46:38.863594+07'),
(10, 3, 2, 1200000, 0, 1200000, 'Saipul Muiz', '2025-03-20 14:46:38.863595+07', 'Saipul Muiz', '2025-03-20 14:46:38.863595+07'),
(11, 3, 3, 1500000, 1000000, 500000, 'Saipul Muiz', '2025-03-20 14:46:38.863596+07', 'Saipul Muiz', '2025-03-21 10:52:36.317187+07'),
(12, 3, 6, 2000000, 0, 2000000, 'Saipul Muiz', '2025-03-20 14:46:38.863603+07', 'Saipul Muiz', '2025-03-20 14:46:38.863603+07');

INSERT INTO "public"."transactions" ("id", "contract_number", "user_id", "otr", "tenor", "admin_fee", "installment_amount", "interest", "asset_name", "created_by", "created_at", "updated_by", "updated_at") VALUES
(3, '480438053804', 3, 1000000, 3, 20000, 1358333.33, 15000, 'Monitor 17inc', '', '2025-03-21 10:52:36.315624+07', '', '2025-03-21 10:52:36.315624+07');

INSERT INTO "public"."users" ("user_id", "full_name", "legal_name", "email", "password", "nik", "birth_place", "birth_date", "salary", "ktp_photo", "selfie_photo", "created_by", "created_at", "updated_by", "updated_at") VALUES
(1, 'Budi', 'Budi', 'budi@gmail.com', '$2a$08$Nhp5rioE2tBU0abe1edX7eMqJOSuoK5AnnDZFJXhYIaSMY1yYUO1i', '2093928392829394', '', '', 0, '', '', '', '2025-03-19 15:36:45.248847+07', '', '2025-03-19 15:36:45.248847+07'),
(2, 'Annisa', 'Annisa', 'annisa@gmail.com', '$2a$08$b73o6mTise5LmM7qcbjIcu9ro2531EDSY3mAWp6qI4m0icRUazV5e', '2093928392829395', '', '', 0, '', '', '', '2025-03-19 15:36:45.248847+07', '', '2025-03-19 15:36:45.248847+07'),
(3, 'Saipul Muiz', 'Saipul Muiz', 'saipulmuiz87@gmail.com', '$2a$08$Ia5QzMUnBzPd4KZYKDpRaekiEx6gLktOtJJd5XTTACir9NN1Ma/Pa', '8493949593929493', 'ciamis', '1998-01-02', 10000000, '', '', 'saipulmuiz3@gmail.com', '2025-03-20 14:44:11.2443+07', 'saipulmuiz3@gmail.com', '2025-03-20 14:44:11.2443+07');



-- Indices
CREATE UNIQUE INDEX uni_transactions_contract_number ON public.transactions USING btree (contract_number);


-- Indices
CREATE UNIQUE INDEX uni_users_nik ON public.users USING btree (nik);
CREATE UNIQUE INDEX idx_users_user_id ON public.users USING btree (user_id);
