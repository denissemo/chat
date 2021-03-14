CREATE KEYSPACE chat
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}
    AND durable_writes = true;

USE chat;

CREATE TABLE users (
    username text,
    password text,
    first_name text,
    last_name text,
    created_at timestamp,
    last_seen timestamp,
    PRIMARY KEY (username)
);

CREATE TABLE contacts (
    username text,
    contact text,
    first_name text,
    last_name text,
    last_seen timestamp,
    PRIMARY KEY (( username), last_seen)
) WITH CLUSTERING ORDER BY (last_seen DESC);

CREATE TABLE chats (
    id uuid,
    name text,
    type text,
    PRIMARY KEY ( id, type )
);

CREATE TABLE messages (
    message_id uuid,
    chat_id uuid,
    type text,
    date timestamp,
    edited timestamp,
    sender text,
    sender_username text,
    text text,
    reply_to_message_id uuid,
    PRIMARY KEY ( (chat_id), date )
) WITH CLUSTERING ORDER BY (date DESC);