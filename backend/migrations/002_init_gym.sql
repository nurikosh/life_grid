CREATE SCHEMA IF NOT EXISTS gym;


-- general exercises
CREATE TABLE IF NOT EXISTS gym.exercises (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    muscle_group VARCHAR(255)
);

-- gym sessions
CREATE TABLE IF NOT EXISTS gym.sessions (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES identity.users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL, -- e.g., "Leg Day", "Upper Body"
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    ended_at TIMESTAMPTZ DEFAULT NULL
);

-- session workout exercises
CREATE TABLE IF NOT EXISTS gym.session_exercises (
    id UUID PRIMARY KEY,
    exercise_id UUID REFERENCES gym.exercises(id) ON DELETE SET NULL,
    session_id UUID REFERENCES gym.sessions(id) ON DELETE CASCADE
);

-- exercise sets
CREATE TABLE IF NOT EXISTS gym.sets (
    id UUID PRIMARY KEY,
    session_exercise_id UUID REFERENCES gym.session_exercises(id) ON DELETE SET NULL,
    reps INTEGER,
    order_num INTEGER NOT NULL DEFAULT 1,
    weight FLOAT 
);

-- workout templates
CREATE TABLE IF NOT EXISTS gym.templates (
    id          UUID PRIMARY KEY,
    user_id     UUID REFERENCES identity.users(id) ON DELETE CASCADE,
    name        VARCHAR(255) NOT NULL,
    notes       TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- exercises in workout templates
CREATE TABLE IF NOT EXISTS gym.template_exercises (
    id            UUID PRIMARY KEY,
    template_id   UUID REFERENCES gym.templates(id) ON DELETE CASCADE,
    exercise_id   UUID REFERENCES gym.exercises(id) ON DELETE CASCADE,
    order_index   INTEGER NOT NULL DEFAULT 0,
    target_sets   INTEGER,
    target_reps   INTEGER,
    target_weight FLOAT
);
