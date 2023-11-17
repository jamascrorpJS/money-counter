package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func Jw(userID int) string {
	key := []byte("key")

	currentTime := time.Now().Local()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: currentTime.Add(time.Minute).Unix(),
		Subject:   fmt.Sprint(userID),
	})

	s, err := t.SignedString(key)
	if err != nil {
		// Обработка ошибки, например, вывод в консоль или логирование
		fmt.Println("Ошибка при подписи токена:", err)
		return ""
	}
	return s
}

func Jws(s string) (*jwt.Token, error) {
	key := []byte("key")
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что используется правильный метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Неправильный метод подписи: %v", token.Header["alg"])
		}

		// Возвращаем секретный ключ для проверки подписи
		return key, nil
	})

	// Проверяем ошибку при раскодировании
	if err != nil {
		// Обработка ошибки, например, вывод в консоль или логирование
		fmt.Println("Ошибка при раскодировании токена:", err)
		return nil, err
	}
	// Выводим раскодированный токен

	fmt.Println("Раскодированный токен:", token.Claims.(jwt.MapClaims)["sub"])
	return token, nil
}
