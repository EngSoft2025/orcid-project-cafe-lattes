'use client';

import { useEffect, useState } from 'react';
import SingleDataCard from './Metrics/SingleDataCard';
import PublicationsByYearChart from './Metrics/PublicationsByYearChart';
import PublicationsByVenueChart from './Metrics/PublicationsByVenueChart';
import { MetricsData } from '../types';

interface MetricsProps {
  metricas: MetricsData | null;
}

export default function Metrics({ metricas }: MetricsProps) {  


  if (!metricas) {
    return (
      <section className="p-6"  id="metricas">
        <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4">
          Metricas
        </h2>
        <p className="text-gray-400 text-center py-12">Nenhuma métrica disponível.</p>
      </section>
    );
  }

  return (
    <section className="bg-background-darker mt-30 text-foreground px-4 sm:px-[5%] lg:px-[15%] py-10 flex flex-col gap-8 pb-50"  id="metricas">
      <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4">
        Metricas
      </h2>

      <div>
        {/* Cards */}
        <div className="grid md:grid-cols-3 gap-6 mb-8">
          <SingleDataCard title='Total de Publicações' value={metricas.publicationsByYear.reduce((acc, cur) => acc + cur.count, 0)}/>
          <SingleDataCard title='Índice H' value={metricas.hIndex}/>
          {/* <SingleDataCard title='Citações' value={metricas.totalCitations}/> */}
        </div>

        {/* Graficos */}
        <div className="grid md:grid-cols-2 gap-6">

          {/* Publications by Year */}
          <PublicationsByYearChart data={metricas.publicationsByYear} />

          {/* Publications by Venue */}
          <PublicationsByVenueChart data={metricas.publicationsByVenue} />
        </div>
      </div>
    </section>
  );
}