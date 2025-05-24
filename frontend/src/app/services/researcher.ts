// Importa o tipo `Pesquisador` do arquivo types.ts
import { Pesquisador } from '@/app/types';


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