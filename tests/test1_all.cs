using System;

/* Пример №1: покрывает объявления, присваивания, if/else, while, do-while, for,
   массивы, goto+метка, вызовы функций, строки/числа (в т.ч. экспонента). */

namespace Demo {
    class Program {

        static int Sum(int a, int b) {
            int res;
            res = a + b;
            return res;
        }

        static void Main(string[] args) {
            int i;
            int n = 5;
            float x;
            string msg;
            bool flag;

            // числа с плавающей точкой и экспонентой
            x = 1.25e-2;

            // массив
            int[] arr = new int[10];

            // заполнение массива (while)
            i = 0;
            while (i < 10) {
                arr[i] = i * 2;
                i = i + 1;
            }

            // do-while
            do {
                n = n - 1;
            } while (n > 0);

            // if/else + goto + метка (внутри for)
            for (i = 0; i <= 3; i = i + 1) {
                if (arr[i] != 0) {
                    Console.WriteLine(arr[i]);
                } else {
                    goto END_LABEL;
                }
            }

            // строки и вызов пользовательской функции
            msg = "sum=";
            Console.WriteLine(msg);
            Console.WriteLine(Sum(2, 3));

            flag = (2 < 3) && (3 != 4) || false;
            if (flag) {
                Console.WriteLine("flag=true");
            }

            END_LABEL:
            Console.WriteLine("done");
        }
    }
}
