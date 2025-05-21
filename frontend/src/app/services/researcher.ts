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

// URL base para a API de dados do pesquisador
const BASE_URL = 'http://localhost:3001';

/**
 * Função para buscar os dados de um pesquisador da API.
 * 
 * @returns {Promise<Pesquisador>} - Retorna uma Promise com os dados do pesquisador.
 * 
 * @throws {Error} - Lança um erro caso a requisição falhe.
 */
export async function getPesquisador(): Promise<Pesquisador> {
  try {
    // Faz a requisição à API para obter os dados do pesquisador
    const res = await fetch(`${BASE_URL}/pesquisador`, {
      cache: 'no-store', // Evita cache para garantir dados atualizados (importante para SSR)
    });

    // Verifica se a resposta foi bem-sucedida
    if (!res.ok) {
      throw new Error(`Erro ao buscar pesquisador: ${res.status}`); // Lança um erro com o código de status
    }

    // Converte o corpo da resposta para JSON e retorna os dados do pesquisador
    return await res.json();
  } catch (error) {
    // Loga o erro no console para depuração
    console.error(error);

    // Propaga o erro para ser tratado pelo chamador
    throw error;
  }
}