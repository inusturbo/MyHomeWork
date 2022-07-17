# NagoyaU Programing

## 令和４年度入学試験　2021（令和３）年8月5日

```c
#include <stdio.h>
#define True 1
#define False 0
#define N 5
#define SIZE 10
#define PERIOD 10

typedef struct {
    int x, y, stepx, stepy, infected, source;
} Person;

Person persons[N];
int map[PERIOD][SIZE][SIZE];

Person initialize(int state[5]) {
    Person person;
    person.x = state[0];
    person.y = state[1];
    person.stepx = state[2];
    person.stepy = state[3];
    person.infected = state[4];
    person.source = -1;

    return person;
}

void print_map(int t) {
    int i, x, y;

    for (y = 0; y < SIZE; y++) {
        for (x = 0; x < SIZE; x++) {
            i = map[t][x][y];
            if (i >= 0) {
                printf("%d", i);
                if (persons[i].infected)
                    printf("+");
                else
                    printf(" ");
            }
            else
                printf(". ");
        }
        printf("\n");
    }
}

int update(int d, int step) {
    d = d + step;
    if (d < 0) d = 0;
    if (d >= SIZE)
        d = SIZE - 1;

    return d;
}

void infect(int t) {
    int i, x, y, xx, yy, n;
    for (i = 0; i < N; i++) {
        if (!persons[i].infected) continue;
        for (xx = -1; xx <= 1; xx++)
            for (yy = -1; yy <=1; yy++) {
                x = update(persons[i].x, xx);
                y = update(persons[i].y, yy);
                n = map[t][x][y];
                if (n >= 0 && !persons[n].infected) {
                    persons[n].infected = True;
                    persons[n].source = i;
                }
            }
    }
}

void move(int t) {
    int i, x, y;
    for (i = 0; i < N; i++) {
        x = update(persons[i].x, persons[i].stepx);
        y = update(persons[i].y, persons[i].stepy);

        if (map[t][x][y] >= 0) {
            x = persons[i].x;
            y = persons[i].y;
        }

        if (persons[i].x == x && persons[i].y == y) {
            persons[i].stepx = -persons[i].stepx;
            persons[i].stepy = -persons[i].stepy;
        }
        map[t][x][y] = i;
        persons[i].x = x;
        persons[i].y = y;
        printf("%d: %d(%d, %d) (%+d, %+d)\n" , t, i, x, y, persons[i].stepx, persons[i].stepy);
    }
}

void print_path(int i) {
    if (persons[i].source > 0) {
        print_path(persons[i].source);
        printf(" -> %d", i);
    } else
        printf("%d", i);
}

int main() {
    int initial_states[N][5] = {
            {0, 0, +1, +1, False},
            {1, 2, 0, +1, True},
            {2, 4, -1, -1, False},
            {3, 5, 0, +1, True},
            {4, 3, 0, +1, False}
    };
    int i, t, x, y;

    for (t = 0; t < PERIOD; t++)
        for (x = 0; x < SIZE; x++)
            for (y = 0; y < SIZE; y++)
                map[t][x][y] = -1;

    for (i = 0; i < N; i++) {
        persons[i] = initialize(initial_states[i]);
        map[0][persons[i].x][persons[i].y] = i;
    }
    printf("time 0\n");
    print_map(0);
    for (t = 0; t < PERIOD - 1; t++) {
        infect(t);
        move(t + 1);
        printf("time %d\n", t + 1);
        print_map(t + 1);
    }
    printf("infection path\n");
    print_path(0);
    printf("\n");


    return 0;
}
```

解答：

（１）A`SIZE - 1` B`i` C`print_path(persons[i] - source)`

（２）(b)

（３）SIZE$\times$SIZEの範囲外の状況で、座標を直す

（４）

```
9: 0(9, 9) (+1, +1)
9: 1(1, 8) (+0, -1)
9: 2(9, 9) (-1, -1)
9: 3(3, 5) (+0, -1)
9: 4(4, 9) (+0, -1)
```

