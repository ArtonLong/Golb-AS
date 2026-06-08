# Golb-AS
for å kjør nettsiden må du først sette opp PostgreSQL: https://www.postgresql.org/
Gå til download -> windows -> Download the installer
bruk installern til å last ned postgres server og pgadmin
åpne pgadmin, sette det opp også lag en nytt database inn på serveren
åpne serveren -> høretrykk på Databases -> Create -> Database
høyetrykk på databasen du lagt og åpne Query Tool

når du er inn på Query Tool lim inn dette:
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INT REFERENCES users(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    summary TEXT,
    body TEXT NOT NULL,
    meta_json JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

da trykker du på Execute Script eller F5

i samme directory som main.go, lag en .env fil med to parameter
DB_URL=postgres://postgres_username:postgres_password@localhost:postgres_port/database_navn
API_SECRET=secret-key

erstatte postgres_username, password, port og database_name med det du lagt for postgres serveren

for å kjør nettsiden:
naviger til Golb-AS dir i terminalen
kjør: go run main.go

første oppstart kan ta litt tid