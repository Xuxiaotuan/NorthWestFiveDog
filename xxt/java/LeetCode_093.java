/**
 * 复原 IP 地址
 * 题目描述：给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
 *
 * 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
 *
 * 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
 *
 * 示例说明请见LeetCode官网。
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/restore-ip-addresses/
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

import java.util.ArrayList;
import java.util.List;

public class scratch {
    // IP地址一共有4段
    private static final int SEG_COUNT = 4;
    // IP地址4段的值分别是多少
    private static int[] segments;
    // 结果集
    private static List<String> result = new ArrayList<>();

    public static List<String> restoreIpAddresses(String s) {
        segments = new int[SEG_COUNT];
        dfs(s, 0, 0);
        return result;
    }

    /**
     * 回溯算法
     *
     * @param s        原始字符串
     * @param segId    当前是第几段IP地址
     * @param segStart 当前第segId段IP地址的起始位置
     */
    private static void dfs(String s, int segId, int segStart) {
        // 如果找到了4段IP地址并且遍历完了字符串，则说明当前结果是一种符合条件的IP地址
        if (segId == SEG_COUNT) {
            if (segStart == s.length()) {
                StringBuffer ipAddr = new StringBuffer();
                for (int i = 0; i < SEG_COUNT; ++i) {
                    ipAddr.append(segments[i]);
                    if (i != SEG_COUNT - 1) {
                        ipAddr.append('.');
                    }
                }
                result.add(ipAddr.toString());
            }
            return;
        }

        // 如果还没有找到4段IP地址，字符串就已经遍历完了，则提前退出
        if (segStart == s.length()) {
            return;
        }

        // 由于不能有前导0，如果当前数字为0，那么这一段IP地址只能是0
        if (s.charAt(segStart) == '0') {
            segments[segId] = 0;
            dfs(s, segId + 1, segStart + 1);
        }

        // 一般情况下，每一种可能都要考虑
        int addr = 0;
        for (int segEnd = segStart; segEnd < s.length(); ++segEnd) {
            addr = addr * 10 + (s.charAt(segEnd) - '0');
            if (addr > 0 && addr <= 0xFF) {
                segments[segId] = addr;
                dfs(s, segId + 1, segEnd + 1);
            } else {
                break;
            }
        }
    }

    public static void main(String[] args) {
        for (String restoreIpAddress : restoreIpAddresses("101023")) {
            System.out.println(restoreIpAddress);
        }
    }
}