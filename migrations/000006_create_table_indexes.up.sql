CREATE INDEX chatusers_chat_id 						ON chatsusers(chat_id);
CREATE INDEX chatsusers_chat_id_user_id 	ON chatsusers(chat_id, user_id);
CREATE INDEX chatsusers_user_id 					ON chatsusers(user_id);
CREATE INDEX messages_created_at_chat_id 	ON messages (created_at, chat_id);