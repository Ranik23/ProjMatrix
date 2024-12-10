// menu.js
function showMenu() {
    document.getElementById('start-container').style.display = 'none';
    document.getElementById('menu-container').style.display = 'flex';
}

function showMatrixInput() {
    const dataInput = document.getElementById('data-input').value;
    const calculationType = document.getElementById('calculation-type').value;

    if (dataInput === 'manual') {
        if (calculationType === 'polynomial') {
            setupPolynomialMode();
        } else if (calculationType === 'linear-form') {
            setupLinearFormMode();
        }
    } else if (dataInput === 'generate') {
        if (calculationType === 'polynomial') {
            showPolynomialGenerationPage();
        } else if (calculationType === 'linear-form') {
            showLinearFormGenerationPage();
        }
    }
}

// Экспортируем функции в window, чтобы они были доступны глобально
window.showMenu = showMenu;
window.showMatrixInput = showMatrixInput;

function showPolynomialGenerationPage() {
    // Скрываем другие элементы страницы
    document.getElementById('start-container').style.display = 'none';
    document.getElementById('menu-container').style.display = 'none';
    document.getElementById('polynomial-generation-container').style.display = 'block';
}

// Экспортируем функцию для глобального доступа
window.showPolynomialGenerationPage = showPolynomialGenerationPage;

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

function showLoadingOverlay() {
    const loadingOverlay = document.getElementById("loading-overlay");
    loadingOverlay.style.display = "flex";
}

function hideLoadingOverlay() {
    const loadingOverlay = document.getElementById("loading-overlay");
    loadingOverlay.style.display = "none";
}
