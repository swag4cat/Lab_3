#include <iostream>
#include <iomanip>
#include <cmath>
#include <vector>

using namespace std;

// Функция расчёта температуры кофе по закону охлаждения Ньютона
double calculateCoffeeTemp(double initialTemp, double ambientTemp, double coolingRate, double time) {
    return ambientTemp + (initialTemp - ambientTemp) * exp(-coolingRate * time);
}

// Функция линейной аппроксимации по методу наименьших квадратов
void approximateLine(const vector<double>& x, const vector<double>& y, double& a, double& b) {
    int n = x.size();
    double sumX = 0, sumY = 0, sumXY = 0, sumX2 = 0;

    for (int i = 0; i < n; ++i) {
        sumX += x[i];
        sumY += y[i];
        sumXY += x[i] * y[i];
        sumX2 += x[i] * x[i];
    }

    a = (n * sumXY - sumX * sumY) / (n * sumX2 - sumX * sumX);
    b = (sumY - a * sumX) / n;
}

// Функция вычисления коэффициента корреляции Пирсона
double korrel(const vector<double>& x, const vector<double>& y) {
    int n = x.size();
    if (n == 0) return 0;

    double sumX = 0, sumY = 0;
    for (int i = 0; i < n; ++i) {
        sumX += x[i];
        sumY += y[i];
    }

    double meanX = sumX / n;
    double meanY = sumY / n;

    double numerator = 0, sumSqX = 0, sumSqY = 0;

    for (int i = 0; i < n; ++i) {
        double dx = x[i] - meanX;
        double dy = y[i] - meanY;
        numerator += dx * dy;
        sumSqX += dx * dx;
        sumSqY += dy * dy;
    }

    if (sumSqX == 0 || sumSqY == 0) return 0; // защита от деления на 0

    return numerator / sqrt(sumSqX * sumSqY);
}

// Функция для безопасного ввода положительного числа
double inputPositive(const string& prompt) {
    double value;
    while (true) {
        cout << prompt;
        cin >> value;
        if (cin.fail() || value <= 0) {
            cin.clear();
            cin.ignore(10000, '\n');
            cout << "Некорректный ввод. Повторите попытку.\n";
        } else {
            return value;
        }
    }
}

int main() {
    // Объявление переменных
    vector<double> times, temps;
    double initialTemp, ambientTemp, coolingRate, maxTime, timeStep;

    // Вывод заголовка

    cout << "╔════════════════════════════════╗\n";
    cout << "║  МОДЕЛИРОВАНИЕ ОСТЫВАНИЯ КОФЕ  ║\n";
    cout << "╚════════════════════════════════╝\n";

    // Ввод начальной температуры кофе
    cout << "Введите начальную температуру кофе (в °C): ";
    cin >> initialTemp;

    // Ввод температуры окружающей среды
    cout << "Введите температуру окружающей среды (в °C): ";
    cin >> ambientTemp;

    // Ввод коэффициента остывания (он должен быть положительным)
    coolingRate = inputPositive("Введите коэффициент остывания (в 1/мин): ");

    // Ввод общего времени моделирования
    maxTime = inputPositive("Введите общее время моделирования (в мин): ");

    // Ввод шага по времени
    timeStep = inputPositive("Введите шаг по времени (в мин): ");

    // Заголовок таблицы с результатами
    cout << "\n╔════════════════════════════════════════╗\n";
    cout << "║        РЕЗУЛЬТАТЫ МОДЕЛИРОВАНИЯ        ║\n";
    cout << "╠═══════════════╦════════════════════════╣\n";
    cout << "║   Время (мин) ║   Температура (°C)     ║\n";
    cout << "╠═══════════════╬════════════════════════╣\n";

    // Устанавливаем точность вывода до 2 знаков после запятой
    cout << fixed << setprecision(2);

    // Цикл моделирования остывания в заданном промежутке времени
    for (double t = 0; t <= maxTime; t += timeStep) {  // Вычисляем текущую температуру с использованием функции
        double currentTemp = calculateCoffeeTemp(initialTemp, ambientTemp, coolingRate, t);
        times.push_back(t);
        temps.push_back(currentTemp);
        cout << "║ " << setw(13) << t << " ║ " << setw(19) << currentTemp << "    ║\n"; // Выводим результаты в таблице
    }

    // Заключительный вывод таблицы
    cout << "╚═══════════════╩════════════════════════╝\n";

    double a, b;
    approximateLine(times, temps, a, b);

    // Выводим значения аппроксимирующей прямой
    cout << "\nАппроксимирующая прямая: y = " << fixed << setprecision(4) << a << " * x + " << b << endl;

    double r = korrel(times, temps);

    // Выводим значение коэффизиента корреляции
    cout << "Коэффициент корреляции: r = " << fixed << setprecision(4) << r << endl;

    return 0;
}
