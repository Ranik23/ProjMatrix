let linearFormCoefficients = [];
let currentLinearCoefficientIndex = 0;

// Функция для инициализации ввода коэффициентов при изменении количества матриц
function resetLinearFormCoefficientInput() {
    const matrixCount = parseInt(document.getElementById('matrix-count').value);

    // Сбросить массив коэффициентов и индекс
    linearFormCoefficients = [];
    currentLinearCoefficientIndex = 0;

    // Обновляем интерфейс: скрываем отображение старых коэффициентов, очищаем ввод
    const coefficientContainer = document.getElementById('linear-coefficient-entry-container');
    const coefficientDisplayContainer = document.getElementById('linear-coefficient-display-container');
    const coefficientInput = document.getElementById('linear-coefficient-input');

    // Скрываем старое отображение и показываем контейнер для ввода
    coefficientContainer.style.display = 'block';
    coefficientDisplayContainer.style.display = 'none';
    coefficientDisplayContainer.textContent = ''; // Очищаем старое отображение
    coefficientInput.value = ''; // Очищаем текущее значение ввода

    // Запустить ввод заново
    updateLinearCoefficientLabel();
}

function initializeLinearFormCoefficientInput() {
    const matrixCount = parseInt(document.getElementById('matrix-count').value);

    if (matrixCount >= 1) {
        resetLinearFormCoefficientInput();
    }
}

// Обработчик ввода коэффициента с помощью Enter
function handleLinearCoefficientEnter(event) {
    if (event.key === 'Enter') {
        const input = document.getElementById('linear-coefficient-input');
        const value = parseInt(input.value);

        if (!isNaN(value)) {
            linearFormCoefficients.push(value);
            currentLinearCoefficientIndex++;

            const matrixCount = parseInt(document.getElementById('matrix-count').value);
            if (currentLinearCoefficientIndex < matrixCount) {
                input.value = '';
                updateLinearCoefficientLabel();
            } else {
                displayLinearFormCoefficients();
            }
        }
    }
}

// Обновить метку для ввода коэффициентов
function updateLinearCoefficientLabel() {
    const label = document.getElementById('linear-coefficient-label');
    label.textContent = `Введите a${currentLinearCoefficientIndex}:`;
}

// Показать введённые коэффициенты
function displayLinearFormCoefficients() {
    const coefficientDisplayContainer = document.getElementById('linear-coefficient-display-container');
    const coefficientContainer = document.getElementById('linear-coefficient-entry-container');

    coefficientContainer.style.display = 'none';
    coefficientDisplayContainer.style.display = 'block';
    coefficientDisplayContainer.textContent = `Коэффициенты: [${linearFormCoefficients.join(', ')}]`;
}

// Экспортируем функции для глобального доступа
window.initializeLinearFormCoefficientInput = initializeLinearFormCoefficientInput;
window.handleLinearCoefficientEnter = handleLinearCoefficientEnter;
window.proceedToNextStep = proceedToNextStep;

function setupLinearFormMode() {
    document.getElementById('menu-container').style.display = 'none';
    document.getElementById('matrix-input-container').style.display = 'none';
    document.getElementById('linear-form-next-button').style.display = 'block';
    document.getElementById('matrix-count-selection').style.display = 'block';
    document.getElementById('linear-matrix-size-selection').style.display = 'block';

    // Устанавливаем значение по умолчанию для количества матриц и размера
    document.getElementById('matrix-count').value = "1";
    document.getElementById('linear-matrix-rows').value = "3";
    document.getElementById('linear-matrix-columns').value = "3";

    // Создаем матрицу по умолчанию
    setupMatrixSizeSelection();
    initializeLinearFormCoefficientInput();
}

// Экспортируем функцию для глобального доступа
window.setupLinearFormMode = setupLinearFormMode;



function showLinearFormGenerationPage() {
    // Скрываем другие элементы страницы
    document.getElementById('start-container').style.display = 'none';
    document.getElementById('menu-container').style.display = 'none';
    document.getElementById('linear-form-generation-container').style.display = 'block';
}

// Функция для сбора данных
function generateLinearFormData() {
    const matrixCount = parseInt(document.getElementById('matrix-count-generate').value);
    const rows = parseInt(document.getElementById('matrix-rows-generate-linear').value);
    const columns = parseInt(document.getElementById('matrix-columns-generate-linear').value);

    if (isNaN(rows) || rows < 1 || isNaN(columns) || columns < 1) {
        alert("Пожалуйста, введите корректный размер матрицы (1 или больше).");
        return;
    }

    if (isNaN(matrixCount) || matrixCount < 1) {
        alert("Пожалуйста, введите корректное количество матриц (1 или больше).");
        return;
    }

    const data = {
        operationType: "generate-linear-form",
        matrixCount: matrixCount,
        matrixSize: { rows: rows, columns: columns }
    };

    console.log("Данные для генерации линейной формы:", JSON.stringify(data));

    // Отправляем данные на сервер
    sendDataToServer("/api/submit", data);
}

window.showLinearFormGenerationPage = showLinearFormGenerationPage;
window.generateLinearFormData = generateLinearFormData;

function submitLinearFormData() {
    const matrixCount = parseInt(document.getElementById('matrix-count').value);
    const rows = parseInt(document.getElementById('linear-matrix-rows').value);
    const columns = parseInt(document.getElementById('linear-matrix-columns').value);

    if (isNaN(matrixCount) || matrixCount < 1 || isNaN(rows) || rows < 1 || isNaN(columns) || columns < 1) {
        alert("Пожалуйста, введите корректные данные.");
        return;
    }

    // Сбор данных матриц
    const matrices = [];
    for (let i = 0; i < matrixCount; i++) {
        const matrixInputs = Array.from(document.querySelectorAll(`#matrix-input-fields-container .matrix-border:nth-of-type(${i + 1}) input`));
        const matrix = matrixInputs.map(input => parseFloat(input.value) || 0); // Формируем одномерный массив
        matrices.push(matrix);
    }

    // Сбор коэффициентов
    const coefficients = linearFormCoefficients.slice(); // Копируем массив коэффициентов

    // Формирование JSON
    const data = {
        operationType: "manual-linear-form",
        matrixCount,
        matrixSize: { rows, columns },
        matrices, // Каждая матрица - одномерный массив
        coefficients,
    };

    console.log("Данные для линейной формы:", JSON.stringify(data));
    sendDataToServer("/api/submit", data);
}

// Экспортируем функцию для глобального доступа
window.submitLinearFormData = submitLinearFormData;


function initializeLinearFormGenerationInputs() {
    const matrixCountInput = document.getElementById('matrix-count-generate');
    const rowsInput = document.getElementById('matrix-rows-generate-linear');
    const columnsInput = document.getElementById('matrix-columns-generate-linear');

    // Устанавливаем значения по умолчанию при загрузке
    matrixCountInput.value = 1;
    rowsInput.value = 3;
    columnsInput.value = 3;

    // Устанавливаем обработчики для автоматической замены значений
    [matrixCountInput, rowsInput, columnsInput].forEach(input => {
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
}

// Вызываем функцию при загрузке страницы или переключении на страницу генерации линейной формы
initializeLinearFormGenerationInputs();
