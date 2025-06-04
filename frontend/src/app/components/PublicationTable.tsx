// components/PublicationTable.tsx

'use client';

import { useEffect, useMemo, useState } from "react";
import { Search, LinkIcon } from "lucide-react";
import type { Publicacao } from "@/app/types";

// hook de debounce genérico (já existia)
function useDebounce<T>(value: T, delay = 400): T {
  const [debounced, setDebounced] = useState(value);
  useEffect(() => {
    const timer = setTimeout(() => setDebounced(value), delay);
    return () => clearTimeout(timer);
  }, [value, delay]);
  return debounced;
}

interface PublicationTableProps {
  publicacoes: Publicacao[];
}

export default function PublicationTable({ publicacoes }: PublicationTableProps) {
  const [paginaAtual, setPaginaAtual] = useState(1);
  const [busca, setBusca] = useState("");
  const debouncedBusca = useDebounce(busca, 500);

  const [anoFiltro, setAnoFiltro] = useState<string>("Todos os anos");
  const [tipoFiltro, setTipoFiltro] = useState<string>("Todos os tipos");
  const [carregando, setCarregando] = useState(true);

  const porPagina = 20;

  // 1) Filtrar publicações por “debouncedBusca”, “anoFiltro” e “tipoFiltro”
  const publicacoesFiltradas = useMemo(() => {
    let arr = [...publicacoes];

    // 1.1) Filtro de busca no título (case‐insensitive)
    if (debouncedBusca.trim()) {
      const termo = debouncedBusca.trim().toLowerCase();
      arr = arr.filter((p) => p.titulo.toLowerCase().includes(termo));
    }

    // 1.2) Filtro de ano
    if (anoFiltro !== "Todos os anos") {
      const anoNum = Number(anoFiltro);
      arr = arr.filter((p) => p.ano === anoNum);
    }

    // 1.3) Filtro de tipo
    if (tipoFiltro !== "Todos os tipos") {
      arr = arr.filter((p) => p.tipo === tipoFiltro);
    }

    return arr;
  }, [publicacoes, debouncedBusca, anoFiltro, tipoFiltro]);

  // 2) Total após filtrar
  const total = publicacoesFiltradas.length;

  // 3) Calcular páginas e itens atuais
  const totalPaginas = Math.ceil(total / porPagina);
  const startIndex = (paginaAtual - 1) * porPagina;
  const endIndex = startIndex + porPagina;
  const paginaItens = publicacoesFiltradas.slice(startIndex, endIndex);

  // 4) Quando qualquer critério mudar (busca/filtros), volta para página 1
  useEffect(() => {
    setPaginaAtual(1);
  }, [debouncedBusca, anoFiltro, tipoFiltro]);

  // 5) Simula um pequeno delay para “carregando...” (opcional)
  useEffect(() => {
    setCarregando(true);
    const t = setTimeout(() => setCarregando(false), 200); // 200ms de “spinner”
    return () => clearTimeout(t);
  }, [publicacoesFiltradas, paginaAtual]);

  return (
    <section
      className="bg-background-darker mt-30 text-foreground px-4 sm:px-[5%] lg:px-[15%] py-10 flex flex-col gap-8 pb-50"
      id="publicacoes"
    >
      {/* Título */}
      <h2 className="text-xl font-semibold relative before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 before:w-1 before:h-6 before:bg-primary-yellow pl-4">
        Publicações
      </h2>

      <div className="rounded-md sm:px-[5%] py-8 text-foreground text-sm">
        {/* Busca e filtros */}
        <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4 mb-4">
          {/* Busca com debounce */}
          <div className="relative w-full md:max-w-md">
            <input
              type="text"
              placeholder="Buscar publicações..."
              value={busca}
              onChange={(e) => setBusca(e.target.value)}
              className="w-full px-4 py-2 bg-background-dark border border-border-color rounded-md placeholder-clara-gray-light text-foreground text-sm"
            />
            <Search className="absolute right-3 top-2.5 text-clara-gray-light" size={18} />
          </div>

          {/* Filtros de ano e tipo */}
          <div className="flex gap-4">
            <select
              value={anoFiltro}
              onChange={(e) => setAnoFiltro(e.target.value)}
              className="bg-background-dark border border-border-color text-sm px-3 py-2 rounded-md text-clara-gray"
            >
              <option>Todos os anos</option>
              {/* Para popular dinamicamente, você poderia gerar options a partir de `publicacoes` */}
              {[...new Set(publicacoes.map((p) => p.ano))]
                .filter((y) => y > 0)
                .sort((a, b) => b - a)
                .map((ano) => (
                  <option key={ano}>{ano}</option>
                ))}
            </select>
            <select
              value={tipoFiltro}
              onChange={(e) => setTipoFiltro(e.target.value)}
              className="bg-background-dark border border-border-color text-sm px-3 py-2 rounded-md text-clara-gray"
            >
              <option>Todos os tipos</option>
              {/* Popula dinamicamente com todos os tipos presentes */}
              {[...new Set(publicacoes.map((p) => p.tipo))]
                .sort()
                .map((tipo) => (
                  <option key={tipo}>{tipo}</option>
                ))}
            </select>
          </div>
        </div>

        {/* Tabela */}
        <div className="overflow-x-auto rounded-md">
          <table className="w-full text-sm">
            <thead className="bg-background-dark text-xs border border-border-color text-clara-gray uppercase">
              <tr className="border-b border-border-color bg-background">
                <th className="px-4 py-3 text-left">Título</th>
                <th className="px-4 py-3 text-left">Tipo</th>
                <th className="px-4 py-3 text-left">Ano</th>
                <th className="px-4 py-3 text-left">Periódico/Conferência</th>
                <th className="px-4 py-3 text-left">Link</th>
              </tr>
            </thead>
            <tbody>
              {carregando ? (
                <tr>
                  <td colSpan={5} className="text-center py-6 text-clara-gray">
                    Carregando...
                  </td>
                </tr>
              ) : paginaItens.length === 0 ? (
                <tr>
                  <td colSpan={5} className="text-center py-6 text-clara-gray">
                    Nenhuma publicação encontrada.
                  </td>
                </tr>
              ) : (
                paginaItens.map((pub) => (
                  <tr
                    key={pub.id}
                    className="border-b border-border-color hover:bg-background-dark"
                  >
                    <td className="px-4 py-3">{pub.titulo}</td>
                    <td className="px-4 py-3">{pub.tipo}</td>
                    <td className="px-4 py-3">{pub.ano || "-"}</td>
                    <td className="px-4 py-3">{pub.periodico || "-"}</td>
                    <td className="px-4 py-3">
                      {pub.doi ? (
                        <a
                          href={pub.doi}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-primary-yellow flex items-center gap-1 hover:underline"
                        >
                          <LinkIcon size={14} /> DOI
                        </a>
                      ) : (
                        <span className="text-clara-gray">—</span>
                      )}
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>

        {/* Paginação */}
        <div className="flex flex-col sm:flex-row items-center justify-between mt-4 text-sm text-clara-gray">
          <span className="mb-2 sm:mb-0">
            {paginaItens.length > 0
              ? `Mostrando ${startIndex + 1}–${
                  startIndex + paginaItens.length
                } de ${total}`
              : `Nenhum resultado`}
          </span>
          <div className="flex gap-2 flex-wrap">
            <button
              className="px-3 py-1 rounded bg-clara-gray-dark text-clara-gray disabled:opacity-50"
              onClick={() => setPaginaAtual((p) => Math.max(p - 1, 1))}
              disabled={paginaAtual === 1}
            >
              Anterior
            </button>
            {Array.from({ length: totalPaginas }, (_, i) => i + 1).map((n) => (
              <button
                key={n}
                className={`px-3 py-1 rounded cursor-pointer ${
                  n === paginaAtual
                    ? "bg-primary-yellow text-black"
                    : "bg-clara-gray-dark text-clara-gray"
                }`}
                onClick={() => setPaginaAtual(n)}
              >
                {n}
              </button>
            ))}
            <button
              className="px-3 py-1 rounded bg-clara-gray-dark text-clara-gray disabled:opacity-50 cursor-pointer"
              onClick={() => setPaginaAtual((p) => Math.min(p + 1, totalPaginas))}
              disabled={paginaAtual === totalPaginas}
            >
              Próxima
            </button>
          </div>
        </div>
      </div>
    </section>
  );
}
