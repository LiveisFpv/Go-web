import os
from flask import Flask, request, jsonify, send_file
from flask_cors import CORS
import requests
from fpdf import FPDF
import io
import datetime
from datetime import datetime

app = Flask(__name__)
CORS(app)

class PDF(FPDF):
    def footer(self):
        self.set_y(-15)
        self.set_font('DejaVu', '', 8)
        self.cell(0, 10, f'Страница {self.page_no()}', 0, 0, 'C')
    def report_info(self, report_data):
        self.set_font('DejaVu', 'B', 12)
        self.ln(5)

        self.set_font('DejaVu', '', 10)
        column_width = 60
        line_height = 8

        for key, value in report_data.items():
            self.cell(column_width, line_height, key + ":", border=1)
            self.cell(0, line_height, str(value), border=1)
            self.ln()
    def dynamic_table(self, data, table_name):
        # Заголовок таблицы
        self.set_font('DejaVu', 'B', 14)
        self.cell(0, 10, table_name, 0, 1, 'C')
        self.ln(5)

        # Отображение шапки таблицы
        headers = data[0].keys()
        column_width = 45
        line_height = 10

        self.set_font('DejaVu', 'B', 10)
        self.set_fill_color(180, 180, 180)
        self.set_text_color(255, 255, 255)
        for header in headers:
            self.cell(column_width, line_height, header.replace("_", " ").capitalize(), 1, 0, 'C', fill=True)
        self.ln()

        # Отображение данных таблицы
        self.set_font('DejaVu', '', 10)
        self.set_text_color(0, 0, 0)
        fill = False

        for row in data:
            for value in row.values():
                x = self.get_x()
                y = self.get_y()
                self.multi_cell(column_width, line_height, str(value), border=1, align='C', fill=fill)
                self.set_xy(x + column_width, y)
            self.ln()
            fill = not fill



@app.route('/generate-pdf', methods=['POST'])
def generate_pdf():
    # Получение JSON-данных из тела запроса
    request_data = request.get_json()
    # Проверка валидности данных
    if not request_data or "url" not in request_data:
        return jsonify({"error": "Invalid data"}), 400
    
    # Получение токена из запроса
    token = request_data.get("token")
    if not token:
        return jsonify({"error": "No authorization token provided"}), 401
    
    # Создаем заголовки с токеном авторизации
    headers = {
        "Authorization": f"Bearer {token}"
    }
    
    # Получение данных с API
    response = requests.get(request_data["url"], headers=headers)
    if response.status_code != 200:
        return jsonify({"error": f"Failed to retrieve data from API: {response.status_code}"}), 403
    
    # Парсинг JSON-данных
    data = response.json()

    table_data = data["data"]
    table_name = request_data["name"].capitalize()
    now=datetime.now()
    report_data = {
        "Дата и время создания": now.strftime("%Y-%m-%d %H:%M:%S"),
        "Создатель": request_data.get("creator", "Иванов Иван Иванович"),
        "Организация": request_data.get("organization", "Ниженский Государственный институт"),
        "Примечание": request_data.get("note", "Создан пример отчета")
    }

    # Генерация PDF
    pdf = PDF()
    # Подключение шрифта
    pdf.add_font('DejaVu', '', './font/DejaVuSans.ttf', uni=True)
    pdf.add_font('DejaVu', 'B', './font/DejaVuSans-Bold.ttf', uni=True)
    pdf.set_font('DejaVu', '', 12)
    pdf.add_page(orientation='L')
    pdf.report_info(report_data)
    pdf.dynamic_table(table_data,table_name)
    url=request_data['url']
    for i in range(1,int(data["pagination"]["total"]/data["pagination"]["page_size"])):
        # Получение следующей страницы
        url=url.replace(f"page={i}",f"page={i+1}")
        response = requests.get(f"{url}", headers=headers)
        if response.status_code != 200:
            return jsonify({"error": f"Failed to retrieve data from API: {response.status_code}"}), 403
        data=response.json()

        table_data = data["data"]
        pdf.add_page(orientation='L')
        pdf.dynamic_table(table_data,table_name)
    
    # Использование BytesIO для сохранения PDF в памяти
    pdf_buffer = io.BytesIO()
    pdf_buffer.write(pdf.output(dest='S').encode('latin1'))  # 'S' возвращает строку, кодируем в 'latin1'
    pdf_buffer.seek(0)  # Устанавливаем указатель в начало
    return send_file(pdf_buffer, as_attachment=True, download_name="report.pdf", mimetype="application/pdf")

if __name__ == '__main__':
    app.run(host='0.0.0.0',port=9000)