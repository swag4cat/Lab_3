#include <iostream>
#include <cmath>      // для sqrt и log
#include <iomanip>    // для форматирования вывода (setw, setprecision)

using namespace std;

// Функция, реализующая график, заданный по частям
double f(double x) {
    if (x >= -4 && x <= 0) {
        // Полукруг: y = sqrt(4 - (x + 2)^2), центр в (-2, 0), радиус 2
        double underRoot = 4 - pow(x + 2, 2);
        return (underRoot >= 0) ? sqrt(underRoot) : NAN; // если подкоренное отрицательно — вернуть NAN
    }
    else if (x > 0 && x < 2) {
        // Функция: y = ln(x) / x
        return log(x) / x;
    }
    else if (x >= 2 && x <= 4) {
        // Постоянная функция: y = 1
        return 1;
    }
    else {
        // Вне области определения
        return NAN;
    }
}

int main() {
    // Интервал и шаг
    double x_start = -4.0;
    double x_end = 4.0;
    double dx = 0.5;

    // Заголовок таблицы
    cout << "\n╔════════════════════════════════════════╗\n";
    cout << "║        ТАБЛИЦА ЗНАЧЕНИЙ ФУНКЦИИ        ║\n";
    cout << "╠═══════════════╦════════════════════════╣\n";
    cout << "║   Значение x  ║     Значение f(x)      ║\n";
    cout << "╠═══════════════╬════════════════════════╣\n";

    // Основной цикл по X от x_start до x_end с шагом dx
    for (double x = x_start; x <= x_end + 1e-6; x += dx) {
        double y = f(x); // вычисляем значение функции
        cout << "║" << setw(9) << fixed << setprecision(2) << x << "      ║ ";
        if (!isnan(y)) {
            // Если значение определено — выводим
            cout << setw(15) << fixed << setprecision(5) << y << "        ║" << endl;
        } else {
            // Если значение не определено — выводим соответствующее сообщение
            cout << setw(10) << "нет значения" << endl;
        }
    }
    cout << "╚═══════════════╩════════════════════════╝\n";
    return 0;
}
