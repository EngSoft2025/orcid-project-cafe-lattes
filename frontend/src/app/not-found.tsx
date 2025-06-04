"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation"; // Import correto para App Router

export default function Custom404() {
  const router = useRouter();

  useEffect(() => {
    // Redireciona para a rota padrão após 2 segundos
    const timer = setTimeout(() => {
      router.push("/"); // Redireciona para a página inicial
    }, 2000);

    return () => clearTimeout(timer); // Limpa o timer ao desmontar o componente
  }, [router]);

  return (
    <div className="flex flex-col items-center justify-center h-screen bg-background-darker text-foreground">
      <h1 className="text-4xl font-bold">404 - Página Não Encontrada</h1>
      <p className="text-lg mt-4">Redirecionando para a página inicial...</p>
    </div>
  );
}