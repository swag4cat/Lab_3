#include <iostream>
#include <vector>
#include <climits>

using namespace std;

int main() {
    int n, m;
    cout << "Введите n и m: ";
    cin >> n >> m;

    vector<int> nums(n);
    cout << "Ввод: ";
    for (int i = 0; i < n; ++i) {
        cin >> nums[i];
    }

    // dp[i] — максимальная разница очков, которую может получить текущий игрок, начиная с позиции i
    vector<long long> dp(n + 1, 0);

    // suffix_sum[i] — сумма элементов от i до конца массива
    vector<long long> suffix_sum(n + 1, 0);

    for (int i = n - 1; i >= 0; --i) {
        suffix_sum[i] = suffix_sum[i + 1] + nums[i];
    }

    for (int i = n - 1; i >= 0; --i) {
        long long max_diff = LLONG_MIN;

        // Игрок может взять от 1 до m чисел
        for (int j = 1; j <= m && i + j <= n; ++j) {
            // Сумма взятых чисел: suffix_sum[i] - suffix_sum[i + j]
            // Разница очков: (взятая сумма) - (оптимальный результат противника)
            max_diff = max(max_diff, (suffix_sum[i] - suffix_sum[i + j]) - dp[i + j]);
        }

        dp[i] = max_diff;
    }

    // Если dp[0] > 0, Павел выиграл
    cout << "Результат: " << (dp[0] > 0 ? 1 : 0) << endl;

    return 0;
}
