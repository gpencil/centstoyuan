package centstoyuan

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ToString 将分（int64）转换为元字符串（保留2位小数）
// 例如：
//   - 100 -> "1.00"
//   - 150 -> "1.50"
//   - 1 -> "0.01"
//   - 0 -> "0.00"
func ToString(cents int64) string {
	yuan := float64(cents) / 100.0
	return fmt.Sprintf("%.2f", yuan)
}

// ToCents 将元字符串转换为分（int64）
// 支持格式：
//   - "1" -> 100分
//   - "1.5" -> 150分
//   - "1.50" -> 150分
//   - "0.01" -> 1分
func ToCents(amountStr string) (int64, error) {
	// 去除首尾空格
	amountStr = strings.TrimSpace(amountStr)

	// 验证金额格式（只允许数字和最多一个小数点）
	if amountStr == "" {
		return 0, fmt.Errorf("金额不能为空")
	}

	// 解析为浮点数
	amountFloat, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, fmt.Errorf("金额格式错误,amountStr:%v,err: %w", amountStr, err)
	}

	// 金额不能为负数
	if amountFloat < 0 {
		return 0, fmt.Errorf("金额不能为负数")
	}

	// 转换为分（乘以100并四舍五入）
	// 使用 math.Round 避免浮点数精度问题
	amountInCents := int64(math.Round(amountFloat * 100))

	// 验证小数位数不超过2位（通过反向计算验证）
	// 例如：1.234 * 100 = 123.4，round后是123，123/100=1.23，与原值1.234不同
	reconstructed := float64(amountInCents) / 100
	diff := math.Abs(amountFloat - reconstructed)
	if diff > 0.001 { // 允许0.001的误差（处理浮点数精度问题）
		return 0, fmt.Errorf("金额最多支持2位小数")
	}

	return amountInCents, nil
}
