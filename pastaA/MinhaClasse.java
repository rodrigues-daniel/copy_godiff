public class MinhaClasse {

    private String nome;
    private int idade;

    public MinhaClasse(String nome, int idade) {
        this.nome = nome;
        this.idade = idade;
    }

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }

    public int getIdade() {
        return idade;
    }

    public void setIdade(int idade) {
        this.idade = idade;
    }

    public void imprimirDados() {
        System.out.println("Nome: " + nome);
        System.out.println("Idade: " + idade);
    }

    public static void main(String[] args) {
        MinhaClasse objeto = new MinhaClasse("Fulano", 30);
        objeto.imprimirDados();
    }
}
