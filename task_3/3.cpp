#include <iostream>
#include <cmath>
#include <numeric> // для std::gcd
using namespace std;

// Проверка сходимости ряда по признаку Даламбера
bool isConvergent(int b) {
    // Если 1/b < 1, т.е. b > 1, значит, ряд сходится
    return b > 1;
}

// Попытка вычислить сумму ряда и вернуть её как несократимую дробь
pair<int, int> getRationalSum(int a, int b) {
    const double EPS = 1e-10; // точность для прекращения суммирования
    const int MAX_DENOMINATOR = 1000000; // для преобразования в дробь

    double sum = 0.0;
    for (int n = 1; ; ++n) {
        double term = pow(n, a) / pow(b, n);
        if (term < EPS) break;
        sum += term;
    }

    // Преобразуем в несократимую дробь
    int numerator = static_cast<int>(round(sum * MAX_DENOMINATOR));
    int gcd_val = gcd(numerator, MAX_DENOMINATOR);
    return {numerator / gcd_val, MAX_DENOMINATOR / gcd_val};
}

int main() {
    int a, b;
    cout << "Введите значение для a: ";
    cin >> a;
    cout << "Введите значение для b: ";
    cin >> b;

    // Если b == 1, тогда знаменатель не убывает, и ряд расходится
    if (b == 1) {
        cout << "infinity" << endl;
        return 0;
    }

    // Проверка сходимости по признаку Даламбера
    if (!isConvergent(b)) {
        cout << "infinity" << endl;
        return 0;
    }

    // Вычисление суммы ряда
    auto [numerator, denominator] = getRationalSum(a, b);

    // Проверка на иррациональность: попробуем сравнить приближение с точной суммой
    double approx = static_cast<double>(numerator) / denominator;
    double exact = 0.0;
    for (int n = 1; n < 100; ++n) {
        exact += pow(n, a) / pow(b, n);
    }

    if (fabs(approx - exact) > 1e-6) {
        cout << "irrational" << endl;
    } else {
        cout << numerator << "/" << denominator << endl;
    }

    return 0;
}
