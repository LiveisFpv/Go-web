import os
from flask import Flask, request, jsonify, send_file
from flask_cors import CORS
import requests
from fpdf import FPDF
import io

app = Flask(__name__)
CORS(app)

class PDF(FPDF):
    def header(self):
        self.set_font('DejaVu', 'B', 12)
        self.cell(0, 10, 'Динамический Отчет', 0, 1, 'C')

    def dynamic_table(self, data):
        # Извлечение ключей из первой записи для создания заголовков таблицы
        headers = data[0].keys()
        self.set_font('DejaVu', 'B', 10)
        for header in headers:
            self.cell(40, 10, header.replace("_", " ").capitalize(), 1, 0, 'C')
        self.ln()

        # Заполнение данных
        self.set_font('DejaVu', '', 10)
        for row in data:
            for value in row.values():
                self.cell(40, 10, str(value), 1, 0, 'C')
            self.ln()

@app.route('/generate-pdf', methods=['POST'])
def generate_pdf():
    # Получение JSON-данных из тела запроса
    request_data = request.get_json()
    # Проверка валидности данных
    if not request_data or "url" not in request_data:
        return jsonify({"error": "Invalid data"}), 400
    
    # Получение данных с API
    response = requests.get(request_data["url"])
    if response.status_code!= 200:
        return jsonify({"error": f"Failed to retrieve data from API: {response.status_code}"}), 403
    
    # Парсинг JSON-данных
    data = response.json()

    table_data = data["data"]
    
    # Генерация PDF
    pdf = PDF()
    font_dir = os.path.join(os.path.dirname(__file__), 'font')
    print(font_dir )
    # Подключение шрифта
    pdf.add_font('DejaVu', '', './font/DejaVuSans.ttf', uni=True)
    pdf.add_font('DejaVu', 'B', './font/DejaVuSans-Bold.ttf', uni=True)
    pdf.set_font('DejaVu', '', 12)
    pdf.add_page(orientation='L')
    
    pdf.dynamic_table(table_data)
    
    # Использование BytesIO для сохранения PDF в памяти
    pdf_buffer = io.BytesIO()
    # Теперь выводим в байтовый поток
    pdf_buffer.write(pdf.output(dest='S').encode('latin1'))  # 'S' возвращает строку, кодируем в 'latin1'
    pdf_buffer.seek(0)  # Устанавливаем указатель в начало
    return send_file(pdf_buffer, as_attachment=True, download_name="report.pdf", mimetype="application/pdf")

if __name__ == '__main__':
    app.run(port=5000)