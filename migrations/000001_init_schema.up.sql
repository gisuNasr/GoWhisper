CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   username VARCHAR(50) UNIQUE NOT NULL,
   first_name VARCHAR(50),
   last_name VARCHAR(50),
   created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
   deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE TABLE devices (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
   identity_key_pub TEXT NOT NULL,
   signed_pre_key_pub TEXT NOT NULL,
   one_time_pre_keys JSONB NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
   deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE TABLE rooms (
     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
     name VARCHAR(100),
     type VARCHAR(20) NOT NULL CHECK (type IN ('pv', 'group')),
     created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
     updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE room_members (
      room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
      user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
      PRIMARY KEY (room_id, user_id)
);

CREATE TABLE messages (
      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
      user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
      device_id UUID NOT NULL REFERENCES devices(id) ON DELETE CASCADE,
      encrypted_payload TEXT NOT NULL,
      created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
      updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);