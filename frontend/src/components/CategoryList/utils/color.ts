import type { Color } from "../../../types"

/**
 * 色の配列
 */
export const colorList: Color[] = ["blue", "red", "orange", "yellow", "green", "purple"]

const colorMap: Record<Color, string> = {
  blue: "bg-blue-800",
  red: "bg-red-800",
  orange: "bg-orange-800",
  yellow: "bg-yellow-800",
  green: "bg-green-800",
  purple: "bg-purple-800",
}

/**
 * 色の文字列(blue等)をtailwindcssの背景用(bg-blue-800等)に書き換える関数
 * @param color 色のテキスト(blue等)
 * @return tailwindcssのテキスト(bg-blue-800等)
 */
export const convertColorText = (color: Color): string => {
  return colorMap[color] || "bg-gray-800"
}
