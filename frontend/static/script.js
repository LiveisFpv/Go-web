const options = "http://127.0.0.1:15432";

document.addEventListener("DOMContentLoaded", async () => {
    try {
        const response = await fetch(`${options}/api/v1/tables`);
        const responseData = await response.json();

        const tablesList = document.getElementById("tables-list");
        responseData.data.forEach(table => {
            const button = document.createElement("button");
            button.textContent = capitalizeWords(table);
            button.addEventListener("click", event => {
                event.preventDefault();
                loadTableData(table);
            });
            tablesList.append(button, document.createElement("br"));
        });
    } catch (error) {
        console.error("Error loading tables:", error);
    }
});

async function loadTableData(tableName) {
    try {
        const [dataResponse, metadataResponse] = await Promise.all([
            fetch(`${options}/api/v1/${tableName}`),
            fetch(`${options}/api/v1/${tableName}/metadata`)
        ]);

        const data = await dataResponse.json();
        const metadata = await metadataResponse.json();

        const tableData = document.getElementById("table-data");
        const tableNameElement = document.getElementById("table-name");
        tableNameElement.textContent = capitalizeWords(tableName);
        tableData.innerHTML = "";

        if (data.data.length > 0) {
            createTableHeader(tableData, data.data[0], metadata.data);
            createTableBody(tableData, data.data, metadata.data, tableName);
        }
    } catch (error) {
        console.error("Error loading table data:", error);
    }
}

function createTableHeader(tableData, sampleRow, metadata) {
    const thead = document.createElement("thead");
    const headerRow = document.createElement("tr");

    const selectAllTh = document.createElement("th");
    const selectAllCheckbox = document.createElement("input");
    selectAllCheckbox.type = "checkbox";
    selectAllCheckbox.addEventListener("click", () => {
        document.querySelectorAll(".row-checkbox").forEach(checkbox => {
            checkbox.checked = selectAllCheckbox.checked;
        });
    });
    selectAllTh.appendChild(selectAllCheckbox);
    headerRow.appendChild(selectAllTh);

    Object.keys(sampleRow).forEach(key => {
        const th = document.createElement("th");
        th.textContent = capitalizeWords(key.replace(/_/g, " "));
        headerRow.appendChild(th);
    });

    const actionsTh = document.createElement("th");
    actionsTh.textContent = "Actions";
    headerRow.appendChild(actionsTh);

    thead.appendChild(headerRow);
    tableData.appendChild(thead);
}

function createTableBody(tableData, rows, metadata, tableName) {
    const tbody = document.createElement("tbody");
    // Кнопка добавления новой записи
    const addRowElement = document.createElement("tr");
    const emptyTd = document.createElement("td");
    addRowElement.appendChild(emptyTd);
    
    for (const _ in rows[0]) {
        const td = document.createElement("td");
        addRowElement.appendChild(td);
    }
    const addTd = document.createElement("td");
    const addButton = document.createElement("button");
    addButton.textContent = "Add";
    addButton.onclick = async function () {
        await openAddModalWithMetadata(tableName);
    };
    addTd.appendChild(addButton);
    addRowElement.appendChild(addTd);
    tbody.appendChild(addRowElement);

    // Обработка строк данных
    rows.forEach(row => {
        const rowElement = document.createElement("tr");
        
        // Получение уникального значения для data-id
        const uniqueField = metadata.find(field => field.unique);
        if (uniqueField && row[uniqueField.name]) {
            rowElement.setAttribute("data-id", row[uniqueField.name]);
        } else {
            console.error("Unique value not found for row:", row);
        }

        const checkboxTd = document.createElement("td");
        const checkbox = document.createElement("input");
        checkbox.type = "checkbox";
        checkbox.classList.add("row-checkbox");
        checkboxTd.appendChild(checkbox);
        rowElement.appendChild(checkboxTd);

        Object.values(row).forEach(cell => {
            const td = document.createElement("td");
            td.textContent = cell;
            rowElement.appendChild(td);
        });

        // Кнопка редактирования
        const editTd = document.createElement("td");
        const editButton = document.createElement("button");
        editButton.textContent = "Edit";
        editButton.classList.add("edit-button");
        editButton.onclick = async function () {
            const table = document.getElementById("table");
            if (table.scrollWidth > table.clientWidth) {
                await openEditModal(row, tableName);
            } else {
                await enableInlineEditing(row, tableName);
            }
        };
        editTd.appendChild(editButton);
        rowElement.appendChild(editTd);

        tbody.appendChild(rowElement);
    });

    tableData.appendChild(tbody);
}

