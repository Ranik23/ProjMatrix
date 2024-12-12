function initializeDefaultInputs() {
    const rowsInput = document.getElementById('matrix-rows-generate');
    const columnsInput = document.getElementById('matrix-columns-generate');
    const degreeInput = document.getElementById('polynomial-degree-generate');

    // Устанавливаем значения по умолчанию при загрузке
    rowsInput.value = 3;
    columnsInput.value = 3;
    degreeInput.value = 1;

    // Устанавливаем обработчики для автоматической замены значений
    [rowsInput, columnsInput].forEach(input => {
        const defaultValue = input.value;

        input.addEventListener('focus', (event) => {
            if (event.target.value === defaultValue) {
                event.target.value = '';
            }
        });

        input.addEventListener('blur', (event) => {
            if (event.target.value === '') {
                event.target.value = defaultValue;
            }
        });
    });

    // Обработчики для степени полинома
    degreeInput.addEventListener('focus', (event) => {
        if (event.target.value === "1") {
            event.target.value = '';
        }
    });

    degreeInput.addEventListener('blur', (event) => {
        if (event.target.value === '') {
            event.target.value = 1;
        }
    });
}

// Вызываем функцию при загрузке страницы или переключении на страницу генерации
initializeDefaultInputs();

function validateMatrixValue(input) {
    // Если введено значение с запятой, заменяем её на точку
    input.value = input.value.replace(",", ".");

    // Проверка корректного формата (целое или вещественное число с точкой)
    if (!/^[-+]?\d+(\.\d+)?$/.test(input.value) && input.value !== "") {
        alert("Введите корректное значение для матрицы (целое или вещественное число через точку).");
        input.value = ""; // Сбрасываем к пустому значению
    }
}

function validatePolynomialDegree(input) {
    const value = parseInt(input.value);
    if (isNaN(value) || value <= 0 || !Number.isInteger(value)) {
        input.value = 1; // Значение по умолчанию
        alert("Введите положительное целое число для степени полинома.");
    }
}

function validateCoefficient(input) {
    const value = input.value;
    if (!/^[-+]?\d+(\.\d+)?$/.test(value)) {
        input.value = 1; // Значение по умолчанию
        alert("Введите корректное значение для коэффициента (целое или вещественное число через точку).");
    }
}

function validatePositiveInteger(input) {
    const value = input.value;
    if (!/^[1-9]\d*$/.test(value)) { // Проверка на положительные целые числа
        input.value = ""; // Сбрасываем к пустому значению
        alert("Введите положительное целое число.");
    }
}

// Экспортируем функции для использования в других файлах
window.validateMatrixValue = validateMatrixValue;
window.validatePolynomialDegree = validatePolynomialDegree;
window.validateCoefficient = validateCoefficient;