（５）0と2は同じ位置に移動しますが、2の方が後の数字なので、0がある位置は2で上書きされます。

## 令和２年度入学試験　2019（令和元）年8月7日

```c
#include <stdio.h>
#define MAXSIZE     1000
#define P           17
#define SENTINEL    -1
#define NOTFOUND    -2
int hfirst[P];          //hfirst[i]是哈希值h的第一个元素 hfirst[i]为空时表示没有这个元素
int element[MAXSIZE];   //存放正整数
int next[MAXSIZE];      //存放列表中下一个元素
int avail = -1;
int maxnode = 0;
int hashfunc(int data) {
    return data % P;
}
int search(int h, int data) {
    int pred = -1;
    if (hfirst[h] == -1) return NOTFOUND;
    if (element[hfirst[h]] == data) return pred;
    pred = hfirst[h];
    while (next[pred] != SENTINEL) {
        if (element[next[pred]] == data) return pred;
        pred = next[pred];
    }
    return NOTFOUND;
}

void insert(int h, int data) {
    int u;
    if (avail != -1) {
        u = avail;
        avail = next[avail];
    } else {
        u = maxnode;
        maxnode = maxnode + 1;
    }
    element[u] = data;
    next[u] = hfirst[h];
    hfirst[h] = u;
}

void delete(int h, int pred) {
    int u;
    if (pred != -1) {
        u = next[pred];
        next[pred] = next[u];
    } else {
        u = hfirst[h];
        hfirst[h] = hfirst[u];
    }
    next[u] = avail; avail = u;
}

void initarrays() {
    int i;
    for (i = 0; i < P; i++) hfirst[i] = -1;
    for (i = 0; i < MAXSIZE; i++) element[i] = 0;
    for (i = 0; i < MAXSIZE; i++) next[i] = SENTINEL;
}

void outputarrays(int maxnode) {
    int i;
    printf("hfirst\n");
    for (i = 0; i < P; i++) printf("%d, ", hfirst[i]);
    printf("\n");
    printf("element\n");
    for (i = 0; i < maxnode; i++) printf("%d, ", element[i]);
    printf("\n");
    printf("next\n");
    for (i = 0; i < maxnode; i++) printf("%d, ", next[i]);
    printf("\n");
}

int main() {
    int data[] = {1, 2, 18, 19, 20};
    int h, i, pred;
    initarrays();
    for (i = 0; i < 5; i++) {
        h = hashfunc(data[i]);
        if (search(h, data[i])==NOTFOUND) insert(h, data[i]);
    }
    outputarrays(maxnode);
    h = hashfunc(1);
    pred = search(h, 1);
    if (pred != NOTFOUND) delete(h, pred);
    outputarrays(maxnode);
    return 0;
}

```

解答：

（１）

```
-1, 0, 1, 2, -1, 3, -1, 4, -1, -1, -1, -1, -1, -1, -1, -1, -1, 
1, 2, 3, 5, 7, 
-1, -1, -1, -1, -1, 
```

（２）

```
-1, 2, 3, 4, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 
1, 2, 18, 19, 20, 
-1, -1, 0, 1, -1, 
```



（３）dataのhash値が多くの値で同じになる場合

（４）(ア)`next[u]` (イ)`hfirst[u]`

（５）

```
-1, -1, 1, 2, -1, 3, -1, 4, -1, -1, -1, -1, -1, -1, -1, -1, -1, 
1, 2, 3, 5, 7, 
-1, -1, -1, -1, -1,
```

## 平成31年度入学試験　2018（平成30）年8月8日

