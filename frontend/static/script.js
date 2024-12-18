let options ="http://127.0.0.1:15432"
document.addEventListener("DOMContentLoaded", function() {
    fetch(`${options}/api/v1/tables`)  // Запрос для получения списка таблиц
        .then(response => response.json())
        .then(responseData => {
            const tablesList = document.getElementById("tables-list");
            // Мы получаем массив названий таблиц из поля 'data' в ответе
            responseData.data.forEach(table => {
                const button = document.createElement("button");
                button.textContent = table.replace(/\b\w/g, char => char.toUpperCase());
                button.addEventListener("click", function(event) {
                    event.preventDefault();
                    loadTableData(table);
                });
                tablesList.appendChild(button);
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
            // Чекбокс для выбора всех элементов
            const selectAllTh = document.createElement("th");
            const selectAllCheckbox = document.createElement("input");
            selectAllCheckbox.type = "checkbox";
            selectAllCheckbox.id = "select-all";
            selectAllCheckbox.onclick = function() {
                const checkboxes = document.querySelectorAll('.row-checkbox');
                checkboxes.forEach(checkbox => {
                    checkbox.checked = selectAllCheckbox.checked;
                });
            };
            selectAllTh.appendChild(selectAllCheckbox);
            headerRow.appendChild(selectAllTh);

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
            const td = document.createElement("td");
                rowElement.appendChild(td);
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

                // Чекбокс для выбора записи для удаления
                const checkboxTd = document.createElement("td");
                const rowCheckbox = document.createElement("input");
                rowCheckbox.type = "checkbox";
                rowCheckbox.classList.add("row-checkbox");
                checkboxTd.appendChild(rowCheckbox);
                rowElement.appendChild(checkboxTd);

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
                    const table = document.getElementById("table");
                    if(table.scrollWidth > table.clientWidth){
                        openEditModal(row,tableName);
                    } else{
                        // Если прокрутки нет, можно редактировать данные прямо в таблице
                        enableInlineEditing(row, tableName);
                    }
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
function enableInlineEditing(row, tableName) {
    // Получаем метаинформацию
    fetch(`${options}/api/v1/${tableName}/metadata`)
        .then(response => response.json())
        .then(metadata => {
            const tableData = document.getElementById("table-data").getElementsByTagName("tbody")[0];
            // Убираем все остальные активные строки редактирования
            const editingRows = tableData.querySelectorAll(".editing");
            editingRows.forEach(editingRow => {
                editingRow.classList.remove("editing");
                editingRow.innerHTML = ""; // Полностью очищаем, чтобы пересоздать
            });

            // Перезаполняем строку редактирования
            const rowElement = document.createElement("tr");
            rowElement.classList.add("editing");
            const td = document.createElement("td");
            rowElement.appendChild(td);
            metadata.data.forEach(column => {
                const td = document.createElement("td");
                const input = document.createElement("input");
                input.type = column.type || "text"; // Устанавливаем тип данных из метаинформации
                input.value = row[column.name] || ""; // Заполняем текущими значениями
                input.id = `edit-${column.name}`; // Уникальный ID для каждого input
                if (column.required) {
                    input.required = true; // Добавляем валидацию для обязательных полей
                }
                td.appendChild(input);
                rowElement.appendChild(td);
            });

            // Кнопки "Сохранить" и "Отмена"
            const actionsTd = document.createElement("td");
            const saveButton = document.createElement("button");
            saveButton.textContent = "Save";
            saveButton.onclick = function () {
                const updatedData = {};

                // Заполняем данные для отправки, основываясь на метаинформации
                metadata.data.forEach(column => {
                    const input = document.getElementById(`edit-${column.name}`);
                    if (input) {
                        updatedData[column.name] = column.type === "number"
                            ? parseFloat(input.value) || null
                            : input.value || null;
                    }
                });

                // Отправляем обновленные данные на сервер
                fetch(`${options}/api/v1/${tableName}`, {
                    method: "PUT",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(updatedData)
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
                        alert("Record updated successfully");
                        loadTableData(tableName); // Перезагружаем таблицу после успешного обновления
                    })
                    .catch(error => console.error("Error updating record:", error));
            };

            const cancelButton = document.createElement("button");
            cancelButton.textContent = "Cancel";
            cancelButton.onclick = function () {
                loadTableData(tableName); // Просто перезагружаем таблицу, чтобы отменить редактирование
            };

            actionsTd.appendChild(saveButton);
            actionsTd.appendChild(cancelButton);
            rowElement.appendChild(actionsTd);

            // Очищаем текущую строку и вставляем редактор
            tableData.innerHTML = ""; // Очищаем тело таблицы
            tableData.appendChild(rowElement);
        })
        .catch(error => console.error("Error loading metadata for inline editing:", error));
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
function openEditModal(rowData, tableName) {
    const modal = document.getElementById("edit-modal");

    // Закрытие модального окна
    document.getElementById("edit-close-modal").onclick = function () {
        modal.style.display = "none";
    };

    // Закрытие при клике на область за пределами окна
    window.onclick = function (event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    };

    // Запрос метаинформации у сервера
    fetch(`${options}/api/v1/${tableName}/metadata`)
        .then(response => response.json())
        .then(metadata => {
            const form = document.getElementById("edit-form");
            form.innerHTML = ""; // Очищаем форму перед добавлением новых полей

            // Создание динамических полей формы на основе метаинформации
            metadata.data.forEach(column => {
                const div = document.createElement("div");
                div.classList.add("card");

                const label = document.createElement("label");
                label.textContent = column.name.replace(/_/g, " ").replace(/\b\w/g, char => char.toUpperCase());

                const input = document.createElement("input");
                input.type = column.type
                input.id = column.name;
                input.value = rowData[column.name] || ""; // Заполняем значением из строки
                if (column.required) {
                    input.required = true;
                }

                div.appendChild(label);
                div.appendChild(input);
                form.appendChild(div);
            });

            // Добавление кнопки внутри формы
            const submitButton = document.createElement("button");
            submitButton.type = "submit";
            submitButton.textContent = "Save Changes";
            form.appendChild(submitButton);
            const deleteButton = document.createElement("button");
            deleteButton.className="delete"
            deleteButton.textContent = "Delete Record";
            form.appendChild(deleteButton);
            deleteButton.onclick = function (event) {
                event.preventDefault(); // Останавливаем обычную отправку формы
                const DeleteData = {};
                metadata.data.forEach(column => {
                    const input = document.getElementById(column.name);
                    if (column.type === "number") {
                        DeleteData[column.name] = parseFloat(input.value) || null;
                    } else {
                        DeleteData[column.name] = input.value || null;
                    }
                });
                if (confirm("Are you sure you want to delete this record?")) {
                    fetch(`${options}/api/v1/${tableName}`, {
                        method: "DELETE",
                        headers: {
                            "Content-Type": "application/json",
                        },
                        body: JSON.stringify(DeleteData),
                    })
                       .then(response => {
                            if (!response.ok) {
                                alert(`HTTP error! status: ${response.status}`);
                            }
                            return response.json();
                        })
                       .then(data => {
                            if (data.error) {
                                alert(`Error deleting record: ${data.error}`);
                            } else {
                                alert("Record deleted successfully");
                                modal.style.display = "none";
                                loadTableData(tableName);
                            }
                        })
                       .catch(error => console.error("Error deleting record:", error));
                }
            };

            // Показать модальное окно
            modal.style.display = "block";

            // Обработчик отправки формы
            form.onsubmit = function (event) {
                event.preventDefault(); // Останавливаем обычную отправку формы

                const updatedData = {};
                metadata.data.forEach(column => {
                    const input = document.getElementById(column.name);
                    if (column.type === "number") {
                        updatedData[column.name] = parseFloat(input.value) || null;
                    } else {
                        updatedData[column.name] = input.value || null;
                    }
                });

                // Отправка данных на сервер
                fetch(`${options}/api/v1/${tableName}`, {
                    method: "PUT", // Используем метод PUT для обновления данных
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify(updatedData),
                })
                    .then(response => {
                        if (!response.ok) {
                            throw new Error(`HTTP error! status: ${response.status}`);
                        }
                        return response.json();
                    })
                    .then(data => {
                        if (data.error != null) {
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
        })
        .catch(error => alert(`Error fetching metadata: ${error}`));
}
