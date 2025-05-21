// Define o tipo `Publicacao`, que representa a estrutura de uma publicação
export type Publicacao = {
  id: number;       // Identificador único da publicação
  titulo: string;   // Título da publicação
  tipo: string;     // Tipo da publicação (ex.: artigo, livro, etc.)
  ano: number;      // Ano de publicação
  periodico: string; // Nome do periódico ou veículo de publicação
  doi: string;      // DOI (Digital Object Identifier) da publicação
};

// URL base para a API de publicações
const BASE_URL = 'http://localhost:3001';

/**
 * Função para buscar publicações filtradas da API.
 * 
 * @param {Object} params - Parâmetros de filtro e paginação.
 * @param {number} [params.page=1] - Número da página (padrão: 1).
 * @param {number} [params.limit=10] - Limite de itens por página (padrão: 10).
 * @param {string} [params.search] - Texto para busca no título das publicações.
 * @param {number} [params.ano] - Ano de publicação para filtrar.
 * @param {string} [params.tipo] - Tipo de publicação para filtrar.
 * 
 * @returns {Promise<{ data: Publicacao[]; total: number }>} - Retorna uma Promise com os dados das publicações e o total de itens.
 * 
 * @throws {Error} - Lança um erro caso a requisição falhe.
 */
export async function getPublicacoesFiltradas({
  page = 1,
  limit = 10,
  search,
  ano,
  tipo,
}: {
  page?: number;
  limit?: number;
  search?: string;
  ano?: number;
  tipo?: string;
}): Promise<{ data: Publicacao[]; total: number }> {
  // Cria um objeto URLSearchParams para construir os parâmetros da query string
  const params = new URLSearchParams();
  params.append('_page', page.toString()); // Define o número da página
  params.append('_limit', limit.toString()); // Define o limite de itens por página

  // Adiciona o parâmetro de busca no título, se fornecido
  if (search?.trim()) {
    params.append('titulo_like', encodeURIComponent(search.trim())); // Busca por substring no título
  }

  // Adiciona o filtro por ano, se fornecido
  if (ano != null) params.append('ano', ano.toString());

  // Adiciona o filtro por tipo, se fornecido
  if (tipo) params.append('tipo', tipo);

  // Constrói a URL final com os parâmetros
  const url = `${BASE_URL}/publicacoes?${params.toString()}`;

  // Faz a requisição à API usando fetch
  const res = await fetch(url, { cache: 'no-store' }); // `cache: 'no-store'` evita cache de respostas

  // Verifica se a resposta foi bem-sucedida
  if (!res.ok) throw new Error('Erro ao buscar publicações');

  // Obtém o total de itens do cabeçalho da resposta
  const total = Number(res.headers.get('X-Total-Count') || 0);

  // Converte o corpo da resposta para JSON
  const data = await res.json();

  // Retorna os dados das publicações e o total de itens
  return { data, total };
}