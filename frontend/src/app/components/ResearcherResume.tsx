'use client';

import { useEffect, useState } from 'react';
import { getPesquisador } from '@/app/services/researcher';
import { Pesquisador} from "@/app/types"
import { CheckCircle } from 'lucide-react';
import Image from 'next/image';

export default function ResearcherResume () {
  const [pesquisador, setPesquisador] = useState<Pesquisador | null>(null);

  useEffect(() => {
    getPesquisador().then(setPesquisador).catch(console.error);
  }, []);

  if (!pesquisador) return <p className="text-clara-gray">Carregando...</p>;

  return (
    <section className="bg-background-darker text-foreground px-4 sm:px-[5%] lg:px-[15%] py-10 flex flex-col gap-8">
      {/* Título */}
      <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4">
        Resumo do Pesquisador
      </h2>

      <div className="flex flex-col md:flex-row gap-6">
        {/* Lado esquerdo - imagem + ORCID */}
        <div className="flex flex-col items-center md:items-start gap-3">
          <div className="relative">
            <Image
              src={pesquisador.foto}
              alt={`Foto de ${pesquisador.nome}`}
              width={130}
              height={130}
              className="rounded-full border-4 border-primary-yellow"
            />
            <CheckCircle className="absolute bottom-0 right-0 bg-background-darker rounded-full text-green-400" size={24} />
          </div>
          <div className="flex items-center justify-center gap-2 text-sm text-primary-green font-mono">
            <Image
              src="/images/orcidImage.png"
              alt="orcidImage"
              width={15}
              height={15}
              className='select-none'
            />
            {pesquisador.id}
          </div>
        </div>

        {/* Lado direito - dados */}
        <div className="flex-1 flex flex-col gap-4">
          <h3 className="text-2xl font-bold">{pesquisador.nome}</h3>

          <p>
            <span className="font-semibold">Área principal:</span>{' '}
            {pesquisador.areaPrincipal}
          </p>

          <div>
            <p className="font-semibold mb-2">Afiliações:</p>
            {pesquisador.afiliacoes.map((af, index) => (
              <div
                key={index}
                className="bg-clara-gray-dark p-4 rounded-md text-sm text-clara-gray-light mb-2"
              >
                <p className="text-foreground font-semibold">
                  {af.instituicao}
                </p>
                <p>{af.descricao}</p>
              </div>
            ))}
          </div>

          <div>
            <p className="font-semibold mb-2">Palavras-chave:</p>
            <div className="flex flex-wrap gap-2">
              {pesquisador.palavrasChave.map((palavra, i) => (
                <span
                  key={i}
                  className="bg-primary-yellow/20 text-primary-yellow text-xs font-semibold px-3 py-1 rounded-full"
                >
                  {palavra}
                </span>
              ))}
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
