# 计算机组成原理笔记

[toc]

## 介绍

本笔记是计算机组成原理的复习笔记。使用的书主要是：坂井修一 的 ＜コンピュータ　アーキテクチャ＞

## はじめに

### 理解度の確認

問1

半加算器（１ビット）

$$S=\overline{X}\cdot Y+X\cdot\overline{Y}=X\oplus Y$$

$$C_{out}=X\cdot Y$$


| $X$   | $Y$       | $S$      | $C_{out}$ |
| ----- | --------- | -------- | --------- |
| 0 | 0 | 0 | 0 |
| 0 | 1   | 1 | 0 |
| 1 | 0 | 1 | 0 |
| 1 | 1 | 0 | 1 |

全加算器（１ビット）

$$S=X\cdot\overline{Y}\cdot \overline{C_{in}}+\overline{X}\cdot Y\cdot \overline{C_{in}}+\overline{X}\cdot\overline{Y}\cdot C_{in}+X\cdot Y\cdot C_{in}=X\oplus Y\oplus C_{in}$$

$$C_{out}=X\cdot Y+Y\cdot C_{in}+X\cdot C_{in}$$

| $X$  | $Y$  | $C_{in}$ | $S$  | $C_{out}$ |
| ---- | ---- | -------- | ---- | --------- |
| 0    | 0    | 0        | 0    | 0         |
| 0    | 0    | 1        | 1    | 0         |
| 0    | 1    | 0        | 1    | 0         |
| 0    | 1    | 1        | 0    | 1         |
| 1    | 0    | 0        | 1    | 0         |
| 1    | 0    | 1        | 0    | 1         |
| 1    | 1    | 0        | 0    | 1         |
| 1    | 1    | 1        | 1    | 1         |

問2

$Q\oplus 1=\overline{Q}$并且$Q\oplus 0=Q$

$S/\overline{A}=0$时，$C_{in}=0$这时是加法器。

$S/\overline{A}=1$时，$C_{in}=1$，Y变成了Y反，所以是减法器。

問3

1. 当clock=0时，P2、P3都是1，G5和G6各自保持原来的状态，也就是与输入无关，输出Q不发生变化。
2. 当clock=1时，clock的上升沿到来前，P2、P3仍然保持1，P4是d反，P1是d。当clock上升沿到来后，P2变成d反，P3变成d，于是Q=d，Q反=d反

問4

因为ALU需要完成两组数据的运算，如果只用1组寄存器，无法为ALU提供两组数据。
