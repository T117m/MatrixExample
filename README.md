# MatrixExample

Простая рабочая имплементация примера использования конкурентности из бумаги ["Google Go!"](https://doc.cat-v.org/programming/go/papers/Paper__aigner_baumgartner.pdf) 2010 года за авторством Martin Aigner и Alexander Baumgartner (мне лень имена переводить), на которую я наткнулся луркая по [cat-v.org](https://cat-v.org). Сам пример на странице 23 выглядел вот так:  
```go
func InverseProduct (a Matrix, b Matrix) {
    a_inv_future := InverseFuture(a);
    b_inv_future := InverseFuture(b);
    a_inv := <−a_inv_future;
    b_inv := <−b_inv_future;
    returnProduct(a_inv, b_inv);
}
func InverseFuture (a Matrix) {
    future := make (chan Matrix);
    go func () { future <− Inverse(a) }();
    returnfuture;
}
```
