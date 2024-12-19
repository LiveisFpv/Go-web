function renderChart(data) {
    const ctx = document.getElementById("reportChart").getContext("2d");

    // Пример данных для графика
    const labels = data.map(item => item.field1); // Например, имена групп
    const values = data.map(item => item.field2); // Например, количество студентов

    new Chart(ctx, {
        type: "bar", // Тип графика: 'bar', 'line', 'pie', 'doughnut' и т.д.
        data: {
            labels: labels,
            datasets: [{
                label: "Example Data",
                data: values,
                backgroundColor: "rgba(75, 192, 192, 0.2)",
                borderColor: "rgba(75, 192, 192, 1)",
                borderWidth: 1
            }]
        },
        options: {
            responsive: true,
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });
}

function exportToPDF() {
    const { jsPDF } = window.jspdf;
    const doc = new jsPDF();
    // Пример добавления текста в PDF
    doc.text("Report Title", 10, 10);

    // Пример добавления таблицы
    const table = document.getElementById("table-data"); // Таблица с данными
    if (table) {
        const data = [...table.rows].map(row => [...row.cells].map(cell => cell.innerText));
        doc.autoTable({
            head: [data[0]], // Заголовки таблицы
            body: data.slice(1), // Данные таблицы
        });
    }

    // Пример добавления диаграммы в PDF
    const canvas = document.getElementById("reportChart"); // Canvas с графиком
    if (canvas) {
        const chartImage = canvas.toDataURL("image/png");
        doc.addImage(chartImage, "PNG", 10, 50, 180, 80);
    }

    // Сохранение файла
    doc.save("report.pdf");
}
