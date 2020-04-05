(define (move item lst)
  (if (null? lst)
    (list item)
    (if (= item 0)
      (append (move (car lst) (cdr lst)) '(0))
      (cons item (move (car lst) (cdr lst))))))
 
(define (MoveZeros lst)
  (if (null? lst)
    lst
    (move (car lst) (cdr lst))))
 
(MoveZeros '(0 1 4 0 2 4 0 5 1 3 0 2 3 1 7 8))