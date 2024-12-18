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
            const tableData = document.getElementById("table-data");
            const tablename=document.getElementById("table-name");
            tablename.textContent = tableName.replace(/\b\w/g, char => char.toUpperCase());
            // Очищаем таблицу перед заполнением
            tableData.innerHTML = "";

            // Заголовки таблицы
            const thead = document.createElement("thead");
            const headerRow = document.createElement("tr");

            // Делаем заголовки на основе ключей первого объекта в массиве данных
            if (data.data.length > 0) {
                var columns=[]
                Object.keys(data.data[0]).forEach(key => {
                    const th = document.createElement("th");
                    th.textContent = key.replace(/_/g, " ").replace(/\b\w/g, char => char.toUpperCase()); // Форматируем название колонок
                    headerRow.appendChild(th);
                    columns.push(key)
                });
                const actionsTh = document.createElement("th");
                actionsTh.textContent = "Actions";
                headerRow.appendChild(actionsTh);
                thead.appendChild(headerRow);
                tableData.appendChild(thead);
            }

            // Данные таблицы
            const tbody = document.createElement("tbody");
            // Кнопка добавления новой записи
            const rowElement = document.createElement("tr");
            for (const _ in data.data[0]) {
                const td = document.createElement("td");
                rowElement.appendChild(td);
            }
            const addTd = document.createElement("td");
            const addButton = document.createElement("button");
            addButton.textContent = "Add";
            addButton.onclick = function() {
                openAddModalWithMetadata(tableName);
            };
            addTd.appendChild(addButton);
            rowElement.appendChild(addTd);
            tbody.appendChild(rowElement);
            //Заносим полученные данные через Json в таблицу
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
                    openEditModal(row,columns,tableName);
                };
                editTd.appendChild(editButton);
                rowElement.appendChild(editTd);

                tbody.appendChild(rowElement);
            });
            tableData.appendChild(tbody);
        })
        .catch(error => {console.error('Error loading table data:', error)
        });
}

function openAddModalWithMetadata(tableName) {
    fetch(`${options}/api/v1/${tableName}/metadata`)
        .then(response => response.json())
        .then(metadata => {
            const modal = document.getElementById("add-modal");
            const form = document.getElementById("add-form");
            form.innerHTML = ""; // Очищаем форму перед добавлением новых полей

            // Создаем поля формы на основе метаинформации
            metadata.data.forEach(column => {
                const div = document.createElement("div");
                div.classList.add("card");

                const label = document.createElement("label");
                label.textContent = column.name.replace(/_/g, " ").replace(/\b\w/g, char => char.toUpperCase());

                const input = document.createElement("input");
                input.type = column.type;
                input.id = column.name;
                if (column.required) {
                    input.required = true;
                }
                div.appendChild(label);
                div.appendChild(input);
                form.appendChild(div);
            });

            const addButton = document.createElement("button");
            addButton.type = "submit";
            addButton.textContent = "Add Record";
            form.appendChild(addButton);
            modal.style.display = "block";

            form.onsubmit = function (event) {
                event.preventDefault();
                const data = {};
                form.querySelectorAll("input").forEach(input => {
                    // Получаем метаинформацию о текущем поле
                    // Преобразуем значение в нужный тип
                    switch (input.type) {
                        case "number":
                            data[input.id] = parseFloat(input.value); // Преобразуем в число
                            break;
                        case "text":
                        case "email":
                        default:
                            data[input.id] = input.value; // Оставляем строку
                    }
                });

                fetch(`${options}/api/v1/${tableName}`, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(data)
                })
                    .then(response => {
                        if (!response.ok) {
                            throw new Error(`HTTP error! status: ${response.status}`);
                        }
                        return response.json();
                    })
                    .then(responseData => {
                        if (responseData.error) {
                            alert(`Error: ${responseData.error}`);
                            return;
                        }
                        alert("Record added successfully");
                        modal.style.display = "none";
                        loadTableData(tableName);
                    })
                    .catch(error => console.error("Error adding record:", error));
            };

            document.getElementById("add-close-modal").onclick = function () {
                modal.style.display = "none";
            };
            window.onclick = function (event) {
                if (event.target == modal) {
                    modal.style.display = "none";
                }
            };
        })
        .catch(error => console.error("Error fetching metadata:", error));
}

// Открыть модальное окно с данными для редактирования
function openEditModal(rowData, columns, tableName) { 
    const modal = document.getElementById("edit-modal");

    // Закрытие модального окна
    document.getElementById("edit-close-modal").onclick = function() {
        modal.style.display = "none";
    };

    // Закрытие при клике на область за пределами окна
    window.onclick = function(event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    };

    const form = document.getElementById("edit-form");
    form.innerHTML = ""; // Очищаем форму перед добавлением новых полей

    // Создание динамических полей формы на основе данных
    columns.forEach(column => {
        if (column !== "Actions") {  // Пропускаем колонку с действиями
            const div = document.createElement("div")
            div.classList.add('card');
            const label = document.createElement("label");
            label.textContent = column;
            const input = document.createElement("input");
            input.type = "text";
            input.id = column;
            input.value = rowData[column] || ""; // Заполняем значением из строки
            div.appendChild(label);
            div.appendChild(input);
            form.appendChild(div);
        }
    });
    // Добавление кнопки внутри формы
    const submitButton = document.createElement("button");
    submitButton.type = "submit";
    submitButton.textContent = "Save Changes";
    form.appendChild(submitButton);
    // Показать модальное окно
    document.getElementById("edit-modal").style.display = "block";

    // Обработчик отправки формы
    form.onsubmit = function(event) {
        event.preventDefault(); // Останавливаем обычную отправку формы

        const updatedData = {};
        columns.forEach(column => {
            if (column !== "Actions") {
                updatedData[column] = document.getElementById(column).value;
            }
        });

        // Отправка данных на сервер
        fetch(`${options}/api/v1/${tableName}/${rowData[columns[0]]}`, {
            method: "PUT",  // Используем метод PUT для обновления данных
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(updatedData)
        })
        .then(response => {
            if (!response.ok) {
                alert(`HTTP error! status: ${response.status}`);
            }
            return response.json()
        })
        .then(data => {
            if (data.error!=null){
                alert(`Error updating data: ${data.error}`);
                return;
            }
            alert(`Record updated`);
            // Закрыть модальное окно после сохранения
            modal.style.display = "none";
            // Перезагрузить таблицу
            loadTableData(tableName);
        })
        .catch(error => alert(`Error updating data: ${error}`));
    };
}
