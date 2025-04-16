#include <iostream>
#include <vector>
#include <tuple>
#include <cmath>
#include <random>
#include <set>

using namespace std;

// Перегрузка оператора вывода для вектора
template <typename T>
ostream& operator<<(ostream& os, vector<T>& result){
    for (auto item : result){
        os << "\t" << item << "\t║";  // Выводим каждый элемент вектора
    }
    // Выводим разделительную линию
    os << "\n╠═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣\n";
    return os;
}

// Функция для возведения числа в степень по модулю
int powMod(int base, int degree, int mod) {
    int result = base % mod;  // Начальное значение - base по модулю mod
    for (int i = 2; i <= degree; i++) {
        result = (result * base) % mod;  // Множим результат на base и берем по модулю
    }
    return result;
}

// Функция для нахождения простых чисел с использованием решета Эратосфена
vector<int> sieveEratos(int N) {
    vector<bool> boolPrime(N + 1, true);  // Массив для пометки простых чисел
    for (int p = 2; p * p <= N; p++) {  // Для каждого числа от 2 до корня из N
        if (boolPrime[p]) {  // Если число простое
            for (int i = p * p; i <= N; i += p) {  // Помечаем все его кратные как составные
                boolPrime[i] = false;
            }
        }
    }

    vector<int> primes;
    for (int p = 2; p <= N; p++) {
        if (boolPrime[p]) {  // Если число простое, добавляем его в список
            primes.push_back(p);
        }
    }

    return primes;
}

// Функция для генерации случайного числа в заданном диапазоне
int randDist(int min, int max) {
    random_device rd;
    mt19937 rand(rd());
    uniform_int_distribution<int> dist(min, max);  // Генерация случайного числа в диапазоне
    return dist(rand);
}

// Функция для выполнения теста Рабина-Миллера
bool rabin(int num, int k) {
    if (num == 2 || num == 3) {
        return true;  // 2 и 3 всегда простые
    }
    if (num < 2 || num % 2 == 0) {
        return false;  // Четные числа, кроме 2, не простые
    }

    int s = 0;
    int d = num - 1;
    while (d % 2 == 0) {  // Разлагаем num - 1 на 2^s * d
        d /= 2;
        s++;
    }

    for (int i = 0; i < k; i++) {  // Повторяем тест k раз
        int a = rand() % (num - 4) + 2;  // Случайное основание для теста
        int x = powMod(a, d, num);
        int y;
        for (int j = 0; j < s; j++) {
            y = powMod(x, 2, num);  // Возведение в квадрат по модулю
            if (y == 1 && x != 1 && x != (num - 1)) {
                return false;  // Если x^2 ≡ 1 и x != 1 и x != num - 1, то число составное
            }
            x = y;
        }
        if (y != 1) {
            return false;  // Если не выполняется условие, то число составное
        }
    }
    return true;  // Если число прошло все проверки, оно, вероятно, простое
}

// Функция для генерации простого числа с заданной битностью
tuple<int, vector<int>> calcN(vector<int> prime, int bit) {
    int maxIndex = prime.size() - 1;
    int maxPow = 1;
    // Находим максимальную степень числа, не превосходящую 2^(bit-1)
    for (; pow(2, maxPow) < pow(2, bit - 1); maxPow++);

    int m = 1;
    vector<int> qList;

    while (true) {
        int randIndex = randDist(0, maxIndex);  // Случайный индекс простого числа
        int alphaDegree = randDist(1, maxPow);  // Случайная степень

        if (pow(prime[randIndex], alphaDegree)) {
            if (m * pow(prime[randIndex], alphaDegree) < pow(2, bit - 1)) {
                m *= pow(prime[randIndex], alphaDegree);  // Умножаем m на простое число в степени alphaDegree
                qList.push_back(prime[randIndex]);  // Добавляем простое число в список qList
            }
        }

        if (m > pow(2, bit - 2)) {
            if (m >= pow(2, bit - 1)) {  // Если m слишком большое, начинаем заново
                m = 1;
                qList.clear();
            } else {
                break;
            }
        }
    }

    int n = 2 * m + 1;  // Генерируем число n

    return make_tuple(n, qList);  // Возвращаем n и список простых чисел qList
}

// Тест Миллера для проверки простоты числа
bool Miller(int n, int t, vector<int> qList) {
    set<int> aSet;
    int a;

    while (aSet.size() != t) {
        a = randDist(2, n - 1);  // Генерация случайных оснований
        aSet.insert(a);  // Добавляем основание в множество
    }

    // Проверка для каждого основания
    for (int aj : aSet) {
        if (powMod(aj, n - 1, n) != 1) {
            return false;  // Если условие не выполняется, число составное
        }
    }

    // Проверка для каждого простого делителя из qList
    for (int q : qList) {
        bool isPrime = true;
        for (int aj : aSet) {
            int res = static_cast<int>(pow(aj, (n - 1) / q)) % n;  // Проверка для каждого основания
            if (res != 1) {
                isPrime = false;
                break;
            }
        }
        if (isPrime) {
            return false;  // Если n проходит проверку, оно вероятно составное
        }
    }

    return true;  // Если все проверки прошли успешно, n вероятно простое
}

int main() {
    int bit = 10, t = 5;
    vector<int> prime = sieveEratos(500);  // Строим таблицу малых простых чисел

    vector<int> qList;

    int n;
    int k = 0;
    bool isPrime;

    vector<int> result;
    vector<string> testVer;
    vector<int> kCont;

    // Генерация простых чисел
    while (result.size() < 10) {
        tie(n, qList) = calcN(prime, bit);  // Генерация числа n и списка простых чисел

        if (Miller(n, t, qList)) {  // Если число прошло тест Миллера
            result.push_back(n);

            // Вероятностный тест Рабина
            if (rabin(n, 3)) {
                testVer.push_back("+");
            } else {
                testVer.push_back("-");
            }
            kCont.push_back(k);
            k = 0;
        } else {
            if (rabin(n, 3)) {
                k++;  // Если число не прошло тест Миллера, увеличиваем счетчик k
            }
        }
    }

    // Выводим таблицу результатов
    cout << "\n╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗\n";
    cout << "║Num\t║";
    for (int i = 1; i < 11; i++)
        cout << "\t" << i << "\t║";
    cout << endl;
    cout << "╠═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣";
    cout << endl;
    cout << "║P\t║" << result;
    cout << "║Test\t║" << testVer;
    cout << "║k\t║";
    for (auto item : kCont)
        cout << "\t" << item << "\t║";
    cout << endl;
    cout << "╚═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝";
}
