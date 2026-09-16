// Package converter реалізує просту конвертацію валют за фіксованим курсом.
//
// Завдання 4 (бонус): реалізуйте ConvertCurrency самостійно або з допомогою
// ШІ-асистента — і порівняйте результат із вимогами в converter_test.go.
package converter

import "errors"

// Помилки, які повертає ConvertCurrency. Перевіряйте їх через errors.Is.
var (
	// ErrNegativeAmount повертається для від'ємної суми.
	ErrNegativeAmount = errors.New("converter: amount must not be negative")
	// ErrNonPositiveRate повертається для нульового або від'ємного курсу.
	ErrNonPositiveRate = errors.New("converter: rate must be positive")
)

// ConvertCurrency конвертує amount за курсом rate.
//
// Правила:
//   - результат дорівнює amount * rate;
//   - якщо amount від'ємний — повертається помилка;
//   - якщо rate від'ємний або дорівнює нулю — повертається помилка.
func ConvertCurrency(amount float64, rate float64) (float64, error) {
	if amount < 0 {
		return 0, ErrNegativeAmount
	}
	if rate <= 0 {
		return 0, ErrNonPositiveRate
	}
	return amount * rate, nil
}
