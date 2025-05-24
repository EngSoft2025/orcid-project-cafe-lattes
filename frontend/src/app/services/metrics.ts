export type YearMetric = {
  year: number;
  count: number;
};

export type VenueMetric = {
  name: string;
  count: number;
};

export type MetricsData = {
  publicationsByYear: YearMetric[];
  hIndex: number;
  totalCitations: number;
  publicationsByVenue: VenueMetric[];
};

const BASE_URL = 'http://localhost:3001'; // ajuste se usar outra porta

export async function getMetrics(): Promise<MetricsData> {
  const response = await fetch(`${BASE_URL}/metrics`, { cache: 'no-store' });

  if (!response.ok) {
    throw new Error('Falha ao buscar métricas');
  }

  const data = await response.json();

  // Validação básica (opcional)
  if (
    !Array.isArray(data.publicationsByYear) ||
    typeof data.hIndex !== 'number' ||
    typeof data.totalCitations !== 'number' ||
    !Array.isArray(data.publicationsByVenue)
  ) {
    throw new Error('Formato de dados inválido');
  }

  return data as MetricsData;
}
