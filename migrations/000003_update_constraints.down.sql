
ALTER TABLE chatsUsers DROP CONSTRAINT cu_chat_id;
ALTER TABLE chatsUsers ADD CONSTRAINT cu_chat_id FOREIGN KEY(chat_id) REFERENCES chats(id);
ALTER TABLE chatsUsers DROP CONSTRAINT cu_user_id;
ALTER TABLE chatsUsers ADD CONSTRAINT cu_user_id FOREIGN KEY(user_id) REFERENCES users(id);