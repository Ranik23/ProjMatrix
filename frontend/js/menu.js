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

function showLoadingOverlay() {
    const loadingOverlay = document.getElementById("loading-overlay");
    loadingOverlay.style.display = "flex";
}

function hideLoadingOverlay() {
    const loadingOverlay = document.getElementById("loading-overlay");
    loadingOverlay.style.display = "none";
}
