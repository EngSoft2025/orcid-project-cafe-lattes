'use client';

import { ArrowUp } from 'lucide-react';
import Link from 'next/link';

export default function Footer() {
  const year = new Date().getFullYear();

  /** rolagem suave até o topo da página */
  const scrollToTop = () =>
    window.scrollTo({ top: 0, behavior: 'smooth' });

  return (
    <footer className="bg-background-darker text-foreground w-full px-4 sm:px-[5%] lg:px-[10%] py-12 flex flex-col gap-8 border-t-1 border-border-color">
      {/* ===== Linha superior ===== */}
      <div className="flex flex-col md:flex-row md:items-start md:justify-between gap-6">
        {/* Bloco esquerdo */}
        <div className="space-y-2">
          <div className="flex gap-2 items-end">
            <h2 className="font-bold text-xl">C.L.A.R.A.</h2>
            <span className="text-primary-yellow text-sm font-medium">ORCID Beacon</span>
          </div>
          <p className="text-sm text-clara-gray-light">
            Versão 1.0.0 — Conectando Pesquisadores e Métricas
          </p>
        </div>

        {/* Bloco direito (links) */}
        <div className="flex items-center gap-6">
          <button
            onClick={scrollToTop}
            className="flex items-center gap-2 border border-border-color rounded-md px-4 py-2 text-sm text-clara-gray-light hover:bg-clara-gray-dark/20 transition"
          >
            Voltar ao topo <ArrowUp size={14} strokeWidth={2} />
          </button>

          <Link
            href="https://info.orcid.org/what-is-orcid/" target="_blank"
            className="text-sm text-clara-gray-light hover:text-primary-yellow-hover transition"
          >
            Sobre ORCID
          </Link>

        </div>
      </div>

      {/* ===== Informações legais (centralizadas) ===== */}
      <div className="text-center space-y-1 text-xs text-clara-gray-light">
        <p>
          © {year} C.L.A.R.A. ORCID Beacon — Este projeto utiliza a API pública do ORCID
        </p>
        <p>
          ORCID é uma marca registrada de ORCID, Inc. Esta aplicação não é endossada pela ORCID, Inc.
        </p>
      </div>
    </footer>
  );
}
