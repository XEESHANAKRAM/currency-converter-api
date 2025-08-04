// static/script.js - Frontend JavaScript logic

// DOM elements
const fromCurrency = document.getElementById('fromCurrency');
const toCurrency = document.getElementById('toCurrency');
const amount = document.getElementById('amount');
const result = document.getElementById('result');
const convertButton = document.getElementById('convertButton');
const swapButton = document.getElementById('swapButton');

// API base URL (adjust for your deployment)
const API_BASE = '/api/v1';

// Convert currency function
async function convertCurrency() {
    const fromValue = fromCurrency.value;
    const toValue = toCurrency.value;
    const amountValue = parseFloat(amount.value);

    // Validation
    if (!amountValue || amountValue <= 0) {
        showError('Please enter a valid amount');
        return;
    }

    if (fromValue === toValue) {
        result.textContent = amountValue.toFixed(2);
        return;
    }

    // Show loading state
    setLoadingState(true);

    try {
        // Make API request
        const response = await fetch(
            `${API_BASE}/convert?from=${fromValue}&to=${toValue}&amount=${amountValue}`
        );

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();

        if (data.success) {
            // Display result with animation
            result.textContent = data.converted_amount.toFixed(2);
            result.classList.add('success-animation');
            setTimeout(() => {
                result.classList.remove('success-animation');
            }, 500);

            // Show exchange rate info
            showSuccessMessage(
                `1 ${fromValue} = ${data.exchange_rate.toFixed(4)} ${toValue}`
            );
        } else {
            showError(data.error || 'Conversion failed');
        }
    } catch (error) {
        console.error('Conversion error:', error);
        showError('Failed to convert currency. Please try again.');
    } finally {
        setLoadingState(false);
    }
}

// Swap currencies function
function swapCurrencies() {
    const temp = fromCurrency.value;
    fromCurrency.value = toCurrency.value;
    toCurrency.value = temp;
    
    // Add swap animation
    swapButton.style.transform = 'rotate(180deg)';
    setTimeout(() => {
        swapButton.style.transform = 'rotate(0deg)';
    }, 300);

    // Auto-convert after swap
    if (amount.value) {
        convertCurrency();
    }
}

// UI helper functions
function setLoadingState(loading) {
    if (loading) {
        convertButton.innerHTML = '<span class="loading"></span> Converting...';
        convertButton.disabled = true;
    } else {
        convertButton.innerHTML = '<i class="fas fa-calculator mr-2"></i> Convert Currency';
        convertButton.disabled = false;
    }
}

function showError(message) {
    // Create and show error notification
    const errorDiv = document.createElement('div');
    errorDiv.className = 'fixed top-4 right-4 bg-red-500 text-white p-4 rounded-lg shadow-lg z-50';
    errorDiv.innerHTML = `<i class="fas fa-exclamation-circle mr-2"></i> ${message}`;
    document.body.appendChild(errorDiv);

    setTimeout(() => {
        errorDiv.remove();
    }, 3000);
}

function showSuccessMessage(message) {
    const successDiv = document.createElement('div');
    successDiv.className = 'fixed top-4 right-4 bg-green-500 text-white p-4 rounded-lg shadow-lg z-50';
    successDiv.innerHTML = `<i class="fas fa-check-circle mr-2"></i> ${message}`;
    document.body.appendChild(successDiv);

    setTimeout(() => {
        successDiv.remove();
    }, 3000);
}

// Event listeners
convertButton.addEventListener('click', convertCurrency);
swapButton.addEventListener('click', swapCurrencies);
amount.addEventListener('input', convertCurrency);
fromCurrency.addEventListener('change', convertCurrency);
toCurrency.addEventListener('change', convertCurrency);

// Initial conversion on page load
document.addEventListener('DOMContentLoaded', () => {
    if (amount.value) {
        convertCurrency();
    }
});