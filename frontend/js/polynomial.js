// polynomial.js

let coefficients = [];
let currentCoefficientIndex = 0;


// Инициализировать ввод коэффициентов
function initializeCoefficientInput() {
    const degree = parseInt(document.getElementById('polynomial-degree').value);
    const coefficientContainer = document.getElementById('coefficient-entry-container');
    const coefficientDisplayContainer = document.getElementById('coefficient-display-container');

    if (degree >= 1) {
        coefficientContainer.style.display = 'block';
        coefficientDisplayContainer.style.display = 'none';
        currentCoefficientIndex = 0;
        coefficients = [];
        updateCoefficientLabel();

        // Установим значение по умолчанию для ввода коэффициентов
        const input = document.getElementById('coefficient-input');
        input.value = 1;

        // Устанавливаем обработчики событий для очистки значения по умолчанию
        input.addEventListener('focus', (event) => {
            if (event.target.value === "1") {
                event.target.value = '';
            }
        });

        input.addEventListener('blur', (event) => {
            if (event.target.value === '') {
                event.target.value = 1;
            }
        });
    }
}

// Экспортируем функцию в window
window.initializeCoefficientInput = initializeCoefficientInput;


// Обработчик для ввода коэффициента
// polynomial.js

function handleCoefficientEnter(event) {
    if (event.key === 'Enter') {
        const input = document.getElementById('coefficient-input');
        const value = parseFloat(input.value);

        if (!isNaN(value)) {
            coefficients.push(value);
            currentCoefficientIndex++;

            if (currentCoefficientIndex <= parseInt(document.getElementById('polynomial-degree').value)) {
                input.value = '';
                updateCoefficientLabel();
            } else {
                displayCoefficients();
            }
        }
    }
}

// Экспортируем функцию для глобального доступа
window.handleCoefficientEnter = handleCoefficientEnter;


// Обновить текст метки для коэффициента
function updateCoefficientLabel() {
    const label = document.getElementById('coefficient-label');
    label.textContent = `Введите a${currentCoefficientIndex}:`;
}

// Показать введенные коэффициенты
function displayCoefficients() {
    const coefficientDisplayContainer = document.getElementById('coefficient-display-container');
    const coefficientContainer = document.getElementById('coefficient-entry-container');

    coefficientContainer.style.display = 'none';
    coefficientDisplayContainer.style.display = 'block';
    coefficientDisplayContainer.textContent = `Коэффициенты: [${coefficients.join(', ')}]`;
}

// polynomial.js
function setupPolynomialMode() {
    document.getElementById('menu-container').style.display = 'none';
    document.getElementById('matrix-input-container').style.display = 'flex';
    document.getElementById('linear-form-next-button').style.display = 'none';

    document.getElementById('matrix-rows').value = "3";
    document.getElementById('matrix-columns').value = "3";

    createMatrixInputs();

    document.getElementById('polynomial-degree-container').style.display = 'block';
    document.getElementById('coefficient-entry-container').style.display = 'block';
    document.getElementById('coefficient-display-container').style.display = 'none';

    // Инициализация ввода коэффициентов
    initializeCoefficientInput();
}

window.setupPolynomialMode = setupPolynomialMode;

function initializePolynomialDegreeInput() {
    const degreeInput = document.getElementById('polynomial-degree');

    // Устанавливаем значение по умолчанию
    degreeInput.value = 1;

    // Обработчики событий для очистки и восстановления значения
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

// Вызываем настройку при загрузке или инициализации
initializePolynomialDegreeInput();

function generatePolynomialData() {
    const rows = parseInt(document.getElementById('matrix-rows-generate').value);
    const columns = parseInt(document.getElementById('matrix-columns-generate').value);
    const degree = parseInt(document.getElementById('polynomial-degree-generate').value);

    // Проверка корректности введенных данных
    if (isNaN(rows) || rows < 1 || isNaN(columns) || columns < 1) {
        alert("Пожалуйста, введите корректный размер матрицы (1 или больше).");
        return;
    }

    if (isNaN(degree) || degree < 1) {
        alert("Пожалуйста, введите корректную степень полинома (1 или больше).");
        return;
    }

    // Формирование JSON
    const data = {
        operationType: "generate-polynomial",
        matrixSize: { rows: rows, columns: columns },
        degree: degree
    };

    console.log("Данные для генерации полинома:", JSON.stringify(data));

    // Отправка данных на сервер
    sendDataToServer("/api/submit", data);
}
window.generatePolynomialData = generatePolynomialData;