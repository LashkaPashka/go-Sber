import pytesseract
from PIL import Image
import re
import json
from typing import List, Dict, Union

def extract_text_from_image(image_path: str) -> str:
    """Извлекает текст из изображения с помощью OCR."""
    image = Image.open(image_path)
    text = pytesseract.image_to_string(image, lang='rus')
    return text

def parse_receipt(text: str) -> Dict[str, Union[float, List[Dict]]]:
    """Анализирует текст чека и возвращает структурированные данные."""
    lines = [line.strip() for line in text.split('\n') if line.strip()]
    products = []
    total_account = 0.0
    current_product = None

    # Паттерны для поиска информации
    price_pattern = re.compile(r'(\d{1,3}(?:[ ,.]\d{3})*)[,.](\d{2})\s*$')  # Ищет цены
    total_pattern = re.compile(r'ИТОГО К ОПЛАТЕ:\s*([\d ,.]+)', re.IGNORECASE)
    quantity_pattern = re.compile(r'(\d+)[=,](\d+)')  # Ищет количество вида "1=2"

    i = 0
    while i < len(lines):
        line = lines[i]
        
        # Проверяем, является ли строка общей суммой
        total_match = total_pattern.search(line)
        if total_match:
            total_str = total_match.group(1).replace(' ', '').replace(',', '.')
            total_account = float(total_str)
            i += 1
            continue

        # Проверяем, содержит ли строка цену
        price_match = price_pattern.search(line)
        if price_match:
            price_str = price_match.group(1).replace(' ', '').replace(',', '.') + '.' + price_match.group(2)
            price = float(price_str)
            
            # Проверяем наличие информации о количестве
            quantity_match = quantity_pattern.search(line)
            number_servings = 1
            if quantity_match:
                number_servings = int(quantity_match.group(2))
                if number_servings == 0:
                    number_servings = 1

            # Получаем название продукта
            product_name = line[:price_match.start()].strip()
            
            # Если название пустое, берем предыдущую строку
            if not product_name and i > 0:
                product_name = lines[i-1].strip()
                # Удаляем предыдущую строку из продуктов, если она была добавлена как отдельный продукт
                if products and products[-1]['name'] == product_name and products[-1]['price'] == 0:
                    products.pop()

            # Создаем продукт
            if product_name:
                products.append({
                    'name': product_name,
                    'price': price / number_servings,
                    'total': price,
                    'numberServings': number_servings
                })
            i += 1
        else:
            # Если строка не содержит цены, это может быть начало нового продукта
            if current_product is None:
                current_product = {'name': line, 'numberServings': 1}
            else:
                current_product['name'] += ' ' + line
            i += 1

    # Фильтруем продукты с нулевой ценой (если есть)
    products = [p for p in products if p.get('price', 0) > 0]

    return {
        'total_account': total_account,
        'products': products,
    }

def receipt_to_json(image_path: str) -> str:
    """Преобразует изображение чека в JSON модель."""
    text = extract_text_from_image(image_path)
    receipt_data = parse_receipt(text)
    return json.dumps(receipt_data, ensure_ascii=False, indent=2)

# Пример использования
if __name__ == "__main__":
    image_path = 'api/static/images/4.jpg'
    try:
        json_output = receipt_to_json(image_path)
        print(json_output)
    except Exception as e:
        print(f"Произошла ошибка: {str(e)}")
        print("Убедитесь, что:")
        print("1. Установлен Tesseract OCR и добавлен в PATH")
        print("2. Установлены все необходимые библиотеки (pillow, pytesseract)")
        print("3. Изображение четкое и хорошо читаемое")