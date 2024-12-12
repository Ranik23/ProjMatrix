function sendDataToServer(url, payload) {
    showLoadingOverlay();

    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(result => {
            console.log("Ответ от сервера:", result);
            proceedToNextStep();
        })
        .catch(error => {
            console.error("Ошибка отправки данных:", error);
            alert("Произошла ошибка при вычислениях. Попробуйте ещё раз.");
        })
        .finally(() => {
            // Скрываем индикатор загрузки
            hideLoadingOverlay();
        });
}
