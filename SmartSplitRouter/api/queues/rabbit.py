import pika
import json

class RabbitMQ:
    def __init__(self):
        self.user = "user"
        self.password = "password"
        self.host = "localhost"
        self.port = 5672
        self.connection = None
        self.channel = None
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
        
    
    def Consumer(self):
        def callback(ch, method, properties, body):
            print(f"Получено сообщение: {body.decode()}")

        connection = pika.BlockingConnection(pika.ConnectionParameters("localhost", 5672))
        channel = connection.channel()

        channel.queue_declare(queue="test_queue")

        channel.basic_consume(queue="test_queue", on_message_callback=callback, auto_ack=True)

        print("Ожидание сообщений. Для выхода нажми CTRL+C")
        channel.start_consuming()