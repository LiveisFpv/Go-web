const options = "http://127.0.0.1:15432";
const py_options = "http://127.0.0.1:9000";
let currentURL = "";
let currentTable = "";
const registerButton = document.getElementById('signup-button');
const loginButton = document.getElementById('login-button');
const logountButton = document.getElementById('logout-button');
const userinfo = document.getElementById('user-info');
const registerWindow = document.getElementById('registerWindow');
const loginWindow = document.getElementById('loginWindow');
const closeButtonsignup = document.getElementById('closeButtonsignup');
const closeButtonlogin = document.getElementById('closeButtonlogin');
var token=""

function auth(){
    logountButton.classList.remove('hidden');
    userinfo.classList.remove('hidden');
    loginButton.classList.add('hidden');
    registerButton.classList.add('hidden');
}

registerButton.addEventListener('click', () => {
  registerWindow.classList.remove('hidden');
});
loginButton.addEventListener('click', () => {
    loginWindow.classList.remove('hidden');
})

closeButtonsignup.addEventListener('click', () => {
    registerWindow.classList.add('hidden');
});
closeButtonlogin.addEventListener('click', () => {
    loginWindow.classList.add('hidden');
});

const signupsubmit =document.getElementById("signup-submit")
const loginsubmit = document.getElementById("login-submit")

signupsubmit.addEventListener('click', async () => {
    const email = document.getElementById("signup-email").value;
    const login = document.getElementById("signup-login").value;
    const password = document.getElementById("signup-password").value;
    const response = await fetch(`${options}/register`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(
            {
                email: email,
                login: login,
                password: password
            }
        )
    });
    const responseData = await response.json()
    if (responseData.error) {
        alert(`Error: ${responseData.error}`);
        return;
    }
    else{
    token=responseData.data.token;
    auth();
    registerWindow.classList.add('hidden');
    }

});

loginsubmit.addEventListener('click', async () => {
    const login = document.getElementById("login").value;
    const password = document.getElementById("login-password").value;
    const response = await fetch(`${options}/auth`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(
            {
                login: login,
                password: password
            }
        )
    });
    const responseData = await response.json()
    if (responseData.error) {
        alert(`Error: ${responseData.error}`);
        return;
    }
    else{
        token=responseData.data.token;
        auth();
        loginWindow.classList.add('hidden');
    }
});

