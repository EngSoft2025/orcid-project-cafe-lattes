// components/Header.tsx

'use client';

import { useState } from "react";
import { Menu, X, Search as SearchIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { fetchRecordData } from "../services/api";

export default function Header() {
  const [menuOpen, setMenuOpen] = useState(false);
  const [searchBy, setSearchBy] = useState<'orcid' | 'name'>('orcid');
  const [searchInput, setSearchInput] = useState("");
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [isVerifying, setIsVerifying] = useState(false);

  const router = useRouter();

  // Mapeia cada item para o ID correspondente no DOM
  const sectionIds: Record<string, string> = {
    "Resumo": "resumo",
    "Publicações": "publicacoes",
    "Colaborações": "colaboracoes",
    "Métricas": "metricas",
    "Contato": "contato"
  };

  const handleScroll = (section: string) => {
    const element = document.getElementById(section);
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' });
      setMenuOpen(false);
    }
  };

  const handleSearch = async () => {
    setErrorMessage(null);

    // === 1) Busca por ORCID ===
    if (searchBy === 'orcid') {
      const orcidId = searchInput.trim();
      if (!orcidId) {
        setErrorMessage("Por favor, insira um ORCID válido.");
        return;
      }
      
      setIsVerifying(true);
      
      try {
        // CHAMA fetchRecordData, que já faz a requisição única
        await fetchRecordData(orcidId);
        // Se não lançou erro, significa que encontrou o pesquisador
        router.push(`/Researcher/${orcidId}`);
      } catch (err) {
        // fetchRecordData lança erro se res.ok for false (por exemplo 404)
        setErrorMessage("Pesquisador não encontrado para este ORCID.");
      } finally {
        setIsVerifying(false);
      }

    // === 2) Busca por NOME ===
    } else {
      const nome = searchInput.trim();
      if (!nome) {
        setErrorMessage("Por favor, insira um nome para buscar.");
        return;
      }
      // Redireciona para uma página genérica de resultados de busca
      router.push(`/search?name=${encodeURIComponent(nome)}`);
    }
  };

  return (
    <header className="w-full bg-background-darker text-foreground px-4 sm:px-[5%] lg:px-[10%] py-6 flex flex-col gap-4">
      {/* Top Bar */}
      <div className="flex items-center justify-between">
        {/* Logo */}
        <div className="flex gap-2 items-end">
          <h1 className="font-bold text-2xl">C.L.A.R.A</h1>
          <p className="text-primary-yellow text-xs font-medium">ORCID Beacon</p>
        </div>

        {/* Desktop Navigation */}
        <ul className="hidden md:flex gap-8 items-center">
          {["Resumo", "Publicações", "Colaborações", "Métricas", "Contato"].map((item) => (
            <li key={item} className="relative group overflow-hidden">
              <button 
                onClick={() => handleScroll(sectionIds[item])}
                className="text-clara-gray hover:text-primary-yellow-hover transition-colors font-semibold text-sm cursor-pointer"
              >
                {item}
                <span className="absolute left-0 bottom-0 w-0 h-[2px] bg-primary-yellow group-hover:w-full transition-all duration-300 ease-in-out"></span>
              </button>
            </li>
          ))}
        </ul>

        {/* Mobile Menu Toggle */}
        <button
          className="md:hidden text-clara-gray"
          onClick={() => setMenuOpen(!menuOpen)}
          aria-label="Abrir menu"
        >
          {menuOpen ? <X size={24} /> : <Menu size={24} />}
        </button>
      </div>

      {/* Mobile Menu */}
      {menuOpen && (
        <ul className="md:hidden flex flex-col gap-4 z-50 shadow-md bg-background-dark px-6 py-4">
          {["Resumo", "Publicações", "Colaborações", "Métricas", "Contato"].map((item) => (
            <li key={item}>
              <button 
                onClick={() => handleScroll(sectionIds[item])}
                className="text-clara-gray hover:text-primary-yellow-hover transition font-medium"
              >
                {item}
              </button>
            </li>
          ))}
        </ul>
      )}

      {/* Search Bar */}
      <div className="w-full flex flex-col my-5">
        {/* Tabs de tipo de busca */}
        <div className="flex overflow-hidden gap-2">
          <button
            className={`px-4 py-2 rounded-t-md cursor-pointer transition-colors border text-sm font-medium ${
              searchBy === 'orcid'
                ? 'bg-primary-yellow text-black border-primary-yellow'
                : 'bg-clara-gray-dark text-foreground border-border-color'
            }`}
            onClick={() => {
              setSearchBy('orcid');
              setErrorMessage(null);
            }}
          >
            ORCID ID
          </button>

          <button
            className={`px-4 py-2 rounded-t-md cursor-pointer transition-colors border text-sm font-medium ${
              searchBy === 'name'
                ? 'bg-primary-yellow text-black border-primary-yellow'
                : 'bg-clara-gray-dark text-foreground border-border-color'
            }`}
            onClick={() => {
              setSearchBy('name');
              setErrorMessage(null);
            }}
          >
            Nome
          </button>
        </div>

        {/* Input + Botão */}
        <div className="mt-2 flex w-full items-center gap-3 md:flex-row">
          <div className="relative flex-1">
            <input
              type="text"
              value={searchInput}
              onChange={(e) => {
                setSearchInput(e.target.value);
                setErrorMessage(null);
              }}
              placeholder={
                searchBy === 'orcid'
                  ? 'Buscar por ORCID ID (xxxx-xxxx-xxxx-xxxx)'
                  : 'Digite um nome'
              }
              className="w-full px-4 py-2 bg-background border border-border-color rounded-r-md rounded-b-md text-sm text-foreground placeholder-clara-gray-light focus:outline-none focus:ring-2 focus:ring-primary-yellow focus:shadow-primary-yellow"
            />
            <SearchIcon className="absolute right-3 top-2.5 text-clara-gray-light" size={18} />
          </div>

          <button
            onClick={handleSearch}
            className={`flex items-center gap-2 px-4 py-2 rounded-md text-sm font-medium transition ${
              isVerifying
                ? 'bg-clara-gray-dark text-clara-gray cursor-not-allowed'
                : 'bg-primary-yellow text-black hover:bg-primary-yellow-hover'
            }`}
            disabled={isVerifying}
          >
            {isVerifying ? "Verificando…" : (
              <>
                <SearchIcon size={16} /> Buscar
              </>
            )}
          </button>
        </div>

        {/* Mensagem de erro, se houver */}
        {errorMessage && (
          <p className="mt-2 text-sm text-red-500">{errorMessage}</p>
        )}
      </div>

      <div className="w-full border border-border-color" />
    </header>
  );
}
