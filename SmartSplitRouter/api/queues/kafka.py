import json
from typing import List
from confluent_kafka import Consumer, KafkaException, KafkaError, Producer
import sys

class Kafka():
    producer: any
    consumer: any
    
    def __init__(self, brokers: List):
        self.producer = Producer({
                    'bootstrap.servers': f'{brokers[0]}',
                    'client.id': 'python-producer'
        })
        
        self.consumer = Consumer({
                'bootstrap.servers': f'{brokers[0]}', 
                'group.id': 'mygroup',
                'auto.offset.reset': 'earliest'
        })

    def serialize_data(self, data: any):
        return json.dumps(data)

    def send_message(self, topic: str, message: str):
        self.producer.produce(topic, value=message)
        self.producer.flush()
    
    def Publisher(self, topic: str, data: any):
        # Сериализуем данные
        serialized_data = self.serialize_data(data)
        
        # Отправляем данные в Kafka
        self.send_message(topic=topic, message=serialized_data)
        
        print(f"Sent data: {serialized_data}")

    def Subscriber(self, topics):
            try:
                # подписываемся на топик
                self.consumer.subscribe(topics)

                while True:
                    msg = self.consumer.poll(timeout=1.0)
                    if msg is None: continue
                    if msg.error():
                        if msg.error().code() == KafkaError._PARTITION_EOF:
                            sys.stderr.write('%% %s [%d] reached end at offset %d\n' %
                                            (msg.topic(), msg.partition(), msg.offset()))
                        elif msg.error():
                            raise KafkaException(msg.error())
                    else:
                        print(f"Received message: {msg.value().decode('utf-8')}")
                        break
            except KeyboardInterrupt:
                pass
            finally:
                self.consumer.close()

   


def main():
    kaf = Kafka(brokers=["localhost:9092"])
    
    kaf.Publisher("test-topic", {"first_name": "Alex", "last_name": "Berl", "age": 12})
    print("\n")
    kaf.Subscriber(["test-topic"])


if __name__ == "__main__":
    main()