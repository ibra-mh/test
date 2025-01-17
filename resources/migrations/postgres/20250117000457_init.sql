DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'paths') THEN
        -- UP 
        CREATE TABLE paths (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            guid UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
            description TEXT,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE,
            deleted_at TIMESTAMP WITH TIME ZONE
        );

        CREATE TABLE contents (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            guid UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
            description TEXT,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE,
            deleted_at TIMESTAMP WITH TIME ZONE
        );

        CREATE TABLE content_paths (
            id SERIAL PRIMARY KEY,
            content_id INTEGER REFERENCES contents(id),
            path_id INTEGER REFERENCES paths(id),
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE,
            deleted_at TIMESTAMP WITH TIME ZONE
        );
    ELSE
        -- DOWN 
        DROP TABLE content_paths;
        DROP TABLE contents;
        DROP TABLE paths;
    END IF;
END $$;