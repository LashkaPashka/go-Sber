import pika
import json

def serialize_data(self, data: any):
        return json.dumps(data)

def Producer(msg: any):
    serialized_data = serialize_data(msg)
    connection = pika.BlockingConnection(pika.ConnectionParameters("localhost", 5672))
    channel = connection.channel()
    
    channel.queue_declare(queue="test-queue")
    
    channel.basic_publish(
        exchange="",
        routing_key="topic-divide",
        body=msg
    )
    
    print("Sent meassage!")
    connection.close
    
def Consumer():
    def callback(ch, method, properties, body):
        print(f"📥 Получено сообщение: {body.decode()}")

    connection = pika.BlockingConnection(pika.ConnectionParameters("localhost", 5672))
    channel = connection.channel()

    channel.queue_declare(queue="test_queue")

    channel.basic_consume(queue="test_queue", on_message_callback=callback, auto_ack=True)

    print("Ожидание сообщений. Для выхода нажми CTRL+C")
    channel.start_consuming()