document.addEventListener("DOMContentLoaded", async () => {
    try {
        const response = await fetch(`${options}/tables`);
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
function openFilter(){
    const filterContainer = document.getElementById("filters-form");
    if (filterContainer.style.display=="none"){
        filterContainer.style.display="block";
    } else {
        filterContainer.style.display="none";
    }
}
async function deleteRows() {
    tableName=document.getElementById("table-name").textContent.toLowerCase();
    // Найти все отмеченные чекбоксы
    const table=document.getElementById("table-data");
    const selectedCheckboxes = table.querySelectorAll(".row-checkbox:checked");

    // Получить идентификаторы строк
    const idsToDelete = Array.from(selectedCheckboxes).map(checkbox => {
        const row = checkbox.closest("tr");
        return row.getAttribute("data-id");
    }).filter(id => id !== null); // Убедиться, что id не пустой

    if (idsToDelete.length === 0) {
        alert("Please, select rows for remove.");
        return;
    }

    // Отправить запрос на backend
    try {
        const response = await fetch(`${options}/api/v1/${tableName}/ids`, { 
            method: 'Delete',
            headers: {
                'Authorization': 'Bearer ' + token,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ ids: idsToDelete }),
        });

        if (!response.ok) {
            throw new Error(`Error remove rows: ${response.statusText}`);
        }
        
        await loadTableData(tableName); // Обновляем таблицу после удаления

        alert("All rows deleted successfully");
    } catch (error) {
        console.error("Error remove rows:", error);
        alert("Error remove rows.");
    }
}
async function generatePDF() {
    const jsonData = {
        url: currentURL.replace('127.0.0.1','backend-go'),
        name: currentTable
    };
    
    const response = await fetch(`${py_options}/generate-pdf`, {
        method: "POST",
        headers: {
            'Authorization': 'Bearer ' + token,
            "Content-Type": "application/json"
        },
        body: JSON.stringify(jsonData)
    });

    if (response.ok) {
        const blob = await response.blob();
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = url;
        a.download = "report.pdf";
        document.body.appendChild(a);
        a.click();
        a.remove();
    } else {
        console.error("Failed to generate PDF");
    }
}
async function loadTableData(tableName,page=1,filters = {}) {
    try {
        const filterParams = new URLSearchParams(filters).toString(); // Преобразуем фильтры в строку параметров
        try{
        search=document.getElementById("search-input").value; //Получаем данные из строки
        } catch{
            search='';
        }
        const url = `${options}/api/v1/${tableName}/?page=${page}${filterParams ? '&' + filterParams : ''}&search=${search}`;
        currentURL = url;
        currentTable = tableName;
        console.log(url)
        const [dataResponse, metadataResponse] = await Promise.all([
            fetch(url, {
                method: "GET",
                headers: {
                    "Authorization": 'Bearer '+token,
                    "Content-Type": "application/json"
                }
            }),
            fetch(`${options}/api/v1/${tableName}/metadata`,{
                method: "GET",
                headers: {
                    "Authorization": 'Bearer '+token,
                    "Content-Type": "application/json"
                }
            })
        ]);

        const data = await dataResponse.json();
        const metadata = await metadataResponse.json();

        const tableData = document.getElementById("table-data");
        const tableNameElement = document.getElementById("table-name");
        tableNameElement.textContent = capitalizeWords(tableName);
        tableData.innerHTML = "";
        // Создаем фильтры на основе метаданных
        if (Object.keys(filters).length == 0) {
            createFilters(metadata.data);  // Добавляем фильтры 
        }
        if (data.data.length > 0) {
            createTableHeader(tableData, metadata.data);
            createTableBody(tableData, data.data, metadata.data, tableName,page,filters);
            paginationLoad(tableName, page, data.pages,filters);
        }
    } catch (error) {
        console.error("Error loading table data:", error);
    }
}

// Функция для динамического создания фильтров
async function createFilters(metadata) {
    const filterContainer = document.getElementById("filters-form");
    filterContainer.innerHTML = `<input type="text" id="search-input" placeholder="Search...">`; // Очищаем старые фильтры
    
    const filters = {}; // Это объект для хранения фильтров

    metadata.forEach(column => {
        if (column.filterable !== false) {
            const filterDiv = document.createElement("div");
            const label = document.createElement("label");
            label.textContent = capitalizeWords(column.name.replace(/_/g, " "));
            const input = document.createElement("input");
            input.type = column.type || "text"; // Используем текстовый input по умолчанию
            input.id = `filter-${column.name}`;
            // if (column.required) {
            //     input.required = true; // Добавляем валидацию для обязательных полей
            // }
            // Сохраняем фильтры в объект
            input.addEventListener("input", () => {
                filters[column.name] = input.value;
            });
            if (column.id){
                filterDiv.style.display ="none";
            }
            filterDiv.appendChild(label);
            filterDiv.appendChild(input);
            filterContainer.appendChild(filterDiv);
        }
    });

    // Кнопка применения фильтров
    const applyButton = document.createElement("button");
    applyButton.textContent = "Apply Filters";
    applyButton.addEventListener("click", async (event) => applyFilters(filters,event));
    filterContainer.appendChild(applyButton);
}
function applyFilters(filters,event) {
    event.preventDefault();
    const tableName = document.getElementById("table-name").textContent.toLowerCase();
    loadTableData(tableName, 1, filters); // Перезагружаем таблицу с фильтрами
}

async function paginationLoad(tableName,pageStart, countPages,filters){
    if (countPages>0){
        const paginationContainer = document.getElementById("pagination");
        paginationContainer.innerHTML = "";
        
        //Созадем переход на первую страницу
        let pageLi= document.createElement("li");
        let pageLink = document.createElement("a");
        pageLink.innerHTML = "&laquo;";
        pageLink.href = "";
        pageLink.addEventListener("click", async (event) => {
            event.preventDefault();
            await loadTableData(tableName, 1,filters);
            // Убираем aria-current у всех ссылок
            const currentLink = paginationContainer.querySelector('a[aria-current="page"]');
            if (currentLink) {
                currentLink.removeAttribute("aria-current");
            }
            document.getElementById(`page-${1}`).setAttribute("aria-current", "page");
        });
        pageLi.appendChild(pageLink);
        paginationContainer.appendChild(pageLi);

        //Создаем кнопки для пагинации с нумерами страниц
        for (let i = Math.max(1,pageStart-1); i <= Math.min(pageStart+1,countPages); i++) {
            const pageLi= document.createElement("li");
            const pageLink = document.createElement("a");
            pageLink.id = `page-${i}`
            pageLink.textContent = i;
            pageLink.href = "";
            pageLink.addEventListener("click", async (event) => {
                event.preventDefault();
                await loadTableData(tableName, i, filters);
                document.getElementById(`page-${i}`).setAttribute("aria-current", "page");
                document.getElementById(`page-${i-1}`).removeAttribute("aria-current");
            });
            pageLi.appendChild(pageLink);
            paginationContainer.appendChild(pageLi);
        }
        document.getElementById(`page-${pageStart}`).setAttribute("aria-current", "page");
        
        //Созадем переход на последнюю страницу
        pageLi= document.createElement("li");
        pageLink = document.createElement("a");
        pageLink.innerHTML = "&raquo;";
        pageLink.href = "";
        pageLink.addEventListener("click", async (event) => {
            event.preventDefault();
            await loadTableData(tableName, countPages,filters);
            // Убираем aria-current у всех ссылок
            const currentLink = paginationContainer.querySelector('a[aria-current="page"]');
            if (currentLink) {
                currentLink.removeAttribute("aria-current");
            }
            document.getElementById(`page-${countPages}`).setAttribute("aria-current", "page");
        });
        pageLi.appendChild(pageLink);
        paginationContainer.appendChild(pageLi);
    }
}

function createTableHeader(tableData, data) {
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

    data.forEach(column => {
        const th = document.createElement("th");
        th.textContent = capitalizeWords(column.name.replace(/_/g, " "));
        headerRow.appendChild(th);
        if (column.id==true){
            th.classList.add("hide-column");
        }
    });

    const actionsTh = document.createElement("th");
    actionsTh.textContent = "Actions";
    headerRow.appendChild(actionsTh);

    thead.appendChild(headerRow);
    tableData.appendChild(thead);
}

function createTableBody(tableData, rows, metadata, tableName,page,filters) {
    const tbody = document.createElement("tbody");
    // Кнопка добавления новой записи
    const addRowElement = document.createElement("tr");
    const emptyTd = document.createElement("td");
    addRowElement.appendChild(emptyTd);
    metadata.forEach(column => {
        const td = document.createElement("td");
        if (column.id==true){
            td.classList.add("hide-column");
        }
        addRowElement.appendChild(td);
    })
    const addTd = document.createElement("td");
    const addButton = document.createElement("button");
    addButton.textContent = "Add";
    addButton.onclick = async function () {
        await openAddModalWithMetadata(tableName,page,filters);
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
        metadata.forEach(column => {
            const td = document.createElement("td");
            td.textContent = row[column.name];
            if (column.id==true){
                td.classList.add("hide-column");
            }
            rowElement.appendChild(td);
        })

        // Кнопка редактирования
        const editTd = document.createElement("td");
        const editButton = document.createElement("button");
        editButton.textContent = "Edit";
        editButton.classList.add("edit-button");
        editButton.onclick = async function () {
            const table = document.getElementById("table");
            if (table.scrollWidth > table.clientWidth) {
                await openEditModal(row, tableName,page,filters);
            } else {
                await enableInlineEditing(row, tableName,page,filters);
            }
        };
        editTd.appendChild(editButton);
        rowElement.appendChild(editTd);

        tbody.appendChild(rowElement);
    });

    tableData.appendChild(tbody);
}

async function enableInlineEditing(row, tableName,page,filters) {
    try {
        // Получаем метаинформацию
        const response = await fetch(`${options}/api/v1/${tableName}/metadata`,{
            method: "GET",
            headers: {
                'Authorization': 'Bearer ' + token,
            }});
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
            if (column.id==true){
                td.classList.add("hide-column");
            }
            td.appendChild(input);
            rowElement.appendChild(td);
        });
        
        // Кнопки "Сохранить" и "Отмена"
        const actionsTd = document.createElement("td");
        const saveButton = document.createElement("button");
            saveButton.textContent = "Save";
        saveButton.onclick = async function (event) {
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
            loadTableData(tableName,page,filters);
        }
        const cancelButton = createButton("Cancel", () => loadTableData(tableName,page,filters));
        actionsTd.append(saveButton, cancelButton);
        rowElement.appendChild(actionsTd);
    } catch (error) {
        loadTableData(tableName,page,filters);
        console.error("Error enabling inline editing:", error);
    }
}

