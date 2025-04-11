import pika
import json
import os

class RabbitMQ:
    def __init__(self):
        self.user = "user"
        self.password = "password"
        self.host = "localhost"
        self.port = 5672
        self.connection = None
        self.channel = None
        self.STORAGE_FILE = "processed.json"
        self.processed_ids = self.load_processed()
        self.connect()

    def connect(self):
        credentials = pika.PlainCredentials(self.user, self.password)
        parameters = pika.ConnectionParameters(host=self.host, port=self.port, credentials=credentials)
        self.connection = pika.BlockingConnection(parameters)
        self.channel = self.connection.channel()

    def close(self):
        if self.connection and not self.connection.is_closed:
            self.connection.close()
            
    def serialize_data(self, data: any) -> str:
        return json.dumps(data)

    def consume(self, queue_name, callback):
        if not self.channel:
            raise Exception("Connection is not established.")
        self.channel.basic_consume(queue=queue_name, on_message_callback=callback, auto_ack=True)
        self.channel.start_consuming()

    def Publish(self, queue_name, message):
        serialized_data = self.serialize_data(message)
        
        if not self.channel:
            raise Exception("Connection is not established.")
        self.channel.queue_declare(queue=queue_name)
        self.channel.basic_publish(exchange='',
                                   routing_key=queue_name,
                                   body=serialized_data,
                                   properties=pika.BasicProperties(
                                       delivery_mode=2,
                                   ))
        print(f"Sent message to queue {queue_name}: {message}")
    # Загрузить уже обработанные ID из файла
    def load_processed(self):
        if os.path.exists(self.STORAGE_FILE):
            with open(self.STORAGE_FILE, "r") as f:
                return set(json.load(f))
        return set()

    # Сохранить новые обработанные ID
    def save_processed(self, processed_ids):
        with open(self.STORAGE_FILE, "w") as f:
            json.dump(list(processed_ids), f)

    def Consumer(self, queue_name) -> dict:        
        self.channel.queue_declare(queue=queue_name)

        method_frame, header_frame, body = self.channel.basic_get(queue=queue_name, auto_ack=False)
        self.channel.queue_purge(queue=queue_name)
        
        if method_frame:
            msg_id = header_frame.message_id

            if msg_id in self.processed_ids:
                print("❗️ Дубликат — отклоняем")
                self.channel.basic_ack(delivery_tag=method_frame.delivery_tag)
            else:
                print("✅ Уникальное сообщение — обрабатываем")
                self.processed_ids.add(msg_id)
                self.save_processed(self.processed_ids)

                # Здесь логика обработки
                self.channel.basic_ack(delivery_tag=method_frame.delivery_tag)
                return body.decode()
        else:
            print("Очередь пуста.")
        
        self.connection.close()
