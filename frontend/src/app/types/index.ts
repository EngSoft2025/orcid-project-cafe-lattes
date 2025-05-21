// Define o tipo `Publicacao`, que representa a estrutura de uma publicação
export type Publicacao = {
  id: number;       // Identificador único da publicação
  titulo: string;   // Título da publicação
  tipo: string;     // Tipo da publicação (ex.: artigo, livro, etc.)
  ano: number;      // Ano de publicação
  periodico: string; // Nome do periódico ou veículo de publicação
  doi: string;      // DOI (Digital Object Identifier) da publicação
};

// Define o tipo `Afiliacao`, que representa uma afiliação de um pesquisador
export type Afiliacao = {
  instituicao: string; // Nome da instituição de afiliação
  descricao: string;   // Descrição da afiliação (ex.: cargo ou departamento)
};

// Define o tipo `Pesquisador`, que representa os dados de um pesquisador
export type Pesquisador = {
  id: string;              // Identificador único do pesquisador
  nome: string;            // Nome completo do pesquisador
  areaPrincipal: string;   // Área principal de atuação do pesquisador
  foto: string;            // URL da foto do pesquisador
  afiliacoes: Afiliacao[]; // Lista de afiliações do pesquisador
  palavrasChave: string[]; // Palavras-chave relacionadas ao pesquisador
};