async function enableInlineEditing(row, tableName) {
    try {
        // Получаем метаинформацию
        const response = await fetch(`${options}/api/v1/${tableName}/metadata`);
        const metadata = await response.json();
        // Получаем уникальное поле row[uniqueField.name]
        const uniqueField = metadata.data.find(field => field.unique);
        // Получаем строку по data-id
        const rowElement = document.querySelector(`#table-data tbody tr[data-id="${row[uniqueField.name]}"]`)
        if (!rowElement) {
            console.error("Row with id " + row[uniqueField.name] + " not found.");
            return;
        }
        const editButtons = document.querySelectorAll('.edit-button');
        editButtons.forEach(button => {
            button.style.display = 'none'; // Скрываем каждую кнопку
        });
        rowElement.innerHTML = "";
        const td = document.createElement("td");
        rowElement.appendChild(td);
        metadata.data.forEach(column => {
            const td = document.createElement("td");
            const input = document.createElement("input");
            input.type = column.type || "text";
            input.value = row[column.name] || "";
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
        saveButton.onclick = async function () {
            const updatedData = {};
            event.preventDefault();
            // Заполняем данные для отправки, основываясь на метаинформации
            metadata.data.forEach(column => {
                const input = document.getElementById(`edit-${column.name}`);
                if (input) {
                    updatedData[column.name] = column.type === "number"
                        ? parseFloat(input.value) || null
                        : input.value || null;
                }
            });
            await updateRow(tableName, updatedData);
            loadTableData(tableName);
        }
        const cancelButton = createButton("Cancel", () => loadTableData(tableName));
        actionsTd.append(saveButton, cancelButton);
        rowElement.appendChild(actionsTd);
    } catch (error) {
        loadTableData(tableName);
        console.error("Error enabling inline editing:", error);
    }
}

async function updateRow(tableName, data) {
    try {
        const response = await fetch(`${options}/api/v1/${tableName}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(data)
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        alert("Record updated successfully");
    } catch (error) {
        console.error("Error updating row:", error);
    }
}

function createButton(text, onClick) {
    const button = document.createElement("button");
    button.textContent = text;
    button.addEventListener("click", onClick);
    return button;
}

function capitalizeWords(string) {
    return string.replace(/\b\w/g, char => char.toUpperCase());
}
async function openAddModalWithMetadata(tableName) {
    try {
        // Получаем метаинформацию
        const response = await fetch(`${options}/api/v1/${tableName}/metadata`);
        const metadata = await response.json();
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

            form.onsubmit = async function (event) {
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
                const response = await fetch(`${options}/api/v1/${tableName}`,{
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(data)
                });
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const responseData = await response.json()
                if (responseData.error) {
                    alert(`Error: ${responseData.error}`);
                    return;
                }
                alert("Record added successfully");
                modal.style.display = "none";
                loadTableData(tableName);
            }
            document.getElementById("add-close-modal").onclick = function () {
                modal.style.display = "none";
            };
            window.onclick = function (event) {
                if (event.target == modal) {
                    modal.style.display = "none";
                }
            };
    } catch (error) {
        loadTableData(tableName);
        console.error("Error opening add modal:", error);
    }
}

// Открыть модальное окно с данными для редактирования
async function openEditModal(rowData, tableName) {
    try{
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
        const response = await fetch(`${options}/api/v1/${tableName}/metadata`);
        const metadata = await response.json();
        const form = document.getElementById("edit-form");
        form.innerHTML = ""; // Очищаем форму перед добавлением новых полей
        metadata.data.forEach(column => {
            const div = document.createElement("div");
            div.classList.add("card");

            const label = document.createElement("label");
            label.textContent = capitalizeWords(column.name.replace(/_/g, " "));
            const input = document.createElement("input");
            input.type = column.type;
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
        deleteButton.onclick = async function (event) {
            const DeleteData = {};
            event.preventDefault();
            metadata.data.forEach(column => {
                const input = document.getElementById(column.name);
                if (column.type === "number") {
                    DeleteData[column.name] = parseFloat(input.value) || null;
                } else {
                    DeleteData[column.name] = input.value || null;
                }
            });
            if (confirm("Are you sure you want to delete this record?")) {
                const response = await fetch(`${options}/api/v1/${tableName}`, {
                    method: "DELETE",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify(DeleteData),
                })
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const responseData = await response.json()
                if (responseData.error) {
                    alert(`Error deleting record: ${responseData.error}`);
                } else {
                    alert("Record deleted successfully");
                    modal.style.display = "none";
                    loadTableData(tableName);
                }

            };
        };

        // Показать модальное окно
        modal.style.display = "block";

        // Обработчик отправки формы
        submitButton.onclick = async function (event) {
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
            await updateRow(tableName, updatedData);
            modal.style.display = "none";
            loadTableData(tableName);
        };
    } catch (error) {
        loadTableData(tableName);
        console.error("Error opening add modal:", error);
    };
}