async function updateRow(tableName, data) {
    try {
        const response = await fetch(`${options}/api/v1/${tableName}/`, {
            method: "PUT",
            headers: { 
                'Authorization': 'Bearer ' + token,
                "Content-Type": "application/json"
            },
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
async function openAddModalWithMetadata(tableName,page,filters) {
    try {
        // Получаем метаинформацию
        const response = await fetch(`${options}/api/v1/${tableName}/metadata`,{
            method: "GET",
            headers: {
                'Authorization': 'Bearer ' + token,
            }});
        const metadata = await response.json();
        const modal = document.getElementById("add-modal");
            const form = document.getElementById("add-form");
            form.innerHTML = ""; // Очищаем форму перед добавлением новых полей

            // Создаем поля формы на основе метаинформации
            metadata.data.forEach(column => {
                const div = document.createElement("div");
                if (column.id != true){
                    div.classList.add("card");
                }
                else{
                    div.style.display = "none";
                }

                const label = document.createElement("label");
                label.textContent = column.name.replace(/_/g, " ").replace(/\b\w/g, char => char.toUpperCase());

                const input = document.createElement("input");
                input.type = column.type;
                input.id = column.name;
                if (column.required && !column.id) {
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
                const response = await fetch(`${options}/api/v1/${tableName}/`,{
                    method: "POST",
                    headers: {
                        'Authorization': 'Bearer ' + token,
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
                loadTableData(tableName,page,filters);
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
        loadTableData(tableName,page,filters);
        console.error("Error opening add modal:", error);
    }
}

// Открыть модальное окно с данными для редактирования
async function openEditModal(rowData, tableName,page,filters) {
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
        const response = await fetch(`${options}/api/v1/${tableName}/metadata`,{
            method: "GET",
            headers: {
                'Authorization': 'Bearer ' + token,
            }});
        const metadata = await response.json();
        const form = document.getElementById("edit-form");
        form.innerHTML = ""; // Очищаем форму перед добавлением новых полей
        metadata.data.forEach(column => {
            const div = document.createElement("div");
            if (column.id != true){
                div.classList.add("card");
            }
            else{
                div.style.display = "none";
            }

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
                const response = await fetch(`${options}/api/v1/${tableName}/`, {
                    method: "DELETE",
                    headers: {
                        'Authorization': 'Bearer ' + token,
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
                    loadTableData(tableName,page,filters);
                }

            };
        };

        // Показать модальное окно
        modal.style.display = "block";

        // Обработчик отправки формы
        submitButton.onsubmit = async function (event) {
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
            loadTableData(tableName,page,filters);
        };
    } catch (error) {
        loadTableData(tableName,page,filters);
        console.error("Error opening add modal:", error);
    };
}
