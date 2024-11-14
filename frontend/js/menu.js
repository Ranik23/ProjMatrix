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

    const data = {
        matrixCount: matrixCount,
        matrixSize: { rows: rows, columns: columns }
    };

    console.log("Данные для генерации линейной формы:", JSON.stringify(data));

    // Переход на страницу результатов после генерации
    proceedToNextStep();
}

window.showLinearFormGenerationPage = showLinearFormGenerationPage;
window.generateLinearFormData = generateLinearFormData;