let options ="http://127.0.0.1:15432"
document.addEventListener("DOMContentLoaded", function() {
    fetch(`${options}/api/v1/tables`)  // Запрос для получения списка таблиц
        .then(response => response.json())
        .then(responseData => {
            const tablesList = document.getElementById("tables-list");
            // Мы получаем массив названий таблиц из поля 'data' в ответе
            responseData.data.forEach(table => {
                const link = document.createElement("a");
                link.href = "#";
                link.textContent = table;
                link.addEventListener("click", function(event) {
                    event.preventDefault();
                    loadTableData(table);
                });
                tablesList.appendChild(link);
                tablesList.appendChild(document.createElement("br"));
            });
        })
        .catch(error => console.error('Error loading tables:', error));
});


// Функция для загрузки данных таблицы
function loadTableData(tableName) {
    fetch(`${options}/api/v1/${tableName}`)
        .then(response => response.json())
        .then(data => {
            const body=document.getElementsByTagName("body")[0];
            body.style.height="100%";
            const tableData = document.getElementById("table-data");

            // Очищаем таблицу перед заполнением
            tableData.innerHTML = "";

            // Заголовки таблицы
            const thead = document.createElement("thead");
            const headerRow = document.createElement("tr");

            // Делаем заголовки на основе ключей первого объекта в массиве данных
            if (data.data.length > 0) {
                Object.keys(data.data[0]).forEach(key => {
                    const th = document.createElement("th");
                    th.textContent = key.replace(/_/g, " ").replace(/\b\w/g, char => char.toUpperCase()); // Форматируем название колонок
                    headerRow.appendChild(th);
                });
                const actionsTh = document.createElement("th");
                actionsTh.textContent = "Actions";
                headerRow.appendChild(actionsTh);
                thead.appendChild(headerRow);
                tableData.appendChild(thead);
            }

            // Данные таблицы
            const tbody = document.createElement("tbody");
            data.data.forEach(row => {
                const rowElement = document.createElement("tr");

                // Заполнение строк таблицы данными
                Object.values(row).forEach(cell => {
                    const td = document.createElement("td");
                    td.textContent = cell;
                    rowElement.appendChild(td);
                });

                // Кнопка для редактирования
                const editTd = document.createElement("td");
                const editButton = document.createElement("button");
                editButton.textContent = "Edit";
                editButton.onclick = function() {
                    alert("Editing not implemented");
                };
                editTd.appendChild(editButton);
                rowElement.appendChild(editTd);

                tbody.appendChild(rowElement);
            });
            tableData.appendChild(tbody);
        })
        .catch(error => {console.error('Error loading table data:', error)
            const body=document.getElementsByTagName("body")[0];
            body.style.height="100vh";
        });
}

