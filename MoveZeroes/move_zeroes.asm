.globl _start
 
.data
 
nums:
        .long   0
        .long   1
        .long   4
        .long   3
        .long   0
        .long   2
        .long   0
 
.text
 
_start:
        pushl   %ebp
        movl    %esp, %ebp
        # moveZeroes(nums, 7)
        pushl   $7
        pushl   $nums
        call    moveZeroes
        addl    $8, %esp
        # finaliza
        movl    $0, %eax
        int     $0x80
 
moveZeroes:
        pushl   %ebp
        movl    %esp, %ebp
        subl    $16, %esp
        # Argumentos
        #    12(%ebp) = nums = endereço do início de um vetor de long
        #    8(%ebp) = n    = quantidade de elementos long a partir de nums
        # Variáveis locais:
        #    i = posição atual válida: invariante aplica-se para subconjunto 0..{i-1}
        #    j = posição de busca futura: posição de leitura da fita de entrada
        movl    $0, -4(%ebp) # i = 0
        movl    $0, -8(%ebp) # j = 0
        jmp     .verifica_entrada_disponivel
 
.ler_valor:
        # Temos que j < n, portanto devemos movimentar nums[j] para nums[i] afim
        # de trazer o valor lido atualmente nums[j] para ser candidato de uma
        # posição válida (será verificado a seguir).
        movl    -4(%ebp), %eax
        leal    0(,%eax,4), %edx
        movl    8(%ebp), %eax
        addl    %eax, %edx
        movl    -8(%ebp), %eax
        leal    0(,%eax,4), %ecx
        movl    8(%ebp), %eax
        addl    %ecx, %eax
        movl    (%eax), %eax
        movl    %eax, (%edx)
        # Move j para o próximo valor a ser lido, se tiver entrada disponível.
        addl    $1, -8(%ebp)
        # Se nums[i] != 0, significa que os elementos 0..i são válidos, e assim
        # podemos deslocar i uma posição a frente.
        movl    -4(%ebp), %eax
        leal    0(,%eax,4), %edx
        movl    8(%ebp), %eax
        addl    %edx, %eax
        movl    (%eax), %eax
        testl   %eax, %eax
        setne   %al
        movzbl  %al, %eax
        addl    %eax, -4(%ebp)
        # Aqui temos garantia da invariante de laço:
        # Todos elementos 0..i-1 são diferentes de zero garantindo a ordem
        # relativa dos elementos não-zero da entrada. Podemos dar continuidade
        # na leitura do próximo valor de entrada.
 
.verifica_entrada_disponivel:
        # Se j < n, temos elementos que ainda não foram lidos da entrada.
        movl    -8(%ebp), %eax
        cmpl    12(%ebp), %eax
        jl      .ler_valor
        jmp     .verifica_elemento_a_ser_zerado
 
.zera_valor_ao_final:
        # Zera nums[j] e move j para próximo elemento.
        movl    8(%ebp), %eax
        movl    $0, (%eax)
        addl    $1, -4(%ebp)
 
.verifica_elemento_a_ser_zerado:
        # Se i < n, temos elementos que devem ser zerados no fim da fita.
        # Ao chegar aqui, temos que os elementos 0..i-1 são válidos de acordo com
        # a invariante descrita acima. Os elementos i..n devem ser definidos para
        # zero, pois nós deslocamos o ponteiro j exatamente (n-i) mais vezes, que
        # é exatamente a quantidade de zeros da lista de entrada.  Assim, torna-se
        # trivial a tarefa de zerar todos os elementos restantes.
        movl    -4(%ebp), %eax
        cmpl    12(%ebp), %eax
        jl      .zera_valor_ao_final
        nop
        leave
        ret