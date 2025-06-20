import { CheckCircle } from 'lucide-react';
import Image from 'next/image';


import { Pesquisador, Afiliacao } from '@/app/types';

interface ResearcherResumeProps {
  pesquisador: Pesquisador | null;
}

export default function ResearcherResume({ pesquisador }: ResearcherResumeProps) {  

  if (!pesquisador)
    return (
      <section
        id="resumo"
        className="bg-background-darker text-foreground px-4 sm:px-[5%] lg:px-[15%] py-10 flex flex-col gap-8"
      >
        {/* Título Skeleton */}
        <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4 animate-pulse">
          Resumo do Pesquisador
        </h2>

        <div className="flex flex-col md:flex-row gap-6">
          {/* Lado esquerdo - imagem e ORCID Skeleton */}
          <div className="flex flex-col items-center md:items-start gap-3">
            <div className="w-[130px] h-[130px] rounded-full border-4 border-primary-yellow animate-pulse" />
            <div className="w-40 h-4 bg-gray-700 rounded animate-pulse" />
          </div>

          {/* Lado direito - dados Skeleton */}
          <div className="flex-1 flex flex-col gap-4">
            <div className="w-3/4 h-8 bg-gray-700 rounded animate-pulse" />
            <div className="w-full h-4 bg-gray-700 rounded animate-pulse" />
            <div className="w-full h-4 bg-gray-700 rounded animate-pulse" />
            <div className="w-full h-4 bg-gray-700 rounded animate-pulse" />
            <div className="w-1/2 h-4 bg-gray-700 rounded animate-pulse" />
          </div>
        </div>
      </section>
    );

  return (
    <section
      id="resumo"
      className="bg-background-darker text-foreground px-4 sm:px-[5%] lg:px-[15%] py-10 flex flex-col gap-8"
    >
      {/* Título */}
      <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4">
        Resumo do Pesquisador
      </h2>

      <div className="flex flex-col md:flex-row gap-6">
        {/* Lado esquerdo - imagem + ORCID */}
        <div className="flex flex-col items-center md:items-start gap-3">
          <div className="relative">
            <Image
              src="/images/orcidImage.png"
              alt="ORCID"
              width={130}
              height={130}
              className="select-none border-4 border-primary-yellow rounded-full"
            />
            <CheckCircle
              className="absolute bottom-0 right-0 bg-background-darker rounded-full text-green-400"
              size={24}
            />
          </div>
          <div className="flex items-center justify-center gap-2 text-sm text-primary-green font-mono">
            {pesquisador.id}
          </div>
        </div>

        {/* Lado direito - dados */}
        <div className="flex-1 flex flex-col gap-4">
          <h3 className="text-2xl font-bold">{pesquisador.nome}</h3>
          {/* <p>
            <span className="font-semibold">Área principal:</span>{' '}
            {pesquisador.areaPrincipal}
          </p> */}
          <div>
            <p className="font-semibold mb-2">Afiliações:</p>
            {pesquisador.afiliacoes.map((af, index) => (
              <div
                key={index}
                className="bg-clara-gray-dark p-4 rounded-md text-sm text-clara-gray-light mb-2"
              >
                <p className="text-foreground font-semibold">{af.instituicao}</p>
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