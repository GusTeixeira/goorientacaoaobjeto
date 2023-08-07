<h1>Go (Golang) - Orientação a Objetos</h1>
<p>Este é um repositório destinado aos meus estudos da linguagem de programação Go (também conhecida como Golang) e sua aplicação em conceitos de Orientação a Objetos. Criada pelos engenheiros do Google em 2009, o Go é uma linguagem moderna, eficiente e concisa, projetada para oferecer simplicidade e alto desempenho.</p>
<h2>Orientação a Objetos em Go</h2>
<p>Ao contrário de outras linguagens de programação populares, como Java ou Python, o Go não possui suporte direto para conceitos de orientação a objetos tradicionais, como classes e herança. No entanto, isso não significa que Go não suporte programação orientada a objetos. Em vez disso, a linguagem adota uma abordagem diferente para alcançar os mesmos princípios da orientação a objetos, focando na composição em vez da herança.</p>
<h3>Tipos e Métodos</h3>
<p>Em Go, é possível definir tipos personalizados que podem conter campos e métodos associados a esses tipos. Os métodos são funções que operam nos dados do tipo definido e podem ser associados a um tipo específico usando uma sintaxe especial. Isso nos permite criar estruturas de dados com comportamentos específicos e funcionalidades encapsuladas.</p>
<pre><code class="language-go">type Pessoa struct {
    Nome string
    Idade int
}

// Método associado ao tipo Pessoa
func (p *Pessoa) FazerAniversario() {
    p.Idade++
}
</code></pre>
<h3>Composição</h3>
<p>Em vez de usar herança, Go utiliza a composição para reutilizar código e compartilhar comportamentos entre tipos. A composição é alcançada por meio da incorporação de um tipo dentro de outro, o que permite que o tipo incorporado herde os campos e métodos do tipo externo.</p>
<pre><code class="language-go">type Animal struct {
    Nome string
}

type Cachorro struct {
    Animal  // Incorporação do tipo Animal
    Raca string
}

func main() {
    c := Cachorro{
        Animal: Animal{Nome: "Bobby"},
        Raca:   "Labrador",
    }
    fmt.Println(c.Nome) // Acessando o campo Nome do tipo Animal através do tipo Cachorro
}
</code></pre>
<h3>Interfaces</h3>
<p>Go possui interfaces, que são coleções de métodos abstratos. Uma interface define um conjunto de comportamentos esperados, e qualquer tipo que implemente esses métodos é considerado uma implementação válida da interface.</p>
<pre><code class="language-go">type FormaGeometrica interface {
    Area() float64
}

type Retangulo struct {
    Base   float64
    Altura float64
}

func (r Retangulo) Area() float64 {
    return r.Base * r.Altura
}
</code></pre>
<h3>Polimorfismo</h3>
<p>Graças às interfaces, Go permite o polimorfismo, onde diferentes tipos podem ser tratados de forma homogênea, desde que implementem os métodos da interface em questão. Isso proporciona maior flexibilidade na escrita de código e torna o sistema mais extensível.</p>
<pre><code class="language-go">func CalcularArea(forma FormaGeometrica) float64 {
    return forma.Area()
}

func main() {
    retangulo := Retangulo{Base: 5, Altura: 3}
    area := CalcularArea(retangulo)
    fmt.Println("Área do retângulo:", area)
}
</code></pre>
<h2>Considerações Finais</h2>
<p>Go pode não seguir a abordagem tradicional de orientação a objetos, mas oferece uma maneira elegante e eficaz de alcançar os mesmos princípios usando tipos, métodos, composição e interfaces. A linguagem valoriza a simplicidade e o desempenho, tornando-a uma excelente escolha para o desenvolvimento de aplicativos modernos e eficientes.</p>