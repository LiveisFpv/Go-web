const modal = document.getElementById("edit-modal");
const editForm = document.getElementById("edit-form");
const editFieldsContainer = document.getElementById("edit-fields");

// Открытие формы редактирования
function editRow(button) {
    const row = button.closest("tr");
    const rowId = row.dataset.id;
    const fields = row.querySelectorAll(".editable");

    // Заполнение полей формы
    document.getElementById("edit-id").value = rowId;
    editFieldsContainer.innerHTML = "";
    fields.forEach(field => {
        const key = field.dataset.key;
        const value = field.textContent;
        const inputField = `
            <label>${key}: <input name="${key}" value="${value}"></label>
        `;
        editFieldsContainer.innerHTML += inputField;
    });

    modal.classList.remove("hidden");
}

// Закрытие модального окна
function closeModal() {
    modal.classList.add("hidden");
}

// Обновление строки через API
editForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const formData = new FormData(editForm);
    const updatedRow = { ID: document.getElementById("edit-id").value };

    formData.forEach((value, key) => {
        updatedRow[key] = value;
    });

    const response = await fetch("/update", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(updatedRow),
    });

    if (response.ok) {
        alert("Record updated successfully!");
        location.reload();
    } else {
        alert("Error updating record");
    }
});
