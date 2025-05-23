'use client';

import { useEffect, useState } from 'react';
import {
  BarChart,
  Bar,
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
  PieChart,
  Pie,
  Cell,
  Legend,
} from 'recharts';
import { getMetrics, MetricsData } from '../services/metrics'; // Importa o serviço e os tipos
import SingleDataCard from './Metrics/SingleDataCard';
import PublicationsByYearChart from './Metrics/PublicationsByYearChart';
import PublicationsByVenueChart from './Metrics/PublicationsByVenueChart';

export default function Metrics() {
  const [metrics, setMetrics] = useState<MetricsData | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchMetrics() {
      setLoading(true);
      try {
        const data = await getMetrics(); // Usa o serviço para buscar as métricas
        setMetrics(data);
      } catch (error) {
        console.error('Erro ao buscar métricas:', error);
        setMetrics(null);
      } finally {
        setLoading(false);
      }
    }
    fetchMetrics();
  }, []);

  if (loading) {
    return (
      <section className="p-6">
        <h2 className="text-white text-2xl font-semibold mb-6">Métricas</h2>

        <div className="grid md:grid-cols-3 gap-6 mb-8">
          {[...Array(3)].map((_, i) => (
            <div key={i} className="bg-gray-700 rounded-lg h-24 animate-pulse" />
          ))}
        </div>

        <div className="grid md:grid-cols-2 gap-6">
          {[...Array(2)].map((_, i) => (
            <div key={i} className="bg-gray-700 rounded-lg h-80 animate-pulse" />
          ))}
        </div>
      </section>
    );
  }

  if (!metrics) {
    return (
      <section className="p-6">
        <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4">
          Metricas
        </h2>
        <p className="text-gray-400 text-center py-12">Nenhuma métrica disponível.</p>
      </section>
    );
  }

  return (
    <section className="bg-background-darker mt-30 text-foreground px-4 sm:px-[5%] lg:px-[15%] py-10 flex flex-col gap-8 pb-50">
      <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4">
        Metricas
      </h2>

      <div>
        {/* Cards */}
        <div className="grid md:grid-cols-3 gap-6 mb-8">
          <SingleDataCard title='Total de Publicações' value={metrics.publicationsByYear.reduce((acc, cur) => acc + cur.count, 0)}/>
          <SingleDataCard title='Índice H' value={metrics.hIndex}/>
          <SingleDataCard title='Citações' value={metrics.totalCitations}/>
        </div>

        {/* Graficos */}
        <div className="grid md:grid-cols-2 gap-6">

          {/* Publications by Year */}
          <PublicationsByYearChart data={metrics.publicationsByYear} />

          {/* Publications by Venue */}
          <PublicationsByVenueChart data={metrics.publicationsByVenue} />
        </div>
      </div>
    </section>
  );
}