```c
#include <stdlib.h>

char* index(char* string, int p) {
    char *tmp;
    int i;
    i = 1;
    tmp = string;

    while ((*tmp !='\0') && (i < p)) {
        i++;
        if ((*tmp & 0x80) == 0)
            tmp++;
        else
            tmp += 3;
    }
    return tmp;
}

int length_b(char* string) {
    int len;
    len = 0;
    while (string[len] != '\0')
        len++;
    return len;
}

char* concat(char* string1, char* string2) {
    char *result, *tmp;
    result = (char*) malloc(sizeof(char) * (length_b(string1) + length_b(string2) + 1));
    tmp = result;
    while (*string1 != '\0') {
        *tmp = *string1;
        tmp++;
        string1++;
    }
    while (*string2 != '\0') {
        *tmp = *string2;
        tmp++;
        string2++;
    }
    *tmp = '\0';
    return result;
}

char* concat_p(char* string1, char* string2, char* l) {
    char *result, *tmp1, *tmp2;
    tmp1 = string1;
    result = (char*) malloc(sizeof(char) * 100);
    tmp2 = result;
    while ((tmp1 != l) && (tmp1 !='\0')) {
        *tmp2 = *tmp1;
        tmp1++;
        tmp2++;
    }
    *tmp2 = '\0';
    return concat(result, string2);
}

char* replace(char* string1, char* string2, int p) {
    char *tmp1, *tmp2;
    tmp1 = index(string1, p);
    tmp2 = concat_p(string1, string2, tmp1);
    if ((*tmp2 & 0x80) == 0)
        tmp2++;
    else
        tmp2 += 3;
    return concat(tmp2, tmp1);
}

int main() {
    char str1[4] = {0x41, 0x42, 0x43, '\0'};
    char str2[10] = {0xE3, 0x81, 0x82, 0xE3, 0x81,0x84,0xE3,0x81,0x86, '\0'};
    char str3[9] = {0xE3, 0x81, 0x82, 0x41,0xE3,0x81,0x84,0x42,'\0'};
    char str4[4] = {0xE3, 0x81, 0x86, '\0'};
    char *r;
    r = index(str1, 1);
    r = index(str2, 2);
    r = index(str3, 3);
    r = replace(str3, str4, 3);
    r = str1;
    return 0;
}
```



解答：

（１）A`*tmp` B`*string1` C`*tmp` D`*string2` E`'\0'`

（２）


| アドレス | r    | r+1  | r+2  | r+3  |
| -------- | ---- | ---- | ---- | ---- |
| 値      | 0x41 | 0x42 | 0x43 | '\0' |

（３）


| アドレス | r    | r+1  | r+2  | r+3  | r+4  | r+5  | r+6  |
| -------- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| 値      | 0xE3 | 0x81 | 0x84 | 0xE3 | 0x81 | 0x86 | '\0' |

（４）

| アドレス | r    | r+1  | r+2  | r+3  | r+4  |
| -------- | ---- | ---- | ---- | ---- | ---- |
| 値      | 0xE3 | 0x81 | 0x84 | 0x42 | '\0' |

（５）



（６）



（７）

## 平成30年度入学試験　2017（平成20）年8月3日

```c
#include <stdio.h>
#include <stdlib.h>

struct cell {
    int num;
    struct cell *next;
};

typedef struct cell * CELL;
CELL head = NULL;

void insert(int i) {
    CELL c = (CELL) malloc(sizeof(struct cell));
    CELL tmp = head;
    c->num = i; c->next = NULL;

    if (head != NULL) {
        while (tmp->next != NULL) tmp = tmp->next;
        tmp->next = c;
    } else {
        head = c;
    }
}

int top() {
    if (head != NULL) {
        return head->num;
    } else {
        return -1;
    }
}

void eliminate() {
    if (head != NULL) {
        head = head->next;
    }
}

void display() {
    CELL tmp = head;
    while (tmp != NULL) {
        printf("%d;" ,tmp->num);
        tmp = tmp->next;
    }
    printf("\n");
}

int main() {
    insert(0); insert(4); insert(9); insert(3);
    display();
    printf("%d\n", top());
    eliminate(); eliminate(); insert(7); insert(2);
    display();
    return 0;
}
```

（１）

```
0;4;9;3;
0
9;3;7;2;
```

（２）(b)