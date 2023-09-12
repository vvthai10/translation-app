CREATE TABLE "translate" (
    "id" bigserial PRIMARY KEY,
    "original_text" varchar NOT NULL,
    "source" varchar NOT NULL,
    "destination" varchar NOT NULL,
    "result_text" varchar NOT NULL
)
