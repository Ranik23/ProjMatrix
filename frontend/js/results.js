// results.js

// Переход на главную страницу
function goToHome() {
    window.location.href = "/";
}

// Эта функция будет использоваться для загрузки и отображения результатов
function displayResults(data) {
    // Здесь можно будет отобразить результаты, полученные с сервера
    console.log("Результаты:", data);
}

// Экспортируем функции для использования на странице результатов
window.goToHome = goToHome;
window.displayResults = displayResults;
